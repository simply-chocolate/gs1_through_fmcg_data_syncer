FROM golang:1.19-alpine as builder

WORKDIR /app

COPY go.mod go.sum ./

RUN go mod download

COPY . ./

RUN go build -v -o gs1_syncer

##################################################
#                    Runner                      #
##################################################

FROM alpine:3.17 as runner

WORKDIR /app

COPY --from=builder /app/gs1_syncer ./

CMD [ "./gs1_syncer" ]
