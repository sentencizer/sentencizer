//go:build js && wasm

package main

import (
	"strings"
	"syscall/js"

	"github.com/sentencizer/sentencizer"
)

func segment(this js.Value, inputs []js.Value) interface{} {
	text := inputs[0].String()
	lang := inputs[1].String()
	segmenter := sentencizer.NewSegmenter(lang)
	return strings.Join(segmenter.Segment(text), "\r")
}

func main() {
	c := make(chan struct{}, 0)
	js.Global().Set("segment", js.FuncOf(segment))
	<-c
}
