package command

import (
	"github.com/urfave/cli"
	"log"
	"os"
)

type Args struct {
	Host string
	Port int64
	Username string
	Password string
	BinlogFileName string
	BinlogPosition int64
}

func ParseArgs() Args {
	app := cli.NewApp()
	app.Flags = []cli.Flag {
		cli.StringFlag{
			Name: "host",
			Value: "127.0.0.1",
			Usage: "host of the database",
		},
		cli.Uint64Flag{
			Name: "port",
			Value: 3306,
			Usage: "port of the database",
		},
	}

	commandArgs := Args{
		Host: "127.0.0.1",
		Port: 3306,
	}

	app.Action = func(c *cli.Context) error {
		host := c.String("host")
		commandArgs.Host = host
		//log.Printf("host: %s", host)

		port := c.Int64("port")
		//log.Printf("port: %d", port)
		commandArgs.Port = port
		return nil
	}

	err := app.Run(os.Args)
	if err != nil {
		log.Fatal(err)
	}
	return commandArgs
}

