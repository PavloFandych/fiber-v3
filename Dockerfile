FROM alpine:3.20.3
WORKDIR /app
COPY /bin/fiber-v3 .
EXPOSE 3000
CMD ["./fiber-v3"]