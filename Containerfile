FROM go-toolset:1.19.9-2 AS build

# Add Maintainer Info
LABEL maintainer="Matt Kimberley <matt.kimberley@redhat.com>"
COPY ./src .

# Build the Go app
RUN go mod init bootcamp && \
    go mod tidy -e && \
    go build .

FROM ubi8/ubi-micro
COPY --from=build /opt/app-root/src/bootcamp /app/main

# Expose port 8080 to the outside world
EXPOSE 8080/tcp
# Command to run the executable
CMD ["/app/main"]