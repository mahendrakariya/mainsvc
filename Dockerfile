FROM golang:1.12.9-stretch

COPY mainsvc mainsvc

CMD ["./mainsvc"]

