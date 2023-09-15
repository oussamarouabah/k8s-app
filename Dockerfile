FROM alpine

COPY ./lister ./lister

ENTRYPOINT [ "./lister" ]