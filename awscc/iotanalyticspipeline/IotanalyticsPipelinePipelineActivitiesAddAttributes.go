package iotanalyticspipeline


type IotanalyticsPipelinePipelineActivitiesAddAttributes struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.49.0/docs/resources/iotanalytics_pipeline#attributes IotanalyticsPipeline#attributes}.
	Attributes *map[string]*string `field:"optional" json:"attributes" yaml:"attributes"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.49.0/docs/resources/iotanalytics_pipeline#name IotanalyticsPipeline#name}.
	Name *string `field:"optional" json:"name" yaml:"name"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.49.0/docs/resources/iotanalytics_pipeline#next IotanalyticsPipeline#next}.
	Next *string `field:"optional" json:"next" yaml:"next"`
}

