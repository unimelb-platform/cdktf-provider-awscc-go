package ivsrecordingconfiguration


type IvsRecordingConfigurationRenditionConfiguration struct {
	// Renditions indicates which renditions are recorded for a stream.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/ivs_recording_configuration#renditions IvsRecordingConfiguration#renditions}
	Renditions *[]*string `field:"optional" json:"renditions" yaml:"renditions"`
	// Resolution Selection indicates which set of renditions are recorded for a stream.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/ivs_recording_configuration#rendition_selection IvsRecordingConfiguration#rendition_selection}
	RenditionSelection *string `field:"optional" json:"renditionSelection" yaml:"renditionSelection"`
}

