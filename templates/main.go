package main

import (
	"log"
	"os"
	"text/template"
)

var tpl *template.Template

type MiumiuSpeak struct {
	Text string
}

func init() {
	tpl = template.Must(template.ParseFiles("index.gohtml"))
}

func main() {
	data := MiumiuSpeak{Text: "wenwenwen"}

	err := tpl.Execute(os.Stdout, data)
	if err != nil {
		log.Fatalln(err)
	}
}
