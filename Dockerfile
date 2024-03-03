FROM golang:1.20


WORKDIR /usr/src/app

COPY go.mod go.sum ./
RUN go mod download && go mod verify

COPY . . 

RUN go build -o ./super-hero-build

EXPOSE 3222

CMD ["./super-hero-build", "serve"]