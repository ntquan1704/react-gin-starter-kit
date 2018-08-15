FROM golang:1.10

ENV GOBIN /go/bin
ENV DEBIAN_FRONTEND noninteractive

RUN apt-get update -yq \
    && apt-get install -y curl gcc g++ make apt-utils

RUN apt-get update -yq \
    && curl -sL https://deb.nodesource.com/setup_10.x | bash \
    && apt-get install nodejs -yq

RUN curl https://raw.githubusercontent.com/golang/dep/master/install.sh | sh

RUN npm install -g yarn

COPY . /go/src/app

WORKDIR /go/src/app

RUN yarn install \
    && go get github.com/olebedev/on \
	&& go get -u github.com/jteeuwen/go-bindata/... \
    && dep ensure 

RUN cd server/ \
    && go build -ldflags '-w -s' -a -installsuffix cgo -o $@ /go/src/app/main
RUN ${GOBIN}/go-bindata -pkg=main -prefix=server/data -o=server/bindata.go server/data/...

CMD [ "./main", "run" ]