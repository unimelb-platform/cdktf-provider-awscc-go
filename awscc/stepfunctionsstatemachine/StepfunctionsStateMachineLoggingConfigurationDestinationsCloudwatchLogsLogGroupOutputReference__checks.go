//go:build !no_runtime_type_checking

package stepfunctionsstatemachine

import (
	"fmt"

	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

func (s *jsiiProxy_StepfunctionsStateMachineLoggingConfigurationDestinationsCloudwatchLogsLogGroupOutputReference) validateGetAnyMapAttributeParameters(terraformAttribute *string) error {
	if terraformAttribute == nil {
		return fmt.Errorf("parameter terraformAttribute is required, but nil was provided")
	}

	return nil
}

func (s *jsiiProxy_StepfunctionsStateMachineLoggingConfigurationDestinationsCloudwatchLogsLogGroupOutputReference) validateGetBooleanAttributeParameters(terraformAttribute *string) error {
	if terraformAttribute == nil {
		return fmt.Errorf("parameter terraformAttribute is required, but nil was provided")
	}

	return nil
}

func (s *jsiiProxy_StepfunctionsStateMachineLoggingConfigurationDestinationsCloudwatchLogsLogGroupOutputReference) validateGetBooleanMapAttributeParameters(terraformAttribute *string) error {
	if terraformAttribute == nil {
		return fmt.Errorf("parameter terraformAttribute is required, but nil was provided")
	}

	return nil
}

func (s *jsiiProxy_StepfunctionsStateMachineLoggingConfigurationDestinationsCloudwatchLogsLogGroupOutputReference) validateGetListAttributeParameters(terraformAttribute *string) error {
	if terraformAttribute == nil {
		return fmt.Errorf("parameter terraformAttribute is required, but nil was provided")
	}

	return nil
}

func (s *jsiiProxy_StepfunctionsStateMachineLoggingConfigurationDestinationsCloudwatchLogsLogGroupOutputReference) validateGetNumberAttributeParameters(terraformAttribute *string) error {
	if terraformAttribute == nil {
		return fmt.Errorf("parameter terraformAttribute is required, but nil was provided")
	}

	return nil
}

func (s *jsiiProxy_StepfunctionsStateMachineLoggingConfigurationDestinationsCloudwatchLogsLogGroupOutputReference) validateGetNumberListAttributeParameters(terraformAttribute *string) error {
	if terraformAttribute == nil {
		return fmt.Errorf("parameter terraformAttribute is required, but nil was provided")
	}

	return nil
}

func (s *jsiiProxy_StepfunctionsStateMachineLoggingConfigurationDestinationsCloudwatchLogsLogGroupOutputReference) validateGetNumberMapAttributeParameters(terraformAttribute *string) error {
	if terraformAttribute == nil {
		return fmt.Errorf("parameter terraformAttribute is required, but nil was provided")
	}

	return nil
}

func (s *jsiiProxy_StepfunctionsStateMachineLoggingConfigurationDestinationsCloudwatchLogsLogGroupOutputReference) validateGetStringAttributeParameters(terraformAttribute *string) error {
	if terraformAttribute == nil {
		return fmt.Errorf("parameter terraformAttribute is required, but nil was provided")
	}

	return nil
}

func (s *jsiiProxy_StepfunctionsStateMachineLoggingConfigurationDestinationsCloudwatchLogsLogGroupOutputReference) validateGetStringMapAttributeParameters(terraformAttribute *string) error {
	if terraformAttribute == nil {
		return fmt.Errorf("parameter terraformAttribute is required, but nil was provided")
	}

	return nil
}

func (s *jsiiProxy_StepfunctionsStateMachineLoggingConfigurationDestinationsCloudwatchLogsLogGroupOutputReference) validateInterpolationForAttributeParameters(property *string) error {
	if property == nil {
		return fmt.Errorf("parameter property is required, but nil was provided")
	}

	return nil
}

func (s *jsiiProxy_StepfunctionsStateMachineLoggingConfigurationDestinationsCloudwatchLogsLogGroupOutputReference) validateResolveParameters(_context cdktf.IResolveContext) error {
	if _context == nil {
		return fmt.Errorf("parameter _context is required, but nil was provided")
	}

	return nil
}

func (j *jsiiProxy_StepfunctionsStateMachineLoggingConfigurationDestinationsCloudwatchLogsLogGroupOutputReference) validateSetComplexObjectIndexParameters(val interface{}) error {
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

func (j *jsiiProxy_StepfunctionsStateMachineLoggingConfigurationDestinationsCloudwatchLogsLogGroupOutputReference) validateSetComplexObjectIsFromSetParameters(val *bool) error {
	if val == nil {
		return fmt.Errorf("parameter val is required, but nil was provided")
	}

	return nil
}

func (j *jsiiProxy_StepfunctionsStateMachineLoggingConfigurationDestinationsCloudwatchLogsLogGroupOutputReference) validateSetInternalValueParameters(val interface{}) error {
	switch val.(type) {
	case cdktf.IResolvable:
		// ok
	case *StepfunctionsStateMachineLoggingConfigurationDestinationsCloudwatchLogsLogGroup:
		val := val.(*StepfunctionsStateMachineLoggingConfigurationDestinationsCloudwatchLogsLogGroup)
		if err := _jsii_.ValidateStruct(val, func() string { return "parameter val" }); err != nil {
			return err
		}
	case StepfunctionsStateMachineLoggingConfigurationDestinationsCloudwatchLogsLogGroup:
		val_ := val.(StepfunctionsStateMachineLoggingConfigurationDestinationsCloudwatchLogsLogGroup)
		val := &val_
		if err := _jsii_.ValidateStruct(val, func() string { return "parameter val" }); err != nil {
			return err
		}
	default:
		if !_jsii_.IsAnonymousProxy(val) {
			return fmt.Errorf("parameter val must be one of the allowed types: cdktf.IResolvable, *StepfunctionsStateMachineLoggingConfigurationDestinationsCloudwatchLogsLogGroup; received %#v (a %T)", val, val)
		}
	}

	return nil
}

func (j *jsiiProxy_StepfunctionsStateMachineLoggingConfigurationDestinationsCloudwatchLogsLogGroupOutputReference) validateSetLogGroupArnParameters(val *string) error {
	if val == nil {
		return fmt.Errorf("parameter val is required, but nil was provided")
	}

	return nil
}

func (j *jsiiProxy_StepfunctionsStateMachineLoggingConfigurationDestinationsCloudwatchLogsLogGroupOutputReference) validateSetTerraformAttributeParameters(val *string) error {
	if val == nil {
		return fmt.Errorf("parameter val is required, but nil was provided")
	}

	return nil
}

func (j *jsiiProxy_StepfunctionsStateMachineLoggingConfigurationDestinationsCloudwatchLogsLogGroupOutputReference) validateSetTerraformResourceParameters(val cdktf.IInterpolatingParent) error {
	if val == nil {
		return fmt.Errorf("parameter val is required, but nil was provided")
	}

	return nil
}

func validateNewStepfunctionsStateMachineLoggingConfigurationDestinationsCloudwatchLogsLogGroupOutputReferenceParameters(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) error {
	if terraformResource == nil {
		return fmt.Errorf("parameter terraformResource is required, but nil was provided")
	}

	if terraformAttribute == nil {
		return fmt.Errorf("parameter terraformAttribute is required, but nil was provided")
	}

	return nil
}

