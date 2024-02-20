package redshiftendpointauthorization

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type RedshiftEndpointAuthorizationConfig struct {
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
	// The target AWS account ID to grant or revoke access for.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/redshift_endpoint_authorization#account RedshiftEndpointAuthorization#account}
	Account *string `field:"required" json:"account" yaml:"account"`
	// The cluster identifier.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/redshift_endpoint_authorization#cluster_identifier RedshiftEndpointAuthorization#cluster_identifier}
	ClusterIdentifier *string `field:"required" json:"clusterIdentifier" yaml:"clusterIdentifier"`
	// Indicates whether to force the revoke action.
	//
	// If true, the Redshift-managed VPC endpoints associated with the endpoint authorization are also deleted.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/redshift_endpoint_authorization#force RedshiftEndpointAuthorization#force}
	Force interface{} `field:"optional" json:"force" yaml:"force"`
	// The virtual private cloud (VPC) identifiers to grant or revoke access to.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/redshift_endpoint_authorization#vpc_ids RedshiftEndpointAuthorization#vpc_ids}
	VpcIds *[]*string `field:"optional" json:"vpcIds" yaml:"vpcIds"`
}

