# Dockerfile_old References: https://docs.docker.com/engine/reference/builder/

# Build Default Args
ARG APP_NAME=media-storage

# Start from the latest golang base image
FROM golang:latest as builder

# Build Args
ARG BUILD_PATH
ARG APP_NAME

# Set the Current Working Directory inside the container
WORKDIR ${BUILD_PATH}

# Copy the source from the current directory to the Working Directory inside the container
COPY src/ .

# Download all dependencies. Dependencies will be cached if the go.mod and go.sum files are not changed
RUN go mod download

# Build the Go app
RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o ${APP_NAME} .


######## Start a new stage from scratch #######
FROM ubuntu:18.04

# Add Maintainer Info
LABEL maintainer="Vadym Titov <vad.titov@gmail.com>"

# Build Args
ARG DATA_PATH
ARG BUILD_PATH
ARG APP_NAME

WORKDIR /app

######## Start a new volume stage #######
# Create Directory
RUN mkdir -p ${DATA_PATH}

# Copy server configuration
COPY server.cfg .

# Copy the Pre-built binary file from the previous stage
COPY --from=builder ${BUILD_PATH}/${APP_NAME} .

# Expose port 8080 to the outside world
EXPOSE 8080

# Command to run the executable
CMD ["./media-storage"]
