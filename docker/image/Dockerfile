FROM ubuntu:20.04

ENV TZ = Europe/Moscow
ENV DEBIAN_FRONTEND = noninteractive

COPY install_dependences.sh /
RUN ["/bin/bash", "install_dependences.sh"]
RUN ["rm", "install_dependences.sh"]

EXPOSE 80
EXPOSE 50050

CMD ["/bin/bash"]