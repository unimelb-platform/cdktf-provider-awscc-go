package rdsdbcluster


type RdsDbClusterAssociatedRoles struct {
	// The Amazon Resource Name (ARN) of the IAM role that is associated with the DB cluster.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/rds_db_cluster#role_arn RdsDbCluster#role_arn}
	RoleArn *string `field:"required" json:"roleArn" yaml:"roleArn"`
	// The name of the feature associated with the AWS Identity and Access Management (IAM) role.
	//
	// For the list of supported feature names, see DBEngineVersion in the Amazon RDS API Reference.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/rds_db_cluster#feature_name RdsDbCluster#feature_name}
	FeatureName *string `field:"optional" json:"featureName" yaml:"featureName"`
}

