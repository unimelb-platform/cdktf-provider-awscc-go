package elasticbeanstalkapplicationversion

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type ElasticbeanstalkApplicationVersionConfig struct {
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
	// The name of the Elastic Beanstalk application that is associated with this application version.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/elasticbeanstalk_application_version#application_name ElasticbeanstalkApplicationVersion#application_name}
	ApplicationName *string `field:"required" json:"applicationName" yaml:"applicationName"`
	// The Amazon S3 bucket and key that identify the location of the source bundle for this version.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/elasticbeanstalk_application_version#source_bundle ElasticbeanstalkApplicationVersion#source_bundle}
	SourceBundle *ElasticbeanstalkApplicationVersionSourceBundle `field:"required" json:"sourceBundle" yaml:"sourceBundle"`
	// A description of this application version.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/elasticbeanstalk_application_version#description ElasticbeanstalkApplicationVersion#description}
	Description *string `field:"optional" json:"description" yaml:"description"`
}

