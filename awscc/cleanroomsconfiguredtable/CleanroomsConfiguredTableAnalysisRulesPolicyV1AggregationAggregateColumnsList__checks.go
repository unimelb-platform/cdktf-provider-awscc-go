//go:build !no_runtime_type_checking

package cleanroomsconfiguredtable

import (
	"fmt"

	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

func (c *jsiiProxy_CleanroomsConfiguredTableAnalysisRulesPolicyV1AggregationAggregateColumnsList) validateGetParameters(index *float64) error {
	if index == nil {
		return fmt.Errorf("parameter index is required, but nil was provided")
	}

	return nil
}

func (c *jsiiProxy_CleanroomsConfiguredTableAnalysisRulesPolicyV1AggregationAggregateColumnsList) validateResolveParameters(_context cdktf.IResolveContext) error {
	if _context == nil {
		return fmt.Errorf("parameter _context is required, but nil was provided")
	}

	return nil
}

func (j *jsiiProxy_CleanroomsConfiguredTableAnalysisRulesPolicyV1AggregationAggregateColumnsList) validateSetInternalValueParameters(val interface{}) error {
	switch val.(type) {
	case cdktf.IResolvable:
		// ok
	case *[]*CleanroomsConfiguredTableAnalysisRulesPolicyV1AggregationAggregateColumns:
		val := val.(*[]*CleanroomsConfiguredTableAnalysisRulesPolicyV1AggregationAggregateColumns)
		for idx_97dfc6, v := range *val {
			if err := _jsii_.ValidateStruct(v, func() string { return fmt.Sprintf("parameter val[%#v]", idx_97dfc6) }); err != nil {
				return err
			}
		}
	case []*CleanroomsConfiguredTableAnalysisRulesPolicyV1AggregationAggregateColumns:
		val_ := val.([]*CleanroomsConfiguredTableAnalysisRulesPolicyV1AggregationAggregateColumns)
		val := &val_
		for idx_97dfc6, v := range *val {
			if err := _jsii_.ValidateStruct(v, func() string { return fmt.Sprintf("parameter val[%#v]", idx_97dfc6) }); err != nil {
				return err
			}
		}
	default:
		if !_jsii_.IsAnonymousProxy(val) {
			return fmt.Errorf("parameter val must be one of the allowed types: cdktf.IResolvable, *[]*CleanroomsConfiguredTableAnalysisRulesPolicyV1AggregationAggregateColumns; received %#v (a %T)", val, val)
		}
	}

	return nil
}

func (j *jsiiProxy_CleanroomsConfiguredTableAnalysisRulesPolicyV1AggregationAggregateColumnsList) validateSetTerraformAttributeParameters(val *string) error {
	if val == nil {
		return fmt.Errorf("parameter val is required, but nil was provided")
	}

	return nil
}

func (j *jsiiProxy_CleanroomsConfiguredTableAnalysisRulesPolicyV1AggregationAggregateColumnsList) validateSetTerraformResourceParameters(val cdktf.IInterpolatingParent) error {
	if val == nil {
		return fmt.Errorf("parameter val is required, but nil was provided")
	}

	return nil
}

func (j *jsiiProxy_CleanroomsConfiguredTableAnalysisRulesPolicyV1AggregationAggregateColumnsList) validateSetWrapsSetParameters(val *bool) error {
	if val == nil {
		return fmt.Errorf("parameter val is required, but nil was provided")
	}

	return nil
}

func validateNewCleanroomsConfiguredTableAnalysisRulesPolicyV1AggregationAggregateColumnsListParameters(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, wrapsSet *bool) error {
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

