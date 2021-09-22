package services

import (
	"os"
	"os/exec"
)

func Terraform_init(path string) {
	cmd0 := "terraform"
	cmd1 := "-chdir=" + path
	cmd2 := "init"
	cmd := exec.Command(cmd0, cmd1, cmd2)
	cmd.Stdin = os.Stdin
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	err := cmd.Run()
	if err != nil {
		ColorPrint(ERROR, "%v", err)
	}

	ColorPrint(INFO, "Finished executing terraform init")
}

func Terraform_apply(path string) {
	cmd0 := "terraform"
	cmd1 := "-chdir=" + path
	cmd2 := "apply"
	cmd3 := "-auto-approve"
	cmd := exec.Command(cmd0, cmd1, cmd2, cmd3)
	cmd.Stdin = os.Stdin
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	err := cmd.Run()
	if err != nil {
		ColorPrint(ERROR, "%v", err)
	}

	ColorPrint(INFO, "Finished executing terraform apply")
}
