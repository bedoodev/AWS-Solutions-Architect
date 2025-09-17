## Introduction to S3

### What is object storage (Object-based storage)?
**<span style="color:#903030">Object storage</span> is a data storage archtitecture that manages data as objects, as opposed to other storage archtitectures.**

- <span style="color:#107010;">S3 provides you with unlimited storage</span>
- <span style="color:#107010;">You do not need to think about the underlying infrastructure</span>

#### S3 Object
**<span style="color:#903030">Objects</span> contain your data. They are like files.**

*<span style="color:#903030">Object</span> may consist of:*

- <span style="color:#903030">Key</span> this is the name of the object
- <span style="color:#903030">Value</span> the data itself made up of a sequence of bytes
- <span style="color:#903030">Version ID</span> version of object, only appear versioning is enabled
- <span style="color:#903030">Metadata</span> additional information attached to the object

#### S3 Bucket
**<span style="color:#903030">Buckets</span> hold S3 objects. Buckets can also have folders which in turn hold objects.**

- You can store an individual object from <span style="color:#903030">0 bytes to 5 Terabytes</span> in size

<div style="border-left:4px solid red; padding:8px; background:#303030;">
<strong></strong> S3 is a universal namespace so <span style="color:#903030">bucket names</span> must be unique!
</div>

----

#### ACL (Access Control List)
S3 <span style="color:#903030">ACL</span> is a **legacy** mechanism that defines access permissions for a bucket or object.

- <span style="color:#903030">Enabled</span> Flexible, allows object sharing, but risky.
- <span style="color:#903030">Disabled</span> Simple, secure, AWS recommended method.

----

#### Bucket Versioning
S3 <span style="color:#903030">Bucket Versioning</span> is a feature that allows you to store multiple versions of the same object.

- When <span style="color:#903030">versioning</span> is enabled and you upload an object again, the old object is not deleted; a new version is created.
- When you delete an object, it is not permanently removed; a <span style="color:#903030">delete marker</span> is added. You can restore previous versions.
- <span style="color:#903030">Enabled</span> New uploads create new versions, and old versions are retained
- <span style="color:#903030">Disabled</span> Versioning is disabled, every new upload completely replaces the old one.
- <span style="color:#903030">Suspended</span> New objects are not versioned, but previous versions remain stored.

----

#### Encryption
S3 <span style="color:#903030">Encryption</span> is the process of encrypting data in buckets or objects to protect it from unauthorized access.

##### Amazon S3 Encryption Types

###### 1. SSE-S3 (Server-Side Encryption with S3 Managed Keys)

- <span style="color:#903030">Description:</span> AWS S3 automatically encrypts data at rest using server-side keys managed entirely by AWS.  
- <span style="color:#903030">Algorithm:</span> AES-256 encryption.  
- <span style="color:#903030">Key Management:</span> Fully managed by AWS, no customer interaction required.  
- <span style="color:#903030">Advantages:</span>
  - Simple to enable and use.
  - No need to manage keys yourself.
- <span style="color:#903030">Disadvantages:</span>
  - Less flexibility and control over encryption keys.

---

###### 2. SSE-KMS (Server-Side Encryption with AWS KMS Keys)

- <span style="color:#903030">Description:</span> Encryption is handled using AWS Key Management Service (KMS). Customers can create, manage, and audit encryption keys.  
- <span style="color:#903030">Key Management:</span> Managed by AWS KMS, with the option for customer-controlled keys.  
- <span style="color:#903030">Advantages:</span>
  - More control over keys.
  - Detailed audit logging and key rotation support.
  - Integration with IAM for fine-grained access control.
- <span style="color:#903030">Disadvantages:</span>
  - Additional costs for KMS usage.
  - Slightly more complex to configure.

---

###### 3. SSE-C (Server-Side Encryption with Customer-Provided Keys)

- <span style="color:#903030">Description:</span> Customers provide their own encryption keys for each upload request. AWS uses the key for encryption but does not store it.  
- <span style="color:#903030">Key Management:</span> Entirely the customer’s responsibility.  
- <span style="color:#903030">Advantages:</span>
  - Full control over encryption keys.
- <span style="color:#903030">Disadvantages:</span>
  - If the key is lost, the data cannot be decrypted.
  - Key management is entirely manual.

---

###### 4. CSE (Client-Side Encryption)

