## Testing

- go test <`folder/module`> -race -v -cover -failfast


## For Profiling

- go test <`folder/module`> -race -v -failfast -coverprofile=`<output file name>`.out

    - To Read the profile we can use `go tool`

- go tool cover -html=<`output file name`>.out -o <`final output file name`>.html

## Note

- `-count=1` can be used to avoid caching of test

- `-json` flag outputs to json

- `-cpu` to use cores

- `-race` for race

- `-v` for verbose

- `-cover` for coverage

- `-failfast` to stop execution at first failure
