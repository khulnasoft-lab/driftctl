provider "aws" {
    region = "us-east-1"
}

terraform {
    required_providers {
        aws = "3.76.1"
    }
}

resource "aws_iam_group" "my_developers" {
    name = "developers"
    path = "/users/"
}
