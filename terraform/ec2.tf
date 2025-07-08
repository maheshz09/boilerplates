terraform {
  required_providers {
    aws = {
      source = "hashicorp/aws"
      version = "6.0.0-beta3"
    }
  }
}

provider "aws" {
  region = ap-south-1
}

# Reference your existing security group by ID
variable "security_group_id" {
  description = "The ID of the existing security group"
  type        = string
}

# Create a new key pair
resource "aws_key_pair" "new_key" {
  key_name   = "my-new-key"
  public_key = file("~/.ssh/id_rsa.pub") // Path to your public key
}

# Create the EC2 instance
resource "aws_instance" "example" {
  ami                    = "ami-0c02fb55956c7d316" // Use a valid AMI for your region
  instance_type          = "t2.micro"
  key_name               = aws_key_pair.new_key.key_name
  vpc_security_group_ids = [var.security_group_id]

  root_block_device {
    volume_size = 25
    volume_type = "gp2"
  }

  tags = {
    Name = "ExampleInstance"
  }
}