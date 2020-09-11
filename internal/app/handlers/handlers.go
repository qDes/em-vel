package handlers

import (
	"em-vel/internal/app/halcmd"
	"log"
	"net/http"
)

func RunCmd(w http.ResponseWriter, r *http.Request) {
	params := r.URL.Query()

	command := params["command"][0]
	val := params["val"][0]
	log.Println("command: ", command, " value: ", val)
	halcmd.HalCmd(command, val)
}