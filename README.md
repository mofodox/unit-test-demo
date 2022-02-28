# unit-test-demo

Code sharing session on Unit Testing in Go

# To get started



- Clone the project
- To run the test: `go test ./... -v` on the root project folder
- To test code coverage:
  ```
  # Create the output file
  go test ./... -coverprofile=<name>.out -v
  
  # View the coverage in the browser
  go tool cover -html=<name>.out
  ```
