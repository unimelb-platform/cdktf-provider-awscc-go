package grafanaworkspace


type GrafanaWorkspaceVpcConfiguration struct {
	// The list of Amazon EC2 security group IDs attached to the Amazon VPC for your Grafana workspace to connect.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/grafana_workspace#security_group_ids GrafanaWorkspace#security_group_ids}
	SecurityGroupIds *[]*string `field:"required" json:"securityGroupIds" yaml:"securityGroupIds"`
	// The list of Amazon EC2 subnet IDs created in the Amazon VPC for your Grafana workspace to connect.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/grafana_workspace#subnet_ids GrafanaWorkspace#subnet_ids}
	SubnetIds *[]*string `field:"required" json:"subnetIds" yaml:"subnetIds"`
}

