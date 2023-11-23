package main

import (
	"fmt"
	"log"
	"os"

	"github.com/urfave/cli/v2"
)

func main() {

	app := &cli.App{
		Name:  "Healtcheker",
		Usage: "La siguiente aplicación se usa para saber si un sitio web está online u offline",
		Flags: []cli.Flag{
			&cli.StringFlag{
				Name:     "domain",
				Aliases:  []string{"d"},
				Usage:    "nombre del dominio a chequear.",
				Required: true,
			},
			&cli.StringFlag{
				Name:     "port",
				Aliases:  []string{"p"},
				Usage:    "Número de puerto a chequear.",
				Required: false,
			},
		},

		Action: func(c *cli.Context) error {
			port := c.String("port")

			if c.String("port") == "" {
				port = "80"
			}

			status := Check(c.String("domain"), port)

			fmt.Println(status)

			return nil
		},
	}

	err := app.Run(os.Args)

	if err != nil {
		log.Fatal(err)
	}
}
