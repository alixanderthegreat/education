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
Scalable, Global Applications delivered by AWS Regions via their availablitiy zones.

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
> least privlege principle implies that you only give a user the least amount of privlege necessary for them to accomplish their work

Root account <- your admin account 
- It is not best practice to work from the root account

Group <- the roles in your organization ; inheiratable 
> It is a good idea to start by creating an `admin` group

- Permissions are defined with the IAM policy, or at creation
- Users can have multiple groups


User <- people in the organization
- Permissions can be set in the IAM policy, or inline
- Managemenent Console
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

[Continue education here...](https://kipuhealth.udemy.com/course/aws-certified-solutions-architect-associate-saa-c03/learn/lecture/26098060#content)
