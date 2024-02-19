package sagemakerdevice


type SagemakerDeviceDevice struct {
	// The name of the device.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/sagemaker_device#device_name SagemakerDevice#device_name}
	DeviceName *string `field:"required" json:"deviceName" yaml:"deviceName"`
	// Description of the device.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/sagemaker_device#description SagemakerDevice#description}
	Description *string `field:"optional" json:"description" yaml:"description"`
	// AWS Internet of Things (IoT) object name.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/sagemaker_device#iot_thing_name SagemakerDevice#iot_thing_name}
	IotThingName *string `field:"optional" json:"iotThingName" yaml:"iotThingName"`
}

