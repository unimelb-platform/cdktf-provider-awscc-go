//go:build no_runtime_type_checking

package kmsreplicakey

// Building without runtime type checking enabled, so all the below just return nil

func (k *jsiiProxy_KmsReplicaKeyTagsList) validateAllWithMapKeyParameters(mapKeyAttributeName *string) error {
	return nil
}

func (k *jsiiProxy_KmsReplicaKeyTagsList) validateGetParameters(index *float64) error {
	return nil
}

func (k *jsiiProxy_KmsReplicaKeyTagsList) validateResolveParameters(_context cdktf.IResolveContext) error {
	return nil
}

func (j *jsiiProxy_KmsReplicaKeyTagsList) validateSetInternalValueParameters(val interface{}) error {
	return nil
}

func (j *jsiiProxy_KmsReplicaKeyTagsList) validateSetTerraformAttributeParameters(val *string) error {
	return nil
}

func (j *jsiiProxy_KmsReplicaKeyTagsList) validateSetTerraformResourceParameters(val cdktf.IInterpolatingParent) error {
	return nil
}

func (j *jsiiProxy_KmsReplicaKeyTagsList) validateSetWrapsSetParameters(val *bool) error {
	return nil
}

func validateNewKmsReplicaKeyTagsListParameters(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, wrapsSet *bool) error {
	return nil
}

