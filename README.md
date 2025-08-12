# Go Coverage Test

A minimal Go project designed to demonstrate Codecov behavior, especially partial line coverage scenarios.

## Project Structure

```
go-coverage-test/
├── cmd/app/main.go              # Demo application
├── internal/mathx/
│   ├── mathx.go                 # Math functions with partial coverage opportunities
│   └── mathx_test.go            # Table-driven tests with intentional gaps
├── Makefile                     # Build and coverage commands
├── go.mod                       # Go module definition
└── .github/workflows/ci.yml     # GitHub Actions CI pipeline
```

## Key Features for Codecov Testing

The `internal/mathx/mathx.go` file contains functions specifically designed to demonstrate partial line coverage:

1. **`IsPositiveAndEven`** - Uses chained `&&` conditions where tests only hit some branches
2. **`ProcessNumbers`** - Has multiple function calls on the same line with different execution paths
3. **`SmartDivide`** - Features inline conditional logic that creates partial coverage scenarios

Each function includes comments explaining why they're useful for observing partial coverage in Codecov.

## Local Development

### Prerequisites

- Go 1.22 or higher
- Make (optional, but recommended)

### Installation

```bash
git clone <this-repo>
cd go-coverage-test
go mod tidy
```

### Running the Demo

```bash
# Check Go version
go version

# Run the demo application
make run

# Or build and run manually
make build
./bin/app
```

### Running Tests

```bash
# Run all tests
make test

# Run tests with verbose output
make test-verbose

# Run benchmarks
make bench
```

### Coverage Analysis

Generate and view coverage reports:

```bash
# Generate coverage profile (creates coverage.out)
make coverage

# Generate function-level coverage report (creates coverage.txt)
make coverfunc

# Generate HTML coverage report (creates coverage.html)
make coverhtml

# View HTML report in browser
open coverage.html  # macOS
# or
xdg-open coverage.html  # Linux
```

The `coverage.out` file is what gets uploaded to Codecov in CI.

### Expected Coverage Results

When you run coverage analysis, you should see:

- **Partial line coverage** in functions with chained conditions
- **Lines marked as partially covered** where only some execution paths are tested
- **Function coverage** showing which functions are called vs. not called

Example expected output from `make coverfunc`:
```
go-coverage-test/internal/mathx/mathx.go:10:    IsPositiveAndEven       85.7%
go-coverage-test/internal/mathx/mathx.go:17:    ProcessNumbers          90.0%
go-coverage-test/internal/mathx/mathx.go:25:    SmartDivide            80.0%
```

## CI/CD Pipeline

This project includes a GitHub Actions workflow (`.github/workflows/ci.yml`) that:

1. Runs on `push` and `pull_request` events
2. Uses Ubuntu latest with Go 1.22.x
3. Executes tests with coverage collection
4. Uploads coverage results to Codecov
5. Fails CI if Codecov upload encounters errors

### Setting Up Codecov

1. Create account at [codecov.io](https://codecov.io)
2. Connect your GitHub repository
3. Add `CODECOV_TOKEN` to your repository secrets (if using private repo)
4. Push code to trigger the CI pipeline

### Codecov Integration

The workflow uploads `coverage.out` using the official Codecov action:

```yaml
- name: Upload coverage to Codecov
  uses: codecov/codecov-action@v4
  with:
    files: coverage.out
    fail_ci_if_error: true
```

## Partial Coverage Examples

### Example 1: Chained Boolean Conditions
```go
// mathx.go:10 - This line demonstrates partial coverage
return n > 0 && isEven(n) && expensiveValidation(n)
```
Tests with negative numbers will only execute the first condition, creating partial coverage.

### Example 2: Multiple Function Calls
```go
// mathx.go:20 - Mixed execution paths on same line
return complexCalculation(a, b) + simpleCalculation(a, b), validateResult(a, b)
```
Some tests trigger validation errors while others don't, creating partial line coverage.

### Example 3: Inline Conditional Logic
```go
// mathx.go:27 - Conditional execution on same line
result := func() float64 { if math.Abs(b) < 1e-10 { return math.Inf(1) } else { return a / b } }()
```
Different test inputs exercise different branches of the inline conditional.

## Troubleshooting

### Coverage Not Generated
- Ensure tests are passing: `make test`
- Check Go version: `go version` (should be 1.22+)
- Verify module setup: `go mod tidy`

### Codecov Upload Issues
- Check GitHub Actions logs for detailed error messages
- Verify `CODECOV_TOKEN` is set correctly (for private repos)
- Ensure `coverage.out` file is generated in CI

### Partial Coverage Not Visible
- Run `make coverhtml` and open the HTML report
- Look for yellow/orange highlighted lines indicating partial coverage
- Check that tests are intentionally not covering all code paths

## Contributing

This is a demonstration project. Feel free to:
- Add more complex partial coverage scenarios
- Improve test coverage strategies
- Enhance the CI pipeline
- Add additional analysis tools

## License

MIT License - see LICENSE file for details.