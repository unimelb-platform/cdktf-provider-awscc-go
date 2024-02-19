//go:build !no_runtime_type_checking

package dataawsccnetworkfirewallrulegroup

import (
	"fmt"

	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

func (d *jsiiProxy_DataAwsccNetworkfirewallRuleGroupRuleGroupRuleVariablesPortSetsMap) validateGetParameters(key *string) error {
	if key == nil {
		return fmt.Errorf("parameter key is required, but nil was provided")
	}

	return nil
}

func (d *jsiiProxy_DataAwsccNetworkfirewallRuleGroupRuleGroupRuleVariablesPortSetsMap) validateInterpolationForAttributeParameters(property *string) error {
	if property == nil {
		return fmt.Errorf("parameter property is required, but nil was provided")
	}

	return nil
}

func (d *jsiiProxy_DataAwsccNetworkfirewallRuleGroupRuleGroupRuleVariablesPortSetsMap) validateResolveParameters(_context cdktf.IResolveContext) error {
	if _context == nil {
		return fmt.Errorf("parameter _context is required, but nil was provided")
	}

	return nil
}

func (j *jsiiProxy_DataAwsccNetworkfirewallRuleGroupRuleGroupRuleVariablesPortSetsMap) validateSetTerraformAttributeParameters(val *string) error {
	if val == nil {
		return fmt.Errorf("parameter val is required, but nil was provided")
	}

	return nil
}

func (j *jsiiProxy_DataAwsccNetworkfirewallRuleGroupRuleGroupRuleVariablesPortSetsMap) validateSetTerraformResourceParameters(val cdktf.IInterpolatingParent) error {
	if val == nil {
		return fmt.Errorf("parameter val is required, but nil was provided")
	}

	return nil
}

func validateNewDataAwsccNetworkfirewallRuleGroupRuleGroupRuleVariablesPortSetsMapParameters(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) error {
	if terraformResource == nil {
		return fmt.Errorf("parameter terraformResource is required, but nil was provided")
	}

	if terraformAttribute == nil {
		return fmt.Errorf("parameter terraformAttribute is required, but nil was provided")
	}

	return nil
}

