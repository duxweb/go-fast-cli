package helper

import "github.com/gookit/color"

func MessageError(content any) error {
	color.Redf("✖ %s\n", content)
	return nil
}

func MessageSuccess(content any) error {
	color.Greenf("✔ %s\n", content)
	return nil
}
