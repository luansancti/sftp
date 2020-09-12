FROM golang
RUN apt update && apt install ssh openssh-server rsyslog -y
ENV GOPATH=/go/src/sftp
WORKDIR app
USER root
RUN mkdir /var/run/sshd
RUN sed 's@session\s*required\s*pam_loginuid.so@session optional pam_loginuid.so@g' -i /etc/pam.d/sshd
COPY ["./conf/sshd_config", "/etc/ssh/sshd_config"]
COPY ["./conf/docker-entrypoint.sh", "/app/docker-entrypoint.sh"]
RUN groupadd -f sftpgroup
RUN mkdir -p /data/users
RUN chmod 755 /data/users
RUN useradd --shell /bin/false --home-dir /data admin
RUN usermod -a -G sftpgroup admin
ENV NOTVISIBLE "in users profile"
RUN echo "export VISIBLE=now" >> /etc/profile
ENV TZ=America/Sao_Paulo
RUN ln -snf /usr/share/zoneinfo/$TZ /etc/localtime && echo $TZ > /etc/timezone
EXPOSE 22
WORKDIR /go/src/sftp
ENTRYPOINT ["bash", "/app/docker-entrypoint.sh"]