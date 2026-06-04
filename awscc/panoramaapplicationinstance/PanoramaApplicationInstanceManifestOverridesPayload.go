package panoramaapplicationinstance


type PanoramaApplicationInstanceManifestOverridesPayload struct {
	// The overrides document.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.49.0/docs/resources/panorama_application_instance#payload_data PanoramaApplicationInstance#payload_data}
	PayloadData *string `field:"optional" json:"payloadData" yaml:"payloadData"`
}

