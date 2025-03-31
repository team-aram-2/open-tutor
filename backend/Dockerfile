FROM golang:1.23-alpine

WORKDIR /app
COPY . /app

RUN go install

CMD ["go", "run", "."]