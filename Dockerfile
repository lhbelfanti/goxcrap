FROM golang:1.22.3-alpine

LABEL maintainer="Lucas Belfanti"

# Install necessary dependencies
RUN apk update && apk add --no-cache \
    chromium \
    chromium-chromedriver \
    nss \
    freetype \
    harfbuzz \
    && rm -rf /var/cache/apk/*

# Create the application directory and set it as the working directory
WORKDIR /app

# Copy only the Go module files to leverage caching and download go dependencies
COPY go.mod go.sum ./
RUN go mod download

# Copy the application source code, the migrations folder and the .env file
COPY cmd/ ./cmd
COPY internal/ ./internal
COPY .env ./

# Build the application and output the binary as 'goxcrap'
RUN go build -o /goxcrap ./cmd/api

# Expose port
EXPOSE 8091

# Run application
CMD [ "/goxcrap" ]
