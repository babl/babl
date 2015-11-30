package main

import (
	"fmt"
	"log"

	"github.com/codegangsta/cli"
	pb "github.com/larskluge/babl/protobuf"
	"github.com/larskluge/babl/shared"
	"golang.org/x/net/context"
)

func appendModuleCommand(cmds *[]cli.Command, module string) {
	*cmds = append(*cmds, cli.Command{
		Name:  module,
		Usage: "MODULE",
		Action: func(c *cli.Context) {
			module := c.Command.Name
			defaultAction(c, module)
		},
		Flags: []cli.Flag{
			cli.StringSliceFlag{
				Name:  "env, e",
				Usage: "Send environment variables, e.g. -e FOO=42",
			},
		},
	})
}

func pingSubCommands() (cmds []cli.Command) {
	for _, module := range shared.Modules() {
		cmds = append(cmds, cli.Command{
			Name: module,
			Action: func(c *cli.Context) {
				module := c.Command.Name
				fmt.Print("ping.. ")
				conn := conn(address(c))
				defer conn.Close()
				connection := pb.Modules[module].Client(conn)
				req := pb.Empty{}
				res, err := connection.Ping(context.Background(), &req)
				if err == nil {
					fmt.Println(res.Val)
				} else {
					log.Fatalf("Failed: %v", err)
				}
			},
		})
	}
	return
}
