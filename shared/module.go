package shared

import (
	"crypto/x509"
	"fmt"
	"log"

	pb "github.com/larskluge/babl/protobuf"
	"golang.org/x/net/context"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"
)

type Module struct {
	Name    string
	Address string
	Env     map[string]string
}

func (m *Module) Call(stdin []byte) (stdout []byte, stderr []byte, exitcode int, err error) {
	conn := m.Connect()
	defer conn.Close()

	connection := pb.Modules[m.Name].Client(conn)
	req := pb.BinRequest{Stdin: stdin, Env: m.Env}
	res, err := connection.IO(context.Background(), &req)
	if err != nil {
		log.Fatalf("Failed: %v", err)
	}
	status := "SUCCESS"
	if res.Exitcode != 0 {
		status = "ERROR"
	}
	log.Printf("Module finished: %s. %d bytes stdout, %d bytes stderr:", status, len(res.Stdout), len(res.Stderr))
	log.Print(string(res.Stderr))
	fmt.Printf("%s", res.Stdout)

	return res.Stdout, res.Stderr, int(res.Exitcode), err
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
	opts := []grpc.DialOption{grpc.WithTransportCredentials(creds)}
	conn, err := grpc.Dial(m.Address, opts...)
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	return conn
}

func (m *Module) Ping() (res *pb.Pong, err error) {
	conn := m.Connect()
	defer conn.Close()

	connection := pb.Modules[m.Name].Client(conn)
	req := pb.Empty{}
	res, err = connection.Ping(context.Background(), &req)
	return
}
