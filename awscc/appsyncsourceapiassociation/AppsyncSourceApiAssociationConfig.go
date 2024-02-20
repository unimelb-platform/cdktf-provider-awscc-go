package appsyncsourceapiassociation

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type AppsyncSourceApiAssociationConfig struct {
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
	// Description of the SourceApiAssociation.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/appsync_source_api_association#description AppsyncSourceApiAssociation#description}
	Description *string `field:"optional" json:"description" yaml:"description"`
	// Identifier of the Merged GraphQLApi to associate. It could be either GraphQLApi ApiId or ARN.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/appsync_source_api_association#merged_api_identifier AppsyncSourceApiAssociation#merged_api_identifier}
	MergedApiIdentifier *string `field:"optional" json:"mergedApiIdentifier" yaml:"mergedApiIdentifier"`
	// Customized configuration for SourceApiAssociation.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/appsync_source_api_association#source_api_association_config AppsyncSourceApiAssociation#source_api_association_config}
	SourceApiAssociationConfig *AppsyncSourceApiAssociationSourceApiAssociationConfig `field:"optional" json:"sourceApiAssociationConfig" yaml:"sourceApiAssociationConfig"`
	// Identifier of the Source GraphQLApi to associate. It could be either GraphQLApi ApiId or ARN.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/appsync_source_api_association#source_api_identifier AppsyncSourceApiAssociation#source_api_identifier}
	SourceApiIdentifier *string `field:"optional" json:"sourceApiIdentifier" yaml:"sourceApiIdentifier"`
}

