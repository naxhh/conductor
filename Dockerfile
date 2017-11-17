FROM debian

RUN mkdir /public
RUN mkdir /projects
COPY ./conductor /conductor
EXPOSE 8080
ENTRYPOINT /conductor
