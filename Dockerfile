FROM golang:1.22 as build

WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download && go mod verify

COPY *.go ./
RUN go build  -v -o /go-task-list



FROM gcr.io/distroless/base-debian12

WORKDIR /

COPY --from=build /go-task-list /go-task-list

EXPOSE 80

USER nonroot:nonroot

CMD [ "/go-task-list" ]