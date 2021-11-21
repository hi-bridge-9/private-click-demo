FROM golang:1.15.7-alpine as build
RUN apk update && apk add git
WORKDIR /temp
COPY cmd/app/deliver ./
COPY cmd/go.mod cmd/go.sum ./
RUN go mod download
RUN go build -o /temp/main

FROM golang:1.15.7-alpine
COPY --from=build /temp/main /main
CMD [ "/main" ]