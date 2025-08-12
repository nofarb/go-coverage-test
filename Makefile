.PHONY: test coverage coverfunc coverhtml clean build run

# Run all tests
test:
	go test ./...

# Generate coverage profile
coverage:
	go test ./... -covermode=count -coverprofile=coverage.out

# Generate function-level coverage report
coverfunc: coverage
	go tool cover -func=coverage.out -o coverage.txt

# Generate HTML coverage report
coverhtml: coverage
	go tool cover -html=coverage.out -o coverage.html

# Build the application
build:
	go build -o bin/app ./cmd/app

# Run the application
run: build
	./bin/app

# Clean generated files
clean:
	rm -f coverage.out coverage.txt coverage.html
	rm -rf bin/

# Install dependencies
deps:
	go mod tidy
	go mod download

# Run tests with verbose output
test-verbose:
	go test -v ./...

# Run benchmarks
bench:
	go test -bench=. ./...