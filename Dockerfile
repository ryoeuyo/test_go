FROM golang:1.23.4

WORKDIR /app
COPY . .

EXPOSE 8081

ENV GOPROXY=https://proxy.golang.org,direct

RUN go mod tidy
RUN go mod download

RUN make build
CMD ["make", "run"]
