//go:build no_runtime_type_checking

package iamuser

// Building without runtime type checking enabled, so all the below just return nil

func (i *jsiiProxy_IamUserTagsList) validateAllWithMapKeyParameters(mapKeyAttributeName *string) error {
	return nil
}

func (i *jsiiProxy_IamUserTagsList) validateGetParameters(index *float64) error {
	return nil
}

func (i *jsiiProxy_IamUserTagsList) validateResolveParameters(_context cdktf.IResolveContext) error {
	return nil
}

func (j *jsiiProxy_IamUserTagsList) validateSetInternalValueParameters(val interface{}) error {
	return nil
}

func (j *jsiiProxy_IamUserTagsList) validateSetTerraformAttributeParameters(val *string) error {
	return nil
}

func (j *jsiiProxy_IamUserTagsList) validateSetTerraformResourceParameters(val cdktf.IInterpolatingParent) error {
	return nil
}

func (j *jsiiProxy_IamUserTagsList) validateSetWrapsSetParameters(val *bool) error {
	return nil
}

func validateNewIamUserTagsListParameters(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, wrapsSet *bool) error {
	return nil
}

