//go:build no_runtime_type_checking

package workspacesconnectionalias

// Building without runtime type checking enabled, so all the below just return nil

func (w *jsiiProxy_WorkspacesConnectionAliasTagsList) validateGetParameters(index *float64) error {
	return nil
}

func (w *jsiiProxy_WorkspacesConnectionAliasTagsList) validateResolveParameters(_context cdktf.IResolveContext) error {
	return nil
}

func (j *jsiiProxy_WorkspacesConnectionAliasTagsList) validateSetInternalValueParameters(val interface{}) error {
	return nil
}

func (j *jsiiProxy_WorkspacesConnectionAliasTagsList) validateSetTerraformAttributeParameters(val *string) error {
	return nil
}

func (j *jsiiProxy_WorkspacesConnectionAliasTagsList) validateSetTerraformResourceParameters(val cdktf.IInterpolatingParent) error {
	return nil
}

func (j *jsiiProxy_WorkspacesConnectionAliasTagsList) validateSetWrapsSetParameters(val *bool) error {
	return nil
}

func validateNewWorkspacesConnectionAliasTagsListParameters(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, wrapsSet *bool) error {
	return nil
}

