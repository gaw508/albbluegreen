test:
	go test ./... -v -race

fmt:
	bash -c 'diff -u <(echo -n) <(gofmt -s -d .)'

vet:
	bash -c 'diff -u <(echo -n) <(go vet ./...)'

mod:
	go mod download

test-all: mod fmt vet test
