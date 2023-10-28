
export GOPRIVATE="github.com/yoshi-jotaeyang/*"
export GONOPROXY="github.com/yoshi-jotaeyang/*"

go mod tidy

go build -o websocket-server.exe