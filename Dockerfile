# Use Go as the base
FROM golang:alpine

# Install Node.js and NPM
RUN apk add --no-cache nodejs npm

WORKDIR /app

COPY . .

# 1. Build the Frontend
WORKDIR /app/web
RUN npm install
RUN npm run build

# 2. Build the Backend
WORKDIR /app
RUN go mod download
RUN go build -o main .

EXPOSE 80

CMD ["./main"]