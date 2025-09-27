# Introduction to S3

## What is object storage (Object-based storage)?
**<span style="color:#903030">Object storage</span> is a data storage archtitecture that manages data as objects, as opposed to other storage archtitectures.**

- <span style="color:#107010;">S3 provides you with unlimited storage</span>
- <span style="color:#107010;">You do not need to think about the underlying infrastructure</span>

### S3 -- Object
**<span style="color:#903030">Objects</span> contain your data. They are like files.**

*<span style="color:#903030">Object</span> may consist of:*

- <span style="color:#903030">Key</span> this is the name of the object
- <span style="color:#903030">Value</span> the data itself made up of a sequence of bytes
- <span style="color:#903030">Version ID</span> version of object, only appear versioning is enabled
- <span style="color:#903030">Metadata</span> additional information attached to the object

### S3 -- Bucket
**<span style="color:#903030">Buckets</span> hold S3 objects. Buckets can also have folders which in turn hold objects.**

- You can store an individual object from <span style="color:#903030">0 bytes to 5 Terabytes</span> in size

#### S3 -- Bucket Naming Rules
> https://**myexamplebucket**.s3.amazonaws.com/photo.jpg
- S3 Bucket names must be <span style="color:#903030">unique</span> globally
- Bucket names must be <span style="color:#903030">3-63 characters</span> long.
- Only <span style="color:#903030">lovercase letters, numbers, dots</span> and <span style="color:#903030">hypens</span> are allowed
- Names can't be formatted as <span style="color:#903030">IP Adresses</span>
- Names can't start with <span style="color:#903030">xn--, sthree-, sthree-configurator</span>
- Names can't end with <span style="color:#903030">-s3alias, --ol-s3</span>

**Examples:** 
    
- <span style="color:#159015">mybucket-123</span>
- <span style="color:#901515">123.456.789.012</span>
- <span style="color:#901515">My-Bucket</span>
- <span style="color:#901515">data.bucket..archive</span> (contains adjacent periods)
- <span style="color:#159015">log-bucket</span>
- <span style="color:#901515">xn--bucketname</span>


---

#### S3 -- Bucket Restrictions and Limitations
- You can  by default create <span style="color:#43ff54">100 buckets</span>
    - You can create a service request to increase to <span style="color:#43ff54">1000 buckets.</span>
- <span style="color:#43ff54">No max bucket size</span> and <span style="color:#43ff54">no limit to the number of objects</span> in a bucket
    - Files can be between <span style="color:#43ff54">0 and 5TBs</span>
    - Files larger than <span style="color:#43ff54">100MB</span> should use <span style="color:#43ff54">multi-part upload</span>

----

##### S3 -- Bucket Types
*Amazon S3 has two types of buckets:*

1. General purpose buckets
    - Organizes data in a <span style="color:#43ff54">flat hierarchy</span>
    - The original S3 bucket type
    - Recommended for most use cases
    - There aren't prefix limits
    - There is a default limit of <span style="color:#43ff54">100 general buckets</span> per account

2. Directory buckets
    - Organizes data <span style="color:#43ff54">folder hierarchy</span>
    - Only to be used with <span style="color:#43ff54">S3 Express One Zone storage class</span>
    - Individual directories can scale horizontally
    - There is a default limit of <span style="color:#43ff54">10 directory buckets</span> per account

----

### S3 -- Storage Classes

1. **S3 Standart (Default)**

    *It's designed for general purpose storage for frequently accessed data*
    - <span style="color:#43ff54">High Durability</span> 11 9's of durability (99.999999999%)
    - <span style="color:#43ff54">High Availability</span> 4 9's of availability (99.99%)
    - <span style="color:#43ff54">Data Redundancy</span> Data stored in 3 or more availability zones (AZs)
    - <span style="color:#43ff54">Retrieval Time</span> Within milliseconds (low latency)
    - <span style="color:#43ff54">High Throughput</span> Optimized for data that is frequently accessed and/or requires real-time access
    - <span style="color:#43ff54">Scalability</span> Easily Scales to storage size and number of requests
    - <span style="color:#43ff54">Use Cases</span> Big Data Analytics, Mobile and Gaming Applications, Content Distrubition
    - <span style="color:#43ff54">Pricing</span> 
        - Storage per GB
        - Per requests
        - No retrieval fee
        - No minimum storage duration charge

