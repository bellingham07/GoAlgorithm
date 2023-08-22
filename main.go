/*
可以引⼊的库和版本相关请参考 “环境说明”
Please refer to the "Environmental Notes" for the libraries and versions that can be introduced.
*/

package main

import (
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
	"strconv"
	"time"
)

func main() {
	go create()
	go listen()
	time.Sleep(10 * time.Second)
}

func request() {
	resp, err := http.Get("http://localhost:8080/add")
	if err != nil {
		errors.New("bad request")
		return
	}
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		errors.New("read fail")
	}
	fmt.Println("resp", string(body))
}

func create() {
	http.HandleFunc("/add", func(w http.ResponseWriter, r *http.Request) {
		valueStr := r.URL.Query().Get("value")
		v, err := strconv.Atoi(valueStr)
		if err != nil {
			http.Error(w, "error", http.StatusBadRequest)
			return
		}
		res := v + 1
		fmt.Fprintf(w, "res:", res)
	})
}

func listen() {
	http.ListenAndServe(":8080", nil)
}
