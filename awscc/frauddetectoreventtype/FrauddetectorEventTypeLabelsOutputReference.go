package frauddetectoreventtype

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/unimelb-platform/cdktf-provider-awscc-go/awscc/jsii"

	"github.com/hashicorp/terraform-cdk-go/cdktf"
	"github.com/unimelb-platform/cdktf-provider-awscc-go/awscc/frauddetectoreventtype/internal"
)

type FrauddetectorEventTypeLabelsOutputReference interface {
	cdktf.ComplexObject
	Arn() *string
	SetArn(val *string)
	ArnInput() *string
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
	CreatedTime() *string
	SetCreatedTime(val *string)
	CreatedTimeInput() *string
	// The creation stack of this resolvable which will be appended to errors thrown during resolution.
	//
	// If this returns an empty array the stack will not be attached.
	// Experimental.
	CreationStack() *[]*string
	Description() *string
	SetDescription(val *string)
	DescriptionInput() *string
	// Experimental.
	Fqn() *string
	Inline() interface{}
	SetInline(val interface{})
	InlineInput() interface{}
	InternalValue() interface{}
	SetInternalValue(val interface{})
	LastUpdatedTime() *string
	SetLastUpdatedTime(val *string)
	LastUpdatedTimeInput() *string
	Name() *string
	SetName(val *string)
	NameInput() *string
	Tags() FrauddetectorEventTypeLabelsTagsList
	TagsInput() interface{}
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
	PutTags(value interface{})
	ResetArn()
	ResetCreatedTime()
	ResetDescription()
	ResetInline()
	ResetLastUpdatedTime()
	ResetName()
	ResetTags()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(_context cdktf.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for FrauddetectorEventTypeLabelsOutputReference
type jsiiProxy_FrauddetectorEventTypeLabelsOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_FrauddetectorEventTypeLabelsOutputReference) Arn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"arn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FrauddetectorEventTypeLabelsOutputReference) ArnInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"arnInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FrauddetectorEventTypeLabelsOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FrauddetectorEventTypeLabelsOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FrauddetectorEventTypeLabelsOutputReference) CreatedTime() *string {
	var returns *string
	_jsii_.Get(
		j,
		"createdTime",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FrauddetectorEventTypeLabelsOutputReference) CreatedTimeInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"createdTimeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FrauddetectorEventTypeLabelsOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FrauddetectorEventTypeLabelsOutputReference) Description() *string {
	var returns *string
	_jsii_.Get(
		j,
		"description",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FrauddetectorEventTypeLabelsOutputReference) DescriptionInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"descriptionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FrauddetectorEventTypeLabelsOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FrauddetectorEventTypeLabelsOutputReference) Inline() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"inline",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FrauddetectorEventTypeLabelsOutputReference) InlineInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"inlineInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FrauddetectorEventTypeLabelsOutputReference) InternalValue() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FrauddetectorEventTypeLabelsOutputReference) LastUpdatedTime() *string {
	var returns *string
	_jsii_.Get(
		j,
		"lastUpdatedTime",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FrauddetectorEventTypeLabelsOutputReference) LastUpdatedTimeInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"lastUpdatedTimeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FrauddetectorEventTypeLabelsOutputReference) Name() *string {
	var returns *string
	_jsii_.Get(
		j,
		"name",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FrauddetectorEventTypeLabelsOutputReference) NameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"nameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FrauddetectorEventTypeLabelsOutputReference) Tags() FrauddetectorEventTypeLabelsTagsList {
	var returns FrauddetectorEventTypeLabelsTagsList
	_jsii_.Get(
		j,
		"tags",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FrauddetectorEventTypeLabelsOutputReference) TagsInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"tagsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FrauddetectorEventTypeLabelsOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FrauddetectorEventTypeLabelsOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}


func NewFrauddetectorEventTypeLabelsOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) FrauddetectorEventTypeLabelsOutputReference {
	_init_.Initialize()

	if err := validateNewFrauddetectorEventTypeLabelsOutputReferenceParameters(terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet); err != nil {
		panic(err)
	}
	j := jsiiProxy_FrauddetectorEventTypeLabelsOutputReference{}

	_jsii_.Create(
		"awscc.frauddetectorEventType.FrauddetectorEventTypeLabelsOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		&j,
	)

	return &j
}

func NewFrauddetectorEventTypeLabelsOutputReference_Override(f FrauddetectorEventTypeLabelsOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) {
	_init_.Initialize()

	_jsii_.Create(
		"awscc.frauddetectorEventType.FrauddetectorEventTypeLabelsOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		f,
	)
}

func (j *jsiiProxy_FrauddetectorEventTypeLabelsOutputReference)SetArn(val *string) {
	if err := j.validateSetArnParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"arn",
		val,
	)
}

func (j *jsiiProxy_FrauddetectorEventTypeLabelsOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_FrauddetectorEventTypeLabelsOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_FrauddetectorEventTypeLabelsOutputReference)SetCreatedTime(val *string) {
	if err := j.validateSetCreatedTimeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"createdTime",
		val,
	)
}

func (j *jsiiProxy_FrauddetectorEventTypeLabelsOutputReference)SetDescription(val *string) {
	if err := j.validateSetDescriptionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"description",
		val,
	)
}

