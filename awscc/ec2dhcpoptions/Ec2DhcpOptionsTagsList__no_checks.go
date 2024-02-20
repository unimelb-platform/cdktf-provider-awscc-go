//go:build no_runtime_type_checking

package ec2dhcpoptions

// Building without runtime type checking enabled, so all the below just return nil

func (e *jsiiProxy_Ec2DhcpOptionsTagsList) validateGetParameters(index *float64) error {
	return nil
}

func (e *jsiiProxy_Ec2DhcpOptionsTagsList) validateResolveParameters(_context cdktf.IResolveContext) error {
	return nil
}

func (j *jsiiProxy_Ec2DhcpOptionsTagsList) validateSetInternalValueParameters(val interface{}) error {
	return nil
}

func (j *jsiiProxy_Ec2DhcpOptionsTagsList) validateSetTerraformAttributeParameters(val *string) error {
	return nil
}

func (j *jsiiProxy_Ec2DhcpOptionsTagsList) validateSetTerraformResourceParameters(val cdktf.IInterpolatingParent) error {
	return nil
}

func (j *jsiiProxy_Ec2DhcpOptionsTagsList) validateSetWrapsSetParameters(val *bool) error {
	return nil
}

func validateNewEc2DhcpOptionsTagsListParameters(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, wrapsSet *bool) error {
	return nil
}

