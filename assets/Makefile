ifeq ($(OS),Windows_NT)
    BROWSER = start
else
	UNAME := $(shell uname -s)
	ifeq ($(UNAME), Linux)
		BROWSER = xdg-open
	endif
	ifeq ($(UNAME), Darwin)
		BROWSER = open
	endif
endif

.PHONY: all clean server

%.wasm: %.go 
	cp -f "$$(go env GOROOT)/misc/wasm/wasm_exec.js" ../assets/
	GOOS=js GOARCH=wasm go build -o ../assets/json.wasm "$<"
	
lampe: ../cmd/wasm/lampe.wasm 
robot: ../cmd/wasm/robot.wasm
fsm  : ../cmd/wasm/fsm.wasm


	

clean:
	rm -f *.wasm
