package datazoneconnection


type DatazoneConnectionProps struct {
	// Athena Properties Input.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.49.0/docs/resources/datazone_connection#athena_properties DatazoneConnection#athena_properties}
	AthenaProperties *DatazoneConnectionPropsAthenaProperties `field:"optional" json:"athenaProperties" yaml:"athenaProperties"`
	// Glue Properties Input.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.49.0/docs/resources/datazone_connection#glue_properties DatazoneConnection#glue_properties}
	GlueProperties *DatazoneConnectionPropsGlueProperties `field:"optional" json:"glueProperties" yaml:"glueProperties"`
	// HyperPod Properties Input.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.49.0/docs/resources/datazone_connection#hyper_pod_properties DatazoneConnection#hyper_pod_properties}
	HyperPodProperties *DatazoneConnectionPropsHyperPodProperties `field:"optional" json:"hyperPodProperties" yaml:"hyperPodProperties"`
	// IAM Properties Input.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.49.0/docs/resources/datazone_connection#iam_properties DatazoneConnection#iam_properties}
	IamProperties *DatazoneConnectionPropsIamProperties `field:"optional" json:"iamProperties" yaml:"iamProperties"`
	// Redshift Properties Input.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.49.0/docs/resources/datazone_connection#redshift_properties DatazoneConnection#redshift_properties}
	RedshiftProperties *DatazoneConnectionPropsRedshiftProperties `field:"optional" json:"redshiftProperties" yaml:"redshiftProperties"`
	// Spark EMR Properties Input.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.49.0/docs/resources/datazone_connection#spark_emr_properties DatazoneConnection#spark_emr_properties}
	SparkEmrProperties *DatazoneConnectionPropsSparkEmrProperties `field:"optional" json:"sparkEmrProperties" yaml:"sparkEmrProperties"`
	// Spark Glue Properties Input.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.49.0/docs/resources/datazone_connection#spark_glue_properties DatazoneConnection#spark_glue_properties}
	SparkGlueProperties *DatazoneConnectionPropsSparkGlueProperties `field:"optional" json:"sparkGlueProperties" yaml:"sparkGlueProperties"`
}