----
2. **S3 Intelligent Tiering**

    *It's designed for data with unknown or changing access patterns, automatically moving objects between storage tiers to optimize cost.*
    - <span style="color:#43ff54">Automatic Tiering</span> Moves data between frequent and infrequent access tiers based on usage
    - <span style="color:#43ff54">Latency & Throughput</span> Millisecond access latency, like S3 Standard
    - <span style="color:#43ff54">Durability</span> 99.999999999% (11 9’s), like all S3 storage classes
    - <span style="color:#43ff54">Availability</span> 99.9% (slightly lower than Standard, similar to Standard-IA)
    - <span style="color:#43ff54">No Retrieval Fee</span> Unlike IA and Glacier classes, you don’t pay to access data
    - <span style="color:#43ff54">Minimum Storage Duration</span> No minimum storage duration (good for unpredictable workloads)
    - <span style="color:#43ff54">Scalability</span> Automatically scales to billions of objects without performance impact
    - <span style="color:#43ff54">Use Cases</span> Data lakes, analytics, logs, user-generated content with unpredictable access patterns
    - <span style="color:#43ff54">Pricing</span> 
        - Storage per GB (slightly higher than IA, but cost optimized automatically)
        - Per requests (PUT, GET, lifecycle transitions)
        - Small monthly monitoring and automation fee
        - No retrieval fee
----

3. **S3 Express One Zone**

    *Amazon S3 Express One Zone delivers consistent single-digit millisecond data access for your most frequently accessed data and latency-sensitive applications.*
    - <span style="color:#43ff54">10x Faster</span> than S3 Standard
    - <span style="color:#43ff54">%50 Lower</span> costs than S3 Standard
    - <span style="color:#43ff54">Single Availability Zone</span> Data is stored in a user selected single AZ
    - Data is stored in <span style="color:#43ff54">Amazon S3 Directory Bucket</span>
    - <span style="color:#43ff54">Use Cases</span> AI/ML, High Performance Computing (HPC)
    - <span style="color:#43ff54">Pricing</span> 
        - Storage per GB (lower than S3 Standard, single-AZ)
        - Per requests (optimized for high request rates)
        - No retrieval fee
        - No minimum storage duration charge
        - Directory bucket pricing model (unique to S3 Express)

----

4. **S3 Standart-IA** 
    
    *It's designed for data that is less frequently accessed but requires rapid access when needed*
    - <span style="color:#43ff54">Hight Durability</span> Like S3 Standard
    - <span style="color:#43ff54">High Availability</span> 3 9's of availability (99.9%)
    - <span style="color:#43ff54">Data Redundancy</span> Like S3 Standard
    - <span style="color:#43ff54">Cost Effective Storage</span> costs 50% less from standard. <span style="color:#ff5045">As long as you don't access a file more than once a month.</span>
    - <span style="color:#43ff54">High Throughput</span> Optimized for rapid access, although the data is accessed less frequently compared to S3 Standard
    - <span style="color:#43ff54">Scalability</span> Easily Scales to storage size and number of requests Like S3 Standard
    - <span style="color:#43ff54">Use Cases</span> Disaster Recovery, Backups, Long-Term data stores where data is not frequently accessed
    - <span style="color:#43ff54">Pricing</span> 
        - Storage per GB
        - Per requests
        - Has a Retrieval fee
        - Has a minimum storage duration charge of <span style="color:#ff5045">30 days</span>

----

5. S3 One-Zone-IA

    *It's designed for data that is less frequently accessed and has additional saving at reduced availability*
    - <span style="color:#43ff54">High Durability</span> Like S3 Standard
    - <span style="color:#43ff54">Low Availability</span> 99.5%, Since its in a single AZ it has even lower availability than Standard-IA
    - <span style="color:#43ff54">Cost Effective Storage</span> costs 20% Less than Standard-IA
    - <span style="color:#43ff54">Data Redundancy</span> Because of data stores in one AZ, there is a risk to loss data in case of AZ disaster
    - <span style="color:#43ff54">Retrieval Time</span> Within milliseconds (Low Latency)
    - <span style="color:#43ff54">Use Cases</span> It's ideal for secondary-backup copies of on-premises data, or for storing data that is recreatable in case of the AZ failure.
    - <span style="color:#43ff54">Pricing</span>
        - Storage per GB
        - Per requests
        - Has a Retrieval fee
        - Has a minimum storage duration charge of <span style="color:#ff5045">30 days</span>


