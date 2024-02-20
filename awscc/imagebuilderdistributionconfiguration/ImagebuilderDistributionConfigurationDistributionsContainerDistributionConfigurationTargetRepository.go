package imagebuilderdistributionconfiguration


type ImagebuilderDistributionConfigurationDistributionsContainerDistributionConfigurationTargetRepository struct {
	// The repository name of target container repository.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/imagebuilder_distribution_configuration#repository_name ImagebuilderDistributionConfiguration#repository_name}
	RepositoryName *string `field:"optional" json:"repositoryName" yaml:"repositoryName"`
	// The service of target container repository.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/imagebuilder_distribution_configuration#service ImagebuilderDistributionConfiguration#service}
	Service *string `field:"optional" json:"service" yaml:"service"`
}

