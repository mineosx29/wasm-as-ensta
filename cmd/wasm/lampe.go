package main

import(
	"syscall/js"
	//"time"
	"fmt"
	"math"
)

type State uint32

const (
	eteint State = iota

	allume
)


const (
	width = 400
	height = 400
)



var document js.Value
var test_command string


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

func wrapGoFunction(fn func()) func(js.Value, []js.Value) interface {} {
    return func(_ js.Value, _ []js.Value) interface {} {
        fn()
        return nil
    }
}

func SetValue(elem string, key string, value string) {
    getElementById(elem).Set(key, value)
}


func main() {
	quit := make(chan struct{}, 0)


	state := eteint
	test := js.Global().Get("document").Call("getElementById", "image")
	test.Set("src", "code2.png")
	lampe := js.Global().Get("document").Call("getElementById", "image_simu")
    lampe.Set("src", "eteint.png")
	button := js.Global().Get("document").Call("getElementById", "ok")
	var canvas2 js.Value = js.
		Global().
		Get("document").
		Call("getElementById", "canvas")
	var canvas_state_out js.Value = js.Global().Get("document").Call("getElementById", "canvas3")
	var context_out = canvas_state_out.Call("getContext", "2d")


	var context  = canvas2.Call("getContext", "2d")
	canvas2.Set("height", height)
	canvas2.Set("width", width)
	context.Call("clearRect", 0, 0, width, height)

	
	var cb js.Func
	cb = js.FuncOf(func(this js.Value, args []js.Value) interface{} {
    	test_command = GetString("in", "value")
    	switch state {
    	case eteint:
    		fmt.Println("On est dans l'état Eteint de l'automate")
    		lampe.Set("src", "eteint.png")
    		//context.Call("clearRect", 10, 130, 150, 100)
    		context.Call("beginPath")
    		context.Set("fillStyle", "grey")
			context.Call("fillRect", 10, 130, 150,100)
			context.Set("font", "48px serif")
			context.Set("fillStyle", "white")
			context.Call("fillText", "Allume", 15, 185)
			context.Set("fillStyle", "green")
			context.Call("fillRect", 10, 20, 150,100)
			context.Set("font", "48px serif")
			context.Set("fillStyle", "white")
			context.Call("fillText", "Eteint", 17, 75)
			context.Call("stroke")
			context.Call("fill")
			context_out.Call("clearRect", 0, 0, width, height)
			context_out.Call("beginPath")
			context_out.Call("arc", 35, 130, 13, 0, 2*math.Pi, false)
			context_out.Set("fillStyle", "grey")
			context_out.Call("fill")
			context_out.Set("lineWidth", 5)
			context_out.Set("font", "30px serif")
			context_out.Set("fillStyle", "black")
			context_out.Call("fillText", "S", 25, 112)
			
    		if test_command == "active" {
				//lampe.Set("src", "allume.png")
				state = allume
			}
		case allume:
			fmt.Println("On est dans l'état Allume de l'automate")
			lampe.Set("src", "allume.png")
			//context.Call("clearRect", 10, 20, 150, 100)
			context.Call("beginPath")
			context.Set("fillStyle", "grey")
			context.Call("fillRect", 10, 20, 150,100)
			context.Set("font", "48px serif")
			context.Set("fillStyle", "white")
			context.Call("fillText", "Eteint", 17, 75)
			context.Set("fillStyle", "green")
			context.Call("fillRect", 10, 130, 150,100)
			context.Set("font", "48px serif")
			context.Set("fillStyle", "white")
			context.Call("fillText", "Allume", 15, 185)
			context.Call("stroke")
			context.Call("fill")
			context_out.Call("clearRect", 0, 0, width, height)
			context_out.Call("beginPath")
			context_out.Call("arc", 35, 130, 13, 0, 2*math.Pi, false)
			context_out.Set("fillStyle", "green")
			context_out.Call("fill")
			context_out.Set("lineWidth", 5)
			context_out.Set("font", "30px serif")
			context_out.Set("fillStyle", "black")
			context_out.Call("fillText", "S", 25, 112)
			if test_command == "desactive" {
				//lampe.Set("src", "allume.png")
				state = eteint
			}
    	}
    	

		

    	
    		return nil
	})

		
	

	button.Call("addEventListener", "click", cb)	





	<-quit



}