----

6. S3 Glacier Instant Retrieval

    *It's designed for rarely accessed data that still needs immediate access in performance-sensitive use cases.*
    - <span style="color:#43ff54">Hight Durability</span> Like S3 Standard
    - <span style="color:#43ff54">High Availability</span> Like S3 Standard-IA
    - <span style="color:#43ff54">Data Redundancy</span> Like S3 Standard
    - <span style="color:#43ff54">Cost Effective Storage</span> costs 68% less from Standard-IA. <span style="color:#ff5045">For long-lived data that is accessed once per quarter</span>
    - <span style="color:#43ff54">High Throughput</span> Optimized for rapid access, although the data is accessed less frequently compared to S3 Standard
    - <span style="color:#43ff54">Scalability</span> Easily Scales to storage size and number of requests Like S3 Standard
    - <span style="color:#43ff54">Use Cases</span> Image Hosting, Online file-sharing applications, medical imaging and health records, news media assets
    - <span style="color:#43ff54">Pricing</span> 
        - Storage per GB
        - Per requests
        - Has Retrieval fee (More expensive than Standard-IA)
        - Has a minimum storage duration charge of <span style="color:#ff5045">90 days</span>

----

7. S3 Glacier Flexible retrieval

    *It's designed for rarely accessed data that no need immediate access.*
    - <span style="color:#43ff54">There are 3 retrieval tiers</span>
        - <span style="color:#43ff54">Expediated Tier</span> 1-5 Mins, For urgent requests, Limited to 250MB
        - <span style="color:#43ff54">Standard Tier</span> 3-5 Hours, No archive size limit
        - <span style="color:#43ff54">Bulk Tier</span> 5-12 Hours, Noarchive size limit even petabytes worth of data
    - <span style="color:#43ff54">Pricing</span> 
        - Storage per GB (very low cost, cheaper than Instant Retrieval)
        - Per requests (PUT, GET, lifecycle transitions)
        - Retrieval fee (depends on retrieval type: Expedited, Standard, Bulk)
        - Has a minimum storage duration charge of <span style="color:#ff5045">90 days</span>


----

8. S3 Glacier Deep Archive

    *It's designed for data that is rarely accessed and meant for long-term archival and compliance needs.*
    - <span style="color:#43ff54">Lowest-Cost Storage</span> Cheapest storage option in Amazon S3
    - <span style="color:#43ff54">Long-Term Retention</span> Ideal for regulatory archives, compliance records, and digital preservation
    - <span style="color:#43ff54">Durability</span> 99.999999999% (11 9’s), like all S3 storage classes
    - <span style="color:#43ff54">Availability</span> Lower than Standard and IA classes, designed for archival workloads
    - <span style="color:#43ff54">Retrieval Times</span> 12 to 48 Hours (not suited for time-sensitive workloads)
    - <span style="color:#43ff54">Use Cases</span> Financial and healthcare compliance, government records, media archives, scientific datasets
    - <span style="color:#43ff54">Pricing</span> 
        - Storage per GB (lowest in S3, cheaper than Flexible Retrieval)
        - Per requests (PUT, GET, lifecycle transitions)
        - Retrieval fee (charged per retrieval, higher latency = lower cost per GB retrieved)
        - Has a minimum storage duration charge of <span style="color:#ff5045">180 days</span>
----

### S3 -- Security

1. S3 Block Public Access 
    
    *Block Public Access is as safety feature that is enabled by default to block all public access to an S3 bucket*

    **S3 Block Public Access Options**

| Option | What it does | When it applies | Important note |
|--------|--------------|-----------------|----------------|
| **1. Block public access to buckets and objects granted through new ACLs** | Prevents new buckets or objects from being granted public access through ACLs. | Applies only to **new ACLs**. Does not affect existing ACLs. | Existing public ACLs remain in effect. New public ACLs cannot be created. |
| **2. Block public access to buckets and objects granted through any ACLs** | Completely ignores all ACLs that grant public access. | Applies to **both new and existing** ACLs. | Public access via ACLs is fully disabled. |
| **3. Block public access to buckets and objects granted through new public bucket or access point policies** | Blocks the creation of new bucket or access point policies that grant public access. | Applies only to **new policies**. Existing policies remain active. | Prevents accidental exposure via new bucket policies. |
| **4. Block public and cross-account access to buckets and objects through any public bucket or access point policies** | Ignores all public and cross-account access granted by bucket or access point policies. | Applies to **both new and existing** policies. | Even if a policy grants public access, AWS will not honor it. |

