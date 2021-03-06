// discover provides node discovery on the command line.
package main

import (
	"flag"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"os"
	"strings"

	discover "github.com/hashicorp/go-discover"
)

func main() {
	var quiet bool
	var help bool
	var json bool
	flag.BoolVar(&quiet, "q", false, "no verbose output")
	flag.BoolVar(&help, "h", false, "print help")
	flag.BoolVar(&json, "json", false, "output response as json")
	flag.Parse()

	d := &discover.Discover{}

	args := flag.Args()
	if help || len(args) == 0 || args[0] != "addrs" {
		fmt.Println("Usage: discover addrs key=val key=val ...")
		fmt.Println(d.Help())
		os.Exit(0)
	}
	args = args[1:]

	var w io.Writer = os.Stderr
	if quiet {
		w = ioutil.Discard
	}
	l := log.New(w, "", 0)

	l.Printf("Registered providers: %v", d.Names())

	addrs, err := d.Addrs(strings.Join(args, " "), l)
	if err != nil {
		l.Fatal(err)
	}
	switch addrs.(type) {
	case string:
		fmt.Println(addrs)
	case []string:
		fmt.Println(strings.Join(addrs.([]string), " "))
	}
}
