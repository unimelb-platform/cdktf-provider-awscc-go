//go:build no_runtime_type_checking

package ec2eip

// Building without runtime type checking enabled, so all the below just return nil

func (e *jsiiProxy_Ec2EipTagsList) validateGetParameters(index *float64) error {
	return nil
}

func (e *jsiiProxy_Ec2EipTagsList) validateResolveParameters(_context cdktf.IResolveContext) error {
	return nil
}

func (j *jsiiProxy_Ec2EipTagsList) validateSetInternalValueParameters(val interface{}) error {
	return nil
}

func (j *jsiiProxy_Ec2EipTagsList) validateSetTerraformAttributeParameters(val *string) error {
	return nil
}

func (j *jsiiProxy_Ec2EipTagsList) validateSetTerraformResourceParameters(val cdktf.IInterpolatingParent) error {
	return nil
}

func (j *jsiiProxy_Ec2EipTagsList) validateSetWrapsSetParameters(val *bool) error {
	return nil
}

func validateNewEc2EipTagsListParameters(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, wrapsSet *bool) error {
	return nil
}

