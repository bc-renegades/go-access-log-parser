FROM golang:1.14-stretch

WORKDIR /app

COPY . .

RUN go mod download
RUN go get github.com/pilu/fresh

#COPY runner.conf /

#EXPOSE 3001

#ENTRYPOINT ["fresh", "-c", "/runner.conf"]