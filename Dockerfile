FROM alpine:latest

RUN addgroup -S quicknotegroup && \
    adduser -S quicknoteuser -G quicknotegroup && \
    mkdir -p /opt/quicknote

WORKDIR /opt/quicknote

COPY --chown=quicknoteuser:quicknotegroup . /opt/quicknote

RUN chmod -R a-w /opt/quicknote && \
    mv /opt/quicknote/Frontend/static /opt/quicknote/static && \
    rm -rf /opt/quicknote/Frontend

USER quicknoteuser

# default port
EXPOSE 3000

CMD ["./QuickNote"]