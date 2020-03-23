FROM golang:1.13-alpine as build
ENV CGO_ENABLED=0
ENV GOOS=linux
WORKDIR /golang
COPY . .
RUN go build -a -tags netgo -ldflags '-w' -o app .

FROM alpine:3.10
LABEL MAINTAINER=1997jirasak@gmail.com
ENV TZ=Asia/Bangkok
WORKDIR /go
RUN apk add --update --no-cache tzdata ca-certificates
COPY --from=build /golang/app app
ENTRYPOINT [ "./app" ]