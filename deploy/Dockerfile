# ===================== FUNCIONAL - DESARROLLO =====================

# ---------- Etapa de build ----------
FROM golang:1.24.3-alpine3.21 AS builder

WORKDIR /app
COPY .  .
RUN go mod download
COPY .air.toml ./

RUN go install github.com/air-verse/air@latest

# ---------- Etapa final ----------
FROM golang:1.24.3-alpine3.21

WORKDIR /app

# Copiamos solo lo necesario
COPY --from=builder /go/bin/air /usr/local/bin/air
COPY --from=builder /app/main.go ./
COPY --from=builder /app/go.mod ./
COPY --from=builder /app/.air.toml ./

CMD ["air", "-c", ".air.toml"]