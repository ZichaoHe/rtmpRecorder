package controller

import (
	"fmt"
	"net/http"
	"rtmpRecorder/serv/stream"
)

func init() {
	fmt.Println("recorder controller init...")
}

func StartRecord(w http.ResponseWriter, r *http.Request) {
	stream.GlbSm.StartRecord("", "", "")
}

func StopRecord(w http.ResponseWriter, r *http.Request) {

}
