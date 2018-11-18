FROM scratch

COPY ./dist/cpm /usr/local/bin/cpm

ENTRYPOINT ["/usr/local/bin/cpm"]
CMD ["--help"]
