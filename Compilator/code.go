package main

import (

	"fmt"

	"syscall/js"

)

type State uint32

const (
	eteint State = iota
	allume
)
func main() {
eteint -> allume [ label = "active"];
eteint -> eteint [ label = "desactive"];
allume -> allume [ label = "active"];
allume -> eteint [ label = "desactive"];
}