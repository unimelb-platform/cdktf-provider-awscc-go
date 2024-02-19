package neptunedbcluster


type NeptuneDbClusterAssociatedRoles struct {
	// The Amazon Resource Name (ARN) of the IAM role that is associated with the DB cluster.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/neptune_db_cluster#role_arn NeptuneDbCluster#role_arn}
	RoleArn *string `field:"required" json:"roleArn" yaml:"roleArn"`
	// The name of the feature associated with the AWS Identity and Access Management (IAM) role.
	//
	// For the list of supported feature names, see DBEngineVersion in the Amazon Neptune API Reference.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/neptune_db_cluster#feature_name NeptuneDbCluster#feature_name}
	FeatureName *string `field:"optional" json:"featureName" yaml:"featureName"`
}

