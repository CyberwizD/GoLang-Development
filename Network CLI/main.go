package main

import (
	"fmt"
	"log"
	"net"
	"os"

	"github.com/urfave/cli"
)

func main() {
	app := cli.NewApp()
	app.Name = "Network Diagnostics Tool"
	app.Usage = "Performs network lookups and diagnostics."

	cFlags := []cli.Flag{
		cli.StringFlag{
			Name:  "host",
			Value: "google.com",
		},
	}

	app.Commands = []cli.Command{
		{
			Name:  "ns",
			Usage: "Lookup Name Servers for a domain.",
			Flags: cFlags,
			Action: func(c *cli.Context) error {
				ns, err := net.LookupNS(c.String("host"))

				if err != nil {
					return err
				}

				for _, n := range ns {
					fmt.Println(n.Host)
				}

				return nil
			},
		},
	}

	err := app.Run(os.Args)

	if err != nil {
		log.Fatal(err)
	}
}
