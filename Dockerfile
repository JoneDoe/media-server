# Dockerfile_old References: https://docs.docker.com/engine/reference/builder/

# Start from the latest golang base image
FROM golang:latest as builder

# Add Maintainer Info
LABEL maintainer="Vadym Titov <vad.titov@gmail.com>"

# Set the Current Working Directory inside the container
WORKDIR /app


######## Start a new volume stage #######
# Build Args
ARG STORAGE_DIR=/app/storage

# Create Log Directory
RUN mkdir -p ${STORAGE_DIR}

# Declare volumes to mount
VOLUME [${STORAGE_DIR}]

# Environment Variables
ENV STORAGE_DIR=${MASTER_NAME}
#ENV LOG_FILE_LOCATION=${STORAGE_DIR}/app.log

# Copy go mod and sum files
#COPY src/go.mod src/go.sum ./

# Copy the source from the current directory to the Working Directory inside the container
COPY src/ .

# Download all dependencies. Dependencies will be cached if the go.mod and go.sum files are not changed
RUN go mod download

# Build the Go app
RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o bin/media-storage .


######## Start a new stage from scratch #######
FROM alpine:latest

RUN apk --no-cache add ca-certificates

# Copy server configuration
COPY server.cfg .

# Copy the Pre-built binary file from the previous stage
COPY --from=builder /app/bin/media-storage .

# Expose port 8080 to the outside world
EXPOSE 8080

# Command to run the executable
CMD ["./media-storage"]