----

2. S3 Access Control List (ACL) (Legacy Method)
    
    *ACLs grant basic Read/Write permissions to other AWS Accounts*

    - You <span style="color:#ff5045">can</span> grant permissions to <span style="color:#ff5045">other AWS accounts</span>
    - You <span style="color:#ff5045">cannot</span> grant permissions to <span style="color:#ff5045">your AWS account</span>
    - You <span style="color:#ff5045">cannot</span> grant conditional permissions
    - You <span style="color:#ff5045">cannot</span> explicitly deny permissions
    - <span style="color:#43ff54">It's only work on user basis</span>

----

3. S3 Bucket Policies

    *S3 Bucket Policy is a resource-based policy to grant Bucket and Objects to other Principals eg. AWS Accounts, Users, AWS Services*

     **S3 Bucket Policies vs IAM Policies**

| Feature | **S3 Bucket Policy** | **IAM Policy** |
|----------|-----------------------|----------------|
| **Where is it defined?** | Attached directly to an **S3 bucket**. | Attached to an IAM **user, group, or role**. |
| **Scope** | **Resource-based policy** → controls access to the bucket and its objects. | **Identity-based policy** → controls what an IAM identity can do. |
| **Who does it apply to?** | Any AWS principal (users, roles, accounts, even anonymous/public). | Only the IAM identity (user, group, or role) it is attached to. |
| **Public access** | Can allow public (anonymous) or cross-account access. | Cannot allow public access; only identities in your AWS account. |
| **Granularity** | Grants access to a **specific bucket** and its objects. | Can grant access to **many AWS services** and multiple buckets in one policy. |
| **Principal** | You can explicitly list **multiple principals** to grant access. | Principal is **implicitly** the IAM entity the policy is attached to. |
| **Policy size limit** | Up to **20 KB**. | Limited by entity type: <br> - Users: **2 KB** <br> - Groups: **5 KB** <br> - Roles: **10 KB** |
| **Block Public Access** | If enabled (default), it overrides and **denies all anonymous access**, even if a bucket policy allows it. | Not affected directly by Block Public Access (applies only to bucket-level settings). |
| **Explicit deny** | Supports `"Effect": "Deny"`. | Supports `"Effect": "Deny"`. |
| **Typical use case** | - Allow cross-account access <br> - Public read access for a bucket | - Control what actions your IAM users/roles can perform in AWS, including S3 |
| **Recommended use** | Cross-account and bucket-level control. | Internal account-level access control. |


---

4. S3 Access Grants

    *Amazon S3 Access Grants lets you map identities in a directory service (Active Directory, Okta...) to access datasets in s3*

    - No need to create IAM users, everyone gets direct access with their corporate account.

    ```mermaid
    flowchart LR
    A[Corporate User] --> B[Access Grants Instance]
    B --> C[Associated Permissions]
    C --> D[S3 Resource]    
    ```

----

5. S3 CORS

    *Amazon S3 allows you to set CORS configuration to a S3 bucket with static website hosting so different origins can perform HTTP requests from your S3 static website*

    ```json
    {
        "CORSRules": [
            {
            "AllowedOrigins": ["https://example.com"],
            "AllowedMethods": ["GET", "PUT"],
            "AllowedHeaders": ["*"],
            "ExposeHeaders": ["x-amz-server-side-encryption"],
            "MaxAgeSeconds": 3600
            }
        ]
    }
    ```
    
    - S3 CORS configuration can be both `JSON` and `XML`, but AWS Console only accept `JSON`.
    - **AllowedOrigins** → Which domains are allowed to send requests. (`*` means open to everyone).  
    - **AllowedMethods** → Which HTTP methods are permitted (`GET`, `PUT`, `POST`, `DELETE`, `HEAD`).  
    - **AllowedHeaders** → Which headers are allowed in requests. `*` = all headers.  
    - **ExposeHeaders** → The response headers that will be exposed to the browser.  
    - **MaxAgeSeconds** → The amount of time (in seconds) the preflight `OPTIONS` response can be cached by the browser.  

