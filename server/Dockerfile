# FROM golang:1.23.4-alpine AS builder

# RUN apk add --no-cache git tzdata
# ENV TZ=Asia/Jakarta
# ENV GOPROXY=https://proxy.golang.org,direct

# WORKDIR /app/server
# COPY . .

# RUN go mod download

# # Build dan beri permission EKSEKUSI langsung di builder
# RUN go build -o /app/server /app/server/cmd/server/main.go && chmod +x /app/server
# RUN go build -o /app/worker /app/server/cmd/worker/worker.go && chmod +x /app/worker

# # --- RUNTIME IMAGE ---
# FROM alpine:latest
# RUN apk --no-cache add ca-certificates

# WORKDIR /app
# COPY --from=builder /app/server /app/server
# COPY --from=builder /app/worker /app/worker
# COPY wait-for-it.sh /app/wait-for-it.sh

# # Tambahkan permission eksekusi skrip
# RUN chmod +x /app/server /app/worker /app/wait-for-it.sh

# ENTRYPOINT ["/app/wait-for-it.sh", "db:3306", "--"]

FROM golang:1.23.4-alpine AS builder

RUN apk add --no-cache git tzdata
ENV TZ=Asia/Jakarta
ENV GOPROXY=https://proxy.golang.org,direct

WORKDIR /build
COPY . .

RUN go mod download

# Build dengan path output yang jelas
RUN go build -o /app/server ./cmd/server/main.go
RUN go build -o /app/worker ./cmd/worker/worker.go

# Set permission di builder
RUN chmod +x /app/server /app/worker

# --- RUNTIME IMAGE ---
FROM alpine:latest
RUN apk --no-cache add ca-certificates bash netcat-openbsd

WORKDIR /app
COPY --from=builder /app/server /app/server
COPY --from=builder /app/worker /app/worker
COPY --from=builder /build/internal/seeders /app/internal/seeders

COPY wait-for-it.sh /app/wait-for-it.sh

# Set permission
RUN chmod +x /app/server /app/worker /app/wait-for-it.sh

ENTRYPOINT ["/app/wait-for-it.sh", "db:3306", "--"]