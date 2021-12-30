package main

import (
	"bytes"
	"fmt"
	"html/template"
	"io/ioutil"
	"net/http"
	"os"
	"path/filepath"
	"strings"
	"time"
)

type Message struct {
	Text string
}

var files []string

func ExamineFiles() {
	err := filepath.Walk("./views", func(path string, info os.FileInfo, err error) error {
		files = append(files, path)
		return nil
	})
	if err != nil {
		panic(err)
	}

}

type FileResponse struct {
	fileio []byte
	status error
}

func ReadFileAsync(path string, resp chan<- FileResponse) {
	fileio, err := ioutil.ReadFile("views/" + path)
	f := FileResponse{
		fileio: fileio,
		status: err,
	}
	resp <- f

}

func welcome(w http.ResponseWriter, r *http.Request) {
	path := r.URL.EscapedPath()
	path = strings.Replace(path, string(filepath.Separator), "", -1)
	if path == "" {
		t := template.New("welcome")

		var tpl bytes.Buffer
		msg := r.URL.Query().Get("message")
		if msg != "" {
			t, _ = t.Parse("<H1> Howdy {{.Text}}</H1>")

		} else {
			t, _ = t.Parse("<H1> Hello World !!</H1>")
		}
		// we can also apply if logic in template , for simplicity , selecting this
		t.Execute(&tpl, Message{Text: msg})
		fmt.Fprint(w, tpl.String())
	} else {
		fchan := make(chan FileResponse)
		go ReadFileAsync(path, fchan)
		select {
		case resp := <-fchan:
			if resp.status != nil {
				w.WriteHeader(http.StatusInternalServerError)
			} else {
				w.WriteHeader(http.StatusOK)
				w.Write(resp.fileio)
			}
		case <-time.After(1 * time.Second):
			w.WriteHeader(http.StatusRequestTimeout)

		}

	}

}

func main() {
	ExamineFiles()
	fmt.Println(files)
	http.Handle("/files", http.FileServer(http.Dir("./views")))
	http.HandleFunc("/", welcome)
	http.ListenAndServe(":5000", nil)
}
