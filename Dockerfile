FROM alpine:latest

RUN addgroup -S quicknotegroup && \
    adduser -S quicknoteuser -G quicknotegroup && \
    mkdir -p /opt/quicknote

WORKDIR /opt/quicknote

COPY --chown=quicknoteuser:quicknotegroup . /opt/quicknote

RUN mv /opt/quicknote/Frontend/static /opt/quicknote/static && \
    rm -rf /opt/quicknote/Frontend

USER quicknoteuser

# default open port
EXPOSE 3000

CMD ["./QuickNote"]
