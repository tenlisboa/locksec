package main

import (
	"bufio"
	"flag"
	"fmt"
	"log"
	"os"

	"github.com/tenlisboa/locksec/internal/lib/extractor"
)

var lockpath = flag.String("f", "package-lock.json", "This is the path of the lockfile from your project")
var ftype = flag.Int("t", 0, "0 to json file and 1 for yml|yaml file")

func main() {
	flag.Parse()

	f, err := os.OpenFile(*lockpath, os.O_RDONLY, 0444)

	if err != nil {
		log.Fatalf("Error on opening lockFile: %v\n", err)
		os.Exit(2)
	}

	scanner := bufio.NewScanner(f)
	extr := extractor.New(extractor.ToFType(*ftype))

	for scanner.Scan() {
		key, value, _ := extr.ExtractLine(scanner.Text())
		if key == "resolved" {
			// TODO: Verify hosts
			fmt.Println(key, value)
		}
	}

	if err := scanner.Err(); err != nil {
		log.Fatalf("Error scanning file: %v\n", err)
		os.Exit(2)
	}
}
