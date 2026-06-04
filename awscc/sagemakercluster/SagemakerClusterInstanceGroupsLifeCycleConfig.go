package sagemakercluster


type SagemakerClusterInstanceGroupsLifeCycleConfig struct {
	// The file name of the entrypoint script of lifecycle scripts under SourceS3Uri. This entrypoint script runs during cluster creation.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.49.0/docs/resources/sagemaker_cluster#on_create SagemakerCluster#on_create}
	OnCreate *string `field:"required" json:"onCreate" yaml:"onCreate"`
	// An Amazon S3 bucket path where your lifecycle scripts are stored.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.49.0/docs/resources/sagemaker_cluster#source_s3_uri SagemakerCluster#source_s3_uri}
	SourceS3Uri *string `field:"required" json:"sourceS3Uri" yaml:"sourceS3Uri"`
}

