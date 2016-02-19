package shared

import (
	"crypto/x509"
	"fmt"
	"strings"
	"time"

	"github.com/larskluge/babl/log"
	pb "github.com/larskluge/babl/protobuf"
	pbm "github.com/larskluge/babl/protobuf/messages"
	"golang.org/x/net/context"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"
)

type Module struct {
	Name    string
	Tag     string
	Address string
	Env     map[string]string
	Debug   bool
}

func NewModule(name_with_tag string, debug bool) *Module {
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
		Env:     make(map[string]string),
		Debug:   debug,
	}
	m.loadDefaults()
	if debug {
		m.Env["BABL_DEBUG"] = "true"
	}
	EnsureModuleExists(m.Name)
	return &m
}

func (m *Module) Fullname() string {
	return fmt.Sprintf("%s:%s", m.Name, m.Tag)
}

func (m *Module) loadDefaults() {
	mod, ok := Config().Defaults[m.Fullname()]
	if ok {
		m.Env = mod.Env
	}
}

func (m *Module) Call(stdin []byte) (stdout []byte, stderr []byte, exitcode int, err error) {
	conn := m.Connect()
	defer conn.Close()

	connection := pb.Modules[m.Name].Client(conn)
	req := pbm.BinRequest{Stdin: stdin, Env: m.Env}
	res, err := connection.IO(context.Background(), &req)
	if err == nil {
		return res.Stdout, res.Stderr, int(res.Exitcode), err
	} else {
		return nil, nil, 255, err
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

	connection := pb.Modules[m.Name].Client(conn)
	req := pbm.Empty{}
	res, err = connection.Ping(context.Background(), &req)
	return
}
