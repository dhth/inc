default: fmt test build
alias r := run
alias b := build
alias f := fmt
alias t := test
alias i := install

run:
  go run .

build:
  go build .

fmt:
  gofmt -w .

test:
  go test .

install:
  go install .