func (j *jsiiProxy_FrauddetectorEventTypeLabelsOutputReference)SetInline(val interface{}) {
	if err := j.validateSetInlineParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"inline",
		val,
	)
}

func (j *jsiiProxy_FrauddetectorEventTypeLabelsOutputReference)SetInternalValue(val interface{}) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_FrauddetectorEventTypeLabelsOutputReference)SetLastUpdatedTime(val *string) {
	if err := j.validateSetLastUpdatedTimeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"lastUpdatedTime",
		val,
	)
}

func (j *jsiiProxy_FrauddetectorEventTypeLabelsOutputReference)SetName(val *string) {
	if err := j.validateSetNameParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"name",
		val,
	)
}

func (j *jsiiProxy_FrauddetectorEventTypeLabelsOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_FrauddetectorEventTypeLabelsOutputReference)SetTerraformResource(val cdktf.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (f *jsiiProxy_FrauddetectorEventTypeLabelsOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		f,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (f *jsiiProxy_FrauddetectorEventTypeLabelsOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
	if err := f.validateGetAnyMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]interface{}

	_jsii_.Invoke(
		f,
		"getAnyMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (f *jsiiProxy_FrauddetectorEventTypeLabelsOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
	if err := f.validateGetBooleanAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		f,
		"getBooleanAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (f *jsiiProxy_FrauddetectorEventTypeLabelsOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
	if err := f.validateGetBooleanMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*bool

	_jsii_.Invoke(
		f,
		"getBooleanMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (f *jsiiProxy_FrauddetectorEventTypeLabelsOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
	if err := f.validateGetListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*string

	_jsii_.Invoke(
		f,
		"getListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (f *jsiiProxy_FrauddetectorEventTypeLabelsOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
	if err := f.validateGetNumberAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *float64

	_jsii_.Invoke(
		f,
		"getNumberAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (f *jsiiProxy_FrauddetectorEventTypeLabelsOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
	if err := f.validateGetNumberListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*float64

	_jsii_.Invoke(
		f,
		"getNumberListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (f *jsiiProxy_FrauddetectorEventTypeLabelsOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
	if err := f.validateGetNumberMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*float64

	_jsii_.Invoke(
		f,
		"getNumberMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (f *jsiiProxy_FrauddetectorEventTypeLabelsOutputReference) GetStringAttribute(terraformAttribute *string) *string {
	if err := f.validateGetStringAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *string

	_jsii_.Invoke(
		f,
		"getStringAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (f *jsiiProxy_FrauddetectorEventTypeLabelsOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
	if err := f.validateGetStringMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*string

	_jsii_.Invoke(
		f,
		"getStringMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (f *jsiiProxy_FrauddetectorEventTypeLabelsOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		f,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (f *jsiiProxy_FrauddetectorEventTypeLabelsOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
	if err := f.validateInterpolationForAttributeParameters(property); err != nil {
		panic(err)
	}
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		f,
		"interpolationForAttribute",
		[]interface{}{property},
		&returns,
	)

	return returns
}

func (f *jsiiProxy_FrauddetectorEventTypeLabelsOutputReference) PutTags(value interface{}) {
	if err := f.validatePutTagsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		f,
		"putTags",
		[]interface{}{value},
	)
}

func (f *jsiiProxy_FrauddetectorEventTypeLabelsOutputReference) ResetArn() {
	_jsii_.InvokeVoid(
		f,
		"resetArn",
		nil, // no parameters
	)
}

func (f *jsiiProxy_FrauddetectorEventTypeLabelsOutputReference) ResetCreatedTime() {
	_jsii_.InvokeVoid(
		f,
		"resetCreatedTime",
		nil, // no parameters
	)
}

func (f *jsiiProxy_FrauddetectorEventTypeLabelsOutputReference) ResetDescription() {
	_jsii_.InvokeVoid(
		f,
		"resetDescription",
		nil, // no parameters
	)
}

func (f *jsiiProxy_FrauddetectorEventTypeLabelsOutputReference) ResetInline() {
	_jsii_.InvokeVoid(
		f,
		"resetInline",
		nil, // no parameters
	)
}

func (f *jsiiProxy_FrauddetectorEventTypeLabelsOutputReference) ResetLastUpdatedTime() {
	_jsii_.InvokeVoid(
		f,
		"resetLastUpdatedTime",
		nil, // no parameters
	)
}

func (f *jsiiProxy_FrauddetectorEventTypeLabelsOutputReference) ResetName() {
	_jsii_.InvokeVoid(
		f,
		"resetName",
		nil, // no parameters
	)
}

func (f *jsiiProxy_FrauddetectorEventTypeLabelsOutputReference) ResetTags() {
	_jsii_.InvokeVoid(
		f,
		"resetTags",
		nil, // no parameters
	)
}

func (f *jsiiProxy_FrauddetectorEventTypeLabelsOutputReference) Resolve(_context cdktf.IResolveContext) interface{} {
	if err := f.validateResolveParameters(_context); err != nil {
		panic(err)
	}
	var returns interface{}

	_jsii_.Invoke(
		f,
		"resolve",
		[]interface{}{_context},
		&returns,
	)

	return returns
}

func (f *jsiiProxy_FrauddetectorEventTypeLabelsOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		f,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

