terraform {
  required_providers {
    aws = {
      source  = "hashicorp/aws"
      version = "3.26.0"
    }
    random = {
      source  = "hashicorp/random"
      version = "3.0.1"
    }
  }
  required_version = ">= 1.1.0"

  cloud {
    organization = "arunvel"

    workspaces {
      name = "gh-actions-demo"
    }
  }
}

provider "aws" {
  region = "ap-south-1"
}


resource "aws_instance" "test" {
  ami           = "ami-06984ea821ac0a879" # Ubuntu 18.04 LTS
  instance_type = "t2.micro"

  tags = {
    Name = "test-instance"
    env = "dev"
  }
}