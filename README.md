
# Testing with Go - examples

## Presentation
https://docs.google.com/presentation/d/1t_RgWzpIm_juVNdw42rhhG7P-ZYbynGoQ8Vj-sGJFrU/

## Coverage
```
# general coverage information. Run inside your application folder:
go test -cover ./... 

# for a single package within the application
go test -cover ./mocking 

# generate coverage report for a package. Will create a file cover.out with coverage data
go test -coverprofile=cover.out ./mocking

# generate report in your terminal using the cover.out file 
# make sure you did “go get golang.org/x/tools/...”
go tool cover -func=cover.out

# generate pretty output with HTML and open it in a browser
go tool cover -html=cover.out

# generate a heat map. You can see it with the HTML preview
go test -covermode=count -coverprofile=cover.out ./mocking
```

## Running tests
```
# all tests within all packages
go test ./...

# verbose mode
go test -v ./...

# single package
go test -v ./sub_tests

# specific test within a package
go test -v ./sub_tests -run=TestSumNumbers
 
# specific sub test within a test
go test -v ./sub_tests -run=TestSumNumbers/test_2
 
# set of tests filtered with regular expression
go test -v ./sub_tests -run=TestSumNumbers/test_"[0-9]"

# check for race conditions
go test -v -race ./...
```
