FROM golang as build
WORKDIR /build
COPY server.go .
ENV CGO_ENABLED=0
RUN go build hello.go

FROM scratch
COPY --from=build /build/server /
CMD [ "/server" ]
EXPOSE 8000
