//go:build no_runtime_type_checking

package ec2routetable

// Building without runtime type checking enabled, so all the below just return nil

func (e *jsiiProxy_Ec2RouteTableTagsList) validateGetParameters(index *float64) error {
	return nil
}

func (e *jsiiProxy_Ec2RouteTableTagsList) validateResolveParameters(_context cdktf.IResolveContext) error {
	return nil
}

func (j *jsiiProxy_Ec2RouteTableTagsList) validateSetInternalValueParameters(val interface{}) error {
	return nil
}

func (j *jsiiProxy_Ec2RouteTableTagsList) validateSetTerraformAttributeParameters(val *string) error {
	return nil
}

func (j *jsiiProxy_Ec2RouteTableTagsList) validateSetTerraformResourceParameters(val cdktf.IInterpolatingParent) error {
	return nil
}

func (j *jsiiProxy_Ec2RouteTableTagsList) validateSetWrapsSetParameters(val *bool) error {
	return nil
}

func validateNewEc2RouteTableTagsListParameters(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, wrapsSet *bool) error {
	return nil
}

