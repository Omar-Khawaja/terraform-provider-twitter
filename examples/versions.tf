terraform {
  required_providers {
    twitter = {
      # this source assumes you have installed the 
      # provider locally as specified in the README
      # Change this if you are pulling from
      # the Terraform registry
      source  = "example.com/example/twitter"
      version = "<version of provider (ex: v0.1.0)>"
    }
  }
  required_version = ">= 0.13"
}
