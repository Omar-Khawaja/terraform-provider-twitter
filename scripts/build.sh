#!/usr/bin/env bash

go build -o ../terraform-provider-twitter_$(git describe --tags --abbrev=0) ..
