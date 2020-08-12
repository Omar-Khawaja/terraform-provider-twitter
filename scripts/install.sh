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

# This remove the leading v (v0.1.0 becomes 0.1.0)
# This is needed for the local directory layout
PLUGIN_VERSION="$(echo $(git describe --tags --abbrev=0) | cut -d'v' -f2)"

TF_PLUGIN_DIR="$HOME/.terraform.d/plugins/example.com/example/twitter/$PLUGIN_VERSION/$OS_ARCH"

if [[ ! -d "$TF_PLUGIN_DIR" ]]; then
    echo "The Terraform plugin directory does not exist"
    echo "creating it now..."
    mkdir -p "$TF_PLUGIN_DIR"
    echo "$TF_PLUGIN_DIR" has been created
fi

echo "Installing plugin..."

go build -o terraform-provider-twitter_$(git describe --tags --abbrev=0) ..
mv terraform-provider-twitter_$(git describe --tags --abbrev=0) $TF_PLUGIN_DIR

echo "Plugin has been installed."
