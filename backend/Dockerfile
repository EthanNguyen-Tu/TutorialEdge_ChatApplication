FROM golang:1.23.2
RUN mkdir /app
ADD . /app/
WORKDIR /app
RUN go mod download
RUN go build -o /app/main ./
CMD ["/app/main"]