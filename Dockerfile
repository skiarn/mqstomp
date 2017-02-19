#Build: `docker build -t amq .`
#Run: `docker run -p 8161:8161 -p 61613:61613 -itd amq`
#Run with Bash `docker run --rm -it amq bash`
#Browse: http://localhost:8161/hawtio/#/login
FROM ubuntu:latest

MAINTAINER skiarn@users.noreply.github.com

ENV AMQ_VERSION 5.9.0
ENV HOME /opt/activemq
ENV USER activemq
ENV JAVA_HOME /usr/lib/jvm/java-8-openjdk-amd64/jre/bin/java

# Install OpenJDK8-jre-headless
RUN apt-get update && \
    apt-get upgrade -y && \
    apt-get install openjdk-8-jre-headless -y && \
    apt-get install curl -y && \
    apt-get clean all

# Start install
RUN mkdir -p $HOME

RUN curl -LO http://archive.apache.org/dist/activemq/apache-activemq/$AMQ_VERSION/apache-activemq-$AMQ_VERSION-bin.tar.gz
RUN tar -xvzf apache-activemq-$AMQ_VERSION-bin.tar.gz
RUN mv apache-activemq-$AMQ_VERSION/* $HOME

RUN groupadd $USER
RUN useradd --system --home $HOME -g $USER $USER
RUN chown -R $USER:$USER $HOME

RUN chmod 755 $HOME/bin/activemq

# Expose all port
#docker run -p 8161:8161 -p 61613:61613 -itd amq
EXPOSE 8161
EXPOSE 61616
EXPOSE 5672
EXPOSE 61613
EXPOSE 1883
EXPOSE 61614


ADD ./run.sh /run.sh
RUN chmod +x /run.sh

CMD ["/run.sh"]
