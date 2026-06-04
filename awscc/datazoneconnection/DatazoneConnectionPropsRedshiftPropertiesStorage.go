package datazoneconnection


type DatazoneConnectionPropsRedshiftPropertiesStorage struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.49.0/docs/resources/datazone_connection#cluster_name DatazoneConnection#cluster_name}.
	ClusterName *string `field:"optional" json:"clusterName" yaml:"clusterName"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.49.0/docs/resources/datazone_connection#workgroup_name DatazoneConnection#workgroup_name}.
	WorkgroupName *string `field:"optional" json:"workgroupName" yaml:"workgroupName"`
}

