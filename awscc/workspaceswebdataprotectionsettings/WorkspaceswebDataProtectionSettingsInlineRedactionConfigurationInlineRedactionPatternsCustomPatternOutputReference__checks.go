//go:build !no_runtime_type_checking

package workspaceswebdataprotectionsettings

import (
	"fmt"

	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

func (w *jsiiProxy_WorkspaceswebDataProtectionSettingsInlineRedactionConfigurationInlineRedactionPatternsCustomPatternOutputReference) validateGetAnyMapAttributeParameters(terraformAttribute *string) error {
	if terraformAttribute == nil {
		return fmt.Errorf("parameter terraformAttribute is required, but nil was provided")
	}

	return nil
}

func (w *jsiiProxy_WorkspaceswebDataProtectionSettingsInlineRedactionConfigurationInlineRedactionPatternsCustomPatternOutputReference) validateGetBooleanAttributeParameters(terraformAttribute *string) error {
	if terraformAttribute == nil {
		return fmt.Errorf("parameter terraformAttribute is required, but nil was provided")
	}

	return nil
}

func (w *jsiiProxy_WorkspaceswebDataProtectionSettingsInlineRedactionConfigurationInlineRedactionPatternsCustomPatternOutputReference) validateGetBooleanMapAttributeParameters(terraformAttribute *string) error {
	if terraformAttribute == nil {
		return fmt.Errorf("parameter terraformAttribute is required, but nil was provided")
	}

	return nil
}

func (w *jsiiProxy_WorkspaceswebDataProtectionSettingsInlineRedactionConfigurationInlineRedactionPatternsCustomPatternOutputReference) validateGetListAttributeParameters(terraformAttribute *string) error {
	if terraformAttribute == nil {
		return fmt.Errorf("parameter terraformAttribute is required, but nil was provided")
	}

	return nil
}

func (w *jsiiProxy_WorkspaceswebDataProtectionSettingsInlineRedactionConfigurationInlineRedactionPatternsCustomPatternOutputReference) validateGetNumberAttributeParameters(terraformAttribute *string) error {
	if terraformAttribute == nil {
		return fmt.Errorf("parameter terraformAttribute is required, but nil was provided")
	}

	return nil
}

func (w *jsiiProxy_WorkspaceswebDataProtectionSettingsInlineRedactionConfigurationInlineRedactionPatternsCustomPatternOutputReference) validateGetNumberListAttributeParameters(terraformAttribute *string) error {
	if terraformAttribute == nil {
		return fmt.Errorf("parameter terraformAttribute is required, but nil was provided")
	}

	return nil
}

func (w *jsiiProxy_WorkspaceswebDataProtectionSettingsInlineRedactionConfigurationInlineRedactionPatternsCustomPatternOutputReference) validateGetNumberMapAttributeParameters(terraformAttribute *string) error {
	if terraformAttribute == nil {
		return fmt.Errorf("parameter terraformAttribute is required, but nil was provided")
	}

	return nil
}

func (w *jsiiProxy_WorkspaceswebDataProtectionSettingsInlineRedactionConfigurationInlineRedactionPatternsCustomPatternOutputReference) validateGetStringAttributeParameters(terraformAttribute *string) error {
	if terraformAttribute == nil {
		return fmt.Errorf("parameter terraformAttribute is required, but nil was provided")
	}

	return nil
}

func (w *jsiiProxy_WorkspaceswebDataProtectionSettingsInlineRedactionConfigurationInlineRedactionPatternsCustomPatternOutputReference) validateGetStringMapAttributeParameters(terraformAttribute *string) error {
	if terraformAttribute == nil {
		return fmt.Errorf("parameter terraformAttribute is required, but nil was provided")
	}

	return nil
}

func (w *jsiiProxy_WorkspaceswebDataProtectionSettingsInlineRedactionConfigurationInlineRedactionPatternsCustomPatternOutputReference) validateInterpolationForAttributeParameters(property *string) error {
	if property == nil {
		return fmt.Errorf("parameter property is required, but nil was provided")
	}

	return nil
}

func (w *jsiiProxy_WorkspaceswebDataProtectionSettingsInlineRedactionConfigurationInlineRedactionPatternsCustomPatternOutputReference) validateResolveParameters(_context cdktf.IResolveContext) error {
	if _context == nil {
		return fmt.Errorf("parameter _context is required, but nil was provided")
	}

	return nil
}

func (j *jsiiProxy_WorkspaceswebDataProtectionSettingsInlineRedactionConfigurationInlineRedactionPatternsCustomPatternOutputReference) validateSetComplexObjectIndexParameters(val interface{}) error {
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

func (j *jsiiProxy_WorkspaceswebDataProtectionSettingsInlineRedactionConfigurationInlineRedactionPatternsCustomPatternOutputReference) validateSetComplexObjectIsFromSetParameters(val *bool) error {
	if val == nil {
		return fmt.Errorf("parameter val is required, but nil was provided")
	}

	return nil
}

func (j *jsiiProxy_WorkspaceswebDataProtectionSettingsInlineRedactionConfigurationInlineRedactionPatternsCustomPatternOutputReference) validateSetInternalValueParameters(val interface{}) error {
	switch val.(type) {
	case cdktf.IResolvable:
		// ok
	case *WorkspaceswebDataProtectionSettingsInlineRedactionConfigurationInlineRedactionPatternsCustomPattern:
		val := val.(*WorkspaceswebDataProtectionSettingsInlineRedactionConfigurationInlineRedactionPatternsCustomPattern)
		if err := _jsii_.ValidateStruct(val, func() string { return "parameter val" }); err != nil {
			return err
		}
	case WorkspaceswebDataProtectionSettingsInlineRedactionConfigurationInlineRedactionPatternsCustomPattern:
		val_ := val.(WorkspaceswebDataProtectionSettingsInlineRedactionConfigurationInlineRedactionPatternsCustomPattern)
		val := &val_
		if err := _jsii_.ValidateStruct(val, func() string { return "parameter val" }); err != nil {
			return err
		}
	default:
		if !_jsii_.IsAnonymousProxy(val) {
			return fmt.Errorf("parameter val must be one of the allowed types: cdktf.IResolvable, *WorkspaceswebDataProtectionSettingsInlineRedactionConfigurationInlineRedactionPatternsCustomPattern; received %#v (a %T)", val, val)
		}
	}

	return nil
}

func (j *jsiiProxy_WorkspaceswebDataProtectionSettingsInlineRedactionConfigurationInlineRedactionPatternsCustomPatternOutputReference) validateSetKeywordRegexParameters(val *string) error {
	if val == nil {
		return fmt.Errorf("parameter val is required, but nil was provided")
	}

	return nil
}

func (j *jsiiProxy_WorkspaceswebDataProtectionSettingsInlineRedactionConfigurationInlineRedactionPatternsCustomPatternOutputReference) validateSetPatternDescriptionParameters(val *string) error {
	if val == nil {
		return fmt.Errorf("parameter val is required, but nil was provided")
	}

	return nil
}

func (j *jsiiProxy_WorkspaceswebDataProtectionSettingsInlineRedactionConfigurationInlineRedactionPatternsCustomPatternOutputReference) validateSetPatternNameParameters(val *string) error {
	if val == nil {
		return fmt.Errorf("parameter val is required, but nil was provided")
	}

	return nil
}

func (j *jsiiProxy_WorkspaceswebDataProtectionSettingsInlineRedactionConfigurationInlineRedactionPatternsCustomPatternOutputReference) validateSetPatternRegexParameters(val *string) error {
	if val == nil {
		return fmt.Errorf("parameter val is required, but nil was provided")
	}

	return nil
}

func (j *jsiiProxy_WorkspaceswebDataProtectionSettingsInlineRedactionConfigurationInlineRedactionPatternsCustomPatternOutputReference) validateSetTerraformAttributeParameters(val *string) error {
	if val == nil {
		return fmt.Errorf("parameter val is required, but nil was provided")
	}

	return nil
}

func (j *jsiiProxy_WorkspaceswebDataProtectionSettingsInlineRedactionConfigurationInlineRedactionPatternsCustomPatternOutputReference) validateSetTerraformResourceParameters(val cdktf.IInterpolatingParent) error {
	if val == nil {
		return fmt.Errorf("parameter val is required, but nil was provided")
	}

	return nil
}

func validateNewWorkspaceswebDataProtectionSettingsInlineRedactionConfigurationInlineRedactionPatternsCustomPatternOutputReferenceParameters(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) error {
	if terraformResource == nil {
		return fmt.Errorf("parameter terraformResource is required, but nil was provided")
	}

	if terraformAttribute == nil {
		return fmt.Errorf("parameter terraformAttribute is required, but nil was provided")
	}

	return nil
}

