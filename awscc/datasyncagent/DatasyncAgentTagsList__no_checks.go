//go:build no_runtime_type_checking

package datasyncagent

// Building without runtime type checking enabled, so all the below just return nil

func (d *jsiiProxy_DatasyncAgentTagsList) validateAllWithMapKeyParameters(mapKeyAttributeName *string) error {
	return nil
}

func (d *jsiiProxy_DatasyncAgentTagsList) validateGetParameters(index *float64) error {
	return nil
}

func (d *jsiiProxy_DatasyncAgentTagsList) validateResolveParameters(_context cdktf.IResolveContext) error {
	return nil
}

func (j *jsiiProxy_DatasyncAgentTagsList) validateSetInternalValueParameters(val interface{}) error {
	return nil
}

func (j *jsiiProxy_DatasyncAgentTagsList) validateSetTerraformAttributeParameters(val *string) error {
	return nil
}

func (j *jsiiProxy_DatasyncAgentTagsList) validateSetTerraformResourceParameters(val cdktf.IInterpolatingParent) error {
	return nil
}

func (j *jsiiProxy_DatasyncAgentTagsList) validateSetWrapsSetParameters(val *bool) error {
	return nil
}

func validateNewDatasyncAgentTagsListParameters(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, wrapsSet *bool) error {
	return nil
}

