terraform {
  backend "s3" {
    encrypt        = true
    bucket         = "my-tf-test-bucket-arunvel1988"
    dynamodb_table = "terraform-state-lock-dynamo"
    key            = "terraform.tfstate"
    region         = "ap-south-1"
  }
}