package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/maxence-charriere/go-app/v9/pkg/app"
)

type hello struct {
	app.Compo
}

func (h *hello) Render() app.UI {
	return app.Div().Body(
		app.H1().Text("点击按钮"),
		app.Button().Text("点击我").OnClick(h.onButtonClicked),
	)
}

func (h *hello) onButtonClicked(ctx app.Context, e app.Event) {
	fmt.Println("button is clicked!")
	log.Println(ctx.JSSrc().Get("value"))
}

func main() {
	app.Route("/", &hello{})
	app.RunWhenOnBrowser()

	http.Handle("/", &app.Handler{
		Name:        "Hello",
		Description: "An Hello World! example",
	})

	if err := http.ListenAndServe(":8000", nil); err != nil {
		log.Fatal(err)
	}
}
