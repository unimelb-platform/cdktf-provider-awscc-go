package emrserverlessapplication


type EmrserverlessApplicationWorkerTypeSpecificationsImageConfiguration struct {
	// The URI of an image in the Amazon ECR registry.
	//
	// This field is required when you create a new application. If you leave this field blank in an update, Amazon EMR will remove the image configuration.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/emrserverless_application#image_uri EmrserverlessApplication#image_uri}
	ImageUri *string `field:"optional" json:"imageUri" yaml:"imageUri"`
}

