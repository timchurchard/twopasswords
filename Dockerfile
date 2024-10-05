# timchurchard/twopasswords docker
FROM golang:1.21-bookworm as builder

WORKDIR /build

COPY . .

RUN go build .


# Wipe and copy minimal stuff (including deps for examples)
FROM golang:1.21-bookworm

RUN go install -v github.com/martinhoefling/goxkcdpwgen@latest

RUN apt update && apt install -y imagemagick qrencode

COPY --from=builder /build/twopasswords /

COPY --from=builder /build/examples /examples

ENV PATH=$PATH:/

CMD ["/twopasswords"]
