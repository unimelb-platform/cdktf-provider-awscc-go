//go:build no_runtime_type_checking

package codeartifactrepository

// Building without runtime type checking enabled, so all the below just return nil

func (c *jsiiProxy_CodeartifactRepositoryTagsList) validateGetParameters(index *float64) error {
	return nil
}

func (c *jsiiProxy_CodeartifactRepositoryTagsList) validateResolveParameters(_context cdktf.IResolveContext) error {
	return nil
}

func (j *jsiiProxy_CodeartifactRepositoryTagsList) validateSetInternalValueParameters(val interface{}) error {
	return nil
}

func (j *jsiiProxy_CodeartifactRepositoryTagsList) validateSetTerraformAttributeParameters(val *string) error {
	return nil
}

func (j *jsiiProxy_CodeartifactRepositoryTagsList) validateSetTerraformResourceParameters(val cdktf.IInterpolatingParent) error {
	return nil
}

func (j *jsiiProxy_CodeartifactRepositoryTagsList) validateSetWrapsSetParameters(val *bool) error {
	return nil
}

func validateNewCodeartifactRepositoryTagsListParameters(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, wrapsSet *bool) error {
	return nil
}

