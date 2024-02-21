FROM golang:1.22.0
WORKDIR /app
COPY go.mod go.sum ./
RUN go mod download
COPY . .
RUN go build -o /publisher
EXPOSE 4545
CMD ["/publisher"]
