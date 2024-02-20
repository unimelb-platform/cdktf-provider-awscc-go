//go:build no_runtime_type_checking

package secretsmanagersecret

// Building without runtime type checking enabled, so all the below just return nil

func (s *jsiiProxy_SecretsmanagerSecretTagsList) validateGetParameters(index *float64) error {
	return nil
}

func (s *jsiiProxy_SecretsmanagerSecretTagsList) validateResolveParameters(_context cdktf.IResolveContext) error {
	return nil
}

func (j *jsiiProxy_SecretsmanagerSecretTagsList) validateSetInternalValueParameters(val interface{}) error {
	return nil
}

func (j *jsiiProxy_SecretsmanagerSecretTagsList) validateSetTerraformAttributeParameters(val *string) error {
	return nil
}

func (j *jsiiProxy_SecretsmanagerSecretTagsList) validateSetTerraformResourceParameters(val cdktf.IInterpolatingParent) error {
	return nil
}

func (j *jsiiProxy_SecretsmanagerSecretTagsList) validateSetWrapsSetParameters(val *bool) error {
	return nil
}

func validateNewSecretsmanagerSecretTagsListParameters(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, wrapsSet *bool) error {
	return nil
}

