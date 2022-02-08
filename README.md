## steps to reproduce
1) git clone this repo
2) go run . ( everything will work fine )
3) delete all code from `graph/schema.resolvers.go` after working
4) run `go run github.com/99designs/gqlgen generate` ( the generation passes without any errors, but does not generate anything )
5) run `go run . ` ( code will not run )
