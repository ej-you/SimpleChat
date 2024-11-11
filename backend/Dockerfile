# -----
# BUILD
# -----

FROM golang:1.22.4-alpine3.20 AS build

# install make util for Makefile
RUN apk add --no-cache make

# set up workdir
RUN cd /go/src
RUN mkdir -p ./SimpleChat/backend
WORKDIR /go/src/SimpleChat/backend

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

# set up DB path
ENV PATH_DB="/root/db/db.sqlite3"


WORKDIR /root
# copy compiled file and Makefile to run app
COPY --from=build /go/src/SimpleChat/backend/Makefile .
COPY --from=build /go/src/SimpleChat/backend/go_app .

# run app
CMD ["make", "prod"]
