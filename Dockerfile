FROM golang:1.19

WORKDIR /app
ARG GETH_VERSION="1.10.21-67109427"
ARG SOLIDITY_VERSION="0.8.16"
LABEL SOLIDITY_VERSION \
      GETH_VERSION \
      "Golang builder dependencies" \
      "Maintainer Bitex"

COPY . .
RUN go mod download && go build codegen/events_handler.go && mv events_handler eh-gen
RUN wget -q "https://github.com/ethereum/solidity/releases/download/v$SOLIDITY_VERSION/solc-static-linux" \
  && chmod +x solc-static-linux \
  && mv solc-static-linux /usr/local/bin/solc \
  && wget -q "https://gethstore.blob.core.windows.net/builds/geth-alltools-linux-amd64-$GETH_VERSION.tar.gz" \
  && tar xvzf "geth-alltools-linux-amd64-$GETH_VERSION.tar.gz" \
  && mv "geth-alltools-linux-amd64-$GETH_VERSION/abigen" /usr/local/bin/abigen \
  && rm -rf "geth-alltools-linux-amd64-$GETH_VERSION" \
  && rm "geth-alltools-linux-amd64-$GETH_VERSION.tar.gz" \
  && mv eh-gen /usr/local/bin/eh-gen
