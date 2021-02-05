package main

import (
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"os"
	"time"
)

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		var res string
		var err error
		switch r.Method {
		case "GET":
			res = fmt.Sprintf("method: %s, query: %s", r.Method, r.URL.Query())

		case "POST":
			var body []byte
			body, err = ioutil.ReadAll(r.Body)
			res = string(body)
			fmt.Printf(res)
		default:
			res = fmt.Sprintf("Not Allow Method: %s", r.Method)
			w.WriteHeader(http.StatusBadRequest)

		}

		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			_, _ = w.Write([]byte("Status Internal Server Error"))
			return
		}
		_, _ = w.Write([]byte(res))
	})

	http.HandleFunc("/with_time_out/", func(w http.ResponseWriter, r *http.Request) {
		var res string
		switch r.Method {
		case "GET":
			time.Sleep(10 * time.Second)
			res = "method: GET"
		default:
			res = fmt.Sprintf("Not Allow Method: %s", r.Method)
			w.WriteHeader(http.StatusBadRequest)

		}
		_, _ = w.Write([]byte(res))
	})

	http.HandleFunc("/upload/", func(w http.ResponseWriter, r *http.Request) {
		err := r.ParseMultipartForm(32 << 20)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			_, _ = w.Write([]byte("Status Internal Server Error"))
			return
		}

		file, handler, err := r.FormFile("uploadfile")
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			_, _ = w.Write([]byte("Status Internal Server Error"))
			return
		}
		defer file.Close()

		_, err = fmt.Fprintf(w, "%v", handler.Header)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			_, _ = w.Write([]byte("Status Internal Server Error"))
			return
		}

		f, err := os.OpenFile("./"+handler.Filename, os.O_WRONLY|os.O_CREATE, 0666)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			_, _ = w.Write([]byte("Status Internal Server Error"))
			return
		}

		defer f.Close()
		_, err = io.Copy(f, file)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			_, _ = w.Write([]byte("Status Internal Server Error"))
			return
		}

		_, _ = w.Write([]byte("success"))
	})

	_ = http.ListenAndServe(":8080", nil)
}
