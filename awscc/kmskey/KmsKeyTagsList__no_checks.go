//go:build no_runtime_type_checking

package kmskey

// Building without runtime type checking enabled, so all the below just return nil

func (k *jsiiProxy_KmsKeyTagsList) validateAllWithMapKeyParameters(mapKeyAttributeName *string) error {
	return nil
}

func (k *jsiiProxy_KmsKeyTagsList) validateGetParameters(index *float64) error {
	return nil
}

func (k *jsiiProxy_KmsKeyTagsList) validateResolveParameters(_context cdktf.IResolveContext) error {
	return nil
}

func (j *jsiiProxy_KmsKeyTagsList) validateSetInternalValueParameters(val interface{}) error {
	return nil
}

func (j *jsiiProxy_KmsKeyTagsList) validateSetTerraformAttributeParameters(val *string) error {
	return nil
}

func (j *jsiiProxy_KmsKeyTagsList) validateSetTerraformResourceParameters(val cdktf.IInterpolatingParent) error {
	return nil
}

func (j *jsiiProxy_KmsKeyTagsList) validateSetWrapsSetParameters(val *bool) error {
	return nil
}

func validateNewKmsKeyTagsListParameters(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, wrapsSet *bool) error {
	return nil
}

