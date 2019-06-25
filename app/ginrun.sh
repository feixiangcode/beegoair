export GOBIN=/usr/local/gobin
export GOPATH=/data/code/go
export GOROOT=/usr/local/go

export BIN_APP_PORT="8080"
export GIN_BIN="gin.gobin"
export GIN_ALL=1
export GIN_PATH="/data/code/go/src/beegoair/"
export GIN_BUILD="/data/code/go/src/beegoair/app/"

#unset BIN_APP_PORT
#unset GIN_BIN
#unset GIN_ALL
#unset GIN_PATH
#unset GIN_BUILD

gin run
