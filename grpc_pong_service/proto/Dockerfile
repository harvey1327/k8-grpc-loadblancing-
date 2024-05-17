FROM golang:1.22-alpine3.18

ENV VERSION "26.1"
RUN apk update && apk --no-cache add curl protoc protobuf git
RUN curl -LO https://github.com/protocolbuffers/protobuf/releases/download/v$VERSION/protoc-$VERSION-linux-x86_64.zip && unzip protoc-$VERSION-linux-x86_64.zip -d /usr/bin && rm protoc-$VERSION-linux-x86_64.zip
RUN go install google.golang.org/protobuf/cmd/protoc-gen-go@latest
RUN go install google.golang.org/grpc/cmd/protoc-gen-go-grpc@latest

CMD ["protoc", "-I=/go/input", "--go_out=/go/output", "--go_opt=module=/proto/generated", "--go-grpc_out=/go/output", "--go-grpc_opt=module=/proto/generated", "/go/input/message.proto"]

#docker build . -t test:v1
#docker run -it --rm -v ./message.proto:/go/input/message.proto -v ${PWD}/generated:/go/output test:v1