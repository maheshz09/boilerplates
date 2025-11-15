variable "region" {
  description = "default region for AWS"
  type        = string
  default     = "ap-south-1"
}

variable "vpc_cidr" {
  description = "CIDR block for the VPC"
  type        = string
  default     = "10.0.0.0/16"
}

variable "public_subnet_cidr" {
  description = "public subnet inside main-vpc"
  type        = string
  default     = "10.0.1.0/24"
}

variable "private_subnet_cidr" {
  description = "private subnet inside main-vpc"
  type        = string
  default     = "10.0.2.0/24"
}
