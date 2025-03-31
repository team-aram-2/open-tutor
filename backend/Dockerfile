FROM golang:1.24

WORKDIR /app
COPY . /app

RUN go install

CMD ["go", "run", "."]