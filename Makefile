build:
	go build -ldflags="-X 'main.Commit=$$(git rev-parse HEAD)' -X 'main.BuildDate=$$(git show -s --format=%ci HEAD)' -X 'main.Builder=$$(git show -s --format=%ae HEAD)'" -o gocrud cmd/*.go 

install:
	sudo mv gocrud /usr/local/bin
