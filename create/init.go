package create

import (
	"fmt"
	"github.com/duxweb/duxtool/helper"
	"github.com/go-git/go-git/v5"
	"github.com/gookit/gcli/v3"
	"github.com/gookit/gcli/v3/interact"
	"github.com/gookit/goutil/fsutil"
	"os"
	"os/exec"
	"path/filepath"
)

func InitCommand(c *gcli.Command, args []string) error {

	dirname := ""
	if len(args) == 0 {
		dirname, _ = interact.ReadLine("<green>?</> project name: ")
	}

	if dirname == "" {
		return helper.MessageError("please enter a project name")
	}

	projectFolderPath := filepath.Join(".", dirname)

	err := fsutil.Mkdir(projectFolderPath, 0755)
	if err != nil {
		return helper.MessageError(err)
	}

	if !fsutil.IsEmptyDir(projectFolderPath) {
		return helper.MessageError("Project directory is not empty")
	}

	_, err = git.PlainClone(projectFolderPath, false, &git.CloneOptions{
		URL:      fmt.Sprintf("https://github.com/%s/%s", "duxweb", "duxgo"),
		Progress: os.Stdout,
	})

	if err != nil {
		return helper.MessageError(err)
	}

	cmd := exec.Command("rm", "-rf", filepath.Join(projectFolderPath, ".git"))
	cmd.Dir = projectFolderPath
	output, err := cmd.CombinedOutput()
	if err != nil {
		return helper.MessageError(output)
	}

	return helper.MessageSuccess("create complete")

}
