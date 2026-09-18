# --- build stage ---
FROM golang:1.27-bookworm AS build
WORKDIR /src

COPY go.mod go.sum ./
RUN go mod download

COPY . .
# CGO を無効化し静的リンクバイナリにする。distroless/static はlibcを持たないため必須。
RUN CGO_ENABLED=0 GOOS=linux go build -o /out/server .

# --- runtime stage ---
FROM gcr.io/distroless/static-debian12
WORKDIR /app
COPY --from=build /out/server ./server

ENV PORT=50051
EXPOSE 50051

ENTRYPOINT ["/app/server"]
