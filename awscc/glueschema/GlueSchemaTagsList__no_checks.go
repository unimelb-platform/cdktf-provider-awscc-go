//go:build no_runtime_type_checking

package glueschema

// Building without runtime type checking enabled, so all the below just return nil

func (g *jsiiProxy_GlueSchemaTagsList) validateGetParameters(index *float64) error {
	return nil
}

func (g *jsiiProxy_GlueSchemaTagsList) validateResolveParameters(_context cdktf.IResolveContext) error {
	return nil
}

func (j *jsiiProxy_GlueSchemaTagsList) validateSetInternalValueParameters(val interface{}) error {
	return nil
}

func (j *jsiiProxy_GlueSchemaTagsList) validateSetTerraformAttributeParameters(val *string) error {
	return nil
}

func (j *jsiiProxy_GlueSchemaTagsList) validateSetTerraformResourceParameters(val cdktf.IInterpolatingParent) error {
	return nil
}

func (j *jsiiProxy_GlueSchemaTagsList) validateSetWrapsSetParameters(val *bool) error {
	return nil
}

func validateNewGlueSchemaTagsListParameters(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, wrapsSet *bool) error {
	return nil
}

