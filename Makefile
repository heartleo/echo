build:
	export CGO_ENABLED=0 && go build -o ./echo ./cmd/echo