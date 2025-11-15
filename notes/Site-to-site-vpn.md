# Site-to-Site VPN Connection: AWS to Azure with Dynamic Routing

## Prerequisites
- **AWS Account** with administrative privileges.
- **Azure Account** with an active subscription.

## Phase 1: Configuring Azure

### Step 1: Create Resource Group
1. Navigate to the Azure portal.
2. Click on **Resource groups**.
3. Click **+ Create**.
4. Enter the **Resource group name** and select the **Region**.
5. Click **Review + create** and then **Create**.

### Step 2: Create Virtual Network and Subnets
1. Navigate to **Virtual networks**.
2. Click **+ Create**.
3. Enter the **Name**, **Address space** (e.g., 10.1.0.0/16), and **Subnets** (e.g., 10.1.1.0/24).
4. Click **Review + create** and then **Create**.

### Step 3: Create Route Table
1. Navigate to **Route tables**.
2. Click **+ Create**.
3. Enter the **Name** and **Region**.
4. Click **Review + create** and then **Create**.

### Step 4: Create Virtual Network Gateway
1. Navigate to **Virtual network gateways**.
2. Click **+ Create**.
3. Enter the **Name**, **Gateway type** (VPN), **VPN type** (Route-based), and **SKU**.
4. Select the **Virtual network** and **Public IP address**.
5. Click **Review + create** and then **Create**.

### Step 5: Create Local Network Gateway
1. Navigate to **Local network gateways**.
2. Click **+ Create**.
3. Enter the **Name**, **IP address** (AWS VPN Gateway IP), and **Address space** (AWS VPC CIDR, e.g., 10.0.0.0/16).
4. Click **Review + create** and then **Create**.

### Step 6: Create Connection
1. Navigate to **Virtual network gateways**.
2. Select the created **Virtual network gateway**.
3. Click **Connections** and then **+ Add**.
4. Enter the **Name**, select **Site-to-site (IPsec)**, and choose the **Local network gateway**.
5. Enter the **Shared key** (same as AWS VPN connection).
6. Click **OK**.

## Phase 2: Configuring AWS

### Step 1: Create VPC
1. Navigate to the AWS Management Console.
2. Click on **VPC**.
3. Click **Create VPC**.
4. Enter the **Name**, **IPv4 CIDR block** (e.g., 10.0.0.0/16), and **Tenancy**.
5. Click **Create VPC**.

### Step 2: Create Subnet
1. Navigate to **Subnets**.
2. Click **Create subnet**.
3. Select the **VPC**.
4. Enter the **Subnet name** and **IPv4 CIDR block** (e.g., 10.0.1.0/24).
5. Click **Create subnet**.

### Step 3: Create Route Table
1. Navigate to **Route tables**.
2. Click **Create route table**.
3. Select the **VPC**.
4. Enter the **Name**.
5. Click **Create route table**.

### Step 4: Create Internet Gateway
1. Navigate to **Internet gateways**.
2. Click **Create internet gateway**.
3. Enter the **Name**.
4. Click **Create internet gateway**.
5. Attach the internet gateway to the **VPC**.

### Step 5: Create Customer Gateway
1. Navigate to **Customer gateways**.
2. Click **Create customer gateway**.
3. Enter the **Name**, **IP Address** (Azure VPN Gateway IP), and **Routing**.
4. Click **Create customer gateway**.

### Step 6: Create Virtual Private Gateway
1. Navigate to **Virtual private gateways**.
2. Click **Create virtual private gateway**.
3. Enter the **Name**.
4. Click **Create virtual private gateway**.
5. Attach the virtual private gateway to the **VPC**.

### Step 7: Create Site-to-Site VPN Connection
1. Navigate to **VPN connections**.
2. Click **Create VPN connection**.
3. Enter the **Name**, **Target gateway** (Virtual private gateway), **Customer gateway**, and **Routing options** (Dynamic).
4. Enter the **Shared key** (same as Azure connection).
5. Click **Create VPN connection**.

### Step 8: Update Route Table
1. Navigate to **Route tables**.
2. Select the **Route table**.
3. Click **Edit routes**.
4. Add a route to the **Azure network** (e.g., 10.1.0.0/16) via the **Virtual private gateway**.

## Phase 3: Connecting Azure and AWS

### Step 1: Edit Route Table in Azure
1. Navigate to **Route tables**.
2. Select the **Route table**.
3. Click **Routes**.
4. Add a route to the **AWS network** (e.g., 10.0.0.0/16) via the **Virtual network gateway**.

### Step 2: Create Network Security Group in Azure
1. Navigate to **Network security groups**.
2. Click **+ Create**.
3. Enter the **Name** and **Region**.
4. Click **Review + create** and then **Create**.

### Step 3: Create Security Group in AWS
1. Navigate to **Security groups**.
2. Click **Create security group**.
3. Enter the **Name** and **Description**.
4. Click **Create security group**.

### Step 4: Create Virtual Machine in Azure
1. Navigate to **Virtual machines**.
2. Click **+ Create**.
3. Enter the **Name**, **Region**, **Image**, and **Size**.
4. Click **Review + create** and then **Create**.

### Step 5: Create EC2 Instance in AWS
1. Navigate to **EC2**.
2. Click **Launch instance**.
3. Enter the **Name**, **AMI**, **Instance type**, and **Key pair**.
4. Click **Launch instance**.

## Conclusion
By following these steps, you will have successfully created a site-to-site VPN connection between AWS and Azure with dynamic routing. This setup ensures secure and reliable communication between your resources hosted on both platforms.

## Diagram
Below is a diagram illustrating the setup:

!AWS to Azure VPN Diagram

- **AWS VPC**: 10.0.0.0/16
  - **Subnet**: 10.0.1.0/24
  - **Virtual Private Gateway**: Attached to VPC
  - **Customer Gateway**: Azure VPN Gateway IP
  - **VPN Connection**: Site-to-site with Azure

- **Azure VNet**: 10.1.0.0/16
  - **Subnet**: 10.1.1.0/24
  - **Virtual Network Gateway**: Route-based, Public IP
  - **Local Network Gateway**: AWS VPN Gateway IP
  - **Connection**: Site-to-site with AWS

### IP Address Configuration
- **AWS VPN Gateway IP**: Public IP assigned by AWS
- **Azure VPN Gateway IP**: Public IP assigned by Azure
- **BGP IPs**:
  - **AWS Tunnel 1**: 169.254.21.1 (AWS), 169.254.21.2 (Azure)
  - **AWS Tunnel 2**: 169.254.22.1 (AWS), 169.254.22.2 (Azure)

