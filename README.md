# go-nm-flamegraph
go-nm-flamegraph converts output of go tool nm into the [flamegraph.pl](https://github.com/brendangregg/FlameGraph) format.

## Installation
```
go install github.com/hyfunc/go-nm-flamegraph@latest
```

## Usage
```
go tool nm -size <go_bin_file> | go-nm-flamegraph > fg.out
flamegraph.pl fg.out > fg.svg
```

## Example
```
go build -o main
go tool nm -size ./main | go-nm-flamegraph | flamegraph.pl > demo.svg
```
