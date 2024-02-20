package sagemakermodelpackage


type SagemakerModelPackageSourceAlgorithmSpecification struct {
	// A list of algorithms that were used to create a model package.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/sagemaker_model_package#source_algorithms SagemakerModelPackage#source_algorithms}
	SourceAlgorithms interface{} `field:"required" json:"sourceAlgorithms" yaml:"sourceAlgorithms"`
}

