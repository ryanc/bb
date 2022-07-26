FROM golang:latest AS build
WORKDIR /src
COPY . .
RUN CGO_ENABLED=0 go build ./cmd/bb

FROM scratch AS bin
COPY --from=build /src/bb /