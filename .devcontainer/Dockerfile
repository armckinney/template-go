FROM armck/ubuntu:24.04

LABEL   deployable="false" \
    description="Development container for template-go; includes Go toolchain and common Go tools."

ARG WORKSPACE=/workspaces/template-go
ENV WORKSPACE=${WORKSPACE}

ARG GO_VERSION=1.25.0

ENV GOROOT=/usr/local/go
ENV GOPATH=/go
ENV PATH=${GOROOT}/bin:${GOPATH}/bin:${PATH}

# make non-interactive for build
ENV DEBIAN_FRONTEND=noninteractive

RUN mkdir /tmp/container \
    && cd /tmp/container \
    # installs
    && apt update && apt install -y --no-install-recommends \
    ca-certificates \
    curl \
    wget \
    git \
    build-essential \
    pkg-config \
    unzip \
    postgresql-client \
    # install go
    && rm -rf /usr/local/go \
    && wget -q "https://go.dev/dl/go${GO_VERSION}.linux-$(dpkg --print-architecture).tar.gz" -O go.tar.gz \
    && tar -C /usr/local -xzf go.tar.gz \
    && rm go.tar.gz \
    # cleanup
    && apt autoremove -y \
    && apt clean -y \
    && rm -rf /var/lib/apt/lists/* \
    && rm -rf /tmp/container

WORKDIR ${WORKSPACE}

# restore dialog for interactive use
ENV DEBIAN_FRONTEND=dialog
