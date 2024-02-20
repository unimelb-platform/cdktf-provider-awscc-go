//go:build no_runtime_type_checking

package omicsworkflow

// Building without runtime type checking enabled, so all the below just return nil

func (o *jsiiProxy_OmicsWorkflowParameterTemplateMap) validateGetParameters(key *string) error {
	return nil
}

func (o *jsiiProxy_OmicsWorkflowParameterTemplateMap) validateInterpolationForAttributeParameters(property *string) error {
	return nil
}

func (o *jsiiProxy_OmicsWorkflowParameterTemplateMap) validateResolveParameters(_context cdktf.IResolveContext) error {
	return nil
}

func (j *jsiiProxy_OmicsWorkflowParameterTemplateMap) validateSetInternalValueParameters(val interface{}) error {
	return nil
}

func (j *jsiiProxy_OmicsWorkflowParameterTemplateMap) validateSetTerraformAttributeParameters(val *string) error {
	return nil
}

func (j *jsiiProxy_OmicsWorkflowParameterTemplateMap) validateSetTerraformResourceParameters(val cdktf.IInterpolatingParent) error {
	return nil
}

func validateNewOmicsWorkflowParameterTemplateMapParameters(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) error {
	return nil
}

