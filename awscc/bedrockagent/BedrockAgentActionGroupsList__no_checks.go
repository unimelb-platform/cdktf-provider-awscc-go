//go:build no_runtime_type_checking

package bedrockagent

// Building without runtime type checking enabled, so all the below just return nil

func (b *jsiiProxy_BedrockAgentActionGroupsList) validateAllWithMapKeyParameters(mapKeyAttributeName *string) error {
	return nil
}

func (b *jsiiProxy_BedrockAgentActionGroupsList) validateGetParameters(index *float64) error {
	return nil
}

func (b *jsiiProxy_BedrockAgentActionGroupsList) validateResolveParameters(_context cdktf.IResolveContext) error {
	return nil
}

func (j *jsiiProxy_BedrockAgentActionGroupsList) validateSetInternalValueParameters(val interface{}) error {
	return nil
}

func (j *jsiiProxy_BedrockAgentActionGroupsList) validateSetTerraformAttributeParameters(val *string) error {
	return nil
}

func (j *jsiiProxy_BedrockAgentActionGroupsList) validateSetTerraformResourceParameters(val cdktf.IInterpolatingParent) error {
	return nil
}

func (j *jsiiProxy_BedrockAgentActionGroupsList) validateSetWrapsSetParameters(val *bool) error {
	return nil
}

func validateNewBedrockAgentActionGroupsListParameters(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, wrapsSet *bool) error {
	return nil
}

