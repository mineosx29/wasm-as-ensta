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
	ouvert State = iota

	fermee
	milieu
)
	state := ouvert

	button := js.Global().Get("document").Call("getElementById", "ok")
	image := js.Global().Get("document").Call("getElementById", "image")
	image.Set("src", "code2.png")
	var cb js.Func
	cb = js.FuncOf(func(this js.Value, args []js.Value) interface{} {
	commande = GetString("in", "value")

		if state == ouvert {

			if commande == "btn1" {
				fmt.Println(" Passage à Etat : ouvert")
				state = ouvert
				}
		}
		if state == ouvert {

			if commande == "btn2" {
				fmt.Println(" Passage à Etat : fermee")
				state = fermee
				}
		}
		if state == ouvert {

			if commande == "btn3" {
				fmt.Println(" Passage à Etat : milieu")
				state = milieu
				}
		}
		if state == fermee {

			if commande == "btn1" {
				fmt.Println(" Passage à Etat : ouvert")
				state = ouvert
				}
		}
		if state == fermee {

			if commande == "btn2" {
				fmt.Println(" Passage à Etat : fermee")
				state = fermee
				}
		}
		if state == fermee {

			if commande == "btn3" {
				fmt.Println(" Passage à Etat : milieu")
				state = milieu
				}
		}
		if state == milieu {

			if commande == "btn1" {
				fmt.Println(" Passage à Etat : ouvert")
				state = ouvert
				}
		}
		if state == milieu {

			if commande == "btn2" {
				fmt.Println(" Passage à Etat : fermee")
				state = fermee
				}
		}
		if state == milieu {

			if commande == "btn3" {
				fmt.Println(" Passage à Etat : milieu")
				state = milieu
				}
		}
			return nil

		})
	button.Call("addEventListener", "click", cb)
	<-quit
}