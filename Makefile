build:
	@export ENV=prod && mkdir -p build && export GIN_MODE=release && go build -o ./build/michi ./michi.go

clean:
	@rm -rf ./build ./dist/

