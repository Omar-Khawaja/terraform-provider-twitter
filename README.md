## About

Demo Terraform provider for Twitter. This code assumes you are using Terraform
0.13.0+

## Build and Installation

Run `make install` to build the plugin and install it in the necessary directory
for Terraform to detect it.

If you want to just build the plugin and move it to a different directory
yourself, run `make build` to build the plugin.

## Usage

You will need to run `terraform init` before using the provider. See the
[examples](./examples) directory for sample Terraform config.

## Note

This provider requires the following environment variables to be set:

- TWITTER_API_KEY
- TWITTER_API_SECRET_KEY
- TWITTER_ACCESS_TOKEN
- TWITTER_ACCESS_TOKEN_SECRET