FROM golang:1.22.1-alpine as builder

WORKDIR /app

COPY ./apps/backend/go.* ./

RUN go mod download

COPY ./apps/backend ./

RUN CGO_ENABLED=0 go build -a --trimpath --installsuffix cgo --ldflags="-s" -o main ./main.go

FROM scratch

COPY --from=builder /app/main ./

CMD [ "./main" ]