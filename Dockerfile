FROM public.ecr.aws/u0v8y0z7/golang:1.17.2 as builder
WORKDIR /app/
COPY . .
RUN rm -rf .git env.yml
RUN export GOPRIVATE=github.com/es-hs
RUN go env -w GOPRIVATE=github.com/es-hs/*
RUN CGO_ENABLED=0 go build -o authz_service .

FROM public.ecr.aws/u0v8y0z7/alpine
ARG GRPC_PORT=50051
ARG RUN_ARGS="-task server"
ENV RUN_ARGS=${RUN_ARGS}
COPY --from=builder /app /app
WORKDIR /app
EXPOSE ${GRPC_PORT}
CMD ["sh", "-c", "./authz_service $RUN_ARGS"]
