//go:build !no_runtime_type_checking

package kinesisfirehosedeliverystream

import (
	"fmt"

	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

func (k *jsiiProxy_KinesisfirehoseDeliveryStreamHttpEndpointDestinationConfigurationProcessingConfigurationProcessorsList) validateGetParameters(index *float64) error {
	if index == nil {
		return fmt.Errorf("parameter index is required, but nil was provided")
	}

	return nil
}

func (k *jsiiProxy_KinesisfirehoseDeliveryStreamHttpEndpointDestinationConfigurationProcessingConfigurationProcessorsList) validateResolveParameters(_context cdktf.IResolveContext) error {
	if _context == nil {
		return fmt.Errorf("parameter _context is required, but nil was provided")
	}

	return nil
}

func (j *jsiiProxy_KinesisfirehoseDeliveryStreamHttpEndpointDestinationConfigurationProcessingConfigurationProcessorsList) validateSetInternalValueParameters(val interface{}) error {
	switch val.(type) {
	case cdktf.IResolvable:
		// ok
	case *[]*KinesisfirehoseDeliveryStreamHttpEndpointDestinationConfigurationProcessingConfigurationProcessors:
		val := val.(*[]*KinesisfirehoseDeliveryStreamHttpEndpointDestinationConfigurationProcessingConfigurationProcessors)
		for idx_97dfc6, v := range *val {
			if err := _jsii_.ValidateStruct(v, func() string { return fmt.Sprintf("parameter val[%#v]", idx_97dfc6) }); err != nil {
				return err
			}
		}
	case []*KinesisfirehoseDeliveryStreamHttpEndpointDestinationConfigurationProcessingConfigurationProcessors:
		val_ := val.([]*KinesisfirehoseDeliveryStreamHttpEndpointDestinationConfigurationProcessingConfigurationProcessors)
		val := &val_
		for idx_97dfc6, v := range *val {
			if err := _jsii_.ValidateStruct(v, func() string { return fmt.Sprintf("parameter val[%#v]", idx_97dfc6) }); err != nil {
				return err
			}
		}
	default:
		if !_jsii_.IsAnonymousProxy(val) {
			return fmt.Errorf("parameter val must be one of the allowed types: cdktf.IResolvable, *[]*KinesisfirehoseDeliveryStreamHttpEndpointDestinationConfigurationProcessingConfigurationProcessors; received %#v (a %T)", val, val)
		}
	}

	return nil
}

func (j *jsiiProxy_KinesisfirehoseDeliveryStreamHttpEndpointDestinationConfigurationProcessingConfigurationProcessorsList) validateSetTerraformAttributeParameters(val *string) error {
	if val == nil {
		return fmt.Errorf("parameter val is required, but nil was provided")
	}

	return nil
}

func (j *jsiiProxy_KinesisfirehoseDeliveryStreamHttpEndpointDestinationConfigurationProcessingConfigurationProcessorsList) validateSetTerraformResourceParameters(val cdktf.IInterpolatingParent) error {
	if val == nil {
		return fmt.Errorf("parameter val is required, but nil was provided")
	}

	return nil
}

func (j *jsiiProxy_KinesisfirehoseDeliveryStreamHttpEndpointDestinationConfigurationProcessingConfigurationProcessorsList) validateSetWrapsSetParameters(val *bool) error {
	if val == nil {
		return fmt.Errorf("parameter val is required, but nil was provided")
	}

	return nil
}

func validateNewKinesisfirehoseDeliveryStreamHttpEndpointDestinationConfigurationProcessingConfigurationProcessorsListParameters(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, wrapsSet *bool) error {
	if terraformResource == nil {
		return fmt.Errorf("parameter terraformResource is required, but nil was provided")
	}

	if terraformAttribute == nil {
		return fmt.Errorf("parameter terraformAttribute is required, but nil was provided")
	}

	if wrapsSet == nil {
		return fmt.Errorf("parameter wrapsSet is required, but nil was provided")
	}

	return nil
}

