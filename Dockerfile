#pull the latest image of golang
FROM golang:latest

#creating the build directory
RUN mkdir /build
WORKDIR /build



COPY go.mod ./
COPY go.sum ./

COPY .. ./

#so we can pull any version of package from github
RUN export GO111MODULE=on

#dowloading all the dependencies
RUN go mod download

# go build jo hai wo current jaha hai waha jake build kardega
RUN go build -o main


#same as jis port pe apis run ho rahi hai
EXPOSE 5775

#starting point of the program
CMD [ "./main" ]

