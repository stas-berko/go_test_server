FROM golang:alpine AS build-env
WORKDIR cmd
RUN apk --no-cache add build-base git gcc
COPY . .
RUN go build  -o goapp cmd/server/main.go

FROM alpine
WORKDIR /app
COPY templates templates
COPY --from=build-env /go/cmd/goapp /app/
ENTRYPOINT ./goapp