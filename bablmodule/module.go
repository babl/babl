package bablmodule

import (
	"bytes"
	"crypto/x509"
	"fmt"
	"strings"
	"time"
	"unicode"

	log "github.com/Sirupsen/logrus"
	"github.com/larskluge/babl-storage/download"
	"github.com/larskluge/babl-storage/upload"
	pb "github.com/larskluge/babl/protobuf"
	pbm "github.com/larskluge/babl/protobuf/messages"
	"golang.org/x/net/context"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"
)

const (
	MaxPayloadSize         = 1024 * 512 // 512kb
	DefaultBablEndpoint    = "babl.sh:4444"
	DefaultStorageEndpoint = "babl.sh:4443"
)

type Module struct {
	Name            string
	Tag             string
	Env             Env
	async           bool
	debug           bool
	IncludePayload  bool
	endpoint        string
	storageEndpoint string
}

type Env map[string]string

func New(name_with_tag string) *Module {
	tag := ""
	parts := strings.Split(name_with_tag, ":")
	name := parts[0]
	if len(parts) > 1 {
		tag = parts[1]
	}

	m := Module{
		Name:           name,
		Tag:            tag,
		Env:            Env{},
		IncludePayload: true,
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

func (m *Module) Endpoint() string {
	if m.endpoint == "" {
		return DefaultBablEndpoint
	} else {
		return m.endpoint
	}
}

func (m *Module) SetEndpoint(e string) {
	m.endpoint = e
}

func (m *Module) StorageEndpoint() string {
	if m.storageEndpoint == "" {
		return DefaultStorageEndpoint
	} else {
		return m.storageEndpoint
	}
}

func (m *Module) SetStorageEndpoint(e string) {
	m.storageEndpoint = e
}

func (m *Module) Owner() string {
	parts := strings.SplitN(m.Name, "/", 2)
	return parts[0]
}

func (m *Module) GrpcModuleName() (res string) {
	mod := strings.SplitN(m.Name, "/", 2)[1]
	words := strings.Split(mod, "-")
	for _, word := range words {
		w := []rune(word)
		w[0] = unicode.ToUpper(w[0])
		res += string(w)
	}
	return res
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

func (m *Module) Call(stdin []byte) ([]byte, []byte, int, string, error) {
	conn := m.Connect()
	defer conn.Close()

	connection := pb.BinaryClient(pb.NewBinaryClient(conn))
	req := pbm.BinRequest{Stdin: stdin, Env: m.Env}

	if len(stdin) > MaxPayloadSize {
		up, err := upload.New(m.StorageEndpoint(), bytes.NewReader(stdin))
		check(err)
		go up.WaitForCompletion()
		req.Stdin = nil
		req.PayloadUrl = up.Url
	}

	res, err := connection.IO(m.GrpcServiceName(), context.Background(), &req)
	if err == nil {
		exitcode := int(res.Exitcode)
		if res.PayloadUrl != "" && m.IncludePayload {
			res.Stdout, err = download.Download(res.PayloadUrl)
			check(err)
		}
		return res.Stdout, res.Stderr, exitcode, res.PayloadUrl, err
	} else {
		return nil, nil, 254, "", err
	}
}

func (m *Module) Connect() *grpc.ClientConn {
	data, err := Asset("ca.pem")
	check(err)
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
	conn, err := grpc.Dial(m.Endpoint(), opts...)
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
