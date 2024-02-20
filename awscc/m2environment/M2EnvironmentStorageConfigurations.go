package m2environment


type M2EnvironmentStorageConfigurations struct {
	// Defines the storage configuration for an Amazon EFS file system.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/m2_environment#efs M2Environment#efs}
	Efs *M2EnvironmentStorageConfigurationsEfs `field:"optional" json:"efs" yaml:"efs"`
	// Defines the storage configuration for an Amazon FSx file system.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/m2_environment#fsx M2Environment#fsx}
	Fsx *M2EnvironmentStorageConfigurationsFsx `field:"optional" json:"fsx" yaml:"fsx"`
}

