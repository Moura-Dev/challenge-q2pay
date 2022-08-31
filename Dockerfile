# build stage
FROM golang:alpine AS build-env
RUN apk --no-cache add build-base

COPY . /src
WORKDIR /src
RUN go get ./...&& go build -o goapp

# final stage
FROM alpine
WORKDIR /app
COPY --from=build-env /src/goapp /app/
CMD ["./start.sh"]
ENTRYPOINT ./goapp