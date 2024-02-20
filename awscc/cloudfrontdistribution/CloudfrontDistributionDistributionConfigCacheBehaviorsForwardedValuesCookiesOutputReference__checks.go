//go:build !no_runtime_type_checking

package cloudfrontdistribution

import (
	"fmt"

	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

func (c *jsiiProxy_CloudfrontDistributionDistributionConfigCacheBehaviorsForwardedValuesCookiesOutputReference) validateGetAnyMapAttributeParameters(terraformAttribute *string) error {
	if terraformAttribute == nil {
		return fmt.Errorf("parameter terraformAttribute is required, but nil was provided")
	}

	return nil
}

func (c *jsiiProxy_CloudfrontDistributionDistributionConfigCacheBehaviorsForwardedValuesCookiesOutputReference) validateGetBooleanAttributeParameters(terraformAttribute *string) error {
	if terraformAttribute == nil {
		return fmt.Errorf("parameter terraformAttribute is required, but nil was provided")
	}

	return nil
}

func (c *jsiiProxy_CloudfrontDistributionDistributionConfigCacheBehaviorsForwardedValuesCookiesOutputReference) validateGetBooleanMapAttributeParameters(terraformAttribute *string) error {
	if terraformAttribute == nil {
		return fmt.Errorf("parameter terraformAttribute is required, but nil was provided")
	}

	return nil
}

func (c *jsiiProxy_CloudfrontDistributionDistributionConfigCacheBehaviorsForwardedValuesCookiesOutputReference) validateGetListAttributeParameters(terraformAttribute *string) error {
	if terraformAttribute == nil {
		return fmt.Errorf("parameter terraformAttribute is required, but nil was provided")
	}

	return nil
}

func (c *jsiiProxy_CloudfrontDistributionDistributionConfigCacheBehaviorsForwardedValuesCookiesOutputReference) validateGetNumberAttributeParameters(terraformAttribute *string) error {
	if terraformAttribute == nil {
		return fmt.Errorf("parameter terraformAttribute is required, but nil was provided")
	}

	return nil
}

func (c *jsiiProxy_CloudfrontDistributionDistributionConfigCacheBehaviorsForwardedValuesCookiesOutputReference) validateGetNumberListAttributeParameters(terraformAttribute *string) error {
	if terraformAttribute == nil {
		return fmt.Errorf("parameter terraformAttribute is required, but nil was provided")
	}

	return nil
}

func (c *jsiiProxy_CloudfrontDistributionDistributionConfigCacheBehaviorsForwardedValuesCookiesOutputReference) validateGetNumberMapAttributeParameters(terraformAttribute *string) error {
	if terraformAttribute == nil {
		return fmt.Errorf("parameter terraformAttribute is required, but nil was provided")
	}

	return nil
}

func (c *jsiiProxy_CloudfrontDistributionDistributionConfigCacheBehaviorsForwardedValuesCookiesOutputReference) validateGetStringAttributeParameters(terraformAttribute *string) error {
	if terraformAttribute == nil {
		return fmt.Errorf("parameter terraformAttribute is required, but nil was provided")
	}

	return nil
}

func (c *jsiiProxy_CloudfrontDistributionDistributionConfigCacheBehaviorsForwardedValuesCookiesOutputReference) validateGetStringMapAttributeParameters(terraformAttribute *string) error {
	if terraformAttribute == nil {
		return fmt.Errorf("parameter terraformAttribute is required, but nil was provided")
	}

	return nil
}

func (c *jsiiProxy_CloudfrontDistributionDistributionConfigCacheBehaviorsForwardedValuesCookiesOutputReference) validateInterpolationForAttributeParameters(property *string) error {
	if property == nil {
		return fmt.Errorf("parameter property is required, but nil was provided")
	}

	return nil
}

func (c *jsiiProxy_CloudfrontDistributionDistributionConfigCacheBehaviorsForwardedValuesCookiesOutputReference) validateResolveParameters(_context cdktf.IResolveContext) error {
	if _context == nil {
		return fmt.Errorf("parameter _context is required, but nil was provided")
	}

	return nil
}

func (j *jsiiProxy_CloudfrontDistributionDistributionConfigCacheBehaviorsForwardedValuesCookiesOutputReference) validateSetComplexObjectIndexParameters(val interface{}) error {
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

func (j *jsiiProxy_CloudfrontDistributionDistributionConfigCacheBehaviorsForwardedValuesCookiesOutputReference) validateSetComplexObjectIsFromSetParameters(val *bool) error {
	if val == nil {
		return fmt.Errorf("parameter val is required, but nil was provided")
	}

	return nil
}

func (j *jsiiProxy_CloudfrontDistributionDistributionConfigCacheBehaviorsForwardedValuesCookiesOutputReference) validateSetForwardParameters(val *string) error {
	if val == nil {
		return fmt.Errorf("parameter val is required, but nil was provided")
	}

	return nil
}

func (j *jsiiProxy_CloudfrontDistributionDistributionConfigCacheBehaviorsForwardedValuesCookiesOutputReference) validateSetInternalValueParameters(val interface{}) error {
	switch val.(type) {
	case cdktf.IResolvable:
		// ok
	case *CloudfrontDistributionDistributionConfigCacheBehaviorsForwardedValuesCookies:
		val := val.(*CloudfrontDistributionDistributionConfigCacheBehaviorsForwardedValuesCookies)
		if err := _jsii_.ValidateStruct(val, func() string { return "parameter val" }); err != nil {
			return err
		}
	case CloudfrontDistributionDistributionConfigCacheBehaviorsForwardedValuesCookies:
		val_ := val.(CloudfrontDistributionDistributionConfigCacheBehaviorsForwardedValuesCookies)
		val := &val_
		if err := _jsii_.ValidateStruct(val, func() string { return "parameter val" }); err != nil {
			return err
		}
	default:
		if !_jsii_.IsAnonymousProxy(val) {
			return fmt.Errorf("parameter val must be one of the allowed types: cdktf.IResolvable, *CloudfrontDistributionDistributionConfigCacheBehaviorsForwardedValuesCookies; received %#v (a %T)", val, val)
		}
	}

	return nil
}

func (j *jsiiProxy_CloudfrontDistributionDistributionConfigCacheBehaviorsForwardedValuesCookiesOutputReference) validateSetTerraformAttributeParameters(val *string) error {
	if val == nil {
		return fmt.Errorf("parameter val is required, but nil was provided")
	}

	return nil
}

func (j *jsiiProxy_CloudfrontDistributionDistributionConfigCacheBehaviorsForwardedValuesCookiesOutputReference) validateSetTerraformResourceParameters(val cdktf.IInterpolatingParent) error {
	if val == nil {
		return fmt.Errorf("parameter val is required, but nil was provided")
	}

	return nil
}

func (j *jsiiProxy_CloudfrontDistributionDistributionConfigCacheBehaviorsForwardedValuesCookiesOutputReference) validateSetWhitelistedNamesParameters(val *[]*string) error {
	if val == nil {
		return fmt.Errorf("parameter val is required, but nil was provided")
	}

	return nil
}

func validateNewCloudfrontDistributionDistributionConfigCacheBehaviorsForwardedValuesCookiesOutputReferenceParameters(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) error {
	if terraformResource == nil {
		return fmt.Errorf("parameter terraformResource is required, but nil was provided")
	}

	if terraformAttribute == nil {
		return fmt.Errorf("parameter terraformAttribute is required, but nil was provided")
	}

	return nil
}

