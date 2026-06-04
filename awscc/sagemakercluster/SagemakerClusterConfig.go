package sagemakercluster

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type SagemakerClusterConfig struct {
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
	// The instance groups of the SageMaker HyperPod cluster.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.49.0/docs/resources/sagemaker_cluster#instance_groups SagemakerCluster#instance_groups}
	InstanceGroups interface{} `field:"required" json:"instanceGroups" yaml:"instanceGroups"`
	// The name of the HyperPod Cluster.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.49.0/docs/resources/sagemaker_cluster#cluster_name SagemakerCluster#cluster_name}
	ClusterName *string `field:"optional" json:"clusterName" yaml:"clusterName"`
	// If node auto-recovery is set to true, faulty nodes will be replaced or rebooted when a failure is detected.
	//
	// If set to false, nodes will be labelled when a fault is detected.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.49.0/docs/resources/sagemaker_cluster#node_recovery SagemakerCluster#node_recovery}
	NodeRecovery *string `field:"optional" json:"nodeRecovery" yaml:"nodeRecovery"`
	// Specifies parameter(s) specific to the orchestrator, e.g. specify the EKS cluster.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.49.0/docs/resources/sagemaker_cluster#orchestrator SagemakerCluster#orchestrator}
	Orchestrator *SagemakerClusterOrchestrator `field:"optional" json:"orchestrator" yaml:"orchestrator"`
	// Custom tags for managing the SageMaker HyperPod cluster as an AWS resource.
	//
	// You can add tags to your cluster in the same way you add them in other AWS services that support tagging.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.49.0/docs/resources/sagemaker_cluster#tags SagemakerCluster#tags}
	Tags interface{} `field:"optional" json:"tags" yaml:"tags"`
	// Specifies an Amazon Virtual Private Cloud (VPC) that your SageMaker jobs, hosted models, and compute resources have access to.
	//
	// You can control access to and from your resources by configuring a VPC.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.49.0/docs/resources/sagemaker_cluster#vpc_config SagemakerCluster#vpc_config}
	VpcConfig *SagemakerClusterVpcConfig `field:"optional" json:"vpcConfig" yaml:"vpcConfig"`
}

