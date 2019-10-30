FROM ubuntu:16.04

# Install libraries
RUN apt-get update -qq \
    && apt-get install -yq cmake \
			   apt-transport-https \ 
			   ca-certificates \
                           fceux \
                           gcc \
                           libboost-all-dev \
                           libjpeg-dev \
                           libsdl2-dev \
                           make \
                           unzip \
                           wget \
                           zlib1g-dev \
                           supervisor \
                           git \
                           curl \
                           redis-server \
                           mercurial
                        

# Install go tarball
ENV GOLANG_VERSION 1.13.3
RUN wget -qO- https://golang.org/dl/go${GOLANG_VERSION}.linux-amd64.tar.gz | tar -C /usr/local -xzf -

# Set PATHS for GOLANG 
ENV HOME /root
RUN mkdir ${HOME}/go
ENV GOPATH ${HOME}/go
ENV PATH ${PATH}:${GOPATH}/bin:/usr/local/go/bin

WORKDIR ${HOME}

# Install revel framework 
RUN go get -u github.com/revel/revel
RUN go get -u github.com/revel/cmd/revel
RUN go get -u github.com/go-sql-driver/mysql
RUN go get -u gopkg.in/olivere/elastic.v7
RUN go get -u gopkg.in/gorp.v1
RUN go get -u github.com/motemen/go-loghttp
RUN go get -u github.com/motemen/go-nuts/roundtime


# Define mountable directories.
VOLUME ${HOME}/go

WORKDIR ${HOME}/go


# clear apt cache and remove unnecessary packages
RUN apt-get autoclean && apt-get -y autoremove


EXPOSE 9000


WORKDIR ${HOME}/go/src/instapp

CMD ["sh","/root/go/src/instapp/run-the-world.sh"]
