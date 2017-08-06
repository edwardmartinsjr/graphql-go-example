FROM golang:1.7.4
COPY . /go/src/github.com/graphql-go-example/
RUN go get github.com/tools/godep
RUN go get github.com/graphql-go/graphql
RUN go get github.com/graphql-go/graphql-go-handler
RUN go get github.com/lib/pq
RUN go get github.com/go-sql-driver/mysql
RUN export PATH=$PATH:$GOPATH/bin
WORKDIR /go/src/github.com/graphql-go-example/
RUN godep save
RUN godep go install
EXPOSE 8080
CMD ["go", "run", "main.go"]

