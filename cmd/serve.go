//go:build !js
// +build !js

package main

import (
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"os/exec"
	"path/filepath"
	"regexp"
	"strings"
	"time"

	esbuild "github.com/evanw/esbuild/pkg/api"
	"github.com/radovskyb/watcher"
)

func init() {
	goRoot, err := exec.Command("go", "env", "GOROOT").Output()
	if err != nil {
		log.Fatalln("failed getting goroot: ", err.Error())
	}
	wasmExecPath := fmt.Sprint(strings.TrimSpace(string(goRoot)), "/misc/wasm/wasm_exec.js")

	source, err := os.Open(wasmExecPath)
	if err != nil {
		log.Fatalln("failed to open wasm_exec.js source", err)
	}
	defer source.Close()

	destination, err := os.Create("./assets/wasm_exec.js")
	if err != nil {
		log.Fatalln("failed to open wasm_exec.js destination", err)
	}
	defer destination.Close()

	_, err = io.Copy(destination, source)

	if err != nil {
		log.Fatalln("failed copying file: ", err)
	} else {
		log.Println("copied wasm_exec.js from current GOROOT")
	}
}

func buildGo() {
	os.Setenv("GOOS", "js")
	os.Setenv("GOARCH", "wasm")
	cmd := exec.Command("go", "build", "-o", "./assets/main.wasm")
	err := cmd.Run()
	if err != nil {
		log.Println(err.Error())
	} else {
		log.Println("built .go sources to ./assets/main.wasm")
	}
}

func buildTs() {
	result := esbuild.Build(esbuild.BuildOptions{
		EntryPoints: []string{"./draw/renderer.ts"},
		Bundle:      true,
		Outfile:     "./assets/renderer.js",
		Write:       true,
	})

	if len(result.Errors) > 0 {
		log.Println(result.Errors)
	} else {
		log.Println("built .ts sources to ./assets/renderer.js")
	}
}

func watchSourceFiles() {
	w := watcher.New()
	w.SetMaxEvents(1)
	r := regexp.MustCompile(`(\.go|\.ts)$`)
	w.AddFilterHook(watcher.RegexFilterHook(r, false))

	go func() {
		for {
			select {
			case event := <-w.Event:
				if strings.HasSuffix(event.Name(), ".go") {
					buildGo()
				} else if strings.HasSuffix(event.Name(), ".ts") {
					buildTs()
				}
			case err := <-w.Error:
				log.Fatalln(err)
			case <-w.Closed:
				return
			}
		}
	}()

	if err := w.AddRecursive("."); err != nil {
		log.Fatalln(err)
	}

	if err := w.Start(time.Millisecond * 100); err != nil {
		log.Fatalln(err)
	}
}

func serve() {
	http.HandleFunc("/", func(w http.ResponseWriter, req *http.Request) {
		if req.RequestURI == "/" {
			req.RequestURI = "/index.html"
		}
		file, err := os.Open(filepath.Join("./assets", req.RequestURI))
		if err == nil {
			io.Copy(w, file)
		}
	})
	log.Println("serving on port 1334...")
	log.Println(http.ListenAndServe(":1334", nil))
}

func main() {
	buildGo()
	buildTs()

	go watchSourceFiles()

	serve()
}
