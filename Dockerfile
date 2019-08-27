FROM golang:1.12.9-stretch
EXPOSE 8000

COPY mainsvc mainsvc

CMD ["./mainsvc"]

