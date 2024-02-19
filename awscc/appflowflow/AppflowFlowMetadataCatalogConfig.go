package appflowflow


type AppflowFlowMetadataCatalogConfig struct {
	// Configurations of glue data catalog of the flow.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/appflow_flow#glue_data_catalog AppflowFlow#glue_data_catalog}
	GlueDataCatalog *AppflowFlowMetadataCatalogConfigGlueDataCatalog `field:"optional" json:"glueDataCatalog" yaml:"glueDataCatalog"`
}

