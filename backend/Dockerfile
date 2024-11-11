# -----
# BUILD
# -----

FROM golang:1.22.4-alpine3.20 AS build

# install make util for Makefile
RUN apk add --no-cache make

# set up workdir
RUN cd /go/src
RUN mkdir -p ./github.com/echo-tokyo/TaskManager/backend_go
WORKDIR /go/src/github.com/echo-tokyo/TaskManager/backend_go

# install dependences
COPY ./go.mod .
COPY ./go.sum .
RUN go mod tidy
RUN go mod download

# copy project files to container
COPY . .

# compile app
RUN make compile

# ---
# RUN
# ---

FROM alpine:3.20 AS run

# install make util for Makefile
RUN apk add --no-cache make

# make dir for logs
RUN mkdir /logs

WORKDIR /root
# copy compiled file and Makefile to run app
COPY --from=build /go/src/github.com/echo-tokyo/TaskManager/backend_go/Makefile .
COPY --from=build /go/src/github.com/echo-tokyo/TaskManager/backend_go/go_app .

# run app
CMD ["make", "prod"]
