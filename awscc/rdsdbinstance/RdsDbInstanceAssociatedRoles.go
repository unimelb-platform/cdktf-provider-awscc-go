package rdsdbinstance


type RdsDbInstanceAssociatedRoles struct {
	// The name of the feature associated with the AWS Identity and Access Management (IAM) role.
	//
	// IAM roles that are associated with a DB instance grant permission for the DB instance to access other AWS services on your behalf.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/rds_db_instance#feature_name RdsDbInstance#feature_name}
	FeatureName *string `field:"required" json:"featureName" yaml:"featureName"`
	// The Amazon Resource Name (ARN) of the IAM role that is associated with the DB instance.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/rds_db_instance#role_arn RdsDbInstance#role_arn}
	RoleArn *string `field:"required" json:"roleArn" yaml:"roleArn"`
}

