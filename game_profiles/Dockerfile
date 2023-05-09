FROM golang:1.19-alpine AS build
RUN apk add build-base openssh
WORKDIR /build
COPY go.mod .
COPY go.sum .
RUN go mod download
RUN go mod tidy
COPY . .
RUN go build -o game_profiles ./cmd

FROM alpine:latest
WORKDIR /app/
COPY --from=build /build/game_profiles /app/game_profiles

# ENV ORIGIN_ALLOWED https://app.bageldb.com
ENV DEBUG true
# ENV ORIGIN_ALLOWED '*'
CMD ["./game_profiles"]