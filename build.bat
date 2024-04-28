@echo off
@REM SET CGO_ENABLED=0
@REM SET GOOS=darwin
@REM SET GOARCH=amd64
@REM go build main.go

SET CGO_ENABLED=0
SET GOOS=linux
SET GOARCH=amd64
go build main.go

set CGO_ENABLED=0
set GOOS=windows
set GOARCH=amd64
go build main.go