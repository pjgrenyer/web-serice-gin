FROM golang:1.18

COPY ./dist/app .

CMD ["./app"]
