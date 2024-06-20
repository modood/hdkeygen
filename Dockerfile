FROM golang:alpine as builder


WORKDIR /build
COPY ./ ./
RUN go mod download
RUN CGO_ENABLED=0 go build -o ./hdkeygen


FROM scratch
WORKDIR /app
COPY --from=builder /build/hdkeygen ./hdkeygen
ENTRYPOINT ["./hdkeygen"]
