Notes
---
## class: `AWS Solutions Architect Associate`
### instructor: `Stephane Maarek`
AWS stands for Amazon Web Services and is a cloud provider.

AWS provides services. 

> A spaghetti bowl.

`N.B. it is helpful to get a free, personal account; keep in mind they have to check you have a valid credit card on file.`

#### History of AWS
2002 launched to serve their own company's needs first, then they extended their business services little bit by little bit until they are now handling $90B per year...

#### Cloud Use Cases
Scalable, Global Applications delivered by AWS Regions via their availability zones.

Compliance, Proximity and Available services should be taking into account when considering using AWS. 

Availability Zones allow for redundancy of data centers and provide multiple points of contact for low latency bandwidth to services.

See more [AWS's regional offerings](aws.amazon.com/about-aws/global-infrastructure/regional-product-services).

#### AWS Dashboard
- _allows user to change availability zone_
- _allows user to look through all services_
- _search bar allows users to search aws_

#### IAM
Identity and Access Management ; Permissions management
Permissions:
_there are over +1000 permission_
> N.B. users should note that AWS uses standard UNIX conventions for the wild card operator: `*` , ie "iam:Get*" would be all of the Get methods

- IAM POLICY is in JSON ; defines user permissions 
    - Version
    - ID
    - Statements
        - sid ; statement-id
        - effect
        - principal ; account/user/role
        - actions
        - resources
        - condition
> least privilege principle implies that you only give a user the least amount of privilege necessary for them to accomplish their work

Root account <- your admin account 
- It is not best practice to work from the root account

Group <- the roles in your organization ; inherited
> It is a good idea to start by creating an `admin` group

- Permissions are defined with the IAM policy, or at creation
- Users can have multiple groups


User <- people in the organization
- Permissions can be set in the IAM policy, or inline
- Management Console
    - Access Keys (id key, secret key)
- CLI (Command Line Interface) ; for scripting 
    - Downloads for WINDOWS , MAC , LINUX
    - `aws configure`
- SDK (Software Development Kit) ; language specific

Password Policies
- ENFORCE STRONG 
- NO RE-USED 
- Expiring passwords
- Multi-factor Authentication ; MFA

AWS has a cloud shell ; super handy

IAM ROLES for SERVICES
- some services will need to have permissions to perform , some examples :
  - ec2
  - lambda
  - cloudformation

IAM Security Tools
  - Credentials Report { user login table }
  - Access Advisor { granular user permissions }
  - I guess it would make sense if this was in real time

Billing and Cost Management
- Bills { charges by services }
- Budgets { alerts for thresholds }

EC2 { Elastic Cloud Compute }
- Hardware Configuration { cpu / mem }
  - General Purpose
  - Compute Optimized
  - Memory Optimized
  - Storage Optimized
- OS { Windows, Linux, macOS}
- Network
- Firewall
- Bootstrap
- Terminate (Delete Instance)

Security Groups
- Inbound rules { DEFAULT: DENIED }
- Outbound rules { DEFAULT: ALLOWED }

SSH Secure Shell
- used to log into the cloud
- SSH is most common
- Instance Connect is AWS in-browser option
- Be sure to have the correct inbound policies

Security Roles
- Set specific policies, resource "roles" 
- Principle of least access extends beyond users, it extends to resources as well

Resource Types
- You can choose from a variety of instance billing/usage types
- Key: more control means more commitment, otherwise more cost

Spot Requests
- Pricing Based Requests (budget_price > current_price)
- Time Block Requests Pricing 
- Spot Fleets

Private and Public IP addresses
- Public Networks (like the internet) are configured to be accessed by the Internet (everyone!). 
- Private Networks are configured to be access in the same network 
- Can create elastic ip addresses (configuring a static ip)
> Best to use a loadbalancer for handling all calls into our domain

Placement Groups
- Cluster concentrates instances into the same availability zone
- Spread spread against availability zones
- Partition spread across racks

EIN Network Interfaces
- virtual network card
- Availability Zone Bound
- Security Groups are used with this resource
- Failover for EC2 instances (recycling)

EC2 Hibernation
- Ram dumped into Elastic Block Store (EBS) - ENCRYPTED 

EBS - Elastic Block Store 
> Network USB sticks
- Availability Zone Bound
- Network Volumes
- Provisions
- Security Groups
- Delete on terminate  

EBS Snapshots
- Tiered 
  - Snapshot
  - Archived Snapshot
    - 24-72 hour to uptime
- Recycle Bin Snapshots ( retention queue ; prevent permanent deletion )
- Fast Snapshot Restore ($$$)

AMI Amazon Machine Image
- Public AMIs
- Make your own
- AWS Marketplace
- Essentially it is a pre-configured machine image with all the software pre-built into the instance launch

EC2 Instance Store
- you can’t store data on an EC2
- but you could temporarily 
- hight performance compared to EBS

EBS Elastic Block Storage
- Cost and performance trade offs
	- gp general purpose
	- Provisioned IOPS input/output operations per second.
	- hd Hard Disk
- EMBS Multi-Attach 
	- io1/io2 family only
	- availability zone bound
- EBS Encryption
	- encryptions is KMS (AES256)
	- amazon handles all encryption/decryption

EFS Elastic File System
- scalable network drive
- pay per use
- Performance Mode
- Throughput Mode
- Store Tiers

Availability and Scalability 
- Availability is the breadth an application can have (multiple Availability Zones in multiple AWSr Regions)
- Scalability is the height and count of resources to be commanded 

ELB Elastic Load Balancer
- AWS owns it
	- Legacy
	- ALB Application Load Balance
	- NLB Network load balancer
	- GLB Gateway load balancer
- Security groups
- Fully integrated into AWS
- Balances upstream requests to downstream resources
- Healthchecks

ASG Auto-scaling Groups
- Min. Size
- Spin up resources based on specs

RDS Relational Database System
- An assortment of SQL supporting databases
- Can create Read Replicas
- RDS Proxy (for pooling requests)
- 
Aurora DB
- An amazon optimized cloud database
- Reader/Writer endpoints
- autoscaling

Elastic Cache
- Like Redis
- Basically, it is cloud based ram.

Route 53 
- Authoritative DNS Record system
- can purchase domains
- Routing policies 

S3
- Super simple storage
- Scalable storage; like, most of the net uses it.
- You can replicate buckets across regions or in the same region
- Tiers of storage costs and retrieval 
- Transition rules for tier or expiry 
- event messages
- Encryption service
- Logging

AWS CloudFront
- edge locations that get requests into the private aws network
- or basically it is a global caching system 

AWS Global Accelerator 
- UDP style delivery of content from the edge locations

AWS Snowball
- Physical import and export of data from AWS

AWS FSx
- Hosted filesystem
  - Scratch
  - persistent

AWS Storage Gateway
- bridging on-premises and cloud storage

AWS Storage Sync
- sync data between services

AWS SQS - Standard Queueing services
- Ququeing service (messages in and messages out)

AWS SNS
- Publication / Subscription model 
- event based "topics"

Kinesis Data Streams
- Real time data streams

AWS Firehose 
- Near real time data stream to storage 
- Makes user of a buffer that flushes into target based on policy

AWS MQ - Message broker Queue 
- 3rd party message queue protocols

AWS ECS/EKS
- Container service (Docker)
- ASG auto scaling groups

AWS AppRunner / App2Container
- Fully managed infrastructure service for web applications

AWS Lambda
- Serverless code execution
- Automatically scales execution
- Can run up to 1000 functions concurrently

AWS API Gateway
- Serverless API 
- Plays well with other components of the serverless framework

AWS Databases
- RDS
- Aurora 
- ElastiCache
- DynamoDB
- S3
- DocumentDB
- Neptune
- Keyspaces
- Quantum Ledger DB
- Timestream

AWS data & analytics
- Athena 
- Redshift
- Opensearch
- Elastic MapReduce
- Quicksight
- Glue
- Lake Formation

AWS Machine Learning
- Rekognition -> content recognition
- Transcribe -> audio recognition (multi-language, and Personal Identifying Information screener)
- Polly -> Speech synthesis (pronunciation lexicons and pronunciation markup)
- Translate 
- Lex & Connect (Think Alexa plus Contact Management)
- Comprehend -> Natural Language Processing
- Comprehend Medical -> Natural Language Processing for medical services
- Sage Maker -> Building Custom Machine Learning Models
- Kendra -> Fully managed document search service
- Personalize -> think, product recommendations
- Textract -> OCR plus machine learning 

AWS Monitoring & Audit:
- Cloudwatch Metrics
- Cloudwatch Logs
- Cloudwatch Alarms
- Amazon EventBridge
- CloudWatch Containers
- CloudWatch Lambda
- CloudWatch Contributor Insights
- CloudWatch Application Insights 
- AWS CloudTrail
- Cloudtrail Insights
- AWS Config 

AWS Organization
- OrgUnits (OUs)
- AWS Directory Service (Windows)
- AWS Control Tower

AWS Security & Encryption 
- AWS KMS (Key Management Service)
- AWS Parameter Store
- AWS Secrets Manger
- AWS Certificates Manager
- AWS WAF (Web Application Firewall)
- AWS Shield (DDoS)
- AWS Firewall
- AWS GuardDuty
- AWS Inspector

AWS Networking 
- VPC
- Internet Gateway
- NAT instances | Going away? 
- NAT Gateways
- AWS Network Access Control List (NACL)
- VPC Peering 
- VPN, Virtual Private Gateway Virtual Customer Gateway
- Direct Connect
- Transit Gateway
- Network Firewall
_[Continue education here...](https://kipuhealth.udemy.com/course/aws-certified-solutions-architect-associate-saa-c03/learn/lecture/13528044#overview)_
