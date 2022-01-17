package app

import (
	"net"

	"github.com/sirupsen/logrus"
	"github.com/urfave/cli"
)

// Gerar vai retornar a aplicação de linha de comando pronta para ser executada.
func Gerar() *cli.App {
	app := cli.NewApp()
	logrus.SetFormatter(&logrus.TextFormatter{ForceColors: true})

	app.Name = "Buscador de IPs e Nome"

	app.Usage = "Busca IPs e nome de servidores na internet."

	flags := []cli.Flag{
		cli.StringFlag{
			Name:  "host",
			Value: "google.com",
		},
	}

	app.Commands = []cli.Command{
		{
			Name:   "ip",
			Usage:  "Busca IPs de servidores na internet",
			Flags:  flags,
			Action: buscarIPs,
		},
		{
			Name:   "servidores",
			Usage:  "Busca o nome dos servidores na internet",
			Flags:  flags,
			Action: buscarServidores,
		},
	}

	return app
}

func buscarIPs(c *cli.Context) {
	host := c.String("host")

	ips, err := net.LookupIP(host)
	if err != nil {
		logrus.Fatal(err)
	}

	logrus.Info("Host: ", host)

	for _, ip := range ips {
		logrus.Info(ip)
	}
}

func buscarServidores(c *cli.Context) {
	host := c.String("host")

	servidores, err := net.LookupNS(host)
	if err != nil {
		logrus.Fatal(err)
	}

	logrus.Info("Host: ", host)

	for _, servidor := range servidores {
		logrus.Info(servidor.Host)
	}
}
