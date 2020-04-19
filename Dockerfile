FROM golang:1.13 as builder

RUN mkdir /build
ADD . /build
WORKDIR /build

RUN go build -v -o emem cmd/server/main.go


COPY --from=builder /build/emem /app/emem

RUN chmod +x /app/emem

WORKDIR /app

CMD /app/emem