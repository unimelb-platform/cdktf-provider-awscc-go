package dataawsccioteventsdetectormodel

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/unimelb-platform/cdktf-provider-awscc-go/awscc/jsii"

	"github.com/hashicorp/terraform-cdk-go/cdktf"
	"github.com/unimelb-platform/cdktf-provider-awscc-go/awscc/dataawsccioteventsdetectormodel/internal"
)

type DataAwsccIoteventsDetectorModelDetectorModelDefinitionStatesOnInputTransitionEventsActionsOutputReference interface {
	cdktf.ComplexObject
	ClearTimer() DataAwsccIoteventsDetectorModelDetectorModelDefinitionStatesOnInputTransitionEventsActionsClearTimerOutputReference
	// the index of the complex object in a list.
	// Experimental.
	ComplexObjectIndex() interface{}
	// Experimental.
	SetComplexObjectIndex(val interface{})
	// set to true if this item is from inside a set and needs tolist() for accessing it set to "0" for single list items.
	// Experimental.
	ComplexObjectIsFromSet() *bool
	// Experimental.
	SetComplexObjectIsFromSet(val *bool)
	// The creation stack of this resolvable which will be appended to errors thrown during resolution.
	//
	// If this returns an empty array the stack will not be attached.
	// Experimental.
	CreationStack() *[]*string
	DynamoDb() DataAwsccIoteventsDetectorModelDetectorModelDefinitionStatesOnInputTransitionEventsActionsDynamoDbOutputReference
	DynamoDBv2() DataAwsccIoteventsDetectorModelDetectorModelDefinitionStatesOnInputTransitionEventsActionsDynamoDBv2OutputReference
	Firehose() DataAwsccIoteventsDetectorModelDetectorModelDefinitionStatesOnInputTransitionEventsActionsFirehoseOutputReference
	// Experimental.
	Fqn() *string
	InternalValue() *DataAwsccIoteventsDetectorModelDetectorModelDefinitionStatesOnInputTransitionEventsActions
	SetInternalValue(val *DataAwsccIoteventsDetectorModelDetectorModelDefinitionStatesOnInputTransitionEventsActions)
	IotEvents() DataAwsccIoteventsDetectorModelDetectorModelDefinitionStatesOnInputTransitionEventsActionsIotEventsOutputReference
	IotSiteWise() DataAwsccIoteventsDetectorModelDetectorModelDefinitionStatesOnInputTransitionEventsActionsIotSiteWiseOutputReference
	IotTopicPublish() DataAwsccIoteventsDetectorModelDetectorModelDefinitionStatesOnInputTransitionEventsActionsIotTopicPublishOutputReference
	Lambda() DataAwsccIoteventsDetectorModelDetectorModelDefinitionStatesOnInputTransitionEventsActionsLambdaOutputReference
	ResetTimer() DataAwsccIoteventsDetectorModelDetectorModelDefinitionStatesOnInputTransitionEventsActionsResetTimerOutputReference
	SetTimer() DataAwsccIoteventsDetectorModelDetectorModelDefinitionStatesOnInputTransitionEventsActionsSetTimerOutputReference
	SetVariable() DataAwsccIoteventsDetectorModelDetectorModelDefinitionStatesOnInputTransitionEventsActionsSetVariableOutputReference
	Sns() DataAwsccIoteventsDetectorModelDetectorModelDefinitionStatesOnInputTransitionEventsActionsSnsOutputReference
	Sqs() DataAwsccIoteventsDetectorModelDetectorModelDefinitionStatesOnInputTransitionEventsActionsSqsOutputReference
	// Experimental.
	TerraformAttribute() *string
	// Experimental.
	SetTerraformAttribute(val *string)
	// Experimental.
	TerraformResource() cdktf.IInterpolatingParent
	// Experimental.
	SetTerraformResource(val cdktf.IInterpolatingParent)
	// Experimental.
	ComputeFqn() *string
	// Experimental.
	GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{}
	// Experimental.
	GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable
	// Experimental.
	GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool
	// Experimental.
	GetListAttribute(terraformAttribute *string) *[]*string
	// Experimental.
	GetNumberAttribute(terraformAttribute *string) *float64
	// Experimental.
	GetNumberListAttribute(terraformAttribute *string) *[]*float64
	// Experimental.
	GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64
	// Experimental.
	GetStringAttribute(terraformAttribute *string) *string
	// Experimental.
	GetStringMapAttribute(terraformAttribute *string) *map[string]*string
	// Experimental.
	InterpolationAsList() cdktf.IResolvable
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

// The jsii proxy struct for DataAwsccIoteventsDetectorModelDetectorModelDefinitionStatesOnInputTransitionEventsActionsOutputReference
type jsiiProxy_DataAwsccIoteventsDetectorModelDetectorModelDefinitionStatesOnInputTransitionEventsActionsOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_DataAwsccIoteventsDetectorModelDetectorModelDefinitionStatesOnInputTransitionEventsActionsOutputReference) ClearTimer() DataAwsccIoteventsDetectorModelDetectorModelDefinitionStatesOnInputTransitionEventsActionsClearTimerOutputReference {
	var returns DataAwsccIoteventsDetectorModelDetectorModelDefinitionStatesOnInputTransitionEventsActionsClearTimerOutputReference
	_jsii_.Get(
		j,
		"clearTimer",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsccIoteventsDetectorModelDetectorModelDefinitionStatesOnInputTransitionEventsActionsOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsccIoteventsDetectorModelDetectorModelDefinitionStatesOnInputTransitionEventsActionsOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsccIoteventsDetectorModelDetectorModelDefinitionStatesOnInputTransitionEventsActionsOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsccIoteventsDetectorModelDetectorModelDefinitionStatesOnInputTransitionEventsActionsOutputReference) DynamoDb() DataAwsccIoteventsDetectorModelDetectorModelDefinitionStatesOnInputTransitionEventsActionsDynamoDbOutputReference {
	var returns DataAwsccIoteventsDetectorModelDetectorModelDefinitionStatesOnInputTransitionEventsActionsDynamoDbOutputReference
	_jsii_.Get(
		j,
		"dynamoDb",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsccIoteventsDetectorModelDetectorModelDefinitionStatesOnInputTransitionEventsActionsOutputReference) DynamoDBv2() DataAwsccIoteventsDetectorModelDetectorModelDefinitionStatesOnInputTransitionEventsActionsDynamoDBv2OutputReference {
	var returns DataAwsccIoteventsDetectorModelDetectorModelDefinitionStatesOnInputTransitionEventsActionsDynamoDBv2OutputReference
	_jsii_.Get(
		j,
		"dynamoDBv2",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsccIoteventsDetectorModelDetectorModelDefinitionStatesOnInputTransitionEventsActionsOutputReference) Firehose() DataAwsccIoteventsDetectorModelDetectorModelDefinitionStatesOnInputTransitionEventsActionsFirehoseOutputReference {
	var returns DataAwsccIoteventsDetectorModelDetectorModelDefinitionStatesOnInputTransitionEventsActionsFirehoseOutputReference
	_jsii_.Get(
		j,
		"firehose",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsccIoteventsDetectorModelDetectorModelDefinitionStatesOnInputTransitionEventsActionsOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsccIoteventsDetectorModelDetectorModelDefinitionStatesOnInputTransitionEventsActionsOutputReference) InternalValue() *DataAwsccIoteventsDetectorModelDetectorModelDefinitionStatesOnInputTransitionEventsActions {
	var returns *DataAwsccIoteventsDetectorModelDetectorModelDefinitionStatesOnInputTransitionEventsActions
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsccIoteventsDetectorModelDetectorModelDefinitionStatesOnInputTransitionEventsActionsOutputReference) IotEvents() DataAwsccIoteventsDetectorModelDetectorModelDefinitionStatesOnInputTransitionEventsActionsIotEventsOutputReference {
	var returns DataAwsccIoteventsDetectorModelDetectorModelDefinitionStatesOnInputTransitionEventsActionsIotEventsOutputReference
	_jsii_.Get(
		j,
		"iotEvents",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsccIoteventsDetectorModelDetectorModelDefinitionStatesOnInputTransitionEventsActionsOutputReference) IotSiteWise() DataAwsccIoteventsDetectorModelDetectorModelDefinitionStatesOnInputTransitionEventsActionsIotSiteWiseOutputReference {
	var returns DataAwsccIoteventsDetectorModelDetectorModelDefinitionStatesOnInputTransitionEventsActionsIotSiteWiseOutputReference
	_jsii_.Get(
		j,
		"iotSiteWise",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsccIoteventsDetectorModelDetectorModelDefinitionStatesOnInputTransitionEventsActionsOutputReference) IotTopicPublish() DataAwsccIoteventsDetectorModelDetectorModelDefinitionStatesOnInputTransitionEventsActionsIotTopicPublishOutputReference {
	var returns DataAwsccIoteventsDetectorModelDetectorModelDefinitionStatesOnInputTransitionEventsActionsIotTopicPublishOutputReference
	_jsii_.Get(
		j,
		"iotTopicPublish",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsccIoteventsDetectorModelDetectorModelDefinitionStatesOnInputTransitionEventsActionsOutputReference) Lambda() DataAwsccIoteventsDetectorModelDetectorModelDefinitionStatesOnInputTransitionEventsActionsLambdaOutputReference {
	var returns DataAwsccIoteventsDetectorModelDetectorModelDefinitionStatesOnInputTransitionEventsActionsLambdaOutputReference
	_jsii_.Get(
		j,
		"lambda",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsccIoteventsDetectorModelDetectorModelDefinitionStatesOnInputTransitionEventsActionsOutputReference) ResetTimer() DataAwsccIoteventsDetectorModelDetectorModelDefinitionStatesOnInputTransitionEventsActionsResetTimerOutputReference {
	var returns DataAwsccIoteventsDetectorModelDetectorModelDefinitionStatesOnInputTransitionEventsActionsResetTimerOutputReference
	_jsii_.Get(
		j,
		"resetTimer",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsccIoteventsDetectorModelDetectorModelDefinitionStatesOnInputTransitionEventsActionsOutputReference) SetTimer() DataAwsccIoteventsDetectorModelDetectorModelDefinitionStatesOnInputTransitionEventsActionsSetTimerOutputReference {
	var returns DataAwsccIoteventsDetectorModelDetectorModelDefinitionStatesOnInputTransitionEventsActionsSetTimerOutputReference
	_jsii_.Get(
		j,
		"setTimer",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsccIoteventsDetectorModelDetectorModelDefinitionStatesOnInputTransitionEventsActionsOutputReference) SetVariable() DataAwsccIoteventsDetectorModelDetectorModelDefinitionStatesOnInputTransitionEventsActionsSetVariableOutputReference {
	var returns DataAwsccIoteventsDetectorModelDetectorModelDefinitionStatesOnInputTransitionEventsActionsSetVariableOutputReference
	_jsii_.Get(
		j,
		"setVariable",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsccIoteventsDetectorModelDetectorModelDefinitionStatesOnInputTransitionEventsActionsOutputReference) Sns() DataAwsccIoteventsDetectorModelDetectorModelDefinitionStatesOnInputTransitionEventsActionsSnsOutputReference {
	var returns DataAwsccIoteventsDetectorModelDetectorModelDefinitionStatesOnInputTransitionEventsActionsSnsOutputReference
	_jsii_.Get(
		j,
		"sns",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsccIoteventsDetectorModelDetectorModelDefinitionStatesOnInputTransitionEventsActionsOutputReference) Sqs() DataAwsccIoteventsDetectorModelDetectorModelDefinitionStatesOnInputTransitionEventsActionsSqsOutputReference {
	var returns DataAwsccIoteventsDetectorModelDetectorModelDefinitionStatesOnInputTransitionEventsActionsSqsOutputReference
	_jsii_.Get(
		j,
		"sqs",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsccIoteventsDetectorModelDetectorModelDefinitionStatesOnInputTransitionEventsActionsOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsccIoteventsDetectorModelDetectorModelDefinitionStatesOnInputTransitionEventsActionsOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}


func NewDataAwsccIoteventsDetectorModelDetectorModelDefinitionStatesOnInputTransitionEventsActionsOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) DataAwsccIoteventsDetectorModelDetectorModelDefinitionStatesOnInputTransitionEventsActionsOutputReference {
	_init_.Initialize()

	if err := validateNewDataAwsccIoteventsDetectorModelDetectorModelDefinitionStatesOnInputTransitionEventsActionsOutputReferenceParameters(terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet); err != nil {
		panic(err)
	}
	j := jsiiProxy_DataAwsccIoteventsDetectorModelDetectorModelDefinitionStatesOnInputTransitionEventsActionsOutputReference{}

	_jsii_.Create(
		"awscc.dataAwsccIoteventsDetectorModel.DataAwsccIoteventsDetectorModelDetectorModelDefinitionStatesOnInputTransitionEventsActionsOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		&j,
	)

	return &j
}

func NewDataAwsccIoteventsDetectorModelDetectorModelDefinitionStatesOnInputTransitionEventsActionsOutputReference_Override(d DataAwsccIoteventsDetectorModelDetectorModelDefinitionStatesOnInputTransitionEventsActionsOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) {
	_init_.Initialize()

	_jsii_.Create(
		"awscc.dataAwsccIoteventsDetectorModel.DataAwsccIoteventsDetectorModelDetectorModelDefinitionStatesOnInputTransitionEventsActionsOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		d,
	)
}

func (j *jsiiProxy_DataAwsccIoteventsDetectorModelDetectorModelDefinitionStatesOnInputTransitionEventsActionsOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_DataAwsccIoteventsDetectorModelDetectorModelDefinitionStatesOnInputTransitionEventsActionsOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_DataAwsccIoteventsDetectorModelDetectorModelDefinitionStatesOnInputTransitionEventsActionsOutputReference)SetInternalValue(val *DataAwsccIoteventsDetectorModelDetectorModelDefinitionStatesOnInputTransitionEventsActions) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_DataAwsccIoteventsDetectorModelDetectorModelDefinitionStatesOnInputTransitionEventsActionsOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_DataAwsccIoteventsDetectorModelDetectorModelDefinitionStatesOnInputTransitionEventsActionsOutputReference)SetTerraformResource(val cdktf.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (d *jsiiProxy_DataAwsccIoteventsDetectorModelDetectorModelDefinitionStatesOnInputTransitionEventsActionsOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		d,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataAwsccIoteventsDetectorModelDetectorModelDefinitionStatesOnInputTransitionEventsActionsOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
	if err := d.validateGetAnyMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]interface{}

	_jsii_.Invoke(
		d,
		"getAnyMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataAwsccIoteventsDetectorModelDetectorModelDefinitionStatesOnInputTransitionEventsActionsOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
	if err := d.validateGetBooleanAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		d,
		"getBooleanAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataAwsccIoteventsDetectorModelDetectorModelDefinitionStatesOnInputTransitionEventsActionsOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
	if err := d.validateGetBooleanMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*bool

	_jsii_.Invoke(
		d,
		"getBooleanMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataAwsccIoteventsDetectorModelDetectorModelDefinitionStatesOnInputTransitionEventsActionsOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
	if err := d.validateGetListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*string

	_jsii_.Invoke(
		d,
		"getListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataAwsccIoteventsDetectorModelDetectorModelDefinitionStatesOnInputTransitionEventsActionsOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
	if err := d.validateGetNumberAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *float64

	_jsii_.Invoke(
		d,
		"getNumberAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataAwsccIoteventsDetectorModelDetectorModelDefinitionStatesOnInputTransitionEventsActionsOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
	if err := d.validateGetNumberListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*float64

	_jsii_.Invoke(
		d,
		"getNumberListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataAwsccIoteventsDetectorModelDetectorModelDefinitionStatesOnInputTransitionEventsActionsOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
	if err := d.validateGetNumberMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*float64

	_jsii_.Invoke(
		d,
		"getNumberMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataAwsccIoteventsDetectorModelDetectorModelDefinitionStatesOnInputTransitionEventsActionsOutputReference) GetStringAttribute(terraformAttribute *string) *string {
	if err := d.validateGetStringAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *string

	_jsii_.Invoke(
		d,
		"getStringAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataAwsccIoteventsDetectorModelDetectorModelDefinitionStatesOnInputTransitionEventsActionsOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
	if err := d.validateGetStringMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*string

	_jsii_.Invoke(
		d,
		"getStringMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataAwsccIoteventsDetectorModelDetectorModelDefinitionStatesOnInputTransitionEventsActionsOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		d,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataAwsccIoteventsDetectorModelDetectorModelDefinitionStatesOnInputTransitionEventsActionsOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
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

func (d *jsiiProxy_DataAwsccIoteventsDetectorModelDetectorModelDefinitionStatesOnInputTransitionEventsActionsOutputReference) Resolve(_context cdktf.IResolveContext) interface{} {
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

func (d *jsiiProxy_DataAwsccIoteventsDetectorModelDetectorModelDefinitionStatesOnInputTransitionEventsActionsOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		d,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

