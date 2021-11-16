FROM golang:1.15.7-alpine as build
RUN apk update && apk add git
WORKDIR /advertiser
COPY cmd/app/deliver .
COPY cmd/go.mod .
COPY cmd/go.sum .
RUN go mod download
RUN go build -o /advertiser/main

FROM golang:1.15.7-alpine
COPY --from=build /advertiser/main /main
CMD [ "/main" ]