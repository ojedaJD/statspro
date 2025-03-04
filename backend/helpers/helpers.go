package helpers

import (
	"encoding/json"
	"fmt"
	"os/exec"
)

func Jsonify(i interface{}) {
	indent, err := json.MarshalIndent(i, "", " ")
	if err != nil {
		return
	}
	fmt.Println(string(indent))
	err = copy(string(indent))
	if err != nil {
		fmt.Println(err)
	}
}

func copy(content string) error {
	cmd := exec.Command("pbcopy")
	in, err := cmd.StdinPipe()
	if err != nil {
		return err
	}
	if err := cmd.Start(); err != nil {
		return err
	}
	if _, err := in.Write([]byte(content)); err != nil {
		return err
	}
	if err := in.Close(); err != nil {
		return err
	}
	return cmd.Wait()
}
