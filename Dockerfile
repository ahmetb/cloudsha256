FROM golang:1-alpine AS build
COPY main.go .
RUN CGO_ENABLED=0 go build -o /out/a.out .
FROM scratch
COPY --from=build /out/a.out /out/a.out
ENTRYPOINT ["/out/a.out"]
