FROM scratch
WORKDIR /
COPY app /
COPY templates templates
CMD ["/app"]