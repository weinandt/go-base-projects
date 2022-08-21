# Tests
Contains examples on testing using just the go standard library. No mocking libraries or assertion frameworks are used.

There are two kinds of tests: 
1. A test for business logic where a stubbed implementation is needed to compile, but not used at runtime.
2. A stubbed implemetation of an interface where canned values set up in the test are returned.

## To run:
`go test ./...`