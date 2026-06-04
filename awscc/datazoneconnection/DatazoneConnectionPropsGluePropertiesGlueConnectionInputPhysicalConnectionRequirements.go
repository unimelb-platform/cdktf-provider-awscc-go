package datazoneconnection


type DatazoneConnectionPropsGluePropertiesGlueConnectionInputPhysicalConnectionRequirements struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.49.0/docs/resources/datazone_connection#availability_zone DatazoneConnection#availability_zone}.
	AvailabilityZone *string `field:"optional" json:"availabilityZone" yaml:"availabilityZone"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.49.0/docs/resources/datazone_connection#security_group_id_list DatazoneConnection#security_group_id_list}.
	SecurityGroupIdList *[]*string `field:"optional" json:"securityGroupIdList" yaml:"securityGroupIdList"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.49.0/docs/resources/datazone_connection#subnet_id DatazoneConnection#subnet_id}.
	SubnetId *string `field:"optional" json:"subnetId" yaml:"subnetId"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.49.0/docs/resources/datazone_connection#subnet_id_list DatazoneConnection#subnet_id_list}.
	SubnetIdList *[]*string `field:"optional" json:"subnetIdList" yaml:"subnetIdList"`
}

