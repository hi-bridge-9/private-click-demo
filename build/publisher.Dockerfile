FROM golang:1.15.7-alpine as build
RUN apk update && apk add git
WORKDIR /publisher
COPY cmd/app/publisher .
COPY cmd/go.mod .
COPY cmd/go.sum .
RUN go mod download
RUN go build -o /publisher/main

FROM golang:1.15.7-alpine
COPY --from=build /publisher/main /main
CMD [ "/main" ]