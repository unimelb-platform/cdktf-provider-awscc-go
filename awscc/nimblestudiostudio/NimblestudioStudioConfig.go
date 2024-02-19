package nimblestudiostudio

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type NimblestudioStudioConfig struct {
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
	// <p>The IAM role that Studio Admins will assume when logging in to the Nimble Studio portal.</p>.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/nimblestudio_studio#admin_role_arn NimblestudioStudio#admin_role_arn}
	AdminRoleArn *string `field:"required" json:"adminRoleArn" yaml:"adminRoleArn"`
	// <p>A friendly name for the studio.</p>.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/nimblestudio_studio#display_name NimblestudioStudio#display_name}
	DisplayName *string `field:"required" json:"displayName" yaml:"displayName"`
	// <p>The studio name that is used in the URL of the Nimble Studio portal when accessed by Nimble Studio users.</p>.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/nimblestudio_studio#studio_name NimblestudioStudio#studio_name}
	StudioName *string `field:"required" json:"studioName" yaml:"studioName"`
	// <p>The IAM role that Studio Users will assume when logging in to the Nimble Studio portal.</p>.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/nimblestudio_studio#user_role_arn NimblestudioStudio#user_role_arn}
	UserRoleArn *string `field:"required" json:"userRoleArn" yaml:"userRoleArn"`
	// <p>Configuration of the encryption method that is used for the studio.</p>.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/nimblestudio_studio#studio_encryption_configuration NimblestudioStudio#studio_encryption_configuration}
	StudioEncryptionConfiguration *NimblestudioStudioStudioEncryptionConfiguration `field:"optional" json:"studioEncryptionConfiguration" yaml:"studioEncryptionConfiguration"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/nimblestudio_studio#tags NimblestudioStudio#tags}.
	Tags *map[string]*string `field:"optional" json:"tags" yaml:"tags"`
}

