module github.com/ruancaetano/grpc-graphql-store/users

go 1.19

require (
	github.com/joho/godotenv v1.4.0
	github.com/lib/pq v1.10.7
	github.com/ruancaetano/grpc-graphql-store/shared v0.0.0-00010101000000-000000000000
	google.golang.org/grpc v1.49.0
	google.golang.org/protobuf v1.28.1
)

require (
	github.com/golang/protobuf v1.5.2 // indirect
	golang.org/x/net v0.0.0-20201021035429-f5854403a974 // indirect
	golang.org/x/sys v0.0.0-20210119212857-b64e53b001e4 // indirect
	golang.org/x/text v0.3.3 // indirect
	google.golang.org/genproto v0.0.0-20200526211855-cb27e3aa2013 // indirect
)

replace github.com/ruancaetano/grpc-graphql-store/shared => ../shared
