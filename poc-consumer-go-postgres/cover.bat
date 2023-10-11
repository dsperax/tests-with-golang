go test -v -coverprofile resultado.out %1
go tool cover -html resultado.out -o resultado.html