----

### S3 -- Encryption
    
1. **Encryption in Transit**

    *When data is encrypted by the sender and then decrypted the receiver.*

    - Data that is secure when moving between locations algorithms: <span style="color:#ff5045">TLS/SSL</span>
    - <span style="color:#ff5045">TLS</span>: Transport Layer Security, An encryption protocol for data integrity between two or more communicating computer application. TLS 1.2 and 1.3 are the current best practises
    - <span style="color:#ff5045">SSL</span>: Secure Sockets Layer, Same as TLS. 1.0, 2.0, 3.0 are deprecated
    
----

2. **Encryption-At-Rest**
    
    ----

    1. **Client-Side Encryption (CSE)**
        
        *When data is encrypted by the client and then sent to the server*
        
        - This provides a guarantee that AWS and no third-party can decrypt your data.

    ----

    2. **Server-Side Encryption (SSE)**
        
        *When data is encrypted by the server. SSE is always-on for all new S3 objects*
        
        ----
        
        1. <span style="color:#ff5045">SSE-S3</span>
            
            *SSE-S3 is when Amazon manages all the of encryption*

            - S3 encrypts each object with an unique Key
            - S3 uses envelope encryption
            - S3 rotates regularly key automatically
            - There is <span style="color: #43ff54">no additional charge</span> for using SSE-S3
            - SSE-S3 uses 256-bit Advanced Encryption Standard (AES-256)

            ```json
                aws s3api put-object \
                    --bucket mybucket \
                    --key myfile \
                    --server-side-encryption AES256 \  // it's default encryption config that's why no need for this option.
                    --body myfile.txt
            ```
        
        ----

        2.  <span style="color:#ff5045">SSE-KMS</span>

            *SSE-KMS is when you use a KMS Key managed by AWS*

            - First, create a KMS managed key, and choose one to encrypt your object
            - KMS can automatically rotate keys
            - KMS key policy controls who can decrypt using the key
            - KMS can help meet regulatory compliance
            - KMS keys must be in the same region as the bucket.

            ```json
                aws s3api put-object \
                    --bucket mybucket \
                    --key myfile \
                    --server-side-encryption 'aws:kms' \
                    --ssekms-key-id 2cf97dd3-e166-418b-a1a9-700422eb57f1
                    --body myfile.txt
            ```
        
        ----
        
        3. <span style="color:#ff5045">SSE-C</span>

            *SSE-C is when you provide your own encryption key that Amazon S3 then uses the apply AES-256 encryption to your data.*
            
            - You need to provide an encryption key to retrieve the object for everytime
            - AWS does not store the encryption key. The key is removed from Amazon S3 after each request.
            - S3 will be store a randomly Hash-based message Authentication Code (HMAC) of your key to validate for future requests.
            - There is <span style="color: #43ff54">no additional charge</span> for using SSE-C
            ```json
            // Create a 32 byte key for instance:
            openssl rand -base64 32 > my_key.b64

            // Calculate MD5 checksum:
            KEY=$(cat my_key.b64)
            KEY_MD5=$(echo -n "$KEY" | openssl md5 -binary | base64)

            // Upload an object
            aws s3api put-object \
              --bucket my-sse-c-bucket \
              --key secret.txt \
              --body ./secret.txt \
              --sse-customer-algorithm AES256 \
              --sse-customer-key "$KEY" \
              --sse-customer-key-md5 "$KEY_MD5"
            ```
        
        ----

        4. <span style="color:#ff5045">DSSE-KMS</span>

            *DSSE (Dual-layer server-side encryption) KMS. It's a SSE-KMS with the inclusion of client-side encryption*

            - Data is encrypted <span style="color:#ff5045">twice</span> with DSSE-KMS
            - The key used for client-side encryption comes from KMS
            - There are <span style="color:#ff5045">additional charge</span> for DSSE and KMS keys
            
            ```json
                aws s3api put-object \
                    --bucket mybucket \
                    --key myfile \
                    --server-side-encryption aws:kms:dsse \
                    --ssekms-key-id 2cf97dd3-e166-418b-a1a9-700422eb57f1
                    --body myfile.txt
            ```
        
        ----
        
        5. <span style="color:#ff5045">S3 Bucket Key</span>
            
            *When you use SSE-KMS, an individual data key used on every object request. In this case S3 has to call AWS KMS everytime a request is made. KMS charges on the number of requests and so this charge can add up*

            - This will <span style="color:#43ff54">reduce request costs by up to 99%</span>
            - This will <span style="color:#43ff54">decrease</span> request traffic improve overall performance
            ```json
            aws s3api put-bucket-encryption \
              --bucket my-bucket \
              --server-side-encryption-configuration '{
                  "Rules": [
                    {
                      "ApplyServerSideEncryptionByDefault": {
                        "SSEAlgorithm": "aws:kms",
                        "KMSMasterKeyID": "arn:aws:kms:region:acct-id:key/key-id"
                      },
                      "BucketKeyEnabled": true
                    }
                  ]
                }'
            ```
