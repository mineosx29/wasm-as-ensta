python3 dsl_auto.py -i fsm.nz -oDot code.gv -oGO fsm.go
dot -Tpng code.gv -o ../assets/code2.png
mv fsm.go ../cmd/wasm/
cd ../assets/
make fsm

# Renommer main.py 
