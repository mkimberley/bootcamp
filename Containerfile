FROM golang:latest

# Add Maintainer Info
LABEL maintainer="Matt Kimberley <mattkimberley84@gmail.com>"
RUN mkdir /app
COPY src/ /app
COPY go.mod /app
WORKDIR /app

# Build the Go app
RUN go build -o main .

# Expose port 8080 to the outside world
EXPOSE 80/tcp

# Command to run the executable
CMD ["/app/main"]