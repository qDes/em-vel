package halcmd

import (
	"bytes"
	"log"
	"os/exec"
	"strings"
)

func HalCmd(command, val string) {
	cmd := exec.Command("halcmd", command, val)
	cmd.Stdin = strings.NewReader("")
	var out bytes.Buffer
	cmd.Stdout = &out
	err := cmd.Run()
	if err != nil {
		log.Fatal(err)
	}
}