----

### S3 -- Data Consistency
    
**What is data consistency?**

- When data being kept in two different place and whether the data exactly match or do not match.
- **Strongly Consistent:** You see instantly what you change.
    - Once a write operation (PUT/UPDATE/DELETE) is successful, all read operations (GET/LIST) immediately return the most up-to-date data.
- **Eventually Consistent:** You see updated version of object eventually, but not instantly.   - A write operation is accepted immediately, but it may take some time for the change to become visible across all replicas in the system.
- Amazon S3 offers <span style="color:#43ff54">strong consistency</span> for all read, write, and delete operations.
> Prior to Jan 2020, S3 did not have strong consistency for all s3 operations.

----

### S3 -- Object Replication

**S3 Object Replication automatically and asynchronously copies objects from one bucket to another.**
- <span style="color:#ff5045">CRR</span>, Cross-Region Replication (live-replication)
- <span style="color:#ff5045">SRR</span>, Same-Region Replication (live-replication)
- Object Replication can help you do the following;
    - Replicate object into different storage Classes
    - Replicate object while retaining Metadata
    - Maintain object copies under different ownership
    - Keep objects stored over multiple AWS regions
    - Replicate objects within 15 minutes
- Only new objects are replicated, old objects need `S3 Batch Replication`
- Versioning must be enabled for both buckets.
- Special permissions are required for SSE-KMS encrypted objects.

----

### S3 -- Versioning
**S3 Versioning allow you to store multiple versions of S3 objects**
- Store all versions of an object in S3 at the same object key address
- By default, versioning is disabled on buckets.
- Once enabled, it cannot be disabled, only suspended on the bucket.
- Fully integrates with S3 lifecycle rules.
- MFA Delete feature provides extra protection against deletion of your data.
- **With versioning you can recover more easily from unintended user actions and application failures**
- **Versioning-enabled buckets can help you recover objects from accidental deletion or overwrite**

----

### S3 -- Object Lifecycle (Lifecycle Management)
**S3 Lifecycle allows you to manage of objects life cycle, archival or deletion of objects.**
- Can bu used together with versioning and can be applied to both current and previous ersions
- Filter based on prefix, object tags or object size.
- Two Type of Actions;
    - <span style="color:#ff5045">Transition Actions:</span> 
        - Moves the object to another class of storage after a certain period 
        > After 30 days →  Standard-IA, After 365 days →  Glacier.
    
    - <span style="color:#ff5045">Expiring Actions:</span>
        - Deletes the objects or old versions of them after a certain period 

    ```json
       {
        "Rules":[
            {
                "ID":"MoveToGlacierAndExpire",
                "Filter": {
                    "And": {
                        "Prefix": "logs",
                        "ObjectSizeGreaterThan": 2048
                    }
                },
                "Status":"Enabled",
                "Transitions":
                [
                    {
                        "Days":30,
                        "StorageClass":"STANDARD_IA"
                    },
                    {
                        "Days":90,
                        "StorageClass":"GLACIER"
                    }
                ],
                "Expiration":
                {
                    "Days":365
                }
            }
        ]
    }
    ```

    ```mermaid
    flowchart LR
    A[Object Created in logs] -->|"After 30 Days"| B[After 30 days: Move to STANDARD_IA]
    B["Move to STANDARD_IA"] -->|"After 90 days"| C[After 90 days: Move to GLACIER]
    C["Move to GLACIER"] -->|"After 365 days"| D["Delete"]
    ```

