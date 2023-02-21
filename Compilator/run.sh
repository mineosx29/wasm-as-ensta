python3 main.py -i fsm2.nz -oDot code.gv -oGO fsm.go
dot -Tpng code.gv -o ../assets/code2.png
mv fsm.go ../cmd/wasm/
cd ../assets/
make fsm


