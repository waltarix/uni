build:
	goreleaser --clean --snapshot --skip-publish

test:
	CGO_ENABLED=0 go test -v ./...

clean:
	$(RM) -r dist

.PHONY: build test clean
