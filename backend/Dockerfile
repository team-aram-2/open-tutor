FROM golang:1.24-alpine

WORKDIR /app
COPY . /app

RUN go install

CMD ["go", "run", "."]