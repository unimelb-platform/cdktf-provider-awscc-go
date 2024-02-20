//go:build no_runtime_type_checking

package iamrole

// Building without runtime type checking enabled, so all the below just return nil

func (i *jsiiProxy_IamRoleTagsList) validateGetParameters(index *float64) error {
	return nil
}

func (i *jsiiProxy_IamRoleTagsList) validateResolveParameters(_context cdktf.IResolveContext) error {
	return nil
}

func (j *jsiiProxy_IamRoleTagsList) validateSetInternalValueParameters(val interface{}) error {
	return nil
}

func (j *jsiiProxy_IamRoleTagsList) validateSetTerraformAttributeParameters(val *string) error {
	return nil
}

func (j *jsiiProxy_IamRoleTagsList) validateSetTerraformResourceParameters(val cdktf.IInterpolatingParent) error {
	return nil
}

func (j *jsiiProxy_IamRoleTagsList) validateSetWrapsSetParameters(val *bool) error {
	return nil
}

func validateNewIamRoleTagsListParameters(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, wrapsSet *bool) error {
	return nil
}

