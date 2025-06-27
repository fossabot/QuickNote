FROM alpine:latest

RUN addgroup -S appgroup && adduser -S appuser -G appgroup && \
    mkdir -p /opt/quicknote

WORKDIR /opt/quicknote

COPY --chown=appuser:appgroup . /opt/quicknote
RUN mv /opt/quicknote/Frontend/static /opt/quicknote/static && \
    rm -rf /opt/quicknote/Frontend

USER appuser

EXPOSE 3000

CMD ["./QuickNote"]
