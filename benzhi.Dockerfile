FROM golang:1.22 AS build
WORKDIR /src
COPY . .
RUN CGO_ENABLED=0 go build -o /out/cokeoven ./cmd/cokeoven

FROM gcr.io/distroless/static-debian12
COPY --from=build /out/cokeoven /cokeoven
ENTRYPOINT ["/cokeoven"]
