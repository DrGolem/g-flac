# go-flac

Go bindings for libFLAC


WIP, learning cgo and audio libs

Example:
```sh
go run examples/flac2raw.go 1.flac out_1.raw
```
play raw samples:
```sh
ffplay -ar 44100 -ac 2 -f s16le out_1.raw
```
