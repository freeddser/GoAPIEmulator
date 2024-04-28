### API Emulator
    http://localhost:8688/
    {
    "message": "hello world"
    }

### build
    ### MAC
    CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build main.go
    CGO_ENABLED=0 GOOS=windows GOARCH=amd64 go build main.go
    
    ### Linux
    CGO_ENABLED=0 GOOS=darwin GOARCH=amd64 go build main.go
    CGO_ENABLED=0 GOOS=windows GOARCH=amd64 go build main.go

    ### Windows
    SET CGO_ENABLED=0
    SET GOOS=darwin
    SET GOARCH=amd64
    go build main.go
    
    SET CGO_ENABLED=0
    SET GOOS=linux
    SET GOARCH=amd64
    go build main.go

### API List
#### /ping
    {"message":"pong"}

### /timeout/:number
    ### default sleep 10s
    {
    "msg": "success",
    "response": "It sleep done:10s"
    }

### /header
    apikey freeddser
    ### api key error
    {
    "msg": "failed",
    "response": "apikey error"
    }
    ### success
    {
    "msg": "success",
    "response": "api works"
    }