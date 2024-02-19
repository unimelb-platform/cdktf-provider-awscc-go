package sagemakermodelpackage


type SagemakerModelPackageDriftCheckBaselinesExplainability struct {
	// Represents a File Source Object.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/sagemaker_model_package#config_file SagemakerModelPackage#config_file}
	ConfigFile *SagemakerModelPackageDriftCheckBaselinesExplainabilityConfigFile `field:"optional" json:"configFile" yaml:"configFile"`
	// Represents a Metric Source Object.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/sagemaker_model_package#constraints SagemakerModelPackage#constraints}
	Constraints *SagemakerModelPackageDriftCheckBaselinesExplainabilityConstraints `field:"optional" json:"constraints" yaml:"constraints"`
}

