FROM registry.access.redhat.com/ubi9/go-toolset:1.21.10-1

WORKDIR /opt/app-root/src

COPY . ./
RUN <<EOF
set -x
go mod download -x
GOOS=linux GOARCH=amd64 go build -v -o /opt/app-root/bin/ghc
chgrp -R 0 /opt/app-root/bin && chmod -R g=u /opt/app-root/bin
EOF

WORKDIR /opt/app-root/bin

EXPOSE 8888

CMD ["ghc"]

# CMD ["sh", "-c", "/opt/app-root/bin/ghc && sleep infinity || { echo 'Application failed to start'; sleep infinity; }"]
