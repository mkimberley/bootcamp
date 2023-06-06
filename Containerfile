FROM golang:latest AS build

# Add Maintainer Info
LABEL maintainer="Matt Kimberley <mattkimberley84@gmail.com>"
RUN mkdir /app
COPY src/ /app
COPY go.mod /app
WORKDIR /app

# Build the Go app
RUN go get github.com/mkimbeley/bootcamp
RUN go build -o main .

# Expose port 8080 to the outside world
EXPOSE 8080/tcp


FROM scratch
COPY --from=build /app/main /app/main
# Command to run the executable
CMD ["/app/main"]