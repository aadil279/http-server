FROM golang:1.25rc1-alpine3.22

WORKDIR /app

COPY . /app

RUN go build -o /bin/main /app/src/main.go


FROM scratch

COPY --from=0 /bin/main /bin/main

CMD ["/bin/main"]
