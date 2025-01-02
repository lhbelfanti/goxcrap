# Build Stage
FROM golang:1.23.4-alpine AS builder

LABEL maintainer="Lucas Belfanti"

# Install necessary dependencies
RUN apk update && apk add --no-cache \
    curl \
    chromium \
    chromium-chromedriver \
    nss \
    freetype \
    harfbuzz \
    tzdata \
    && rm -rf /var/cache/apk/*

# Set timezone
ENV TZ=America/Argentina/Buenos_Aires

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


# Final Stage
FROM alpine:3.21.0 AS final

# Install necessary dependencies
RUN apk update && apk add --no-cache \
    curl \
    chromium \
    chromium-chromedriver \
    nss \
    freetype \
    harfbuzz \
    tzdata \
    && rm -rf /var/cache/apk/*

# Set timezone
ENV TZ=America/Argentina/Buenos_Aires

# Create the application directory and set it as the working directory
WORKDIR /app

# Copy the built image from builder stage
COPY --from=builder /goxcrap /goxcrap

# Expose port
EXPOSE ${API_PORT}

# Run application
CMD [ "/goxcrap" ]

