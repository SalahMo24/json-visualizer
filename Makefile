# The default target when you just run `make`
.PHONY: run

run:
	/Users/salah/go/bin/templ generate
	@go run cmd/main.go