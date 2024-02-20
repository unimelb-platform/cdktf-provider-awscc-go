//go:build !no_runtime_type_checking

package kinesisfirehosedeliverystream

import (
	"fmt"

	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

func (k *jsiiProxy_KinesisfirehoseDeliveryStreamExtendedS3DestinationConfigurationDynamicPartitioningConfigurationRetryOptionsOutputReference) validateGetAnyMapAttributeParameters(terraformAttribute *string) error {
	if terraformAttribute == nil {
		return fmt.Errorf("parameter terraformAttribute is required, but nil was provided")
	}

	return nil
}

func (k *jsiiProxy_KinesisfirehoseDeliveryStreamExtendedS3DestinationConfigurationDynamicPartitioningConfigurationRetryOptionsOutputReference) validateGetBooleanAttributeParameters(terraformAttribute *string) error {
	if terraformAttribute == nil {
		return fmt.Errorf("parameter terraformAttribute is required, but nil was provided")
	}

	return nil
}

func (k *jsiiProxy_KinesisfirehoseDeliveryStreamExtendedS3DestinationConfigurationDynamicPartitioningConfigurationRetryOptionsOutputReference) validateGetBooleanMapAttributeParameters(terraformAttribute *string) error {
	if terraformAttribute == nil {
		return fmt.Errorf("parameter terraformAttribute is required, but nil was provided")
	}

	return nil
}

func (k *jsiiProxy_KinesisfirehoseDeliveryStreamExtendedS3DestinationConfigurationDynamicPartitioningConfigurationRetryOptionsOutputReference) validateGetListAttributeParameters(terraformAttribute *string) error {
	if terraformAttribute == nil {
		return fmt.Errorf("parameter terraformAttribute is required, but nil was provided")
	}

	return nil
}

func (k *jsiiProxy_KinesisfirehoseDeliveryStreamExtendedS3DestinationConfigurationDynamicPartitioningConfigurationRetryOptionsOutputReference) validateGetNumberAttributeParameters(terraformAttribute *string) error {
	if terraformAttribute == nil {
		return fmt.Errorf("parameter terraformAttribute is required, but nil was provided")
	}

	return nil
}

func (k *jsiiProxy_KinesisfirehoseDeliveryStreamExtendedS3DestinationConfigurationDynamicPartitioningConfigurationRetryOptionsOutputReference) validateGetNumberListAttributeParameters(terraformAttribute *string) error {
	if terraformAttribute == nil {
		return fmt.Errorf("parameter terraformAttribute is required, but nil was provided")
	}

	return nil
}

func (k *jsiiProxy_KinesisfirehoseDeliveryStreamExtendedS3DestinationConfigurationDynamicPartitioningConfigurationRetryOptionsOutputReference) validateGetNumberMapAttributeParameters(terraformAttribute *string) error {
	if terraformAttribute == nil {
		return fmt.Errorf("parameter terraformAttribute is required, but nil was provided")
	}

	return nil
}

func (k *jsiiProxy_KinesisfirehoseDeliveryStreamExtendedS3DestinationConfigurationDynamicPartitioningConfigurationRetryOptionsOutputReference) validateGetStringAttributeParameters(terraformAttribute *string) error {
	if terraformAttribute == nil {
		return fmt.Errorf("parameter terraformAttribute is required, but nil was provided")
	}

	return nil
}

func (k *jsiiProxy_KinesisfirehoseDeliveryStreamExtendedS3DestinationConfigurationDynamicPartitioningConfigurationRetryOptionsOutputReference) validateGetStringMapAttributeParameters(terraformAttribute *string) error {
	if terraformAttribute == nil {
		return fmt.Errorf("parameter terraformAttribute is required, but nil was provided")
	}

	return nil
}

func (k *jsiiProxy_KinesisfirehoseDeliveryStreamExtendedS3DestinationConfigurationDynamicPartitioningConfigurationRetryOptionsOutputReference) validateInterpolationForAttributeParameters(property *string) error {
	if property == nil {
		return fmt.Errorf("parameter property is required, but nil was provided")
	}

	return nil
}

func (k *jsiiProxy_KinesisfirehoseDeliveryStreamExtendedS3DestinationConfigurationDynamicPartitioningConfigurationRetryOptionsOutputReference) validateResolveParameters(_context cdktf.IResolveContext) error {
	if _context == nil {
		return fmt.Errorf("parameter _context is required, but nil was provided")
	}

	return nil
}

func (j *jsiiProxy_KinesisfirehoseDeliveryStreamExtendedS3DestinationConfigurationDynamicPartitioningConfigurationRetryOptionsOutputReference) validateSetComplexObjectIndexParameters(val interface{}) error {
	switch val.(type) {
	case *string:
		// ok
	case string:
		// ok
	case *float64:
		// ok
	case float64:
		// ok
	case *int:
		// ok
	case int:
		// ok
	case *uint:
		// ok
	case uint:
		// ok
	case *int8:
		// ok
	case int8:
		// ok
	case *int16:
		// ok
	case int16:
		// ok
	case *int32:
		// ok
	case int32:
		// ok
	case *int64:
		// ok
	case int64:
		// ok
	case *uint8:
		// ok
	case uint8:
		// ok
	case *uint16:
		// ok
	case uint16:
		// ok
	case *uint32:
		// ok
	case uint32:
		// ok
	case *uint64:
		// ok
	case uint64:
		// ok
	default:
		return fmt.Errorf("parameter val must be one of the allowed types: *string, *float64; received %#v (a %T)", val, val)
	}

	return nil
}

func (j *jsiiProxy_KinesisfirehoseDeliveryStreamExtendedS3DestinationConfigurationDynamicPartitioningConfigurationRetryOptionsOutputReference) validateSetComplexObjectIsFromSetParameters(val *bool) error {
	if val == nil {
		return fmt.Errorf("parameter val is required, but nil was provided")
	}

	return nil
}

func (j *jsiiProxy_KinesisfirehoseDeliveryStreamExtendedS3DestinationConfigurationDynamicPartitioningConfigurationRetryOptionsOutputReference) validateSetDurationInSecondsParameters(val *float64) error {
	if val == nil {
		return fmt.Errorf("parameter val is required, but nil was provided")
	}

	return nil
}

func (j *jsiiProxy_KinesisfirehoseDeliveryStreamExtendedS3DestinationConfigurationDynamicPartitioningConfigurationRetryOptionsOutputReference) validateSetInternalValueParameters(val interface{}) error {
	switch val.(type) {
	case cdktf.IResolvable:
		// ok
	case *KinesisfirehoseDeliveryStreamExtendedS3DestinationConfigurationDynamicPartitioningConfigurationRetryOptions:
		val := val.(*KinesisfirehoseDeliveryStreamExtendedS3DestinationConfigurationDynamicPartitioningConfigurationRetryOptions)
		if err := _jsii_.ValidateStruct(val, func() string { return "parameter val" }); err != nil {
			return err
		}
	case KinesisfirehoseDeliveryStreamExtendedS3DestinationConfigurationDynamicPartitioningConfigurationRetryOptions:
		val_ := val.(KinesisfirehoseDeliveryStreamExtendedS3DestinationConfigurationDynamicPartitioningConfigurationRetryOptions)
		val := &val_
		if err := _jsii_.ValidateStruct(val, func() string { return "parameter val" }); err != nil {
			return err
		}
	default:
		if !_jsii_.IsAnonymousProxy(val) {
			return fmt.Errorf("parameter val must be one of the allowed types: cdktf.IResolvable, *KinesisfirehoseDeliveryStreamExtendedS3DestinationConfigurationDynamicPartitioningConfigurationRetryOptions; received %#v (a %T)", val, val)
		}
	}

	return nil
}

func (j *jsiiProxy_KinesisfirehoseDeliveryStreamExtendedS3DestinationConfigurationDynamicPartitioningConfigurationRetryOptionsOutputReference) validateSetTerraformAttributeParameters(val *string) error {
	if val == nil {
		return fmt.Errorf("parameter val is required, but nil was provided")
	}

	return nil
}

func (j *jsiiProxy_KinesisfirehoseDeliveryStreamExtendedS3DestinationConfigurationDynamicPartitioningConfigurationRetryOptionsOutputReference) validateSetTerraformResourceParameters(val cdktf.IInterpolatingParent) error {
	if val == nil {
		return fmt.Errorf("parameter val is required, but nil was provided")
	}

	return nil
}

func validateNewKinesisfirehoseDeliveryStreamExtendedS3DestinationConfigurationDynamicPartitioningConfigurationRetryOptionsOutputReferenceParameters(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) error {
	if terraformResource == nil {
		return fmt.Errorf("parameter terraformResource is required, but nil was provided")
	}

	if terraformAttribute == nil {
		return fmt.Errorf("parameter terraformAttribute is required, but nil was provided")
	}

	return nil
}

