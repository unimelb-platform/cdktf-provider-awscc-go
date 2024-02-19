package eksnodegroup


type EksNodegroupRemoteAccess struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/eks_nodegroup#ec_2_ssh_key EksNodegroup#ec_2_ssh_key}.
	Ec2SshKey *string `field:"required" json:"ec2SshKey" yaml:"ec2SshKey"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/eks_nodegroup#source_security_groups EksNodegroup#source_security_groups}.
	SourceSecurityGroups *[]*string `field:"optional" json:"sourceSecurityGroups" yaml:"sourceSecurityGroups"`
}

