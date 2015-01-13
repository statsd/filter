
build:
	gox -os="linux darwin" -arch=amd64 cmd/filter/main.go

clean:
	git clean -f

.PHONY: build clean