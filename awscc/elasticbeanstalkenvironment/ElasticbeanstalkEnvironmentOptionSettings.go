package elasticbeanstalkenvironment


type ElasticbeanstalkEnvironmentOptionSettings struct {
	// A unique namespace that identifies the option's associated AWS resource.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/elasticbeanstalk_environment#namespace ElasticbeanstalkEnvironment#namespace}
	Namespace *string `field:"required" json:"namespace" yaml:"namespace"`
	// The name of the configuration option.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/elasticbeanstalk_environment#option_name ElasticbeanstalkEnvironment#option_name}
	OptionName *string `field:"required" json:"optionName" yaml:"optionName"`
	// A unique resource name for the option setting. Use it for a time–based scaling configuration option.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/elasticbeanstalk_environment#resource_name ElasticbeanstalkEnvironment#resource_name}
	ResourceName *string `field:"optional" json:"resourceName" yaml:"resourceName"`
	// The current value for the configuration option.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/elasticbeanstalk_environment#value ElasticbeanstalkEnvironment#value}
	Value *string `field:"optional" json:"value" yaml:"value"`
}

