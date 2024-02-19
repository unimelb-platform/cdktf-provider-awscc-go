package datasynclocationnfs


type DatasyncLocationNfsOnPremConfig struct {
	// ARN(s) of the agent(s) to use for an NFS location.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/datasync_location_nfs#agent_arns DatasyncLocationNfs#agent_arns}
	AgentArns *[]*string `field:"required" json:"agentArns" yaml:"agentArns"`
}

