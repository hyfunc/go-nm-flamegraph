package main

import (
	"bufio"
	"flag"
	"fmt"
	"io"
	"log"
	"os"
	"strconv"
	"strings"
)

type SymbolType string // https://pkg.go.dev/cmd/nm

const (
	Text             SymbolType = "T"
	StaticText       SymbolType = "t"
	ReadOnly         SymbolType = "R"
	StaticReadOnly   SymbolType = "r"
	Data             SymbolType = "D"
	StaticData       SymbolType = "d"
	BSSSegment       SymbolType = "B"
	StaticBSSSegment SymbolType = "b"
	Constant         SymbolType = "C"
	Undefined        SymbolType = "U"
)

var symbolTypes = map[SymbolType]struct{}{
	Text:             {},
	StaticText:       {},
	ReadOnly:         {},
	StaticReadOnly:   {},
	Data:             {},
	StaticData:       {},
	BSSSegment:       {},
	StaticBSSSegment: {},
	Constant:         {},
	Undefined:        {},
}

func main() {
	outputFile := flag.String("o", "", "Output file (default: stdout)")
	flag.Parse()

	var out io.Writer = os.Stdout
	if *outputFile != "" {
		file, err := os.Create(*outputFile)
		if err != nil {
			log.Fatalf("Error creating output file: %v\n", err)
		}
		defer file.Close()
		out = file
	}

	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		line := scanner.Text()
		fields := strings.Fields(line)

		if len(fields) < 2 {
			log.Println("Not enough fields:", line)
			continue
		}

		idxType := -1
		for i := 0; i < len(fields); i++ {
			if _, ok := symbolTypes[SymbolType(fields[i])]; ok {
				idxType = i
				break
			}
		}
		if idxType == -1 {
			log.Println("invalid type:", line)
			continue
		}
		if !(idxType == 1 || idxType == 2) {
			log.Println("expected type position:", line)
			continue
		}

		size, err := strconv.Atoi(fields[idxType-1])
		if err != nil {
			log.Println("invalid size:", line)
			continue
		}

		symbol := strings.Join(fields[idxType+1:], " ")

		lastDot := strings.LastIndex(symbol, ".")
		if lastDot == -1 {
			continue
		}
		pkgPath := symbol[:lastDot]

		stack := strings.ReplaceAll(pkgPath, "/", ";")
		fmt.Fprintf(out, "%s %d\n", stack, size)
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
}
