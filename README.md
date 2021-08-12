## About

Demo Terraform provider for Twitter. This code assumes you are using Terraform 0.13.0+

## Usage

This provider is available on the [Terraform Registry](https://registry.terraform.io/providers/Omar-Khawaja/twitter/latest). See the [examples](./examples/registry) directory for a sample Terraform config to publish your first tweet!

## Note

This provider requires the following environment variables to be set:

- TWITTER_API_KEY
- TWITTER_API_SECRET_KEY
- TWITTER_ACCESS_TOKEN
- TWITTER_ACCESS_TOKEN_SECRET

## Local Development

- This section is only relevant if you are trying to modify this provider and
  run it locally

### Build and Installation

Run `make install` to build the plugin and install it in the necessary directory for Terraform to detect it.

If you want to just build the plugin and move it to a different directory
yourself, run `make build` to build the plugin.

See the [local_development](./examples/local_development) directory for a sample Terraform config that uses a local copy of the provider.