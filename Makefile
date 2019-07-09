test:
	go test ./... -v -race

fmt:
	bash -c 'diff -u <(echo -n) <(gofmt -s -d .)'

vet:
	bash -c 'diff -u <(echo -n) <(go vet ./...)'

test-all: fmt vet test
