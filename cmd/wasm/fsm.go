package main

import (

	"fmt"

	"syscall/js"

)

type State uint32

var commande string

var document js.Value

func getElementById(elem string) js.Value {

	document = js.Global().Get("document")
	return document.Call("getElementById", elem)
}
func getElementValue(elem string, value string) js.Value {

	return getElementById(elem).Get(value)
}

func GetString(elem string, value string) string {

	return getElementValue(elem, value).String()
}

func main() {
	quit := make(chan struct{}, 0)
	const (
	eteint State = iota

	allume
)
	state := eteint
	button := js.Global().Get("document").Call("getElementById", "ok")
	image := js.Global().Get("document").Call("getElementById", "image")
	image.Set("src", "code2.png")
	var cb js.Func
	cb = js.FuncOf(func(this js.Value, args []js.Value) interface{} {
	commande = GetString("in", "value")

		if state == eteint {

			if commande == "active" {
				fmt.Println(" Passage à Etat : allume")
				state = allume
				}
		}
		if state == eteint {

			if commande == "yo" {
				fmt.Println(" Passage à Etat : eteint")
				state = eteint
				}
		}
		if state == eteint {

			if commande == "hein" {
				fmt.Println(" Passage à Etat : allume")
				state = allume
				}
		}
		if state == allume {

			if commande == "yo" {
				fmt.Println(" Passage à Etat : eteint")
				state = eteint
				}
		}
		if state == allume {

			if commande == "active" {
				fmt.Println(" Passage à Etat : allume")
				state = allume
				}
		}
		if state == allume {

			if commande == "hein" {
				fmt.Println(" Passage à Etat : eteint")
				state = eteint
				}
		}
			return nil

		})
	button.Call("addEventListener", "click", cb)
	<-quit
}