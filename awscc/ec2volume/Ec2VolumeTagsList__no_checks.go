//go:build no_runtime_type_checking

package ec2volume

// Building without runtime type checking enabled, so all the below just return nil

func (e *jsiiProxy_Ec2VolumeTagsList) validateGetParameters(index *float64) error {
	return nil
}

func (e *jsiiProxy_Ec2VolumeTagsList) validateResolveParameters(_context cdktf.IResolveContext) error {
	return nil
}

func (j *jsiiProxy_Ec2VolumeTagsList) validateSetInternalValueParameters(val interface{}) error {
	return nil
}

func (j *jsiiProxy_Ec2VolumeTagsList) validateSetTerraformAttributeParameters(val *string) error {
	return nil
}

func (j *jsiiProxy_Ec2VolumeTagsList) validateSetTerraformResourceParameters(val cdktf.IInterpolatingParent) error {
	return nil
}

func (j *jsiiProxy_Ec2VolumeTagsList) validateSetWrapsSetParameters(val *bool) error {
	return nil
}

func validateNewEc2VolumeTagsListParameters(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, wrapsSet *bool) error {
	return nil
}

