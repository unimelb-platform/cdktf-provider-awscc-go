package iotcommand


type IotCommandPayload struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.49.0/docs/resources/iot_command#content IotCommand#content}.
	Content *string `field:"optional" json:"content" yaml:"content"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.49.0/docs/resources/iot_command#content_type IotCommand#content_type}.
	ContentType *string `field:"optional" json:"contentType" yaml:"contentType"`
}

