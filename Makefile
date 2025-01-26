buildexamples:
	go run . -c -f ./examples/function.asm -o ./examples/function.bin
	go run . -c -f ./examples/hello_world.asm -o ./examples/hello_world.bin

runexamples:
	@echo "Running function example"
	go run . -r -f ./examples/function.bin
	@echo
	@echo "Running hello_world example"
	go run . -r -f ./examples/hello_world.bin
