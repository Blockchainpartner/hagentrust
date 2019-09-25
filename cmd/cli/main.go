package main

import (
	"fmt"
	"log"
	"os"

	"github.com/Blockchainpartner/hagentrust/pkg/server"
	"github.com/urfave/cli"
)

var (
	commands = []cli.Command{
		{
			Name:  "api",
			Usage: "",
			Flags: []cli.Flag{
				cli.StringFlag{
					Name:  "port, p",
					Usage: "Port where the API is served.",
					Value: "8080",
				},
			},
			Action: func(c *cli.Context) error {
				port := c.String("port")
				fmt.Println("[HagentrustCLI] Running API on port ", port)
				return server.Init(port)
			},
		},
	}
)

func main() {
	app := cli.NewApp()
	app.Commands = commands
	err := app.Run(os.Args)
	if err != nil {
		log.Fatal(err)
	}
}
