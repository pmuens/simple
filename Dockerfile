FROM golang:latest

RUN go get github.com/cespare/reflex

WORKDIR /go/src

CMD reflex --regex '\.go$' --decoration none go fmt {} &> /dev/null & disown
#RUN echo "This is a test." | wc -

### meta
# Commands on how to build and run the container

# docker build -t docker .
# docker run -it -v $PWD:/go/src go-tinkering bash


# resources which will help to run refelx on start
#https://github.com/circleci/frontend/blob/master/Dockerfile
#https://github.com/circleci/frontend/blob/master/docker-entrypoint.sh
