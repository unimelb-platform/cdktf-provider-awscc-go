//go:build no_runtime_type_checking

package ec2instance

// Building without runtime type checking enabled, so all the below just return nil

func (e *jsiiProxy_Ec2InstanceVolumesList) validateAllWithMapKeyParameters(mapKeyAttributeName *string) error {
	return nil
}

func (e *jsiiProxy_Ec2InstanceVolumesList) validateGetParameters(index *float64) error {
	return nil
}

func (e *jsiiProxy_Ec2InstanceVolumesList) validateResolveParameters(_context cdktf.IResolveContext) error {
	return nil
}

func (j *jsiiProxy_Ec2InstanceVolumesList) validateSetInternalValueParameters(val interface{}) error {
	return nil
}

func (j *jsiiProxy_Ec2InstanceVolumesList) validateSetTerraformAttributeParameters(val *string) error {
	return nil
}

func (j *jsiiProxy_Ec2InstanceVolumesList) validateSetTerraformResourceParameters(val cdktf.IInterpolatingParent) error {
	return nil
}

func (j *jsiiProxy_Ec2InstanceVolumesList) validateSetWrapsSetParameters(val *bool) error {
	return nil
}

func validateNewEc2InstanceVolumesListParameters(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, wrapsSet *bool) error {
	return nil
}

