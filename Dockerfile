FROM golang:1.26.1-alpine AS builder

WORKDIR /app

# ローカルで作成済みの go.mod と go.sum ファイルを コンテナ内の /app ディレクトリにコピーして依存関係をダウンロード
COPY go.mod go.sum ./
RUN go mod download

# ソースコードをコピーしてビルド
COPY . .

# main.go をコンパイルして、api-sever という名前の単一の実行ファイルを作成
RUN CGO_ENABLED=0 GOOS=linux go build -o api-server main.go

# Goの環境を持たない、超軽量なまっさらなOS環境を使用
FROM alpine:latest

WORKDIR /root/

# ビルダーステージで作られた実行ファイル（api-server）だけをコピー
COPY --from=builder /app/api-server .

# ポートを開放
EXPOSE 8080

# コンテナ起動時に実行されるコマンド
CMD ["./api-server"]
