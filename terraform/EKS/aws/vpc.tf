resource "aws_vpc" "main_vpc" {
  cidr_block           = var.vpc_cidr
  enable_dns_support   = true
  enable_dns_hostnames = true
  tags = {
    Name = "oncloudev"
  }
}

resource "aws_internet_gateway" "main_igw" {
  vpc_id = aws_vpc.main_vpc.id
  tags = {
    Name = "oncloudev"
  }
}

resource "aws_subnet" "main_subnet" {
  vpc_id                  = aws_vpc.main_vpc.id
  cidr_block              = var.public_subnet_cidr
  map_public_ip_on_launch = true
  availability_zone       = "${var.region}a"
  tags = {
    Name        = "oncloudev"
    description = "public subnet"
  }
}

resource "aws_subnet" "main_private_subnet" {
  vpc_id                  = aws_vpc.main_vpc.id
  cidr_block              = var.private_subnet_cidr
  map_public_ip_on_launch = false
  availability_zone       = "${var.region}b"
  tags = {
    Name        = "oncloudev"
    description = "private subnet"
  }
}

resource "aws_route_table" "main_public_routetable" {
  vpc_id = aws_vpc.main_vpc.id

  route {
    cidr_block = "0.0.0.0/0"
    gateway_id = aws_internet_gateway.main_igw.id
  }

  tags = {
    Name        = "oncloudev"
    description = "public-route-table"
  }
}

resource "aws_route_table_association" "public_routetable_association" {
  subnet_id      = aws_subnet.main_subnet.id
  route_table_id = aws_route_table.main_public_routetable.id
}

resource "aws_eip" "nat" {
  domain = "vpc"
}

resource "aws_nat_gateway" "main_private_nat" {
  allocation_id = aws_eip.nat.id
  subnet_id     = aws_subnet.main_subnet.id

  tags = {
    Name        = "oncloudev"
    description = "NAT for private subnet"
  }

  depends_on = [aws_internet_gateway.main_igw]
}

resource "aws_route_table" "main_private_routetable" {
  vpc_id = aws_vpc.main_vpc.id

  route {
    cidr_block     = "0.0.0.0/0"
    nat_gateway_id = aws_nat_gateway.main_private_nat.id
  }

  tags = {
    Name        = "oncloudev"
    description = "private-route-table"
  }
}

resource "aws_route_table_association" "private_assoc" {
  subnet_id      = aws_subnet.main_private_subnet.id
  route_table_id = aws_route_table.main_private_routetable.id
}
