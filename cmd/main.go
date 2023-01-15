package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"github.com/urfave/cli"
)

func main() {
	app := cli.NewApp()
	app.Name = "redis-cli"
	app.Usage = "A simple command line tool for interacting with Redis in protobuf binary format"
	app.Commands = []cli.Command{
		{
			Name:    "set",
			Aliases: []string{"s"},
			Usage:   "Set a key-value pair in Redis",
			Action: func(c *cli.Context) error {
				return nil
			},
			Flags: []cli.Flag{
				cli.StringFlag{
					Name:  "key, k",
					Value: "",
					Usage: "The key to set in Redis",
				},
				cli.StringFlag{
					Name:  "field1, f1",
					Value: "",
					Usage: "The value of field1 in protobuf message",
				},
				cli.StringFlag{
					Name:  "field2, f2",
					Value: "",
					Usage: "The value of field2 in protobuf message",
				},
			},
		},
	}

	for {
		fmt.Print("> ")
		input, _ := bufio.NewReader(os.Stdin).ReadString('\n')
		input = strings.TrimSpace(input)

		if input == "exit" {
			break
		}

		// pass the input to the cli app and let it handle the command
		if err := app.Run(strings.Fields(input)); err != nil {
			fmt.Println(err)
		}
	}
}
