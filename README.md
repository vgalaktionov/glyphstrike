# Glyphstrike

> A Roguelike implementation in Go.

![CI/CD](https://github.com/vgalaktionov/glyphstrike/actions/workflows/ci.yaml/badge.svg)

## Playing

WASM version of the latest succesful build is running [here](https://vgalaktionov.github.io/glyphstrike).

Mac binaries: 
[Apple Silicon](https://vgalaktionov.github.io/glyphstrike/bin/glyphstrike-arm64-darwin)  [x64](https://vgalaktionov.github.io/glyphstrike/bin/glyphstrike-amd64-darwin)

Linux binaries: 
[x64](https://vgalaktionov.github.io/glyphstrike/bin/glyphstrike-amd64-linux)

Windows binaries: 
[x64](https://vgalaktionov.github.io/glyphstrike/bin/glyphstrike-amd64.exe)

## Development

### Requirements:

Go 1.19 or newer, NodeJS 18 or newer.

### Useful commands:

Running in development with the default console renderer:

```bash
go run main.go
```

Running in development with the WASM canvas renderer, serving on localhost:1334 :

```bash
go run cmd/serve.go
```

Running all tests

```bash
go test ./... -v
```

Running benchmarks

```bash
# Run all benchmarks
go test ./... -bench=. -benchmem -timeout 2m

# Run whole game benchmarks with profiling
go test ./game -bench=. -benchmem -memprofile memprofile.out -cpuprofile profile.out

# Inspect profiling data
go tool pprof profile.out
go tool pprof memprofile.out
```

## Building

Building for Linux

```bash
GOOS=linux GOARCH=amd64 go build -o bin/glyphstrike-amd64-linux main.go
```

Building for Mac

```bash
# Intel
GOOS=darwin GOARCH=amd64 go build -o bin/glyphstrike-amd64-darwin main.go
# Apple Silicon
GOOS=darwin GOARCH=arm64 go build -o bin/glyphstrike-arm64-darwin main.go
```

Building for Windows

```bash
GOOS=windows GOARCH=amd64 go build -o bin/glyphstrike-amd64.exe main.go
```

Building for Web

```bash
GOOS=js GOARCH=wasm go build -o assets/main.wasm
esbuild draw/renderer.ts --bundle --outfile=assets/renderer.js
```
