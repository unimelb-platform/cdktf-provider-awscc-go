package sagemakerdomain


type SagemakerDomainDomainSettings struct {
	// A collection of settings that are required to start docker-proxy server.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/sagemaker_domain#docker_settings SagemakerDomain#docker_settings}
	DockerSettings *SagemakerDomainDomainSettingsDockerSettings `field:"optional" json:"dockerSettings" yaml:"dockerSettings"`
	// A collection of settings that update the current configuration for the RStudioServerPro Domain-level app.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/sagemaker_domain#r_studio_server_pro_domain_settings SagemakerDomain#r_studio_server_pro_domain_settings}
	RStudioServerProDomainSettings *SagemakerDomainDomainSettingsRStudioServerProDomainSettings `field:"optional" json:"rStudioServerProDomainSettings" yaml:"rStudioServerProDomainSettings"`
	// The security groups for the Amazon Virtual Private Cloud that the Domain uses for communication between Domain-level apps and user apps.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/sagemaker_domain#security_group_ids SagemakerDomain#security_group_ids}
	SecurityGroupIds *[]*string `field:"optional" json:"securityGroupIds" yaml:"securityGroupIds"`
}

