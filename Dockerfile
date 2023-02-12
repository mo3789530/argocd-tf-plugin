FROM golang:1.19-alpine as dev

ENV ROOT=/go/src/app
ENV CGO_ENABLED 0
WORKDIR ${ROOT}

RUN apk update 
RUN apk add make git

COPY go.mod go.sum ./
COPY . ${ROOT}

RUN go mod download

CMD ["go", "run", "main.go"]

FROM golang:1.19-alpine as builder

ENV ROOT=/go/src/app
ENV CGO_ENABLED 0
WORKDIR ${ROOT}

RUN apk update 
RUN apk add make git

COPY go.mod go.sum ./
COPY . ${ROOT}

RUN go mod download

RUN CGO_ENABLED=0 GOOS=linux go build -o $ROOT/binary


FROM scratch as prod

ENV ROOT=/go/src/app
WORKDIR ${ROOT}
COPY --from=builder ${ROOT}/binary ${ROOT}

EXPOSE 8080
CMD ["/go/src/app/binary"]