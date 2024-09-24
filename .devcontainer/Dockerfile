ARG GO_VERSION=1.23.1
ARG ALPINE_VERSION=3.20

FROM golang:${GO_VERSION}-alpine${ALPINE_VERSION}
ARG USERNAME=vscode
ARG USER_UID=1000
ARG USER_GID=1000
ARG GOPLS_VERSION=latest

RUN adduser ${USERNAME} -s /bin/sh -D -u ${USER_UID} ${USER_GID} && \
    mkdir -p /etc/sudoers.d && \
    echo ${USERNAME} ALL=\(root\) NOPASSWD:ALL > /etc/sudoers.d/${USERNAME} && \
    chmod 0440 /etc/sudoers.d/${USERNAME}

RUN apk add -q --update --progress --no-cache \
    git sudo openssh-client curl make gpg jq curl \
    unzip git httpie fish starship

RUN go install golang.org/x/tools/gopls@latest && \
    go install github.com/cweill/gotests/gotests@v1.6.0 && \
    go install github.com/fatih/gomodifytags@v1.16.0 && \
    go install github.com/josharian/impl@v1.1.0 && \
    go install github.com/haya14busa/goplay/cmd/goplay@v1.0.0 && \
    go install github.com/go-delve/delve/cmd/dlv@latest && \
    curl -sSfL https://raw.githubusercontent.com/golangci/golangci-lint/master/install.sh | sh -s -- -b $(go env GOPATH)/bin v1.61.0

USER ${USERNAME}

ENV ENV="/home/${USERNAME}/.ashrc" \
    EDITOR=vi \
    LANG=en_US.UTF-8 \
    PATH=/home/vscode/.local/bin:$PATH

RUN echo "exec `which fish`" > "/home/${USERNAME}/.ashrc" && \
    mkdir -p "/home/${USERNAME}/.config/fish/" && \
    echo "starship init fish | source" >> "/home/${USERNAME}/.config/fish/config.fish"

USER root
RUN chown -R $USER_UID:$USER_GID /go/pkg