package main

import (
	"encoding/hex"
	"fmt"
	"io/ioutil"
	"os"
	"strings"
)

func main() {
	bs, err := ioutil.ReadAll(os.Stdin)
	if err != nil {
		fmt.Printf("read from stdin: %v\n", err)
		os.Exit(1)
	}
	s := strings.TrimSpace(string(bs))
	bs, err = hex.DecodeString(s)
	if err != nil {
		fmt.Printf("decode hex string: %v\n", err)
		os.Exit(1)
	}
	for _, b := range bs {
		fmt.Fprintf(os.Stdout, "%c", b)
	}
}
