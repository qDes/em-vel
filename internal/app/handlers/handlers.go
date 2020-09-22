package handlers

import (
	"em-vel/internal/app/halcmd"
	"fmt"
	"log"
	"net/http"
	"time"
)

func RunCmd(w http.ResponseWriter, r *http.Request) {
	params := r.URL.Query()

	command := params["command"][0]
	val := params["val"][0]
	log.Println("command: ", command, " value: ", val)
	halcmd.HalCmd(command, val)
}

func Checker(w http.ResponseWriter, r *http.Request) {
	fmt.Println("check data")
}

func ChangeVeloParams(w http.ResponseWriter, r *http.Request) {
	params := r.URL.Query()
	for k, v := range params {
		cmd := halcmd.GetCmd(k)
		if cmd != "sasi" {
			halcmd.HalCmd(cmd, v[0])
			time.Sleep(100 * time.Millisecond)
		}
		fmt.Println(k, cmd, v)
	}
}
