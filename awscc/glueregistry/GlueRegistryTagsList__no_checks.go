//go:build no_runtime_type_checking

package glueregistry

// Building without runtime type checking enabled, so all the below just return nil

func (g *jsiiProxy_GlueRegistryTagsList) validateAllWithMapKeyParameters(mapKeyAttributeName *string) error {
	return nil
}

func (g *jsiiProxy_GlueRegistryTagsList) validateGetParameters(index *float64) error {
	return nil
}

func (g *jsiiProxy_GlueRegistryTagsList) validateResolveParameters(_context cdktf.IResolveContext) error {
	return nil
}

func (j *jsiiProxy_GlueRegistryTagsList) validateSetInternalValueParameters(val interface{}) error {
	return nil
}

func (j *jsiiProxy_GlueRegistryTagsList) validateSetTerraformAttributeParameters(val *string) error {
	return nil
}

func (j *jsiiProxy_GlueRegistryTagsList) validateSetTerraformResourceParameters(val cdktf.IInterpolatingParent) error {
	return nil
}

func (j *jsiiProxy_GlueRegistryTagsList) validateSetWrapsSetParameters(val *bool) error {
	return nil
}

func validateNewGlueRegistryTagsListParameters(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, wrapsSet *bool) error {
	return nil
}

