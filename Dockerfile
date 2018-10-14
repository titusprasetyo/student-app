FROM golang:1.8-alpine
ADD . /go/src/student-app
RUN go install student-app

FROM alpine:latest
COPY --from=0 /go/bin/student-app .
ENV PORT 8080
CMD ["./student-app"]