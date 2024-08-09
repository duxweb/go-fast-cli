package annotation

import (
	"github.com/duxweb/duxtool/helper"
	"github.com/gookit/gcli/v3"
	"github.com/gookit/gcli/v3/interact"
	"os"
)

func AnnotationCommand(c *gcli.Command, args []string) error {
	dirname := ""
	if len(args) == 0 {
		dirname, _ = interact.ReadLine("<green>?</> apps directory name: ")
	}

	if dirname == "" {
		return helper.MessageError("please enter a app directory name")
	}
	exeDir, _ := os.Getwd()
	Run(exeDir)

	return nil
}
