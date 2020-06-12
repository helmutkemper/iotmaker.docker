FROM golang:latest

#RUN apt-get update && apt-get install -y ca-certificates git-core ssh
#RUN apt-get update && apt-get install -y git
#RUN apt-get update

#ADD keys/my_key_rsa /root/.ssh/id_rsa
#RUN chmod 700 /root/.ssh/id_rsa
#RUN echo "Host github.com\n\tStrictHostKeyChecking no\n" >> /root/.ssh/config
#RUN git config --global url.ssh://git@github.com/.insteadOf https://github.com/

RUN mkdir /app
RUN chmod 700 /app

RUN mkdir /app/lixo
RUN chmod 700 /app/lixo

ADD ./ /app

WORKDIR /app

#RUN go get -u
#RUN go build -o ./main ./main.go

#RUN adduser -S -D -H -h /app appuser
#USER appuser

#EXPOSE 8080

RUN echo 'ping localhost &' > /bootstrap.sh
RUN echo 'sleep infinity' >> /bootstrap.sh
RUN chmod +x /bootstrap.sh

CMD ["/bootstrap.sh"]