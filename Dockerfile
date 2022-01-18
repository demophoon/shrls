# Go Shrls
FROM golang:1.16-alpine AS builder
WORKDIR /app

COPY ./go.mod ./
COPY ./go.sum ./
RUN go mod download

COPY ./shrls/*.go ./
RUN CGO_ENABLED=0 GOOS=linux go build -ldflags "-s" -a --installsuffix cgo -o shrls

# Shrls Frontend
FROM node:17-alpine as frontend
WORKDIR /app

COPY . /app
RUN npm install
RUN npm run build

# Final Artifact
FROM scratch
COPY --from=builder /app/shrls ./
COPY --from=frontend /app/dist/ /dist/

CMD ["./shrls"]
