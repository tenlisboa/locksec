package main

import (
	"bufio"
	"flag"
	"fmt"
	"log"
	"net/url"
	"os"

	"github.com/tenlisboa/locksec/internal/lib/extractor"
	"golang.org/x/exp/slices"
)

var lockpath = flag.String("f", "package-lock.json", "This is the path of the lockfile from your project")
var ftype = flag.Int("t", 0, "0 to json file and 1 for yml|yaml file")
var key = flag.String("k", "resolved", "The key where the registry host is located in the file")

var knownHosts = []string{
	"registry.yarnpkg.com",
	"registry.npmjs.org",
	"npm.pkg.github.com",
}

func main() {
	flag.Parse()

	f, err := os.OpenFile(*lockpath, os.O_RDONLY, 0444)

	if err != nil {
		log.Fatalf("Error on opening lockFile: %v\n", err)
		os.Exit(2)
	}

	scanner := bufio.NewScanner(f)
	extr := extractor.New(extractor.ToFType(*ftype))

	fmt.Printf("------------------------ \nStart scanning the file: %s ...\n------------------------ \n\n", *lockpath)

	line := 0
	for scanner.Scan() {
		line++
		k, value, _ := extr.ExtractLine(scanner.Text())
		if k != *key {
			continue
		}
		regurl, _ := url.Parse(value)
		if !slices.Contains(knownHosts, regurl.Host) {
			fmt.Printf("Suspicious host found at line: %d -> %s\n\n", line, value)
		}
	}

	if err := scanner.Err(); err != nil {
		log.Fatalf("Error scanning file: %v\n", err)
		os.Exit(2)
	}

	fmt.Println("Finished search successfully.")
}
