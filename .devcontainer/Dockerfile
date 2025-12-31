FROM armck/ubuntu:24.04

LABEL   deployable="false" \
    description="Development container for template-go; includes Go toolchain and common Go tools."

ARG GO_VERSION=1.25.0

ENV GOPATH=/go
ENV PATH=/usr/local/go/bin:${GOPATH}/bin:${PATH}

ARG WORKSPACE=/workspaces/template-go
ENV WORKSPACE=${WORKSPACE}

# make non-interactive for build
ENV DEBIAN_FRONTEND=noninteractive

RUN mkdir /tmp/container \
    && cd /tmp/container \
    # indirect installs
    && apt update && apt install -y --no-install-recommends \
    wget \
    build-essential \
    # installs
    && apt install -y \
    postgresql-client \
    # install go
    && curl -L "https://go.dev/dl/go${GO_VERSION}.linux-$(dpkg --print-architecture).tar.gz" | tar -C /usr/local -xz \
    # cleanup
    && apt autoremove -y \
    && apt clean -y \
    && rm -rf /var/lib/apt/lists/* \
    && rm -rf /tmp/container

WORKDIR ${WORKSPACE}

# install go dependencies
COPY go.mod go.sum ./
RUN go mod download

# restore dialog for interactive use
ENV DEBIAN_FRONTEND=dialog
