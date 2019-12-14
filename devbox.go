package main

import (
	"fmt"
	"github.com/urfave/cli/v2"
	"log"
	"os"
)

func main() {
	app := &cli.App{
		Name:  "devbox",
		Usage: "Swiss Knife full of PHP Development Environments",
		Version: "2.1.0",
		Commands: []*cli.Command{
			{
				Name:    "init",
				Usage:   "Install all necessary software to the machine",
				Action: func(c *cli.Context) error {
					fmt.Println("Init")
					return nil
				},
			},
			{
				Name:    "create",
				Usage:   "Create a project from scratch using one of available templates",
				Action: func(c *cli.Context) error {
					fmt.Println("Create")
					return nil
				},
			},
			{
				Name:    "list",
				Usage:   "Print a list of available projects",
				Action: func(c *cli.Context) error {
					fmt.Println("List")
					return nil
				},
			},
			{
				Name:    "stop",
				Usage:   "Stop a particular project or all projects at once",
				Action: func(c *cli.Context) error {
					fmt.Println("Stop")
					return nil
				},
			},
			{
				Name:    "start",
				Usage:   "Start a particular project or all projects at once",
				Action: func(c *cli.Context) error {
					fmt.Println("Start")
					return nil
				},
			},
			{
				Name:    "kill",
				Usage:   "Kill a particular project or all projects at once",
				Action: func(c *cli.Context) error {
					fmt.Println("Kill")
					return nil
				},
			},
			{
				Name:    "ssh",
				Usage:   "Tunnel to the App container of the project",
				Action: func(c *cli.Context) error {
					fmt.Println("SSH")
					return nil
				},
			},
		},

		Action: func(c *cli.Context) error {
			fmt.Println("No action selected.")
			return nil
		},
	}

	err := app.Run(os.Args)
	if err != nil {
		log.Fatal(err)
	}
}
