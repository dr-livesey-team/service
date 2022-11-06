FROM ubuntu:20.04

ENV TZ = Europe/Moscow
ENV DEBIAN_FRONTEND = noninteractive

COPY install_dependences.sh /
RUN ["/bin/bash", "install_dependences.sh"]

EXPOSE 80

CMD ["/bin/bash"]