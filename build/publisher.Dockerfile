FROM golang:1.15.7-alpine as build
RUN apk update && apk add git
WORKDIR /temp
COPY cmd/app/publisher ./
COPY go.mod go.sum ./
RUN go mod download
RUN go build -o ./main

FROM golang:1.15.7-alpine
COPY --from=build /temp/main /main
COPY key_pair/ ./key_pair
COPY web/html/publisher/ ./web
COPY web/image/publisher/ ./web
CMD [ "/main" ]