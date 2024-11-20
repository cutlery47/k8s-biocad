FROM golang:1.23.3-alpine3.20

COPY main.go .

EXPOSE 32777

CMD ["go", "run", "main.go"]