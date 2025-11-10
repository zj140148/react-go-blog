# 选择官方 Go 镜像，版本可根据你本地 Go 版本调整
FROM golang:1.25.1-alpine

# 设置工作目录
WORKDIR /app

# 安装必要工具（可选，用于热更新、调试）
RUN apk add --no-cache git bash curl

# 拷贝依赖文件
COPY go.mod go.sum ./

# 下载依赖
RUN go mod download

# 拷贝后端源码
COPY . .

# 编译程序（可选，开发阶段可以不编译直接用 go run）
RUN go build -o main .

# 暴露端口（根据你的后端实际端口）
EXPOSE 5000

# 容器启动命令（开发用 go run，可热更新）
CMD ["go", "run", "."]
# 如果用于生产，可改成：
# CMD ["./main"]
