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

# Copy the application source code and the .env file
COPY cmd/ ./cmd
COPY internal/ ./internal

# Build the application and output the binary as 'goxcrap'
RUN CGO_ENABLED=0 GOOS=linux go build -o /goxcrap ./cmd/api

# Expose port
EXPOSE 4000

# Run application
CMD [ "/goxcrap" ]

