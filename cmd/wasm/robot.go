package main

import(
	"fmt"
	"syscall/js"
	"math"
)

type State uint32

const (
	off State = iota

	monte

	descend

	droite

	gauche
)

const (
	width = 400
	height = 400
	
)



var document js.Value
var test_command string
var x int
var y int

/*func clearCircle(ctx string, x int, y int , radius int) js.Value {
	ctx.Call("save")
	ctx.Call("beginPath")

}*/


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
	quit:= make(chan struct{}, 0)

	state:= off
	button := js.Global().Get("document").Call("getElementById", "ok")
	test := js.Global().Get("document").Call("getElementById", "image")
	test.Set("src", "code.png")

	var canvas js.Value = js.Global().Get("document").Call("getElementById", "canvas")
	var canvas_simulation js.Value = js.Global().Get("document").Call("getElementById", "canvas2")
	var canvas_state_out js.Value = js.Global().Get("document").Call("getElementById", "canvas3")
	var context_simulation  = canvas_simulation.Call("getContext", "2d")
	var context_out = canvas_state_out.Call("getContext", "2d")
	//var context = canvas.Call("getContext", "2d")

	canvas_simulation.Set("height", height)
	canvas_simulation.Set("width", width)

	canvas.Set("height", height)
	canvas.Set("width", width)
	context_simulation.Call("clearRect", 0, 0, width, height)
	context_simulation.Call("beginPath")
	context_simulation.Set("fillStyle", "green")
	context_simulation.Call("fillRect", 10, 10, 50, 50)
	context_out.Call("beginPath")
	context_out.Call("arc", 35, 130, 13, 0, 2*math.Pi, false)
	context_out.Set("fillStyle", "grey")
	context_out.Call("fill")
	context_out.Set("lineWidth", 5)
	context_out.Set("font", "30px serif")
	context_out.Set("fillStyle", "black")
	context_out.Call("fillText", "U", 23, 112)
	context_out.Call("beginPath")
	context_out.Call("arc", 75, 130, 13, 0, 2*math.Pi, false)
	context_out.Set("fillStyle", "grey")
	context_out.Call("fill")
	context_out.Set("lineWidth", 5)
	context_out.Set("font", "30px serif")
	context_out.Set("fillStyle", "black")
	context_out.Call("fillText", "D", 63, 112)
	context_out.Call("beginPath")
	context_out.Call("arc", 115, 130, 13, 0, 2*math.Pi, false)
	context_out.Set("fillStyle", "grey")
	context_out.Call("fill")
	context_out.Set("lineWidth", 5)
	context_out.Set("font", "30px serif")
	context_out.Set("fillStyle", "black")
	context_out.Call("fillText", "L", 105, 112)
	context_out.Call("beginPath")
	context_out.Call("arc", 155, 130, 13, 0, 2*math.Pi, false)
	context_out.Set("fillStyle", "grey")
	context_out.Call("fill")
	context_out.Set("lineWidth", 5)
	context_out.Set("font", "30px serif")
	context_out.Set("fillStyle", "black")
	context_out.Call("fillText", "R", 145, 112)


	//context_out.Set("strokeStyle", "#003300")
	//context_out.Call("stroke")
	//context_out.Call("clearRect", 0, 0, width, height)

	var cb js.Func 
	cb = js.FuncOf(func(this js.Value, args []js.Value) interface{} {
		test_command = GetString("in", "value")

		switch state {
		case off:
			fmt.Println("On est dans l'état Off de l'automate")
			context_out.Call("clearRect", 0, 0, width, height)
			context_out.Call("beginPath")
			context_out.Call("arc", 35, 130, 13, 0, 2*math.Pi, false)
			context_out.Set("fillStyle", "grey")
			context_out.Call("fill")
			context_out.Set("lineWidth", 5)
			context_out.Set("font", "30px serif")
			context_out.Set("fillStyle", "black")
			context_out.Call("fillText", "U", 23, 112)
			context_out.Call("beginPath")
			context_out.Call("arc", 75, 130, 13, 0, 2*math.Pi, false)
			context_out.Set("fillStyle", "grey")
			context_out.Call("fill")
			context_out.Set("lineWidth", 5)
			context_out.Set("font", "30px serif")
			context_out.Set("fillStyle", "black")
			context_out.Call("fillText", "D", 63, 112)
			context_out.Call("beginPath")
			context_out.Call("arc", 115, 130, 13, 0, 2*math.Pi, false)
			context_out.Set("fillStyle", "grey")
			context_out.Call("fill")
			context_out.Set("lineWidth", 5)
			context_out.Set("font", "30px serif")
			context_out.Set("fillStyle", "black")
			context_out.Call("fillText", "L", 105, 112)
			context_out.Call("beginPath")
			context_out.Call("arc", 155, 130, 13, 0, 2*math.Pi, false)
			context_out.Set("fillStyle", "grey")
			context_out.Call("fill")
			context_out.Set("lineWidth", 5)
			context_out.Set("font", "30px serif")
			context_out.Set("fillStyle", "black")
			context_out.Call("fillText", "R", 145, 112)
		



			if test_command == "descends"  {
				state = descend
				
			}

			if test_command == "monte" {
				state = monte
				
			}
			
			if test_command == "droite" {
				state = droite
				
			}

			if test_command == "gauche" {
				state = gauche
				
			}

			
		
		case descend:
			fmt.Println("On est dans l'état descend de l'automate")
			for condition := 0; condition < 10; condition++ {
				context_simulation.Call("clearRect", 0, 0, width, height)
				context_out.Call("clearRect", 0, 0, width, height)
				context_simulation.Call("beginPath")
				context_simulation.Set("fillStyle", "green")
				context_simulation.Call("fillRect", 10, (10+x)+condition, 50, 50)
		/*		context_out.Call("beginPath")
				context_out.Call("arc", 75, 130, 13, 0, 2*math.Pi, false)
				context_out.Set("fillStyle", "green")
				context_out.Call("fill")
				context_out.Set("lineWidth", 5)
				context_out.Set("font", "30px serif")
				context_out.Set("fillStyle", "black")
				context_out.Call("fillText", "D", 63, 112)*/
				context_out.Call("beginPath")
				context_out.Call("arc", 35, 130, 13, 0, 2*math.Pi, false)
				context_out.Set("fillStyle", "grey")
				context_out.Call("fill")
				context_out.Set("lineWidth", 5)
				context_out.Set("font", "30px serif")
				context_out.Set("fillStyle", "black")
				context_out.Call("fillText", "U", 23, 112)
				context_out.Call("beginPath")
				context_out.Call("arc", 75, 130, 13, 0, 2*math.Pi, false)
				context_out.Set("fillStyle", "green")
				context_out.Call("fill")
				context_out.Set("lineWidth", 5)
				context_out.Set("font", "30px serif")
				context_out.Set("fillStyle", "black")
				context_out.Call("fillText", "D", 63, 112)
				context_out.Call("beginPath")
				context_out.Call("arc", 115, 130, 13, 0, 2*math.Pi, false)
				context_out.Set("fillStyle", "grey")
				context_out.Call("fill")
				context_out.Set("lineWidth", 5)
				context_out.Set("font", "30px serif")
				context_out.Set("fillStyle", "black")
				context_out.Call("fillText", "L", 105, 112)
				context_out.Call("beginPath")
				context_out.Call("arc", 155, 130, 13, 0, 2*math.Pi, false)
				context_out.Set("fillStyle", "grey")
				context_out.Call("fill")
				context_out.Set("lineWidth", 5)
				context_out.Set("font", "30px serif")
				context_out.Set("fillStyle", "black")
				context_out.Call("fillText", "R", 145, 112)

				x = x + 1

			}
			state = off
			
		case monte:
			fmt.Println("On est dans l'état monte de l'automate")
			for condition := 0; condition < 10; condition++ {
				context_simulation.Call("clearRect", 0, 0, width, height)
				context_out.Call("clearRect", 0, 0, width, height)
				context_simulation.Call("beginPath")
				context_simulation.Set("fillStyle", "green")
				context_simulation.Call("fillRect", 10, (10+x)-condition, 50, 50)
				context_out.Call("beginPath")
				context_out.Call("arc", 35, 130, 13, 0, 2*math.Pi, false)
				context_out.Set("fillStyle", "green")
				context_out.Call("fill")
				context_out.Set("lineWidth", 5)
				context_out.Set("font", "30px serif")
				context_out.Set("fillStyle", "black")
				context_out.Call("fillText", "U", 23, 112)
				context_out.Call("beginPath")
				context_out.Call("arc", 75, 130, 13, 0, 2*math.Pi, false)
				context_out.Set("fillStyle", "grey")
				context_out.Call("fill")
				context_out.Set("lineWidth", 5)
				context_out.Set("font", "30px serif")
				context_out.Set("fillStyle", "black")
				context_out.Call("fillText", "D", 63, 112)
				context_out.Call("beginPath")
				context_out.Call("arc", 115, 130, 13, 0, 2*math.Pi, false)
				context_out.Set("fillStyle", "grey")
				context_out.Call("fill")
				context_out.Set("lineWidth", 5)
				context_out.Set("font", "30px serif")
				context_out.Set("fillStyle", "black")
				context_out.Call("fillText", "L", 105, 112)
				context_out.Call("beginPath")
				context_out.Call("arc", 155, 130, 13, 0, 2*math.Pi, false)
				context_out.Set("fillStyle", "grey")
				context_out.Call("fill")
				context_out.Set("lineWidth", 5)
				context_out.Set("font", "30px serif")
				context_out.Set("fillStyle", "black")
				context_out.Call("fillText", "R", 145, 112)

				x = x - 1

			}
			
			state = off
		case droite:
			fmt.Println("On est dans l'état droite de l'automate")
			for condition := 0; condition < 10; condition++ {
				context_simulation.Call("clearRect", 0, 0, width, height)
		
				context_simulation.Call("beginPath")
				context_simulation.Set("fillStyle", "green")
				context_simulation.Call("fillRect", (10+y)+condition, (10+x), 50, 50)
				context_out.Call("clearRect", 0, 0, width, height)
				context_out.Call("beginPath")
				context_out.Call("arc", 35, 130, 13, 0, 2*math.Pi, false)
				context_out.Set("fillStyle", "grey")
				context_out.Call("fill")
				context_out.Set("lineWidth", 5)
				context_out.Set("font", "30px serif")
				context_out.Set("fillStyle", "black")
				context_out.Call("fillText", "U", 23, 112)
				context_out.Call("beginPath")
				context_out.Call("arc", 75, 130, 13, 0, 2*math.Pi, false)
				context_out.Set("fillStyle", "grey")
				context_out.Call("fill")
				context_out.Set("lineWidth", 5)
				context_out.Set("font", "30px serif")
				context_out.Set("fillStyle", "black")
				context_out.Call("fillText", "D", 63, 112)
				context_out.Call("beginPath")
				context_out.Call("arc", 115, 130, 13, 0, 2*math.Pi, false)
				context_out.Set("fillStyle", "grey")
				context_out.Call("fill")
				context_out.Set("lineWidth", 5)
				context_out.Set("font", "30px serif")
				context_out.Set("fillStyle", "black")
				context_out.Call("fillText", "L", 105, 112)
				context_out.Call("beginPath")
				context_out.Call("arc", 155, 130, 13, 0, 2*math.Pi, false)
				context_out.Set("fillStyle", "green")
				context_out.Call("fill")
				context_out.Set("lineWidth", 5)
				context_out.Set("font", "30px serif")
				context_out.Set("fillStyle", "black")
				context_out.Call("fillText", "R", 145, 112)

				y = y + 1

			}
			state = off

		case gauche:
			fmt.Println("On est dans l'état droite de l'automate")
			for condition := 0; condition < 10; condition++ {
				context_simulation.Call("clearRect", 0, 0, width, height)
		
				context_simulation.Call("beginPath")
				context_simulation.Set("fillStyle", "green")
				context_simulation.Call("fillRect", (10+y)-condition, (10+x), 50, 50)
				context_out.Call("clearRect", 0, 0, width, height)
				context_out.Call("beginPath")
				context_out.Call("arc", 35, 130, 13, 0, 2*math.Pi, false)
				context_out.Set("fillStyle", "grey")
				context_out.Call("fill")
				context_out.Set("lineWidth", 5)
				context_out.Set("font", "30px serif")
				context_out.Set("fillStyle", "black")
				context_out.Call("fillText", "U", 23, 112)
				context_out.Call("beginPath")
				context_out.Call("arc", 75, 130, 13, 0, 2*math.Pi, false)
				context_out.Set("fillStyle", "grey")
				context_out.Call("fill")
				context_out.Set("lineWidth", 5)
				context_out.Set("font", "30px serif")
				context_out.Set("fillStyle", "black")
				context_out.Call("fillText", "D", 63, 112)
				context_out.Call("beginPath")
				context_out.Call("arc", 115, 130, 13, 0, 2*math.Pi, false)
				context_out.Set("fillStyle", "green")
				context_out.Call("fill")
				context_out.Set("lineWidth", 5)
				context_out.Set("font", "30px serif")
				context_out.Set("fillStyle", "black")
				context_out.Call("fillText", "L", 105, 112)
				context_out.Call("beginPath")
				context_out.Call("arc", 155, 130, 13, 0, 2*math.Pi, false)
				context_out.Set("fillStyle", "grey")
				context_out.Call("fill")
				context_out.Set("lineWidth", 5)
				context_out.Set("font", "30px serif")
				context_out.Set("fillStyle", "black")
				context_out.Call("fillText", "R", 145, 112)

				y = y - 1

			}
			state = off

			
		}

		  		return nil
	})


	button.Call("addEventListener", "click", cb)	

	<-quit

}
