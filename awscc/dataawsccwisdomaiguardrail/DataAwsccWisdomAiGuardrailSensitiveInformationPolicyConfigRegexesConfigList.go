package dataawsccwisdomaiguardrail

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/unimelb-platform/cdktf-provider-awscc-go/awscc/jsii"

	"github.com/hashicorp/terraform-cdk-go/cdktf"
	"github.com/unimelb-platform/cdktf-provider-awscc-go/awscc/dataawsccwisdomaiguardrail/internal"
)

type DataAwsccWisdomAiGuardrailSensitiveInformationPolicyConfigRegexesConfigList interface {
	cdktf.ComplexList
	// The creation stack of this resolvable which will be appended to errors thrown during resolution.
	//
	// If this returns an empty array the stack will not be attached.
	// Experimental.
	CreationStack() *[]*string
	// Experimental.
	Fqn() *string
	// The attribute on the parent resource this class is referencing.
	TerraformAttribute() *string
	SetTerraformAttribute(val *string)
	// The parent resource.
	TerraformResource() cdktf.IInterpolatingParent
	SetTerraformResource(val cdktf.IInterpolatingParent)
	// whether the list is wrapping a set (will add tolist() to be able to access an item via an index).
	WrapsSet() *bool
	SetWrapsSet(val *bool)
	// Creating an iterator for this complex list.
	//
	// The list will be converted into a map with the mapKeyAttributeName as the key.
	// Experimental.
	AllWithMapKey(mapKeyAttributeName *string) cdktf.DynamicListTerraformIterator
	// Experimental.
	ComputeFqn() *string
	Get(index *float64) DataAwsccWisdomAiGuardrailSensitiveInformationPolicyConfigRegexesConfigOutputReference
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(_context cdktf.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for DataAwsccWisdomAiGuardrailSensitiveInformationPolicyConfigRegexesConfigList
type jsiiProxy_DataAwsccWisdomAiGuardrailSensitiveInformationPolicyConfigRegexesConfigList struct {
	internal.Type__cdktfComplexList
}

func (j *jsiiProxy_DataAwsccWisdomAiGuardrailSensitiveInformationPolicyConfigRegexesConfigList) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsccWisdomAiGuardrailSensitiveInformationPolicyConfigRegexesConfigList) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsccWisdomAiGuardrailSensitiveInformationPolicyConfigRegexesConfigList) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsccWisdomAiGuardrailSensitiveInformationPolicyConfigRegexesConfigList) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsccWisdomAiGuardrailSensitiveInformationPolicyConfigRegexesConfigList) WrapsSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"wrapsSet",
		&returns,
	)
	return returns
}


func NewDataAwsccWisdomAiGuardrailSensitiveInformationPolicyConfigRegexesConfigList(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, wrapsSet *bool) DataAwsccWisdomAiGuardrailSensitiveInformationPolicyConfigRegexesConfigList {
	_init_.Initialize()

	if err := validateNewDataAwsccWisdomAiGuardrailSensitiveInformationPolicyConfigRegexesConfigListParameters(terraformResource, terraformAttribute, wrapsSet); err != nil {
		panic(err)
	}
	j := jsiiProxy_DataAwsccWisdomAiGuardrailSensitiveInformationPolicyConfigRegexesConfigList{}

	_jsii_.Create(
		"awscc.dataAwsccWisdomAiGuardrail.DataAwsccWisdomAiGuardrailSensitiveInformationPolicyConfigRegexesConfigList",
		[]interface{}{terraformResource, terraformAttribute, wrapsSet},
		&j,
	)

	return &j
}

func NewDataAwsccWisdomAiGuardrailSensitiveInformationPolicyConfigRegexesConfigList_Override(d DataAwsccWisdomAiGuardrailSensitiveInformationPolicyConfigRegexesConfigList, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, wrapsSet *bool) {
	_init_.Initialize()

	_jsii_.Create(
		"awscc.dataAwsccWisdomAiGuardrail.DataAwsccWisdomAiGuardrailSensitiveInformationPolicyConfigRegexesConfigList",
		[]interface{}{terraformResource, terraformAttribute, wrapsSet},
		d,
	)
}

func (j *jsiiProxy_DataAwsccWisdomAiGuardrailSensitiveInformationPolicyConfigRegexesConfigList)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_DataAwsccWisdomAiGuardrailSensitiveInformationPolicyConfigRegexesConfigList)SetTerraformResource(val cdktf.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (j *jsiiProxy_DataAwsccWisdomAiGuardrailSensitiveInformationPolicyConfigRegexesConfigList)SetWrapsSet(val *bool) {
	if err := j.validateSetWrapsSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"wrapsSet",
		val,
	)
}

func (d *jsiiProxy_DataAwsccWisdomAiGuardrailSensitiveInformationPolicyConfigRegexesConfigList) AllWithMapKey(mapKeyAttributeName *string) cdktf.DynamicListTerraformIterator {
	if err := d.validateAllWithMapKeyParameters(mapKeyAttributeName); err != nil {
		panic(err)
	}
	var returns cdktf.DynamicListTerraformIterator

	_jsii_.Invoke(
		d,
		"allWithMapKey",
		[]interface{}{mapKeyAttributeName},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataAwsccWisdomAiGuardrailSensitiveInformationPolicyConfigRegexesConfigList) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		d,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataAwsccWisdomAiGuardrailSensitiveInformationPolicyConfigRegexesConfigList) Get(index *float64) DataAwsccWisdomAiGuardrailSensitiveInformationPolicyConfigRegexesConfigOutputReference {
	if err := d.validateGetParameters(index); err != nil {
		panic(err)
	}
	var returns DataAwsccWisdomAiGuardrailSensitiveInformationPolicyConfigRegexesConfigOutputReference

	_jsii_.Invoke(
		d,
		"get",
		[]interface{}{index},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataAwsccWisdomAiGuardrailSensitiveInformationPolicyConfigRegexesConfigList) Resolve(_context cdktf.IResolveContext) interface{} {
	if err := d.validateResolveParameters(_context); err != nil {
		panic(err)
	}
	var returns interface{}

	_jsii_.Invoke(
		d,
		"resolve",
		[]interface{}{_context},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataAwsccWisdomAiGuardrailSensitiveInformationPolicyConfigRegexesConfigList) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		d,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

