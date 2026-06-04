package datazoneconnection


type DatazoneConnectionPropsGlueProperties struct {
	// Glue Connection Input.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.49.0/docs/resources/datazone_connection#glue_connection_input DatazoneConnection#glue_connection_input}
	GlueConnectionInput *DatazoneConnectionPropsGluePropertiesGlueConnectionInput `field:"optional" json:"glueConnectionInput" yaml:"glueConnectionInput"`
}

