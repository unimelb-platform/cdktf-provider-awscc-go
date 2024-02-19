package kendrafaq

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type KendraFaqConfig struct {
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
	// Index ID.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/kendra_faq#index_id KendraFaq#index_id}
	IndexId *string `field:"required" json:"indexId" yaml:"indexId"`
	// FAQ name.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/kendra_faq#name KendraFaq#name}
	Name *string `field:"required" json:"name" yaml:"name"`
	// FAQ role ARN.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/kendra_faq#role_arn KendraFaq#role_arn}
	RoleArn *string `field:"required" json:"roleArn" yaml:"roleArn"`
	// FAQ S3 path.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/kendra_faq#s3_path KendraFaq#s3_path}
	S3Path *KendraFaqS3Path `field:"required" json:"s3Path" yaml:"s3Path"`
	// FAQ description.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/kendra_faq#description KendraFaq#description}
	Description *string `field:"optional" json:"description" yaml:"description"`
	// FAQ file format.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/kendra_faq#file_format KendraFaq#file_format}
	FileFormat *string `field:"optional" json:"fileFormat" yaml:"fileFormat"`
	// Tags for labeling the FAQ.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/kendra_faq#tags KendraFaq#tags}
	Tags interface{} `field:"optional" json:"tags" yaml:"tags"`
}

