//go:build no_runtime_type_checking

package gameliftscript

// Building without runtime type checking enabled, so all the below just return nil

func (g *jsiiProxy_GameliftScriptTagsList) validateGetParameters(index *float64) error {
	return nil
}

func (g *jsiiProxy_GameliftScriptTagsList) validateResolveParameters(_context cdktf.IResolveContext) error {
	return nil
}

func (j *jsiiProxy_GameliftScriptTagsList) validateSetInternalValueParameters(val interface{}) error {
	return nil
}

func (j *jsiiProxy_GameliftScriptTagsList) validateSetTerraformAttributeParameters(val *string) error {
	return nil
}

func (j *jsiiProxy_GameliftScriptTagsList) validateSetTerraformResourceParameters(val cdktf.IInterpolatingParent) error {
	return nil
}

func (j *jsiiProxy_GameliftScriptTagsList) validateSetWrapsSetParameters(val *bool) error {
	return nil
}

func validateNewGameliftScriptTagsListParameters(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, wrapsSet *bool) error {
	return nil
}

