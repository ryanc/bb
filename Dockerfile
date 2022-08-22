FROM golang:latest AS build
WORKDIR /src
COPY . .
RUN CGO_ENABLED=0 go build ./cmd/bb

FROM scratch AS bin
COPY --from=build /etc/ssl/certs/ca-certificates.crt /etc/ssl/certs/
COPY --from=build /src/bb /
ENTRYPOINT ["/bb"]
