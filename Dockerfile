FROM golang
WORKDIR /go/src/app

COPY . .

RUN go build -o ./shoutrrrd main.go

ENV SHOUTRRR_CONFIGURATION_FILE=/services.yaml
ENV SHOUTRRRD_HOST=:80
VOLUME [ "/services.yaml" ]
EXPOSE 80

CMD ["./shoutrrrd"]