FROM golang:1.22

WORKDIR /app

COPY go.mod go.sum ./

RUN go mod download



ENV PATH="/go/bin:${PATH}"

COPY . .

RUN go build -o main .

EXPOSE 8080

CMD ["air"]