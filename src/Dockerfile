FROM golang:1.10 AS builder

WORKDIR $GOPATH/src/ 
COPY . ./
WORKDIR $GOPATH/src/main
RUN go get github.com/tools/godep
RUN godep get
RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o /main .

FROM scratch
COPY --from=busybox:1.28 /bin/busybox /bin/busybox
COPY --from=builder /main ./
EXPOSE 8081
ENTRYPOINT ["./main" ]
