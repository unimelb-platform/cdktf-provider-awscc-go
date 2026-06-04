//go:build no_runtime_type_checking

package ssoinstance

// Building without runtime type checking enabled, so all the below just return nil

func (s *jsiiProxy_SsoInstanceTagsList) validateAllWithMapKeyParameters(mapKeyAttributeName *string) error {
	return nil
}

func (s *jsiiProxy_SsoInstanceTagsList) validateGetParameters(index *float64) error {
	return nil
}

func (s *jsiiProxy_SsoInstanceTagsList) validateResolveParameters(_context cdktf.IResolveContext) error {
	return nil
}

func (j *jsiiProxy_SsoInstanceTagsList) validateSetInternalValueParameters(val interface{}) error {
	return nil
}

func (j *jsiiProxy_SsoInstanceTagsList) validateSetTerraformAttributeParameters(val *string) error {
	return nil
}

func (j *jsiiProxy_SsoInstanceTagsList) validateSetTerraformResourceParameters(val cdktf.IInterpolatingParent) error {
	return nil
}

func (j *jsiiProxy_SsoInstanceTagsList) validateSetWrapsSetParameters(val *bool) error {
	return nil
}

func validateNewSsoInstanceTagsListParameters(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, wrapsSet *bool) error {
	return nil
}

