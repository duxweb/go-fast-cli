package main

import (
	"github.com/duxweb/duxtool/annotation"
	"github.com/duxweb/duxtool/create"
	"github.com/duxweb/duxtool/gen"
	"github.com/gookit/gcli/v3"
)

func main() {

	app := gcli.NewApp()
	app.Version = "0.0.1"
	app.Desc = "duxgo command line tool"

	app.Add(&gcli.Command{
		Name: "create",
		Desc: "Initialize duxgo project",
		Func: create.InitCommand,
	})

	app.Add(&gcli.Command{
		Name: "annotation",
		Desc: "duxgo annotation index",
		Func: annotation.AnnotationCommand,
	})

	app.Add(&gcli.Command{
		Name: "gen",
		Desc: "duxgo code generator",
		Subs: []*gcli.Command{
			{
				Name: "app",
				Desc: "Generate application module",
				Func: gen.GenAppCommand,
			},
		},
	})

	app.Run(nil)
}