- <span style="color:#903030">Description:</span> Data is encrypted by the client before it is uploaded to S3. AWS only stores the already-encrypted data.  
- <span style="color:#903030">Key Management:</span> Managed entirely by the customer on the client side.  
- <span style="color:#903030">Advantages:</span>
  - Maximum security since AWS never sees the encryption keys.
- <span style="color:#903030">Disadvantages:</span>
  - Key management and encryption logic are fully the customer’s responsibility.
  - Higher complexity to implement.

---

#### Amazon S3 Storage Classes
Amazon S3 provides multiple <span style="color:#903030">storage classes</span> designed for different use cases. Each class balances cost, availability, and durability differently.

---

###### 1. S3 Standard

- <span style="color:#903030">Use Cases:</span>
  - Frequently accessed data
  - Websites, content delivery, big data analytics
- <span style="color:#903030">Advantages:</span>
  - High durability (99.999999999%)
  - High availability (99.99%)
  - Low latency and high throughput
- <span style="color:#903030">Disadvantages:</span>
  - Most expensive storage class compared to others

---

###### 2. S3 Intelligent-Tiering

- <span style="color:#903030">Use Cases:</span>
  - Data with unknown or changing access patterns
  - Long-term storage where access frequency varies
- <span style="color:#903030">Advantages:</span>
  - Automatically moves data between frequent and infrequent access tiers
  - No performance impact
  - Cost savings without manual intervention
- <span style="color:#903030">Disadvantages:</span>
  - Small monitoring and automation fee
  - Slightly higher cost than infrequent access for predictable workloads

---

###### 3. S3 Standard-IA (Infrequent Access)

- <span style="color:#903030">Use Cases:</span>
  - Data accessed less frequently but needs quick access when required
  - Backups, disaster recovery
- <span style="color:#903030">Advantages:</span>
  - Lower cost than Standard
  - High durability and availability
  - Millisecond access
- <span style="color:#903030">Disadvantages:</span>
  - Retrieval fee per GB
  - Minimum storage duration charge (30 days)

---

###### 4. S3 One Zone-IA

- <span style="color:#903030">Use Cases:</span>
  - Infrequently accessed data that can be re-created easily
  - Secondary backups, non-critical data
- <span style="color:#903030">Advantages:</span>
  - ~20% cheaper than Standard-IA
  - High durability within a single AZ
- <span style="color:#903030">Disadvantages:</span>
  - Stored in only one Availability Zone (less resilient)
  - Retrieval fee applies

---

###### 5. S3 Glacier Instant Retrieval

- <span style="color:#903030">Use Cases:</span>
  - Archival data that is rarely accessed but must be retrieved immediately
  - Medical images, compliance records
- <span style="color:#903030">Advantages:</span>
  - Low cost similar to archive
  - Millisecond retrieval time
- <span style="color:#903030">Disadvantages:</span>
  - Retrieval fee
  - Minimum storage duration (90 days)

---

###### 6. S3 Glacier Flexible Retrieval (formerly Glacier)

- <span style="color:#903030">Use Cases:</span>
  - Archival data accessed occasionally (hours retrieval time is acceptable)
  - Long-term backups
- <span style="color:#903030">Advantages:</span>
  - Very low storage cost
  - Options for expedited, standard, or bulk retrieval
- <span style="color:#903030">Disadvantages:</span>
  - Retrieval time ranges from minutes to hours
  - Retrieval cost
  - Minimum storage duration (90 days)

---

###### 7. S3 Glacier Deep Archive

- <span style="color:#903030">Use Cases:</span>
  - Rarely accessed archival data (compliance, regulatory, legal)
  - Data retention for 7–10 years
- <span style="color:#903030">Advantages:</span>
  - Lowest storage cost in S3
  - Suitable for "cold storage" and compliance needs
- <span style="color:#903030">Disadvantages:</span>
  - Retrieval time 12–48 hours
  - Retrieval cost
  - Minimum storage duration (180 days)

---

###### 8. S3 Reduced Redundancy Storage (RRS) [Deprecated]

- <span style="color:#903030">Use Cases:</span>
  - Non-critical, easily reproducible data (e.g., thumbnails)
- <span style="color:#903030">Advantages:</span>
  - Cheaper than S3 Standard
- <span style="color:#903030">Disadvantages:</span>
  - Lower durability (99.99% vs 99.999999999%)
  - Deprecated by AWS, not recommended for new use

