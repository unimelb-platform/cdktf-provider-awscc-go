//go:build no_runtime_type_checking

package ec2host

// Building without runtime type checking enabled, so all the below just return nil

func (e *jsiiProxy_Ec2HostTagsList) validateAllWithMapKeyParameters(mapKeyAttributeName *string) error {
	return nil
}

func (e *jsiiProxy_Ec2HostTagsList) validateGetParameters(index *float64) error {
	return nil
}

func (e *jsiiProxy_Ec2HostTagsList) validateResolveParameters(_context cdktf.IResolveContext) error {
	return nil
}

func (j *jsiiProxy_Ec2HostTagsList) validateSetInternalValueParameters(val interface{}) error {
	return nil
}

func (j *jsiiProxy_Ec2HostTagsList) validateSetTerraformAttributeParameters(val *string) error {
	return nil
}

func (j *jsiiProxy_Ec2HostTagsList) validateSetTerraformResourceParameters(val cdktf.IInterpolatingParent) error {
	return nil
}

func (j *jsiiProxy_Ec2HostTagsList) validateSetWrapsSetParameters(val *bool) error {
	return nil
}

func validateNewEc2HostTagsListParameters(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, wrapsSet *bool) error {
	return nil
}

