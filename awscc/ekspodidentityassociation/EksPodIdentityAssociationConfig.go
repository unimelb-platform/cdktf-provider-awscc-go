package ekspodidentityassociation

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type EksPodIdentityAssociationConfig struct {
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
	// The cluster that the pod identity association is created for.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/eks_pod_identity_association#cluster_name EksPodIdentityAssociation#cluster_name}
	ClusterName *string `field:"required" json:"clusterName" yaml:"clusterName"`
	// The Kubernetes namespace that the pod identity association is created for.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/eks_pod_identity_association#namespace EksPodIdentityAssociation#namespace}
	Namespace *string `field:"required" json:"namespace" yaml:"namespace"`
	// The IAM role ARN that the pod identity association is created for.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/eks_pod_identity_association#role_arn EksPodIdentityAssociation#role_arn}
	RoleArn *string `field:"required" json:"roleArn" yaml:"roleArn"`
	// The Kubernetes service account that the pod identity association is created for.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/eks_pod_identity_association#service_account EksPodIdentityAssociation#service_account}
	ServiceAccount *string `field:"required" json:"serviceAccount" yaml:"serviceAccount"`
	// An array of key-value pairs to apply to this resource.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/eks_pod_identity_association#tags EksPodIdentityAssociation#tags}
	Tags interface{} `field:"optional" json:"tags" yaml:"tags"`
}

