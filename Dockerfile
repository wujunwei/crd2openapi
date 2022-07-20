FROM  golang:1.18 as builder

WORKDIR /workspace
COPY .  .
RUN go env -w GOPROXY=goproxy.io
RUN CGO_ENABLED=0 go build  -o ./convert main.go


From alpine:latest

WORKDIR app
COPY --from=builder /workspace/convert  /app/converter

ENTRYPOINT ["/app/converter"]