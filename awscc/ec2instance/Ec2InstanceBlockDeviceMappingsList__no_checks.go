//go:build no_runtime_type_checking

package ec2instance

// Building without runtime type checking enabled, so all the below just return nil

func (e *jsiiProxy_Ec2InstanceBlockDeviceMappingsList) validateGetParameters(index *float64) error {
	return nil
}

func (e *jsiiProxy_Ec2InstanceBlockDeviceMappingsList) validateResolveParameters(_context cdktf.IResolveContext) error {
	return nil
}

func (j *jsiiProxy_Ec2InstanceBlockDeviceMappingsList) validateSetInternalValueParameters(val interface{}) error {
	return nil
}

func (j *jsiiProxy_Ec2InstanceBlockDeviceMappingsList) validateSetTerraformAttributeParameters(val *string) error {
	return nil
}

func (j *jsiiProxy_Ec2InstanceBlockDeviceMappingsList) validateSetTerraformResourceParameters(val cdktf.IInterpolatingParent) error {
	return nil
}

func (j *jsiiProxy_Ec2InstanceBlockDeviceMappingsList) validateSetWrapsSetParameters(val *bool) error {
	return nil
}

func validateNewEc2InstanceBlockDeviceMappingsListParameters(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, wrapsSet *bool) error {
	return nil
}

