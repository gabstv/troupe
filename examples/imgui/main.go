package main

import (
	"io/ioutil"

	"github.com/gabstv/primen"
	"github.com/gabstv/primen/core"
	"github.com/gabstv/primen/core/ui/imgui"
	"github.com/gabstv/primen/dom"
)

func main() {
	fb, _ := ioutil.ReadFile("ui.xml")
	engine := primen.NewEngine(&primen.NewEngineInput{
		Width:     800,
		Height:    600,
		Resizable: true,
		OnReady: func(e primen.Engine) {
			node, err := dom.ParseXMLText(string(fb))
			if err != nil {
				panic(err)
			}
			imgui.AddUI(node.(dom.ElementNode))
			e.AddEventListener("test", func(eventName string, e core.Event) {
				println(e.Data.(string))
			})
		},
	})
	engine.SetDebugTPS(true)
	imgui.Setup(engine)
	engine.Run()
}