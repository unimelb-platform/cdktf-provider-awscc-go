//go:build no_runtime_type_checking

package ec2flowlog

// Building without runtime type checking enabled, so all the below just return nil

func (e *jsiiProxy_Ec2FlowLogTagsList) validateGetParameters(index *float64) error {
	return nil
}

func (e *jsiiProxy_Ec2FlowLogTagsList) validateResolveParameters(_context cdktf.IResolveContext) error {
	return nil
}

func (j *jsiiProxy_Ec2FlowLogTagsList) validateSetInternalValueParameters(val interface{}) error {
	return nil
}

func (j *jsiiProxy_Ec2FlowLogTagsList) validateSetTerraformAttributeParameters(val *string) error {
	return nil
}

func (j *jsiiProxy_Ec2FlowLogTagsList) validateSetTerraformResourceParameters(val cdktf.IInterpolatingParent) error {
	return nil
}

func (j *jsiiProxy_Ec2FlowLogTagsList) validateSetWrapsSetParameters(val *bool) error {
	return nil
}

func validateNewEc2FlowLogTagsListParameters(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, wrapsSet *bool) error {
	return nil
}

