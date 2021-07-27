#!/bin/sh
ARCH=$1
ISECNET_GO_VERSION=$2

if [ "$ARCH" = "armhf" ]; then
  ARCH="arm"
fi
if [ "$ARCH" = "aarch64" ]; then
  ARCH="arm64"
fi

echo "Arch: $ARCH, Isecnet-go Version: $ISECNET_GO_VERSION"
curl -L https://github.com/thspinto/isecnet-go/releases/download/${ISECNET_GO_VERSION}/isecnet-go-${ISECNET_GO_VERSION}-linux-${ARCH}.tar.gz | tar -xz
chmod +x ./isecnet-go
