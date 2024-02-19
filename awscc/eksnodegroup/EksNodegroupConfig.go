package eksnodegroup

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type EksNodegroupConfig struct {
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
	// Name of the cluster to create the node group in.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/eks_nodegroup#cluster_name EksNodegroup#cluster_name}
	ClusterName *string `field:"required" json:"clusterName" yaml:"clusterName"`
	// The Amazon Resource Name (ARN) of the IAM role to associate with your node group.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/eks_nodegroup#node_role EksNodegroup#node_role}
	NodeRole *string `field:"required" json:"nodeRole" yaml:"nodeRole"`
	// The subnets to use for the Auto Scaling group that is created for your node group.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/eks_nodegroup#subnets EksNodegroup#subnets}
	Subnets *[]*string `field:"required" json:"subnets" yaml:"subnets"`
	// The AMI type for your node group.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/eks_nodegroup#ami_type EksNodegroup#ami_type}
	AmiType *string `field:"optional" json:"amiType" yaml:"amiType"`
	// The capacity type of your managed node group.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/eks_nodegroup#capacity_type EksNodegroup#capacity_type}
	CapacityType *string `field:"optional" json:"capacityType" yaml:"capacityType"`
	// The root device disk size (in GiB) for your node group instances.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/eks_nodegroup#disk_size EksNodegroup#disk_size}
	DiskSize *float64 `field:"optional" json:"diskSize" yaml:"diskSize"`
	// Force the update if the existing node group's pods are unable to be drained due to a pod disruption budget issue.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/eks_nodegroup#force_update_enabled EksNodegroup#force_update_enabled}
	ForceUpdateEnabled interface{} `field:"optional" json:"forceUpdateEnabled" yaml:"forceUpdateEnabled"`
	// Specify the instance types for a node group.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/eks_nodegroup#instance_types EksNodegroup#instance_types}
	InstanceTypes *[]*string `field:"optional" json:"instanceTypes" yaml:"instanceTypes"`
	// The Kubernetes labels to be applied to the nodes in the node group when they are created.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/eks_nodegroup#labels EksNodegroup#labels}
	Labels *map[string]*string `field:"optional" json:"labels" yaml:"labels"`
	// An object representing a node group's launch template specification.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/eks_nodegroup#launch_template EksNodegroup#launch_template}
	LaunchTemplate *EksNodegroupLaunchTemplate `field:"optional" json:"launchTemplate" yaml:"launchTemplate"`
	// The unique name to give your node group.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/eks_nodegroup#nodegroup_name EksNodegroup#nodegroup_name}
	NodegroupName *string `field:"optional" json:"nodegroupName" yaml:"nodegroupName"`
	// The AMI version of the Amazon EKS-optimized AMI to use with your node group.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/eks_nodegroup#release_version EksNodegroup#release_version}
	ReleaseVersion *string `field:"optional" json:"releaseVersion" yaml:"releaseVersion"`
	// The remote access (SSH) configuration to use with your node group.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/eks_nodegroup#remote_access EksNodegroup#remote_access}
	RemoteAccess *EksNodegroupRemoteAccess `field:"optional" json:"remoteAccess" yaml:"remoteAccess"`
	// The scaling configuration details for the Auto Scaling group that is created for your node group.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/eks_nodegroup#scaling_config EksNodegroup#scaling_config}
	ScalingConfig *EksNodegroupScalingConfig `field:"optional" json:"scalingConfig" yaml:"scalingConfig"`
	// The metadata, as key-value pairs, to apply to the node group to assist with categorization and organization.
	//
	// Follows same schema as Labels for consistency.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/eks_nodegroup#tags EksNodegroup#tags}
	Tags *map[string]*string `field:"optional" json:"tags" yaml:"tags"`
	// The Kubernetes taints to be applied to the nodes in the node group when they are created.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/eks_nodegroup#taints EksNodegroup#taints}
	Taints interface{} `field:"optional" json:"taints" yaml:"taints"`
	// The node group update configuration.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/eks_nodegroup#update_config EksNodegroup#update_config}
	UpdateConfig *EksNodegroupUpdateConfig `field:"optional" json:"updateConfig" yaml:"updateConfig"`
	// The Kubernetes version to use for your managed nodes.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/eks_nodegroup#version EksNodegroup#version}
	Version *string `field:"optional" json:"version" yaml:"version"`
}

