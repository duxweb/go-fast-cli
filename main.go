package main

import (
	"fmt"
	"log"
	"os"

	. "github.com/dave/jennifer/jen"
	"github.com/urfave/cli/v2"
)

func main() {
	app := &cli.App{
		Name: "dux",
		Commands: []*cli.Command{
			{
				Name:  "generate:app",
				Usage: "generate app application",
				Flags: []cli.Flag{
					&cli.StringFlag{
						Name:  "name",
						Usage: "app name",
					},
				},
				Action: func(cCtx *cli.Context) error {
					fmt.Println("added task: ", cCtx.Args().First())
					f := NewFile("main")
					f.Func().Id("main").Params().Block(
						Qual("fmt", "Println").Call(Lit("Hello, world")),
					)
					fmt.Printf("%#v", f)
					return nil
				},
			},
		},
	}

	if err := app.Run(os.Args); err != nil {
		log.Fatal(err)
	}
}
