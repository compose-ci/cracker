FROM --platform=$BUILDPLATFORM ubuntu:22.04 AS builder

LABEL maintainer="@k33g_org"

ARG TARGETOS
ARG TARGETARCH

ARG GO_VERSION=${GO_VERSION}
ARG TINYGO_VERSION=${TINYGO_VERSION}

ARG PLUGIN_PATH=${PLUGIN_PATH}
ARG RUNNER_PATH=${RUNNER_PATH}


ARG DEBIAN_FRONTEND=noninteractive

ENV LANG=en_US.UTF-8
ENV LANGUAGE=en_US.UTF-8
ENV LC_COLLATE=C
ENV LC_CTYPE=en_US.UTF-8

# ------------------------------------
# Install Tools
# ------------------------------------
RUN <<EOF
apt-get update 
apt-get install -y wget 

apt-get clean autoclean
apt-get autoremove --yes
rm -rf /var/lib/{apt,dpkg,cache,log}/
EOF

# ------------------------------------
# Install Go
# ------------------------------------
RUN <<EOF
wget https://golang.org/dl/go${GO_VERSION}.linux-${TARGETARCH}.tar.gz
tar -xvf go${GO_VERSION}.linux-${TARGETARCH}.tar.gz
mv go /usr/local
rm go${GO_VERSION}.linux-${TARGETARCH}.tar.gz
EOF

# ------------------------------------
# Set Environment Variables for Go
# ------------------------------------
ENV PATH="/usr/local/go/bin:${PATH}"
#ENV GOPATH="/home/${USER_NAME}/go"
#ENV GOROOT="/usr/local/go"

# ------------------------------------
# Install TinyGo
# ------------------------------------
RUN <<EOF
wget https://github.com/tinygo-org/tinygo/releases/download/v${TINYGO_VERSION}/tinygo_${TINYGO_VERSION}_${TARGETARCH}.deb
dpkg -i tinygo_${TINYGO_VERSION}_${TARGETARCH}.deb
rm tinygo_${TINYGO_VERSION}_${TARGETARCH}.deb
EOF

# ------------------------------------
# Build Wasm Function Server
# ------------------------------------
WORKDIR /app/
COPY ${RUNNER_PATH}/main.go /app/tmp/main.go
COPY ${RUNNER_PATH}/go.mod /app/tmp/go.mod

RUN <<EOF
cd /app/tmp
go mod tidy
go mod download
CGO_ENABLED=0 GOOS=${TARGETOS} GOARCH=${TARGETARCH} go build -ldflags="-s -w" -o cracker-runner main.go
EOF

# ------------------------------------
# Build The plugin
# ------------------------------------
COPY ${PLUGIN_PATH}/main.go /app/tmp/main.go
COPY ${PLUGIN_PATH}/go.mod /app/tmp/go.mod

RUN <<EOF
cd /app/tmp
go mod tidy
go mod download
tinygo build -scheduler=none --no-debug \
  -o plugin.wasm \
  -target wasi main.go
EOF

FROM --platform=$TARGETPLATFORM scratch
#FROM --platform=$TARGETPLATFORM gcr.io/distroless/static-debian12
#FROM --platform=$TARGETPLATFORM ubuntu:22.04

#WORKDIR /app/
COPY --from=builder /app/tmp/cracker-runner /
COPY --from=builder /app/tmp/plugin.wasm /

EXPOSE 8080
#CMD ["/extism-runner", "/plugin.wasm", "say_hello" , "8080"]






