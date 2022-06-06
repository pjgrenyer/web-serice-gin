FROM golang:1.18

# apt-get install libpq-dev

COPY ./dist/app .

CMD ["./app"]