----

### S3 -- Transfer Acceleration
**Transfer Acceleration is a bucket-level feature that provides fast and secure transfer of files over long distances between your end users and an S3 Bucket.**
- <span style="color:#ff5045">Advantages</span>
    - Faster upload/download
    - Optimized data transfer using AWS's global infrastructure
- <span style="color:#ff5045">Disadvantages</span>
    - Extra charge
    - If the client and bucket are in the same region, performance gains are minimal and the feature may be unnecessary.

----

### S3 -- Presigned URLs
**Presigned URLs provides a temporary access to download or upload object data via URL**
```json
    aws s3 presign s3://mybucket/myobject \
        --expires-in 300
```
- *Anatomy of Presigned URLs*
    ```bash
     https://mybucket.s3.amazonaws.com/myobject
     ?X-Amz-Algorithm=AWS4-HMAC-SHA256
     &X-Amz-Credential=YOUR_AWS_ACCESS_KEY%2F20231125%2Fus-east-1...
     &X-Amz-Date=20231125T123456Z
     &X-Amz-Expires=300
     &X-Amz-SignedHeaders=host
     &X-Amz-Signature=GENERATED_SIGNATURE
    ```
----

### S3 -- Access Points
**Access Points simplfy managing data access at scale for shared datasets in S3**

- Supposing that there are 3 application using same bucket   
    - Application A will only read
    - Application B will only write
    - Application C will only be able to view its own folders

    Access Points manage these different permissions easier than bucket policies.

```json
{
    "Version":"2012-10-17",		 	 	 
    "Statement": 
    [
        {
            "Effect": "Allow",
            "Principal": {
                "AWS": "arn:aws:iam::123456789012:user/A"
            },
            "Action": ["s3:GetObject"],
            "Resource": "arn:aws:s3:us-west-2:123456789012:accesspoint/my-access-point/object/demo/*"
        },
        {
            "Effect": "Allow",
            "Principal": {
                "AWS": "arn:aws:iam::123456789012:user/B"
            },
            "Action": ["s3:PutObject"],
            "Resource": "arn:aws:s3:us-west-2:123456789012:accesspoint/my-access-point/object/demo/*"
        }
    ]
}
```

----

### S3 -- Mountpoint

**Mountpoint for Amazon S3 allows you to mount an S3 bucket to your <span style="color:#ff5045">Linux</span> local file system**

- You can use <span style="color:#ff5045">ls, cat, cp</span> commands in S3 bucket.
- Mountpoint can:
    - Read files up to <span style="color:#ff5045">5TB</span> in size
    - List and read existing files
    - Create new files
- Mountpoint cannot:
    - Modify existing files
    - Delete directories
    - Support symbolic links
    - Support file locking
- Mountpoint can be used in following storage classes:
    - S3 Standard
    - S3 Standard IA
    - S3 One-Zone IA
    - S3 Glacier Instant Retrieval

----

### S3 -- Batch Operations

**Amazon S3 Batch Operations allows you to perform bulk operations on millions of objects in S3**

- Possible Operations:
    - <span style="color:#ff5045">COPY</span> Copy objects to another bucket.
    - <span style="color:#ff5045">PUT Object ACLs</span> Change Permissions.
    - <span style="color:#ff5045">PUT Object Tagging</span> Add or Update tags.
    - <span style="color:#ff5045">PUT Object Lock</span> Set retention or legal hold
    - <span style="color:#ff5045">Invoke AWS Lambda</span> Run a custom function for each object.

----

### S3 -- Amazon Inventory

**Amazon S3 Inventory takes inventory of object in an S3 bucket on a repeating schedule so you have an audit history of object changes.**

----

### S3 -- Event Notifications

**S3 Event Notifications allows your bucket to notify other AWS services about S3 event data.**

----

### S3 -- Multipart Upload

**Amazon S3 supports multi-part upload so you can upload a single object in a set of parts.**

- Advantages:
    - Improved throughput
    - In case of network fail you just need to reupload the missing parts
    - Once you start a multi-part upload, you can upload parts at any time, there is no expire time to upload.
- <span style="color:#ff5045">Recommended for files that are +100MB</span>
