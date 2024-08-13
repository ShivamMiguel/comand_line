package app

import (
	"fmt"
	"log"
	"net"

	"github.com/urfave/cli"
)

func Gerar() *cli.App {
	app := cli.NewApp()
	app.Name = "command line application"
	app.Usage = "Find Ips in server in internet"

	flags := []cli.Flag{
		&cli.StringFlag{
			Name:  "host",
			Value: "devbook.com.br",
		},
	}

	app.Commands = []cli.Command{
		{
            Name:   "ip",
            Usage:  "Find IPs in a server",
            Flags: flags, 
            Action: buscarIps,
        },
		{
			Name:   "servers",
            Usage:  "Find servers name",
            Flags: flags, 
            Action: buscarServers,
		},
	}

	return app
}

func buscarIps(c *cli.Context){
	host := c.String("host")
   
	ips, erro := net.LookupIP(host)
	if erro!= nil {
        log.Fatal(erro)
    }

	for _, ip := range ips {
        fmt.Println(ip)
    }
}

func buscarServers(c *cli.Context){
	host := c.String("host")
   
	servers, erro := net.LookupNS(host)
	if erro!= nil {
        log.Fatal(erro)
    }

	for _, server := range servers {
        fmt.Println(server.Host)
    }
}