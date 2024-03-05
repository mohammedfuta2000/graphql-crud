 $ go get github.com/99designs/gqlgen
 $ printf '// +build tools\npackage tools\nimport _ "github.com/99designs/gqlgen"' | gofmt > tools.go
 $ go run github.com/99designs/gqlgen init
 delete graphql todo dummy code
 write schema.graphqls
 $ go run github.com/99designs/gqlgen generate
update code in schema.resolvers.go
 $ go run server.go

 most already built application arent expecting graphql
 there are no tests or validations in this project


 get job and get jobs do not return ID field
