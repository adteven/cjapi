FROM golang:1.18.4

RUN go env -w GO111MODULE=on
RUN go env -w GOPROXY=https://goproxy.cn,direct
# Install beego and the bee dev tool*
RUN go get github.com/beego/beego/v2@latest && go install github.com/beego/bee/v2@master

# Expose the application on port 8080
EXPOSE 8080

# Set the entry point of the container to the bee command that runs the
# application and watches for changes
CMD ["bee", "run"]