FROM golang:1.23

WORKDIR /app
COPY . /app

RUN go install

CMD ["go", "run", "."]