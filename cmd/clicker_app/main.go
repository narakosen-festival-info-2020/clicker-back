package main

import (
	"fmt"
	"os"

	"github.com/narakosen-festival-info-2020/clicker-back/pkg/api"

	"github.com/facebookgo/pidfile"
)

func savePid() {
	dir, err := os.Getwd()
	if err != nil {
		fmt.Println(err.Error())
		return
	}

	path := fmt.Sprintf("%s/%s", dir, "tmp/pids/app.pid")
	pidfile.SetPidfilePath(path)
	err = pidfile.Write()
	if err != nil {
		fmt.Println(err.Error())
		return
	}
}

func removePid() {
	dir, err := os.Getwd()
	if err != nil {
		fmt.Println(err.Error())
		return
	}

	path := fmt.Sprintf("%s/%s", dir, "tmp/pids/app.pid")
	err = os.Remove(path)
	if err != nil {
		fmt.Println(err.Error())
		return
	}
}

func main() {
	savePid()
	defer removePid()

	server := api.Generate()
	server.Up("/clicker")
}
