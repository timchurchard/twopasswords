# timchurchard/twopasswords docker
FROM golang:1.21-bookworm as builder

WORKDIR /build

COPY . .

RUN go build .

RUN go install -v github.com/martinhoefling/goxkcdpwgen@latest && which goxkcdpwgen

COPY ./scripts/generate-addr.sh /


# Wipe and copy minimal stuff
FROM golang:1.21-bookworm

COPY --from=builder /build/twopasswords /

COPY --from=builder /go/bin/goxkcdpwgen /

COPY --from=builder /generate-addr.sh /

ENV PATH=$PATH:/

CMD ["/twopasswords"]
