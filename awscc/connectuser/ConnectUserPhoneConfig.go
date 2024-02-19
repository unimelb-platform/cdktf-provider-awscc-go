package connectuser


type ConnectUserPhoneConfig struct {
	// The phone type.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/connect_user#phone_type ConnectUser#phone_type}
	PhoneType *string `field:"required" json:"phoneType" yaml:"phoneType"`
	// The After Call Work (ACW) timeout setting, in seconds.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/connect_user#after_contact_work_time_limit ConnectUser#after_contact_work_time_limit}
	AfterContactWorkTimeLimit *float64 `field:"optional" json:"afterContactWorkTimeLimit" yaml:"afterContactWorkTimeLimit"`
	// The Auto accept setting.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/connect_user#auto_accept ConnectUser#auto_accept}
	AutoAccept interface{} `field:"optional" json:"autoAccept" yaml:"autoAccept"`
	// The phone number for the user's desk phone.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/connect_user#desk_phone_number ConnectUser#desk_phone_number}
	DeskPhoneNumber *string `field:"optional" json:"deskPhoneNumber" yaml:"deskPhoneNumber"`
}

