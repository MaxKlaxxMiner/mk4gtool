package main

import (
	"github.com/lxn/walk"
	. "github.com/lxn/walk/declarative"
	"strings"
)

func main() {
	var inTE, outTE *walk.TextEdit

	_, _ = MainWindow{
		Title:   "SCREAMO",
		MinSize: Size{640, 480},
		Size:    Size{640, 480},
		Layout:  VBox{},
		Children: []Widget{
			HSplitter{
				Children: []Widget{
					TextEdit{AssignTo: &inTE},
					TextEdit{AssignTo: &outTE, ReadOnly: true},
				},
			},
			PushButton{
				Text: "SCREAM",
				OnClicked: func() {
					_ = outTE.SetText(strings.ToUpper(inTE.Text()))
				},
			},
		},
	}.Run()
}
