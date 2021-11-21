FROM mysql:8.0

RUN apt update
RUN apt -y install locales-all

ENV LANG ja_JP.UTF-8
ENV LANGUAGE ja_JP:ja
ENV LC_ALL ja_JP.UTF-8

ARG UID=1000
ARG GID=1000
ARG USER=docker_user
ARG GROUPNAME=docker_group

RUN groupadd -g $GID $GROUPNAME \
    && useradd --uid $UID --gid $GID $USER
USER $USER
