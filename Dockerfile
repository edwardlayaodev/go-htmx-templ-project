FROM golang:1.22

WORKDIR /usr/src/app

# copy env file
COPY .env ./

# pre-copy/cache go.mod for pre-downloading dependencies and only redownloading them in subsequent builds if they change
COPY go.mod ./
RUN go mod download && go mod verify

COPY . .
RUN go install github.com/githubnemo/CompileDaemon@latest

# builds main.go from ., after build run the binary ./main
# somehow there is an error on vcs status, we disable it for now via buildvcs=false
CMD ["CompileDaemon", "--build=go build -buildvcs=false -o main .", "--command=./main"]