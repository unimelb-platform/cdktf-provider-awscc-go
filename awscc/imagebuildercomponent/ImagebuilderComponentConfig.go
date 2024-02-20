package imagebuildercomponent

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type ImagebuilderComponentConfig struct {
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
	// The name of the component.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/imagebuilder_component#name ImagebuilderComponent#name}
	Name *string `field:"required" json:"name" yaml:"name"`
	// The platform of the component.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/imagebuilder_component#platform ImagebuilderComponent#platform}
	Platform *string `field:"required" json:"platform" yaml:"platform"`
	// The version of the component.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/imagebuilder_component#version ImagebuilderComponent#version}
	Version *string `field:"required" json:"version" yaml:"version"`
	// The change description of the component.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/imagebuilder_component#change_description ImagebuilderComponent#change_description}
	ChangeDescription *string `field:"optional" json:"changeDescription" yaml:"changeDescription"`
	// The data of the component.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/imagebuilder_component#data ImagebuilderComponent#data}
	Data *string `field:"optional" json:"data" yaml:"data"`
	// The description of the component.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/imagebuilder_component#description ImagebuilderComponent#description}
	Description *string `field:"optional" json:"description" yaml:"description"`
	// The KMS key identifier used to encrypt the component.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/imagebuilder_component#kms_key_id ImagebuilderComponent#kms_key_id}
	KmsKeyId *string `field:"optional" json:"kmsKeyId" yaml:"kmsKeyId"`
	// The operating system (OS) version supported by the component.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/imagebuilder_component#supported_os_versions ImagebuilderComponent#supported_os_versions}
	SupportedOsVersions *[]*string `field:"optional" json:"supportedOsVersions" yaml:"supportedOsVersions"`
	// The tags associated with the component.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/imagebuilder_component#tags ImagebuilderComponent#tags}
	Tags *map[string]*string `field:"optional" json:"tags" yaml:"tags"`
	// The uri of the component.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/imagebuilder_component#uri ImagebuilderComponent#uri}
	Uri *string `field:"optional" json:"uri" yaml:"uri"`
}

