//go:build !no_runtime_type_checking

package appflowflow

import (
	"fmt"

	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

func (a *jsiiProxy_AppflowFlowDestinationFlowConfigListDestinationConnectorPropertiesSnowflakeErrorHandlingConfigOutputReference) validateGetAnyMapAttributeParameters(terraformAttribute *string) error {
	if terraformAttribute == nil {
		return fmt.Errorf("parameter terraformAttribute is required, but nil was provided")
	}

	return nil
}

func (a *jsiiProxy_AppflowFlowDestinationFlowConfigListDestinationConnectorPropertiesSnowflakeErrorHandlingConfigOutputReference) validateGetBooleanAttributeParameters(terraformAttribute *string) error {
	if terraformAttribute == nil {
		return fmt.Errorf("parameter terraformAttribute is required, but nil was provided")
	}

	return nil
}

func (a *jsiiProxy_AppflowFlowDestinationFlowConfigListDestinationConnectorPropertiesSnowflakeErrorHandlingConfigOutputReference) validateGetBooleanMapAttributeParameters(terraformAttribute *string) error {
	if terraformAttribute == nil {
		return fmt.Errorf("parameter terraformAttribute is required, but nil was provided")
	}

	return nil
}

func (a *jsiiProxy_AppflowFlowDestinationFlowConfigListDestinationConnectorPropertiesSnowflakeErrorHandlingConfigOutputReference) validateGetListAttributeParameters(terraformAttribute *string) error {
	if terraformAttribute == nil {
		return fmt.Errorf("parameter terraformAttribute is required, but nil was provided")
	}

	return nil
}

func (a *jsiiProxy_AppflowFlowDestinationFlowConfigListDestinationConnectorPropertiesSnowflakeErrorHandlingConfigOutputReference) validateGetNumberAttributeParameters(terraformAttribute *string) error {
	if terraformAttribute == nil {
		return fmt.Errorf("parameter terraformAttribute is required, but nil was provided")
	}

	return nil
}

func (a *jsiiProxy_AppflowFlowDestinationFlowConfigListDestinationConnectorPropertiesSnowflakeErrorHandlingConfigOutputReference) validateGetNumberListAttributeParameters(terraformAttribute *string) error {
	if terraformAttribute == nil {
		return fmt.Errorf("parameter terraformAttribute is required, but nil was provided")
	}

	return nil
}

func (a *jsiiProxy_AppflowFlowDestinationFlowConfigListDestinationConnectorPropertiesSnowflakeErrorHandlingConfigOutputReference) validateGetNumberMapAttributeParameters(terraformAttribute *string) error {
	if terraformAttribute == nil {
		return fmt.Errorf("parameter terraformAttribute is required, but nil was provided")
	}

	return nil
}

func (a *jsiiProxy_AppflowFlowDestinationFlowConfigListDestinationConnectorPropertiesSnowflakeErrorHandlingConfigOutputReference) validateGetStringAttributeParameters(terraformAttribute *string) error {
	if terraformAttribute == nil {
		return fmt.Errorf("parameter terraformAttribute is required, but nil was provided")
	}

	return nil
}

func (a *jsiiProxy_AppflowFlowDestinationFlowConfigListDestinationConnectorPropertiesSnowflakeErrorHandlingConfigOutputReference) validateGetStringMapAttributeParameters(terraformAttribute *string) error {
	if terraformAttribute == nil {
		return fmt.Errorf("parameter terraformAttribute is required, but nil was provided")
	}

	return nil
}

func (a *jsiiProxy_AppflowFlowDestinationFlowConfigListDestinationConnectorPropertiesSnowflakeErrorHandlingConfigOutputReference) validateInterpolationForAttributeParameters(property *string) error {
	if property == nil {
		return fmt.Errorf("parameter property is required, but nil was provided")
	}

	return nil
}

func (a *jsiiProxy_AppflowFlowDestinationFlowConfigListDestinationConnectorPropertiesSnowflakeErrorHandlingConfigOutputReference) validateResolveParameters(_context cdktf.IResolveContext) error {
	if _context == nil {
		return fmt.Errorf("parameter _context is required, but nil was provided")
	}

	return nil
}

func (j *jsiiProxy_AppflowFlowDestinationFlowConfigListDestinationConnectorPropertiesSnowflakeErrorHandlingConfigOutputReference) validateSetBucketNameParameters(val *string) error {
	if val == nil {
		return fmt.Errorf("parameter val is required, but nil was provided")
	}

	return nil
}

func (j *jsiiProxy_AppflowFlowDestinationFlowConfigListDestinationConnectorPropertiesSnowflakeErrorHandlingConfigOutputReference) validateSetBucketPrefixParameters(val *string) error {
	if val == nil {
		return fmt.Errorf("parameter val is required, but nil was provided")
	}

	return nil
}

func (j *jsiiProxy_AppflowFlowDestinationFlowConfigListDestinationConnectorPropertiesSnowflakeErrorHandlingConfigOutputReference) validateSetComplexObjectIndexParameters(val interface{}) error {
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

func (j *jsiiProxy_AppflowFlowDestinationFlowConfigListDestinationConnectorPropertiesSnowflakeErrorHandlingConfigOutputReference) validateSetComplexObjectIsFromSetParameters(val *bool) error {
	if val == nil {
		return fmt.Errorf("parameter val is required, but nil was provided")
	}

	return nil
}

func (j *jsiiProxy_AppflowFlowDestinationFlowConfigListDestinationConnectorPropertiesSnowflakeErrorHandlingConfigOutputReference) validateSetFailOnFirstErrorParameters(val interface{}) error {
	if val == nil {
		return fmt.Errorf("parameter val is required, but nil was provided")
	}
	switch val.(type) {
	case *bool:
		// ok
	case bool:
		// ok
	case cdktf.IResolvable:
		// ok
	default:
		if !_jsii_.IsAnonymousProxy(val) {
			return fmt.Errorf("parameter val must be one of the allowed types: *bool, cdktf.IResolvable; received %#v (a %T)", val, val)
		}
	}

	return nil
}

func (j *jsiiProxy_AppflowFlowDestinationFlowConfigListDestinationConnectorPropertiesSnowflakeErrorHandlingConfigOutputReference) validateSetInternalValueParameters(val interface{}) error {
	switch val.(type) {
	case cdktf.IResolvable:
		// ok
	case *AppflowFlowDestinationFlowConfigListDestinationConnectorPropertiesSnowflakeErrorHandlingConfig:
		val := val.(*AppflowFlowDestinationFlowConfigListDestinationConnectorPropertiesSnowflakeErrorHandlingConfig)
		if err := _jsii_.ValidateStruct(val, func() string { return "parameter val" }); err != nil {
			return err
		}
	case AppflowFlowDestinationFlowConfigListDestinationConnectorPropertiesSnowflakeErrorHandlingConfig:
		val_ := val.(AppflowFlowDestinationFlowConfigListDestinationConnectorPropertiesSnowflakeErrorHandlingConfig)
		val := &val_
		if err := _jsii_.ValidateStruct(val, func() string { return "parameter val" }); err != nil {
			return err
		}
	default:
		if !_jsii_.IsAnonymousProxy(val) {
			return fmt.Errorf("parameter val must be one of the allowed types: cdktf.IResolvable, *AppflowFlowDestinationFlowConfigListDestinationConnectorPropertiesSnowflakeErrorHandlingConfig; received %#v (a %T)", val, val)
		}
	}

	return nil
}

func (j *jsiiProxy_AppflowFlowDestinationFlowConfigListDestinationConnectorPropertiesSnowflakeErrorHandlingConfigOutputReference) validateSetTerraformAttributeParameters(val *string) error {
	if val == nil {
		return fmt.Errorf("parameter val is required, but nil was provided")
	}

	return nil
}

func (j *jsiiProxy_AppflowFlowDestinationFlowConfigListDestinationConnectorPropertiesSnowflakeErrorHandlingConfigOutputReference) validateSetTerraformResourceParameters(val cdktf.IInterpolatingParent) error {
	if val == nil {
		return fmt.Errorf("parameter val is required, but nil was provided")
	}

	return nil
}

func validateNewAppflowFlowDestinationFlowConfigListDestinationConnectorPropertiesSnowflakeErrorHandlingConfigOutputReferenceParameters(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) error {
	if terraformResource == nil {
		return fmt.Errorf("parameter terraformResource is required, but nil was provided")
	}

	if terraformAttribute == nil {
		return fmt.Errorf("parameter terraformAttribute is required, but nil was provided")
	}

	return nil
}

