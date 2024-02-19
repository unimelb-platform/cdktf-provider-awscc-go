//go:build !no_runtime_type_checking

package networkfirewallfirewallpolicy

import (
	"fmt"

	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

func (n *jsiiProxy_NetworkfirewallFirewallPolicyFirewallPolicyOutputReference) validateGetAnyMapAttributeParameters(terraformAttribute *string) error {
	if terraformAttribute == nil {
		return fmt.Errorf("parameter terraformAttribute is required, but nil was provided")
	}

	return nil
}

func (n *jsiiProxy_NetworkfirewallFirewallPolicyFirewallPolicyOutputReference) validateGetBooleanAttributeParameters(terraformAttribute *string) error {
	if terraformAttribute == nil {
		return fmt.Errorf("parameter terraformAttribute is required, but nil was provided")
	}

	return nil
}

func (n *jsiiProxy_NetworkfirewallFirewallPolicyFirewallPolicyOutputReference) validateGetBooleanMapAttributeParameters(terraformAttribute *string) error {
	if terraformAttribute == nil {
		return fmt.Errorf("parameter terraformAttribute is required, but nil was provided")
	}

	return nil
}

func (n *jsiiProxy_NetworkfirewallFirewallPolicyFirewallPolicyOutputReference) validateGetListAttributeParameters(terraformAttribute *string) error {
	if terraformAttribute == nil {
		return fmt.Errorf("parameter terraformAttribute is required, but nil was provided")
	}

	return nil
}

func (n *jsiiProxy_NetworkfirewallFirewallPolicyFirewallPolicyOutputReference) validateGetNumberAttributeParameters(terraformAttribute *string) error {
	if terraformAttribute == nil {
		return fmt.Errorf("parameter terraformAttribute is required, but nil was provided")
	}

	return nil
}

func (n *jsiiProxy_NetworkfirewallFirewallPolicyFirewallPolicyOutputReference) validateGetNumberListAttributeParameters(terraformAttribute *string) error {
	if terraformAttribute == nil {
		return fmt.Errorf("parameter terraformAttribute is required, but nil was provided")
	}

	return nil
}

func (n *jsiiProxy_NetworkfirewallFirewallPolicyFirewallPolicyOutputReference) validateGetNumberMapAttributeParameters(terraformAttribute *string) error {
	if terraformAttribute == nil {
		return fmt.Errorf("parameter terraformAttribute is required, but nil was provided")
	}

	return nil
}

func (n *jsiiProxy_NetworkfirewallFirewallPolicyFirewallPolicyOutputReference) validateGetStringAttributeParameters(terraformAttribute *string) error {
	if terraformAttribute == nil {
		return fmt.Errorf("parameter terraformAttribute is required, but nil was provided")
	}

	return nil
}

func (n *jsiiProxy_NetworkfirewallFirewallPolicyFirewallPolicyOutputReference) validateGetStringMapAttributeParameters(terraformAttribute *string) error {
	if terraformAttribute == nil {
		return fmt.Errorf("parameter terraformAttribute is required, but nil was provided")
	}

	return nil
}

func (n *jsiiProxy_NetworkfirewallFirewallPolicyFirewallPolicyOutputReference) validateInterpolationForAttributeParameters(property *string) error {
	if property == nil {
		return fmt.Errorf("parameter property is required, but nil was provided")
	}

	return nil
}

func (n *jsiiProxy_NetworkfirewallFirewallPolicyFirewallPolicyOutputReference) validatePutPolicyVariablesParameters(value *NetworkfirewallFirewallPolicyFirewallPolicyPolicyVariables) error {
	if value == nil {
		return fmt.Errorf("parameter value is required, but nil was provided")
	}
	if err := _jsii_.ValidateStruct(value, func() string { return "parameter value" }); err != nil {
		return err
	}

	return nil
}

func (n *jsiiProxy_NetworkfirewallFirewallPolicyFirewallPolicyOutputReference) validatePutStatefulEngineOptionsParameters(value *NetworkfirewallFirewallPolicyFirewallPolicyStatefulEngineOptions) error {
	if value == nil {
		return fmt.Errorf("parameter value is required, but nil was provided")
	}
	if err := _jsii_.ValidateStruct(value, func() string { return "parameter value" }); err != nil {
		return err
	}

	return nil
}

func (n *jsiiProxy_NetworkfirewallFirewallPolicyFirewallPolicyOutputReference) validatePutStatefulRuleGroupReferencesParameters(value interface{}) error {
	if value == nil {
		return fmt.Errorf("parameter value is required, but nil was provided")
	}
	switch value.(type) {
	case cdktf.IResolvable:
		// ok
	case *[]*NetworkfirewallFirewallPolicyFirewallPolicyStatefulRuleGroupReferences:
		value := value.(*[]*NetworkfirewallFirewallPolicyFirewallPolicyStatefulRuleGroupReferences)
		for idx_cd4240, v := range *value {
			if err := _jsii_.ValidateStruct(v, func() string { return fmt.Sprintf("parameter value[%#v]", idx_cd4240) }); err != nil {
				return err
			}
		}
	case []*NetworkfirewallFirewallPolicyFirewallPolicyStatefulRuleGroupReferences:
		value_ := value.([]*NetworkfirewallFirewallPolicyFirewallPolicyStatefulRuleGroupReferences)
		value := &value_
		for idx_cd4240, v := range *value {
			if err := _jsii_.ValidateStruct(v, func() string { return fmt.Sprintf("parameter value[%#v]", idx_cd4240) }); err != nil {
				return err
			}
		}
	default:
		if !_jsii_.IsAnonymousProxy(value) {
			return fmt.Errorf("parameter value must be one of the allowed types: cdktf.IResolvable, *[]*NetworkfirewallFirewallPolicyFirewallPolicyStatefulRuleGroupReferences; received %#v (a %T)", value, value)
		}
	}

	return nil
}

func (n *jsiiProxy_NetworkfirewallFirewallPolicyFirewallPolicyOutputReference) validatePutStatelessCustomActionsParameters(value interface{}) error {
	if value == nil {
		return fmt.Errorf("parameter value is required, but nil was provided")
	}
	switch value.(type) {
	case cdktf.IResolvable:
		// ok
	case *[]*NetworkfirewallFirewallPolicyFirewallPolicyStatelessCustomActions:
		value := value.(*[]*NetworkfirewallFirewallPolicyFirewallPolicyStatelessCustomActions)
		for idx_cd4240, v := range *value {
			if err := _jsii_.ValidateStruct(v, func() string { return fmt.Sprintf("parameter value[%#v]", idx_cd4240) }); err != nil {
				return err
			}
		}
	case []*NetworkfirewallFirewallPolicyFirewallPolicyStatelessCustomActions:
		value_ := value.([]*NetworkfirewallFirewallPolicyFirewallPolicyStatelessCustomActions)
		value := &value_
		for idx_cd4240, v := range *value {
			if err := _jsii_.ValidateStruct(v, func() string { return fmt.Sprintf("parameter value[%#v]", idx_cd4240) }); err != nil {
				return err
			}
		}
	default:
		if !_jsii_.IsAnonymousProxy(value) {
			return fmt.Errorf("parameter value must be one of the allowed types: cdktf.IResolvable, *[]*NetworkfirewallFirewallPolicyFirewallPolicyStatelessCustomActions; received %#v (a %T)", value, value)
		}
	}

	return nil
}

func (n *jsiiProxy_NetworkfirewallFirewallPolicyFirewallPolicyOutputReference) validatePutStatelessRuleGroupReferencesParameters(value interface{}) error {
	if value == nil {
		return fmt.Errorf("parameter value is required, but nil was provided")
	}
	switch value.(type) {
	case cdktf.IResolvable:
		// ok
	case *[]*NetworkfirewallFirewallPolicyFirewallPolicyStatelessRuleGroupReferences:
		value := value.(*[]*NetworkfirewallFirewallPolicyFirewallPolicyStatelessRuleGroupReferences)
		for idx_cd4240, v := range *value {
			if err := _jsii_.ValidateStruct(v, func() string { return fmt.Sprintf("parameter value[%#v]", idx_cd4240) }); err != nil {
				return err
			}
		}
	case []*NetworkfirewallFirewallPolicyFirewallPolicyStatelessRuleGroupReferences:
		value_ := value.([]*NetworkfirewallFirewallPolicyFirewallPolicyStatelessRuleGroupReferences)
		value := &value_
		for idx_cd4240, v := range *value {
			if err := _jsii_.ValidateStruct(v, func() string { return fmt.Sprintf("parameter value[%#v]", idx_cd4240) }); err != nil {
				return err
			}
		}
	default:
		if !_jsii_.IsAnonymousProxy(value) {
			return fmt.Errorf("parameter value must be one of the allowed types: cdktf.IResolvable, *[]*NetworkfirewallFirewallPolicyFirewallPolicyStatelessRuleGroupReferences; received %#v (a %T)", value, value)
		}
	}

	return nil
}

func (n *jsiiProxy_NetworkfirewallFirewallPolicyFirewallPolicyOutputReference) validateResolveParameters(_context cdktf.IResolveContext) error {
	if _context == nil {
		return fmt.Errorf("parameter _context is required, but nil was provided")
	}

	return nil
}

func (j *jsiiProxy_NetworkfirewallFirewallPolicyFirewallPolicyOutputReference) validateSetComplexObjectIndexParameters(val interface{}) error {
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

func (j *jsiiProxy_NetworkfirewallFirewallPolicyFirewallPolicyOutputReference) validateSetComplexObjectIsFromSetParameters(val *bool) error {
	if val == nil {
		return fmt.Errorf("parameter val is required, but nil was provided")
	}

	return nil
}

func (j *jsiiProxy_NetworkfirewallFirewallPolicyFirewallPolicyOutputReference) validateSetInternalValueParameters(val interface{}) error {
	switch val.(type) {
	case cdktf.IResolvable:
		// ok
	case *NetworkfirewallFirewallPolicyFirewallPolicy:
		val := val.(*NetworkfirewallFirewallPolicyFirewallPolicy)
		if err := _jsii_.ValidateStruct(val, func() string { return "parameter val" }); err != nil {
			return err
		}
	case NetworkfirewallFirewallPolicyFirewallPolicy:
		val_ := val.(NetworkfirewallFirewallPolicyFirewallPolicy)
		val := &val_
		if err := _jsii_.ValidateStruct(val, func() string { return "parameter val" }); err != nil {
			return err
		}
	default:
		if !_jsii_.IsAnonymousProxy(val) {
			return fmt.Errorf("parameter val must be one of the allowed types: cdktf.IResolvable, *NetworkfirewallFirewallPolicyFirewallPolicy; received %#v (a %T)", val, val)
		}
	}

	return nil
}

func (j *jsiiProxy_NetworkfirewallFirewallPolicyFirewallPolicyOutputReference) validateSetStatefulDefaultActionsParameters(val *[]*string) error {
	if val == nil {
		return fmt.Errorf("parameter val is required, but nil was provided")
	}

	return nil
}

func (j *jsiiProxy_NetworkfirewallFirewallPolicyFirewallPolicyOutputReference) validateSetStatelessDefaultActionsParameters(val *[]*string) error {
	if val == nil {
		return fmt.Errorf("parameter val is required, but nil was provided")
	}

	return nil
}

func (j *jsiiProxy_NetworkfirewallFirewallPolicyFirewallPolicyOutputReference) validateSetStatelessFragmentDefaultActionsParameters(val *[]*string) error {
	if val == nil {
		return fmt.Errorf("parameter val is required, but nil was provided")
	}

	return nil
}

func (j *jsiiProxy_NetworkfirewallFirewallPolicyFirewallPolicyOutputReference) validateSetTerraformAttributeParameters(val *string) error {
	if val == nil {
		return fmt.Errorf("parameter val is required, but nil was provided")
	}

	return nil
}

func (j *jsiiProxy_NetworkfirewallFirewallPolicyFirewallPolicyOutputReference) validateSetTerraformResourceParameters(val cdktf.IInterpolatingParent) error {
	if val == nil {
		return fmt.Errorf("parameter val is required, but nil was provided")
	}

	return nil
}

func (j *jsiiProxy_NetworkfirewallFirewallPolicyFirewallPolicyOutputReference) validateSetTlsInspectionConfigurationArnParameters(val *string) error {
	if val == nil {
		return fmt.Errorf("parameter val is required, but nil was provided")
	}

	return nil
}

func validateNewNetworkfirewallFirewallPolicyFirewallPolicyOutputReferenceParameters(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) error {
	if terraformResource == nil {
		return fmt.Errorf("parameter terraformResource is required, but nil was provided")
	}

	if terraformAttribute == nil {
		return fmt.Errorf("parameter terraformAttribute is required, but nil was provided")
	}

	return nil
}

