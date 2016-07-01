package shared

import (
	"crypto/x509"
	"fmt"
	"strings"
	"time"

	"github.com/larskluge/babl/log"
	pb "github.com/larskluge/babl/protobuf"
	pbm "github.com/larskluge/babl/protobuf/messages"
	"github.com/serenize/snaker"
	"golang.org/x/net/context"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"
)

type Module struct {
	Name    string
	Tag     string
	Address string
	Env     Env
	async   bool
	debug   bool
}

type Env map[string]string

func NewModule(name_with_tag string) *Module {
	tag := ""
	parts := strings.Split(name_with_tag, ":")
	name := parts[0]
	if len(parts) > 1 {
		tag = parts[1]
	}

	m := Module{
		Name:    name,
		Tag:     tag,
		Address: "babl.sh:4444",
		Env:     Env{},
	}
	m.loadDefaults()
	if !CheckModuleName(m.Name) {
		log.Fatal("Module name format incorrect")
	}
	return &m
}

func (m *Module) Fullname() string {
	return fmt.Sprintf("%s:%s", m.Name, m.Tag)
}

func (m *Module) Owner() string {
	parts := strings.SplitN(m.Name, "/", 2)
	return parts[0]
}

func (m *Module) GrpcModuleName() string {
	mod := strings.SplitN(m.Name, "/", 2)[1]
	snake := strings.Replace(mod, "-", "_", -1)
	return snaker.SnakeToCamel(snake)
}

func (m *Module) GrpcServiceName() string {
	return fmt.Sprintf("babl.%s.%s", m.Owner(), m.GrpcModuleName())
}

func (m *Module) KafkaTopicName(method string) string {
	return fmt.Sprintf("%s.%s", m.GrpcServiceName(), method)
}

func (m *Module) loadDefaults() {
	mod, ok := Config().Defaults[m.Fullname()]
	if ok {
		m.Env = mod.Env
	}
}

func (m *Module) GetAsync() bool {
	return m.async
}

func (m *Module) SetAsync(val bool) {
	m.async = val
	if val {
		m.Env["BABL_ASYNC"] = "true"
	} else {
		delete(m.Env, "BABL_ASYNC")
	}
}

func (m *Module) GetDebug() bool {
	return m.debug
}

func (m *Module) SetDebug(val bool) {
	m.debug = val
	if val {
		m.Env["BABL_DEBUG"] = "true"
	} else {
		delete(m.Env, "BABL_DEBUG")
	}
}

func (m *Module) Call(stdin []byte) (stdout []byte, stderr []byte, exitcode int, err error) {
	conn := m.Connect()
	defer conn.Close()

	connection := pb.BinaryClient(pb.NewBinaryClient(conn))

	req := pbm.BinRequest{Stdin: stdin, Env: m.Env}
	res, err := connection.IO(m.GrpcServiceName(), context.Background(), &req)
	if err == nil {
		exitcode := int(res.Exitcode)
		if exitcode != 0 {
			log.SetOutput(log.FatalWriter)
		}
		return res.Stdout, res.Stderr, exitcode, err
	} else {
		return nil, nil, 254, err
	}
}

func (m *Module) Connect() *grpc.ClientConn {
	data, err := Asset("data/ca.pem")
	if err != nil {
		log.Fatal("asset not found")
	}
	sn := "babl.test.youtube.com"
	cp := x509.NewCertPool()
	if !cp.AppendCertsFromPEM(data) {
		log.Fatal("credentials: failed to append certificates")
	}

	creds := credentials.NewClientTLSFromCert(cp, sn)
	opts := []grpc.DialOption{
		grpc.WithTransportCredentials(creds),
		grpc.WithTimeout(5 * time.Minute),
	}
	conn, err := grpc.Dial(m.Address, opts...)
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	return conn
}

func (m *Module) Ping() (res *pbm.Pong, err error) {
	conn := m.Connect()
	defer conn.Close()

	connection := pb.BinaryClient(pb.NewBinaryClient(conn))
	req := pbm.Empty{}
	res, err = connection.Ping(m.GrpcServiceName(), context.Background(), &req)
	return
}
