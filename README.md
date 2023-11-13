# grpc-user

## Folder Structure

- `business`: Define all business of service
- `handler`: Implement all rpc endpoint of grpc service
- `internal`: Save all third-party that use on service
- `micro`: Init connection to other service and function
- `pb`: Save all generated `.go` file from protobuf
- `proto`: Define grpc service
  - `types`: Define request and response of rpc
  - `service.proto`: Define rpc endpoint
- `types`: Exported types that can be use by another service
- `client.go` Define function connect to this service
- `server.go` main file to start service

## Install proto

```bash
# install tools
go install google.golang.org/protobuf/cmd/protoc-gen-go@latest
go install google.golang.org/grpc/cmd/protoc-gen-go-grpc@latest
```

```bash
cd <root folder>
generate model: protoc --proto_path=proto --go_out=plugins=grpc:. path/to/model.proto
EX: generate model: protoc --proto_path=proto --go_out=plugins=grpc:. proto/users/types/user.proto

generate service: protoc --proto_path=proto --go_out=plugins=grpc:. path/to/service.proto
EX: generate model: protoc --proto_path=proto --go_out=plugins=grpc:. proto/service.proto
```

### New version protoc version

```bash
Generate model:
protoc --proto_path=proto --go_out=. --go-grpc_out=require_unimplemented_servers=false:. proto/game/common/game.proto

EX:
protoc --proto_path=proto --go_out=. --go-grpc_out=require_unimplemented_servers=false:. proto/game/types/get_all_game_room.proto

Generate service:
protoc --proto_path=proto --go_out=. --go-grpc_out=require_unimplemented_servers=false:. path/to/service.proto

EX:
protoc --proto_path=proto --go_out=. --go-grpc_out=require_unimplemented_servers=false:. proto/service.proto
```

## How to connect to this service

```go
import (
    "fmt"
    "context"
    userClient "github.com/ponny-io/dogemega-grpc-user"
    "github.com/ponny-io/dogemega-grpc-user/pb/user"
)

func main(){

    client := userClient.NewMicroUserServiceClient("localhost:30000")

    // call rpc
    res, err := client.SVC.CreateUser(context.Background(), &user.CreateUserRequest{
        Data: &user.CreateUserRequestData{

        }
    })
    if err != nil {
        return
    }

    fmt.Println(res)
}
```

## Docker compose

### Setup env

```bash
#Exp:
MONGODB_HOST="localhost:27018"
MONGODB_DATABASE="user-database"
MONGODB_USERNAME="root"
MONGODB_PASSWORD="example"
```

```bash
docker-compose up -d
```

## Start service

```bash
go run server.go
```
