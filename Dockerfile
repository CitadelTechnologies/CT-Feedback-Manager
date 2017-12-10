FROM golang:1.9

WORKDIR /go/src/ct-feedback-manager
COPY . .

RUN go-wrapper download
RUN go-wrapper install

EXPOSE 80

CMD ["go-wrapper", "run"]
