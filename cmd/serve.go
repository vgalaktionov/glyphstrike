//go:build !js
// +build !js

package main

import (
	"io"
	"log"
	"net/http"
	"os"
	"os/exec"
	"path/filepath"
)

func init() {
	os.Setenv("GOOS", "js")
	os.Setenv("GOARCH", "wasm")
	cmd := exec.Command("go", "build", "-o", "./assets/main.wasm")
	err := cmd.Run()
	if err != nil {
		log.Fatal(err.Error())
	}
}

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, req *http.Request) {
		if req.RequestURI == "/" {
			req.RequestURI = "/index.html"
		}
		file, err := os.Open(filepath.Join("./assets", req.RequestURI))
		if err == nil {
			io.Copy(w, file)
		}
	})
	log.Println("Serving on port 1334...")
	log.Println(http.ListenAndServe(":1334", nil))
}
