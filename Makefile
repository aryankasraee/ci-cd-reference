MODULES := api worker

.PHONY: test lint build
test:
	@for m in $(MODULES); do (cd modules/$$m && go test -race ./...) || exit 1; done
lint:
	@for m in $(MODULES); do (cd modules/$$m && golangci-lint run ./...) || exit 1; done
build:
	@for m in $(MODULES); do (cd modules/$$m && go build ./...) || exit 1; done
