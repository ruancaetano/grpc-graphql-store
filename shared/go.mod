module github.com/ruancaetano/grpc-graphql-store/shared

replace github.com/ruancaetano/grpc-graphql-store/users => ../users

go 1.19

require (
	github.com/ruancaetano/grpc-graphql-store/users v0.0.0-00010101000000-000000000000
	google.golang.org/grpc v1.49.0
)

require (
	github.com/golang/protobuf v1.5.2 // indirect
	golang.org/x/net v0.0.0-20211112202133-69e39bad7dc2 // indirect
	golang.org/x/sys v0.0.0-20210615035016-665e8c7367d1 // indirect
	golang.org/x/text v0.3.6 // indirect
	google.golang.org/genproto v0.0.0-20200526211855-cb27e3aa2013 // indirect
	google.golang.org/protobuf v1.28.1 // indirect
)
