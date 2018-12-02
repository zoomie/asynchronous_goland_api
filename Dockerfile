FROM scratch
WORKDIR /
COPY app /
COPY src /
COPY pkg /
CMD ["/app"]