FROM alpine:latest-git

LABEL Vendor="Frog" \
    Name="mynote" \
    Version="1.0.0" \
    Date="12/15/2016"

COPY bin /

EXPOSE 8080

ENTRYPOINT ["/mynote"]
