FROM golang:1.12
LABEL version="v1.0.14"

# Copy sources
ENV PROJECT_DIR /go/src/github.com/Blockchainpartner/hagentrust
RUN mkdir -p ${PROJECT_DIR}
COPY . ${PROJECT_DIR}
WORKDIR ${PROJECT_DIR}