package main

import (

	"fmt"

	"syscall/js"

)

type State uint32

var commande string

var state State

func main() {
	state = eteint

	const (
	eteint State = iota
	allume
)
	switch state {
		case eteint:
			if commande == "active" {
				state = allume
}
		case allume:
			if commande == "desactive" {
				state = eteint
}
}
}