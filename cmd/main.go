package main

import (
	"fmt"
	"os"

	"github.com/google/uuid"
	"github.com/urfave/cli/v2"
)

func main() {
	app := &cli.App{
		Commands: []*cli.Command{
			{
				Name: "uuid",
				Action: func(c *cli.Context) error {
					uuid := uuid.New().String()
					fmt.Print(uuid)
					return nil
				},
			},
		},
	}

	app.Run(os.Args)
}
