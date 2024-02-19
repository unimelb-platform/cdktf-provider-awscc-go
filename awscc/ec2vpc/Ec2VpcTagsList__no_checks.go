//go:build no_runtime_type_checking

package ec2vpc

// Building without runtime type checking enabled, so all the below just return nil

func (e *jsiiProxy_Ec2VpcTagsList) validateGetParameters(index *float64) error {
	return nil
}

func (e *jsiiProxy_Ec2VpcTagsList) validateResolveParameters(_context cdktf.IResolveContext) error {
	return nil
}

func (j *jsiiProxy_Ec2VpcTagsList) validateSetInternalValueParameters(val interface{}) error {
	return nil
}

func (j *jsiiProxy_Ec2VpcTagsList) validateSetTerraformAttributeParameters(val *string) error {
	return nil
}

func (j *jsiiProxy_Ec2VpcTagsList) validateSetTerraformResourceParameters(val cdktf.IInterpolatingParent) error {
	return nil
}

func (j *jsiiProxy_Ec2VpcTagsList) validateSetWrapsSetParameters(val *bool) error {
	return nil
}

func validateNewEc2VpcTagsListParameters(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, wrapsSet *bool) error {
	return nil
}

