package dataawscclexbot

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/unimelb-platform/cdktf-provider-awscc-go/awscc/jsii"

	"github.com/hashicorp/terraform-cdk-go/cdktf"
	"github.com/unimelb-platform/cdktf-provider-awscc-go/awscc/dataawscclexbot/internal"
)

type DataAwsccLexBotBotLocalesIntentsSlotsValueElicitationSettingPromptSpecificationPromptAttemptsSpecificationMap interface {
	cdktf.ComplexMap
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
	// Experimental.
	ComputeFqn() *string
	Get(key *string) DataAwsccLexBotBotLocalesIntentsSlotsValueElicitationSettingPromptSpecificationPromptAttemptsSpecificationOutputReference
	// Experimental.
	InterpolationForAttribute(property *string) cdktf.IResolvable
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(_context cdktf.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for DataAwsccLexBotBotLocalesIntentsSlotsValueElicitationSettingPromptSpecificationPromptAttemptsSpecificationMap
type jsiiProxy_DataAwsccLexBotBotLocalesIntentsSlotsValueElicitationSettingPromptSpecificationPromptAttemptsSpecificationMap struct {
	internal.Type__cdktfComplexMap
}

func (j *jsiiProxy_DataAwsccLexBotBotLocalesIntentsSlotsValueElicitationSettingPromptSpecificationPromptAttemptsSpecificationMap) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsccLexBotBotLocalesIntentsSlotsValueElicitationSettingPromptSpecificationPromptAttemptsSpecificationMap) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsccLexBotBotLocalesIntentsSlotsValueElicitationSettingPromptSpecificationPromptAttemptsSpecificationMap) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsccLexBotBotLocalesIntentsSlotsValueElicitationSettingPromptSpecificationPromptAttemptsSpecificationMap) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}


func NewDataAwsccLexBotBotLocalesIntentsSlotsValueElicitationSettingPromptSpecificationPromptAttemptsSpecificationMap(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) DataAwsccLexBotBotLocalesIntentsSlotsValueElicitationSettingPromptSpecificationPromptAttemptsSpecificationMap {
	_init_.Initialize()

	if err := validateNewDataAwsccLexBotBotLocalesIntentsSlotsValueElicitationSettingPromptSpecificationPromptAttemptsSpecificationMapParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_DataAwsccLexBotBotLocalesIntentsSlotsValueElicitationSettingPromptSpecificationPromptAttemptsSpecificationMap{}

	_jsii_.Create(
		"awscc.dataAwsccLexBot.DataAwsccLexBotBotLocalesIntentsSlotsValueElicitationSettingPromptSpecificationPromptAttemptsSpecificationMap",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewDataAwsccLexBotBotLocalesIntentsSlotsValueElicitationSettingPromptSpecificationPromptAttemptsSpecificationMap_Override(d DataAwsccLexBotBotLocalesIntentsSlotsValueElicitationSettingPromptSpecificationPromptAttemptsSpecificationMap, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"awscc.dataAwsccLexBot.DataAwsccLexBotBotLocalesIntentsSlotsValueElicitationSettingPromptSpecificationPromptAttemptsSpecificationMap",
		[]interface{}{terraformResource, terraformAttribute},
		d,
	)
}

func (j *jsiiProxy_DataAwsccLexBotBotLocalesIntentsSlotsValueElicitationSettingPromptSpecificationPromptAttemptsSpecificationMap)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_DataAwsccLexBotBotLocalesIntentsSlotsValueElicitationSettingPromptSpecificationPromptAttemptsSpecificationMap)SetTerraformResource(val cdktf.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (d *jsiiProxy_DataAwsccLexBotBotLocalesIntentsSlotsValueElicitationSettingPromptSpecificationPromptAttemptsSpecificationMap) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		d,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataAwsccLexBotBotLocalesIntentsSlotsValueElicitationSettingPromptSpecificationPromptAttemptsSpecificationMap) Get(key *string) DataAwsccLexBotBotLocalesIntentsSlotsValueElicitationSettingPromptSpecificationPromptAttemptsSpecificationOutputReference {
	if err := d.validateGetParameters(key); err != nil {
		panic(err)
	}
	var returns DataAwsccLexBotBotLocalesIntentsSlotsValueElicitationSettingPromptSpecificationPromptAttemptsSpecificationOutputReference

	_jsii_.Invoke(
		d,
		"get",
		[]interface{}{key},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataAwsccLexBotBotLocalesIntentsSlotsValueElicitationSettingPromptSpecificationPromptAttemptsSpecificationMap) InterpolationForAttribute(property *string) cdktf.IResolvable {
	if err := d.validateInterpolationForAttributeParameters(property); err != nil {
		panic(err)
	}
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		d,
		"interpolationForAttribute",
		[]interface{}{property},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataAwsccLexBotBotLocalesIntentsSlotsValueElicitationSettingPromptSpecificationPromptAttemptsSpecificationMap) Resolve(_context cdktf.IResolveContext) interface{} {
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

func (d *jsiiProxy_DataAwsccLexBotBotLocalesIntentsSlotsValueElicitationSettingPromptSpecificationPromptAttemptsSpecificationMap) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		d,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

