provider "aws" {
  region = "ap-south-1"
}


resource "aws_instance" "test" {
  ami           = "ami-06984ea821ac0a879" # Ubuntu 18.04 LTS
  instance_type = "t2.micro"

  tags = {
    Name = "test-instance"
    env = "dev2"
    test = "s3"
  }
}