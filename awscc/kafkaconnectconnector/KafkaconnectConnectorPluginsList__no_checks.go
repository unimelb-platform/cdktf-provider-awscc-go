//go:build no_runtime_type_checking

package kafkaconnectconnector

// Building without runtime type checking enabled, so all the below just return nil

func (k *jsiiProxy_KafkaconnectConnectorPluginsList) validateGetParameters(index *float64) error {
	return nil
}

func (k *jsiiProxy_KafkaconnectConnectorPluginsList) validateResolveParameters(_context cdktf.IResolveContext) error {
	return nil
}

func (j *jsiiProxy_KafkaconnectConnectorPluginsList) validateSetInternalValueParameters(val interface{}) error {
	return nil
}

func (j *jsiiProxy_KafkaconnectConnectorPluginsList) validateSetTerraformAttributeParameters(val *string) error {
	return nil
}

func (j *jsiiProxy_KafkaconnectConnectorPluginsList) validateSetTerraformResourceParameters(val cdktf.IInterpolatingParent) error {
	return nil
}

func (j *jsiiProxy_KafkaconnectConnectorPluginsList) validateSetWrapsSetParameters(val *bool) error {
	return nil
}

func validateNewKafkaconnectConnectorPluginsListParameters(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, wrapsSet *bool) error {
	return nil
}

