package sagemakerinferenceexperiment


type SagemakerInferenceExperimentDataStorageConfigContentType struct {
	// The list of all content type headers that SageMaker will treat as CSV and capture accordingly.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/sagemaker_inference_experiment#csv_content_types SagemakerInferenceExperiment#csv_content_types}
	CsvContentTypes *[]*string `field:"optional" json:"csvContentTypes" yaml:"csvContentTypes"`
	// The list of all content type headers that SageMaker will treat as JSON and capture accordingly.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/sagemaker_inference_experiment#json_content_types SagemakerInferenceExperiment#json_content_types}
	JsonContentTypes *[]*string `field:"optional" json:"jsonContentTypes" yaml:"jsonContentTypes"`
}

