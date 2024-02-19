package elasticbeanstalkapplication


type ElasticbeanstalkApplicationResourceLifecycleConfig struct {
	// The ARN of an IAM service role that Elastic Beanstalk has permission to assume.
	//
	// The ServiceRole property is required the first time that you provide a ResourceLifecycleConfig for the application. After you provide it once, Elastic Beanstalk persists the Service Role with the application, and you don't need to specify it again. You can, however, specify it in subsequent updates to change the Service Role to another value.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/elasticbeanstalk_application#service_role ElasticbeanstalkApplication#service_role}
	ServiceRole *string `field:"optional" json:"serviceRole" yaml:"serviceRole"`
	// Defines lifecycle settings for application versions.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/elasticbeanstalk_application#version_lifecycle_config ElasticbeanstalkApplication#version_lifecycle_config}
	VersionLifecycleConfig *ElasticbeanstalkApplicationResourceLifecycleConfigVersionLifecycleConfig `field:"optional" json:"versionLifecycleConfig" yaml:"versionLifecycleConfig"`
}

