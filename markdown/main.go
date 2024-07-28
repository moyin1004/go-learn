package main

import (
	"bytes"
	"fmt"
	"os"
	"time"

	"github.com/yuin/goldmark"
	"github.com/yuin/goldmark/extension"
	"github.com/yuin/goldmark/parser"
	"github.com/yuin/goldmark/renderer/html"
)

func main() {
	data, err := os.ReadFile("./深度学习入门学习笔记.md")
	if err != nil {
		fmt.Println(err)
		return
	}
	start := time.Now().UnixMicro()
	md := goldmark.New(
		goldmark.WithExtensions(extension.GFM, extension.DefinitionList, extension.Footnote),
		goldmark.WithParserOptions(parser.WithAttribute(), parser.WithAutoHeadingID(), parser.WithBlockParsers()),
		goldmark.WithRendererOptions(html.WithUnsafe()),
	)
	buf := bytes.NewBuffer([]byte{})
	err = md.Convert(data, buf)
	if err != nil {
		fmt.Println(err)
		return
	}
	err = os.WriteFile("./text.html", buf.Bytes(), os.ModeAppend)
	fmt.Println(len(data), buf.Len(), err, time.Now().UnixMicro()-start)
}
