package route53profilesprofileresourceassociation

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type Route53ProfilesProfileResourceAssociationConfig struct {
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
	// The name of an association between the  Profile and resource.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.49.0/docs/resources/route53profiles_profile_resource_association#name Route53ProfilesProfileResourceAssociation#name}
	Name *string `field:"required" json:"name" yaml:"name"`
	// The ID of the  profile that you associated the resource to that is specified by ResourceArn.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.49.0/docs/resources/route53profiles_profile_resource_association#profile_id Route53ProfilesProfileResourceAssociation#profile_id}
	ProfileId *string `field:"required" json:"profileId" yaml:"profileId"`
	// The arn of the resource that you associated to the  Profile.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.49.0/docs/resources/route53profiles_profile_resource_association#resource_arn Route53ProfilesProfileResourceAssociation#resource_arn}
	ResourceArn *string `field:"required" json:"resourceArn" yaml:"resourceArn"`
	// A JSON-formatted string with key-value pairs specifying the properties of the associated resource.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.49.0/docs/resources/route53profiles_profile_resource_association#resource_properties Route53ProfilesProfileResourceAssociation#resource_properties}
	ResourceProperties *string `field:"optional" json:"resourceProperties" yaml:"resourceProperties"`
}

