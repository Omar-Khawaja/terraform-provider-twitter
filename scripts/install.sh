#!/usr/bin/env bash

case "$OSTYPE" in
    darwin*)
        OS="darwin"
        ;;
    linux*)
        OS="linux"
        ;;
    *)
        echo "script only works on MacOS or Linux"
        exit 1
esac

case "$(uname -m)" in
    x86_64)
        ARCH="amd64"
        ;;
    *)
        echo "script only works with 64 bit architecture"
        exit 1
        ;;
esac

OS_ARCH="$OS"_"$ARCH"

TF_PLUGIN_DIR="$HOME/.terraform.d/plugins/example.com/local/twitter/${VERSION}/$OS_ARCH"

if [[ ! -d "$TF_PLUGIN_DIR" ]]; then
    echo "The Terraform plugin directory does not exist"
    echo "creating it now..."
    mkdir -p "$TF_PLUGIN_DIR"
    echo "$TF_PLUGIN_DIR" has been created
fi

echo "Installing plugin..."

go build -o terraform-provider-twitter_"${VERSION}" ..
mv terraform-provider-twitter_"${VERSION}" "$TF_PLUGIN_DIR"

echo "Plugin has been installed."
