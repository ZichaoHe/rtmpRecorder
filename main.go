package main

import (
	"net/http"
	"rtmpRecorder/controller"
)

func init() {

}

func main() {
	http.HandleFunc("/v1/test", controller.Test)
	http.ListenAndServe(":8080", nil)
}
