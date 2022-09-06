FROM golang:latest AS build
WORKDIR /src
COPY . .
RUN CGO_ENABLED=0 go build ./cmd/bb

FROM scratch AS bin
COPY --from=build /etc/ssl/certs/ca-certificates.crt /etc/ssl/certs/
COPY --from=build /src/bb /
ADD https://raw.githubusercontent.com/golang/go/master/lib/time/zoneinfo.zip /zoneinfo.zip
ENV ZONEINFO /zoneinfo.zip
ENTRYPOINT ["/bb"]
