package emrserverlessapplication


type EmrserverlessApplicationWorkerTypeSpecifications struct {
	// The image configuration.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/emrserverless_application#image_configuration EmrserverlessApplication#image_configuration}
	ImageConfiguration *EmrserverlessApplicationWorkerTypeSpecificationsImageConfiguration `field:"optional" json:"imageConfiguration" yaml:"imageConfiguration"`
}

