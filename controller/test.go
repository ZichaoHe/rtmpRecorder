package controller

import (
	"fmt"
	"net/http"
)

func init() {

}

func Test(w http.ResponseWriter, r *http.Request) {
	fmt.Println(">>hzc>>:", r.URL)
	w.Write([]byte("hello world"))
}
