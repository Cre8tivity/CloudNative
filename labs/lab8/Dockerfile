FROM golang:1.15-alpine AS build

WORKDIR /src/
COPY main /src/
COPY go.mod /src/
COPY go.sum /src/

RUN apk update && apk add git
RUN go mod download

RUN CGO_ENABLED=0 go build -o /bin/mwebserver

FROM scratch
COPY --from=build /bin/mwebserver /bin/mwebserver
ENTRYPOINT ["/bin/mwebserver"]