FROM golang:1.17-alpine AS build

WORKDIR /src/

COPY main/server.go /src/
COPY movieapi /src/
COPY go.mod /src/
COPY go.sum /src/

RUN apk update && apk add git
RUN go mod download
RUN go get "github.com/Cre8tivity/CloudNative/labs/lab7/lab5COPY/movieapi"

EXPOSE 50051

RUN CGO_ENABLED=0 go build -o /bin/server

FROM scratch
COPY --from=build /bin/server /bin/server
ENTRYPOINT ["/bin/server"]