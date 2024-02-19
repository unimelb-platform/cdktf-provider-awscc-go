package s3accesspoint


type S3AccessPointPublicAccessBlockConfiguration struct {
	// Specifies whether Amazon S3 should block public access control lists (ACLs) for buckets in this account.
	//
	// Setting this element to TRUE causes the following behavior:
	// - PUT Bucket acl and PUT Object acl calls fail if the specified ACL is public.
	//  - PUT Object calls fail if the request includes a public ACL.
	// . - PUT Bucket calls fail if the request includes a public ACL.
	// Enabling this setting doesn't affect existing policies or ACLs.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/s3_access_point#block_public_acls S3AccessPoint#block_public_acls}
	BlockPublicAcls interface{} `field:"optional" json:"blockPublicAcls" yaml:"blockPublicAcls"`
	// Specifies whether Amazon S3 should block public bucket policies for buckets in this account.
	//
	// Setting this element to TRUE causes Amazon S3 to reject calls to PUT Bucket policy if the specified bucket policy allows public access. Enabling this setting doesn't affect existing bucket policies.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/s3_access_point#block_public_policy S3AccessPoint#block_public_policy}
	BlockPublicPolicy interface{} `field:"optional" json:"blockPublicPolicy" yaml:"blockPublicPolicy"`
	// Specifies whether Amazon S3 should ignore public ACLs for buckets in this account.
	//
	// Setting this element to TRUE causes Amazon S3 to ignore all public ACLs on buckets in this account and any objects that they contain. Enabling this setting doesn't affect the persistence of any existing ACLs and doesn't prevent new public ACLs from being set.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/s3_access_point#ignore_public_acls S3AccessPoint#ignore_public_acls}
	IgnorePublicAcls interface{} `field:"optional" json:"ignorePublicAcls" yaml:"ignorePublicAcls"`
	// Specifies whether Amazon S3 should restrict public bucket policies for this bucket.
	//
	// Setting this element to TRUE restricts access to this bucket to only AWS services and authorized users within this account if the bucket has a public policy.
	// Enabling this setting doesn't affect previously stored bucket policies, except that public and cross-account access within any public bucket policy, including non-public delegation to specific accounts, is blocked.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/s3_access_point#restrict_public_buckets S3AccessPoint#restrict_public_buckets}
	RestrictPublicBuckets interface{} `field:"optional" json:"restrictPublicBuckets" yaml:"restrictPublicBuckets"`
}

