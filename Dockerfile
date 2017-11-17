FROM debian

COPY ./conductor /conductor
EXPOSE 8080
ENTRYPOINT /conductor
