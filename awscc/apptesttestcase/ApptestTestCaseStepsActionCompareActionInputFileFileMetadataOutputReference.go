package apptesttestcase

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/unimelb-platform/cdktf-provider-awscc-go/awscc/jsii"

	"github.com/hashicorp/terraform-cdk-go/cdktf"
	"github.com/unimelb-platform/cdktf-provider-awscc-go/awscc/apptesttestcase/internal"
)

type ApptestTestCaseStepsActionCompareActionInputFileFileMetadataOutputReference interface {
	cdktf.ComplexObject
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
	DatabaseCdc() ApptestTestCaseStepsActionCompareActionInputFileFileMetadataDatabaseCdcOutputReference
	DatabaseCdcInput() interface{}
	DataSets() ApptestTestCaseStepsActionCompareActionInputFileFileMetadataDataSetsList
	DataSetsInput() interface{}
	// Experimental.
	Fqn() *string
	InternalValue() interface{}
	SetInternalValue(val interface{})
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
	PutDatabaseCdc(value *ApptestTestCaseStepsActionCompareActionInputFileFileMetadataDatabaseCdc)
	PutDataSets(value interface{})
	ResetDatabaseCdc()
	ResetDataSets()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(_context cdktf.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for ApptestTestCaseStepsActionCompareActionInputFileFileMetadataOutputReference
type jsiiProxy_ApptestTestCaseStepsActionCompareActionInputFileFileMetadataOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_ApptestTestCaseStepsActionCompareActionInputFileFileMetadataOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApptestTestCaseStepsActionCompareActionInputFileFileMetadataOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApptestTestCaseStepsActionCompareActionInputFileFileMetadataOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApptestTestCaseStepsActionCompareActionInputFileFileMetadataOutputReference) DatabaseCdc() ApptestTestCaseStepsActionCompareActionInputFileFileMetadataDatabaseCdcOutputReference {
	var returns ApptestTestCaseStepsActionCompareActionInputFileFileMetadataDatabaseCdcOutputReference
	_jsii_.Get(
		j,
		"databaseCdc",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApptestTestCaseStepsActionCompareActionInputFileFileMetadataOutputReference) DatabaseCdcInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"databaseCdcInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApptestTestCaseStepsActionCompareActionInputFileFileMetadataOutputReference) DataSets() ApptestTestCaseStepsActionCompareActionInputFileFileMetadataDataSetsList {
	var returns ApptestTestCaseStepsActionCompareActionInputFileFileMetadataDataSetsList
	_jsii_.Get(
		j,
		"dataSets",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApptestTestCaseStepsActionCompareActionInputFileFileMetadataOutputReference) DataSetsInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"dataSetsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApptestTestCaseStepsActionCompareActionInputFileFileMetadataOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApptestTestCaseStepsActionCompareActionInputFileFileMetadataOutputReference) InternalValue() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApptestTestCaseStepsActionCompareActionInputFileFileMetadataOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApptestTestCaseStepsActionCompareActionInputFileFileMetadataOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}


func NewApptestTestCaseStepsActionCompareActionInputFileFileMetadataOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) ApptestTestCaseStepsActionCompareActionInputFileFileMetadataOutputReference {
	_init_.Initialize()

	if err := validateNewApptestTestCaseStepsActionCompareActionInputFileFileMetadataOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_ApptestTestCaseStepsActionCompareActionInputFileFileMetadataOutputReference{}

	_jsii_.Create(
		"awscc.apptestTestCase.ApptestTestCaseStepsActionCompareActionInputFileFileMetadataOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewApptestTestCaseStepsActionCompareActionInputFileFileMetadataOutputReference_Override(a ApptestTestCaseStepsActionCompareActionInputFileFileMetadataOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"awscc.apptestTestCase.ApptestTestCaseStepsActionCompareActionInputFileFileMetadataOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		a,
	)
}

func (j *jsiiProxy_ApptestTestCaseStepsActionCompareActionInputFileFileMetadataOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_ApptestTestCaseStepsActionCompareActionInputFileFileMetadataOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_ApptestTestCaseStepsActionCompareActionInputFileFileMetadataOutputReference)SetInternalValue(val interface{}) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_ApptestTestCaseStepsActionCompareActionInputFileFileMetadataOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_ApptestTestCaseStepsActionCompareActionInputFileFileMetadataOutputReference)SetTerraformResource(val cdktf.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (a *jsiiProxy_ApptestTestCaseStepsActionCompareActionInputFileFileMetadataOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		a,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (a *jsiiProxy_ApptestTestCaseStepsActionCompareActionInputFileFileMetadataOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
	if err := a.validateGetAnyMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]interface{}

	_jsii_.Invoke(
		a,
		"getAnyMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_ApptestTestCaseStepsActionCompareActionInputFileFileMetadataOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
	if err := a.validateGetBooleanAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		a,
		"getBooleanAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_ApptestTestCaseStepsActionCompareActionInputFileFileMetadataOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
	if err := a.validateGetBooleanMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*bool

	_jsii_.Invoke(
		a,
		"getBooleanMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_ApptestTestCaseStepsActionCompareActionInputFileFileMetadataOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
	if err := a.validateGetListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*string

	_jsii_.Invoke(
		a,
		"getListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_ApptestTestCaseStepsActionCompareActionInputFileFileMetadataOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
	if err := a.validateGetNumberAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *float64

	_jsii_.Invoke(
		a,
		"getNumberAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_ApptestTestCaseStepsActionCompareActionInputFileFileMetadataOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
	if err := a.validateGetNumberListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*float64

	_jsii_.Invoke(
		a,
		"getNumberListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_ApptestTestCaseStepsActionCompareActionInputFileFileMetadataOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
	if err := a.validateGetNumberMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*float64

	_jsii_.Invoke(
		a,
		"getNumberMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_ApptestTestCaseStepsActionCompareActionInputFileFileMetadataOutputReference) GetStringAttribute(terraformAttribute *string) *string {
	if err := a.validateGetStringAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *string

	_jsii_.Invoke(
		a,
		"getStringAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_ApptestTestCaseStepsActionCompareActionInputFileFileMetadataOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
	if err := a.validateGetStringMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*string

	_jsii_.Invoke(
		a,
		"getStringMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_ApptestTestCaseStepsActionCompareActionInputFileFileMetadataOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		a,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (a *jsiiProxy_ApptestTestCaseStepsActionCompareActionInputFileFileMetadataOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
	if err := a.validateInterpolationForAttributeParameters(property); err != nil {
		panic(err)
	}
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		a,
		"interpolationForAttribute",
		[]interface{}{property},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_ApptestTestCaseStepsActionCompareActionInputFileFileMetadataOutputReference) PutDatabaseCdc(value *ApptestTestCaseStepsActionCompareActionInputFileFileMetadataDatabaseCdc) {
	if err := a.validatePutDatabaseCdcParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"putDatabaseCdc",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_ApptestTestCaseStepsActionCompareActionInputFileFileMetadataOutputReference) PutDataSets(value interface{}) {
	if err := a.validatePutDataSetsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"putDataSets",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_ApptestTestCaseStepsActionCompareActionInputFileFileMetadataOutputReference) ResetDatabaseCdc() {
	_jsii_.InvokeVoid(
		a,
		"resetDatabaseCdc",
		nil, // no parameters
	)
}

func (a *jsiiProxy_ApptestTestCaseStepsActionCompareActionInputFileFileMetadataOutputReference) ResetDataSets() {
	_jsii_.InvokeVoid(
		a,
		"resetDataSets",
		nil, // no parameters
	)
}

func (a *jsiiProxy_ApptestTestCaseStepsActionCompareActionInputFileFileMetadataOutputReference) Resolve(_context cdktf.IResolveContext) interface{} {
	if err := a.validateResolveParameters(_context); err != nil {
		panic(err)
	}
	var returns interface{}

	_jsii_.Invoke(
		a,
		"resolve",
		[]interface{}{_context},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_ApptestTestCaseStepsActionCompareActionInputFileFileMetadataOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		a,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

