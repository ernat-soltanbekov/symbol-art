.PHONY: build run suggest analyze audit test test-verbose golden help

build:
	go build -o symbol-art .

# make run TEXT="Your text, auditor"
run:
	go run . "$(TEXT)"

suggest:
	go run . "Code never sleeps!" standard --suggest

analyze:
	go run . "Bug? What bug?" thinkertoy --analyze

audit:
	go run . "Good luck, auditor!" shadow --analyze --suggest

test:
	go test ./...

test-verbose:
	go test -v ./...

golden:
	go test -run ^TestGolden$

help:
	@echo "Доступные команды:"
	@echo '  make run TEXT="Your text"'
	@echo "  make suggest"
	@echo "  make analyze"
	@echo "  make audit"
	@echo "  make test"
	@echo "  make golden"