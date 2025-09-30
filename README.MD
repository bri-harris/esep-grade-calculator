# ESEP Grade Calculator

This library is a simple utility to calculate a student's grades.

Runtime requirements:
Golang 1.24+

### Installation Steps:
1. Navigate to https://go.dev/doc/install, select and download the Go installation package for your operating system.
2. Follow the installation prompts.
3. Verify that you've installed Go by opening a command prompt and typing the following command: `$ go version`
4. Confirm that the command prints the installed version of Go.

### Project Structure:
- go.mod: Module Definition
- grade_calculator.go: Grade Calculator implementation
- grade_calculator_tests.go: Unit Tests

### Unit Test Usage
Run unit tests:
```
cd grade_calculator/
go test ./...
```

Run unit tests with code coverage:
```
cd grade_calculator/
go test -coverprofile=coverage.txt ./...
```