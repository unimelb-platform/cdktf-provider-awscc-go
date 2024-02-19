package elasticloadbalancingv2truststorerevocation

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type Elasticloadbalancingv2TrustStoreRevocationConfig struct {
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
	// The attributes required to create a trust store revocation.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/elasticloadbalancingv2_trust_store_revocation#revocation_contents Elasticloadbalancingv2TrustStoreRevocation#revocation_contents}
	RevocationContents interface{} `field:"optional" json:"revocationContents" yaml:"revocationContents"`
	// The Amazon Resource Name (ARN) of the trust store.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/elasticloadbalancingv2_trust_store_revocation#trust_store_arn Elasticloadbalancingv2TrustStoreRevocation#trust_store_arn}
	TrustStoreArn *string `field:"optional" json:"trustStoreArn" yaml:"trustStoreArn"`
}

