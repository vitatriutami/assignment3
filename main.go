package main

import (
	"github.com/husfuu/go-auto-reload/handler"
	"github.com/husfuu/go-auto-reload/helper"
	"github.com/husfuu/go-auto-reload/service"
	"net/http"
	"time"
)

func main() {
	go autoUpdate()
	http.HandleFunc("/", handler.Handler)

	err := http.ListenAndServe(":8000", nil)
	if err != nil {
		panic(err)
	}
}

func autoUpdate() {
	disaster := helper.GetDisaster()
	for range time.Tick(15 * time.Second) {
		service.UpdateJSON(disaster)
	}
}
