package sagemakercluster


type SagemakerClusterInstanceGroupsInstanceStorageConfigsEbsVolumeConfig struct {
	// The size in gigabytes (GB) of the additional EBS volume to be attached to the instances in the SageMaker HyperPod cluster instance group.
	//
	// The additional EBS volume is attached to each instance within the SageMaker HyperPod cluster instance group and mounted to /opt/sagemaker.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.49.0/docs/resources/sagemaker_cluster#volume_size_in_gb SagemakerCluster#volume_size_in_gb}
	VolumeSizeInGb *float64 `field:"optional" json:"volumeSizeInGb" yaml:"volumeSizeInGb"`
}

