//go:build no_runtime_type_checking

package ec2subnet

// Building without runtime type checking enabled, so all the below just return nil

func (e *jsiiProxy_Ec2SubnetTagsList) validateGetParameters(index *float64) error {
	return nil
}

func (e *jsiiProxy_Ec2SubnetTagsList) validateResolveParameters(_context cdktf.IResolveContext) error {
	return nil
}

func (j *jsiiProxy_Ec2SubnetTagsList) validateSetInternalValueParameters(val interface{}) error {
	return nil
}

func (j *jsiiProxy_Ec2SubnetTagsList) validateSetTerraformAttributeParameters(val *string) error {
	return nil
}

func (j *jsiiProxy_Ec2SubnetTagsList) validateSetTerraformResourceParameters(val cdktf.IInterpolatingParent) error {
	return nil
}

func (j *jsiiProxy_Ec2SubnetTagsList) validateSetWrapsSetParameters(val *bool) error {
	return nil
}

func validateNewEc2SubnetTagsListParameters(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, wrapsSet *bool) error {
	return nil
}

