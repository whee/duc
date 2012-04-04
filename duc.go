// Copyright 2012, Brian Hetro <whee@smaertness.net>
// Use of this source code is governed by the ISC license
// that can be found in the LICENSE file.

package main

import (
	"bytes"
	"fmt"
	"github.com/whee/ddg"
	"os"
	"text/template"
)

const zci = `
* {{.AbstractText}} Source: {{.AbstractSource}}, {{.AbstractURL}}

{{.Definition}}
`

var t = template.Must(template.New("zci").Parse(zci))

func lookup(term string, ch chan []byte) {
	b := new(bytes.Buffer)
	r, err := ddg.ZeroClick(term)
	if err == nil {
		err = t.Execute(b, r)
	}
	if err != nil {
		fmt.Fprintf(b, "Error looking up %s: %v\n", term, err)
	}
	ch <- b.Bytes()
}

func main() {
	ch := make([]chan []byte, len(os.Args)-1)
	for i := range ch {
		ch[i] = make(chan []byte)
	}
	for i, term := range os.Args[1:] {
		go lookup(term, ch[i])
	}
	for i := range ch {
		fmt.Print(string(<-ch[i]))
	}
}
