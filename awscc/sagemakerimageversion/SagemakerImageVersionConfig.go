package sagemakerimageversion

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type SagemakerImageVersionConfig struct {
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
	// The registry path of the container image on which this image version is based.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/sagemaker_image_version#base_image SagemakerImageVersion#base_image}
	BaseImage *string `field:"required" json:"baseImage" yaml:"baseImage"`
	// The name of the image this version belongs to.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/sagemaker_image_version#image_name SagemakerImageVersion#image_name}
	ImageName *string `field:"required" json:"imageName" yaml:"imageName"`
	// The alias of the image version.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/sagemaker_image_version#alias SagemakerImageVersion#alias}
	Alias *string `field:"optional" json:"alias" yaml:"alias"`
	// List of aliases for the image version.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/sagemaker_image_version#aliases SagemakerImageVersion#aliases}
	Aliases *[]*string `field:"optional" json:"aliases" yaml:"aliases"`
	// Indicates Horovod compatibility.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/sagemaker_image_version#horovod SagemakerImageVersion#horovod}
	Horovod interface{} `field:"optional" json:"horovod" yaml:"horovod"`
	// Indicates SageMaker job type compatibility.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/sagemaker_image_version#job_type SagemakerImageVersion#job_type}
	JobType *string `field:"optional" json:"jobType" yaml:"jobType"`
	// The machine learning framework vended in the image version.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/sagemaker_image_version#ml_framework SagemakerImageVersion#ml_framework}
	MlFramework *string `field:"optional" json:"mlFramework" yaml:"mlFramework"`
	// Indicates CPU or GPU compatibility.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/sagemaker_image_version#processor SagemakerImageVersion#processor}
	Processor *string `field:"optional" json:"processor" yaml:"processor"`
	// The supported programming language and its version.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/sagemaker_image_version#programming_lang SagemakerImageVersion#programming_lang}
	ProgrammingLang *string `field:"optional" json:"programmingLang" yaml:"programmingLang"`
	// The maintainer description of the image version.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/sagemaker_image_version#release_notes SagemakerImageVersion#release_notes}
	ReleaseNotes *string `field:"optional" json:"releaseNotes" yaml:"releaseNotes"`
	// The availability of the image version specified by the maintainer.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/sagemaker_image_version#vendor_guidance SagemakerImageVersion#vendor_guidance}
	VendorGuidance *string `field:"optional" json:"vendorGuidance" yaml:"vendorGuidance"`
}

