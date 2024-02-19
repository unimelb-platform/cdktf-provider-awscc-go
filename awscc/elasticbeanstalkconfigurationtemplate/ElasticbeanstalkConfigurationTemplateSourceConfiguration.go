package elasticbeanstalkconfigurationtemplate


type ElasticbeanstalkConfigurationTemplateSourceConfiguration struct {
	// The name of the application associated with the configuration.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/elasticbeanstalk_configuration_template#application_name ElasticbeanstalkConfigurationTemplate#application_name}
	ApplicationName *string `field:"required" json:"applicationName" yaml:"applicationName"`
	// The name of the configuration template.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/elasticbeanstalk_configuration_template#template_name ElasticbeanstalkConfigurationTemplate#template_name}
	TemplateName *string `field:"required" json:"templateName" yaml:"templateName"`
}

