package apsworkspace


type ApsWorkspaceLoggingConfiguration struct {
	// CloudWatch log group ARN.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/aps_workspace#log_group_arn ApsWorkspace#log_group_arn}
	LogGroupArn *string `field:"optional" json:"logGroupArn" yaml:"logGroupArn"`
}

