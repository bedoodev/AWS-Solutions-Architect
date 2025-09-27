# S3 Static Website CloudFormation Template — What It Does & How to Use It

This template provisions an **Amazon S3 bucket configured for static website hosting** and makes the site **publicly readable** with a bucket policy. It also sets **modern ownership controls** and outputs the website endpoint URL.

---

## ✅ What the Template Creates

1. **S3 Bucket** (`S3Bucket`)
   - **BucketName:** Taken from the `BucketName` parameter (must be globally unique).
   - **OwnershipControls → `BucketOwnerEnforced`:** Disables ACLs and ensures the bucket owner owns all objects.
   - **PublicAccessBlockConfiguration:**
     - `BlockPublicAcls: true` and `IgnorePublicAcls: true` → ACL-based public access is blocked/ignored.
     - `BlockPublicPolicy: false`, `RestrictPublicBuckets: false` → Allows a **bucket policy** to grant public access (required for a public website).
   - **WebsiteConfiguration:** Sets `IndexDocument: index.html` and `ErrorDocument: error.html`.
   - **DeletionPolicy / UpdateReplacePolicy: `Retain`** → The bucket is not deleted or replaced automatically when the stack changes or is removed (prevents accidental data loss).

2. **Bucket Policy** (`BucketPolicy`)
   - Grants **public read** access (`s3:GetObject`) to all objects in the bucket.
   - Enables anonymous users (browsers) to fetch your `index.html`, images, JS, CSS, etc.

3. **Output**
   - **WebsiteURL:** The S3 **website** endpoint (HTTP) for your bucket.

---

## 🔧 How to Deploy & Use

### 1) Prepare a unique bucket name
Pick a globally unique name, e.g. `my-static-site-123456`.

### 2) Deploy the stack
Using AWS CLI:
```bash
aws cloudformation deploy \
  --template-file template.yaml \
  --stack-name s3-static-site \
  --parameter-overrides BucketName=my-static-site-123456 \
  --capabilities CAPABILITY_NAMED_IAM
```

### 3) copy an index.html to bucket
```bash
aws s3 cp index.html s3://my-static-site-123456/index.html
```
