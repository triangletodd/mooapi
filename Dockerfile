FROM golang:alpine AS builder
ENV GOPATH=""
COPY . .
RUN go test -v
RUN go build -o /mooapi

FROM scratch
COPY --from=builder /mooapi /mooapi

ENTRYPOINT ["/mooapi"]
