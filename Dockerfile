# # The build stage
# FROM golang:1.16-buster as builder
# WORKDIR /app
# COPY . .
# RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o go-api main.go

# # The run stage
# FROM scratch
# WORKDIR /app
# COPY --from=builder /app/go-api .
# EXPOSE 3000
# CMD ["./go-api"]



FROM golang:1.18-bullseye as base

RUN adduser \
  --disabled-password \
  --gecos "" \
  --home "/nonexistent" \
  --shell "/sbin/nologin" \
  --no-create-home \
  --uid 65532 \
  small-user

WORKDIR $GOPATH/src/smallest-golang/app/

COPY . .

RUN go mod download
RUN go mod verify

RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o /main .

FROM scratch

COPY --from=base /usr/share/zoneinfo /usr/share/zoneinfo
COPY --from=base /etc/ssl/certs/ca-certificates.crt /etc/ssl/certs/
COPY --from=base /etc/passwd /etc/passwd
COPY --from=base /etc/group /etc/group

COPY --from=base /main .

USER small-user:small-user

CMD ["./main"]

# docker build -t i_go_basic_build_scratch -f Dockerfile_basic_build_scratch .
# docker run -it --rm -p 3004:3000 --name=go-api i_go_basic_build_scratch