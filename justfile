PROJECT_NAME := "bootstrap-go-web"
alias b:=build
alias r:=run
build:
    @echo "Building..."
    @go build -o bootstrap-go-web .

run: build
    @echo "Running..."
    @ENV=local ./bootstrap-go-web