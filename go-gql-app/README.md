step -1 
    go mod init your-appname

step-2 

    create a file tools.go
    inside the files 

    //go:build tools
    // +build tools

    package tools

    import (
    _ "github.com/99designs/gqlgen"
    _ "github.com/99designs/gqlgen/graphql/introspection"
    )
step 3

go get -u  github.com/99designs/gqlgen
go run github.com/99designs/gqlgen init