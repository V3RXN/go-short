FROM golang:1.19 AS build

WORKDIR /src 

COPY go.mod .
COPY go.sum .
RUN go mod download && \ 
    go mod verify

COPY . .
RUN go build -v -o ./dist/go-short-app

FROM golang:1.19 AS production

WORKDIR /app

COPY --from=build /src/dist/go-short-app ./go-short-app

CMD ./go-short-app