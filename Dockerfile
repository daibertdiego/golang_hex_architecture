FROM golang:1.17.6-alpine

EXPOSE 9000

# Install MongoDB or use another Docker container with MongoDB

RUN mkdir /app
COPY go.mod .
COPY go.sum .
RUN go mod download
COPY . .

RUN go build cmd/main.go
RUN mv main /usr/local/bin
CMD [ "main" ]