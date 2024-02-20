package elasticbeanstalkapplication

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type ElasticbeanstalkApplicationConfig struct {
	// Experimental.
	Connection interface{} `field:"optional" json:"connection" yaml:"connection"`
	// Experimental.
	Count interface{} `field:"optional" json:"count" yaml:"count"`
	// Experimental.
	DependsOn *[]cdktf.ITerraformDependable `field:"optional" json:"dependsOn" yaml:"dependsOn"`
	// Experimental.
	ForEach cdktf.ITerraformIterator `field:"optional" json:"forEach" yaml:"forEach"`
	// Experimental.
	Lifecycle *cdktf.TerraformResourceLifecycle `field:"optional" json:"lifecycle" yaml:"lifecycle"`
	// Experimental.
	Provider cdktf.TerraformProvider `field:"optional" json:"provider" yaml:"provider"`
	// Experimental.
	Provisioners *[]interface{} `field:"optional" json:"provisioners" yaml:"provisioners"`
	// A name for the Elastic Beanstalk application.
	//
	// If you don't specify a name, AWS CloudFormation generates a unique physical ID and uses that ID for the application name.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/elasticbeanstalk_application#application_name ElasticbeanstalkApplication#application_name}
	ApplicationName *string `field:"optional" json:"applicationName" yaml:"applicationName"`
	// Your description of the application.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/elasticbeanstalk_application#description ElasticbeanstalkApplication#description}
	Description *string `field:"optional" json:"description" yaml:"description"`
	// Specifies an application resource lifecycle configuration to prevent your application from accumulating too many versions.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/elasticbeanstalk_application#resource_lifecycle_config ElasticbeanstalkApplication#resource_lifecycle_config}
	ResourceLifecycleConfig *ElasticbeanstalkApplicationResourceLifecycleConfig `field:"optional" json:"resourceLifecycleConfig" yaml:"resourceLifecycleConfig"`
}

