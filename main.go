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
		fmt.Fprintf(os.Stderr, "read from stdin: %v\n", err)
		os.Exit(1)
	}
	s := strings.TrimSpace(string(bs))
	bs, err = hex.DecodeString(s)
	if err != nil {
		fmt.Fprintf(os.Stderr, "decode hex string: %v\n", err)
		os.Exit(1)
	}
	os.Stdout.Write(bs)
}
