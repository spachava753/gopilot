data "aws_availability_zones" "available" {}

locals {
  vpc_name = "gopilot-vpc"
  vpc_cidr = "10.0.0.0/16"
  azs      = slice(data.aws_availability_zones.available.names, 0, 3)
  tags     = {
    "ManagedBy" : "gopilot",
    "Terraform" : "true"
  }
}

terraform {
  required_providers {
    aws = {
      source  = "hashicorp/aws"
      version = "5.7.0"
    }
  }
}

# Configure Terraform AWS provider
# https://registry.terraform.io/providers/hashicorp/aws/latest/docs#authentication-and-configuration
provider "aws" {
  region = "us-east-1"
}

module "vpc" {
  source  = "terraform-aws-modules/vpc/aws"
  version = "5.0.0"

  name = local.vpc_name
  cidr = local.vpc_cidr

  azs            = local.azs
  public_subnets = [for k, v in local.azs : cidrsubnet(local.vpc_cidr, 8, k)]
  tags           = local.tags
}

module "ssh_security_group" {
  source              = "terraform-aws-modules/security-group/aws//modules/ssh"
  version             = "~> 5.0"
  name                = "gopilot-ssh-sg"
  vpc_id              = module.vpc.vpc_id
  ingress_cidr_blocks = ["0.0.0.0/0"]
  tags                = local.tags
}

module "key_pair" {
  source = "terraform-aws-modules/key-pair/aws"

  key_name           = "gopilot-key"
  create_private_key = true

  tags = local.tags
}

module "ec2_instance" {
  source  = "terraform-aws-modules/ec2-instance/aws"
  version = "5.2.1"

  name = "gopilot-control-plane"

  instance_type          = "t2.micro"
  monitoring             = true
  vpc_security_group_ids = [module.vpc.default_security_group_id]
  subnet_id              = module.vpc.public_subnets[0]

  key_name = module.key_pair.key_pair_name

  tags = local.tags
}