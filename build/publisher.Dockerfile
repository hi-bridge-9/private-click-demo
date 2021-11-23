FROM golang:1.15.7-alpine as build
RUN apk update && apk add git
WORKDIR /temp
COPY cmd/app/publisher ./
COPY cmd/go.mod cmd/go.sum ./
RUN go mod download
RUN go build -o /temp/main

FROM golang:1.15.7-alpine
COPY --from=build /temp/main /main
COPY key_pair ./
CMD [ "/main" ]