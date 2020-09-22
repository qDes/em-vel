package halcmd

import (
	"bytes"
	"log"
	"os/exec"
	"strings"
)

func HalCmd(command, val string) {
	cmd := exec.Command("halcmd", "setp", command, val)
	cmd.Stdin = strings.NewReader("")
	var out bytes.Buffer
	cmd.Stdout = &out
	err := cmd.Run()
	if err != nil {
		log.Fatal(err)
	}
}

func GetCmd(par string) string{
	switch par {
	case "friction":
		return "inner-loop-velo.0.friction"
	case "calib":
		return "inner-loop-velo.0.out-offset"
	case "shaker_freq":
		return "inner-loop-velo.0.shaker-freq"
	case "m_inner":
		return "inner-loop-velo.0.m-inner"
	case "f_set":
		return "inner-loop-velo.0.out-lim"
	default:
		return "sasi"
	}


}