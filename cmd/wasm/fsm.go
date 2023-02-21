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
	off State = iota

	monte
	descend
	gauche
	droite
)
	state := off
	button := js.Global().Get("document").Call("getElementById", "ok")
	var cb js.Func
	cb = js.FuncOf(func(this js.Value, args []js.Value) interface{} {
	commande = GetString("in", "value")

		if state == off {

			if commande == "haut" {
				fmt.Println(" Passage à Etat : monte")
				state = monte
				}
		}
		if state == off {

			if commande == "bas" {
				fmt.Println(" Passage à Etat : descend")
				state = descend
				}
		}
		if state == monte {

			if commande == "stop" {
				fmt.Println(" Passage à Etat : off")
				state = off
				}
		}
		if state == descend {

			if commande == "stop" {
				fmt.Println(" Passage à Etat : off")
				state = off
				}
		}
			return nil

		})
	button.Call("addEventListener", "click", cb)
	<-quit
}