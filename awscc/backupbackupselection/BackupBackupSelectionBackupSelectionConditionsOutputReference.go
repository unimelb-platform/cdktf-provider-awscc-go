package backupbackupselection

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/unimelb-platform/cdktf-provider-awscc-go/awscc/jsii"

	"github.com/hashicorp/terraform-cdk-go/cdktf"
	"github.com/unimelb-platform/cdktf-provider-awscc-go/awscc/backupbackupselection/internal"
)

type BackupBackupSelectionBackupSelectionConditionsOutputReference interface {
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
	// Experimental.
	Fqn() *string
	InternalValue() interface{}
	SetInternalValue(val interface{})
	StringEquals() BackupBackupSelectionBackupSelectionConditionsStringEqualsList
	StringEqualsInput() interface{}
	StringLike() BackupBackupSelectionBackupSelectionConditionsStringLikeList
	StringLikeInput() interface{}
	StringNotEquals() BackupBackupSelectionBackupSelectionConditionsStringNotEqualsList
	StringNotEqualsInput() interface{}
	StringNotLike() BackupBackupSelectionBackupSelectionConditionsStringNotLikeList
	StringNotLikeInput() interface{}
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
	PutStringEquals(value interface{})
	PutStringLike(value interface{})
	PutStringNotEquals(value interface{})
	PutStringNotLike(value interface{})
	ResetStringEquals()
	ResetStringLike()
	ResetStringNotEquals()
	ResetStringNotLike()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(_context cdktf.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for BackupBackupSelectionBackupSelectionConditionsOutputReference
type jsiiProxy_BackupBackupSelectionBackupSelectionConditionsOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_BackupBackupSelectionBackupSelectionConditionsOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BackupBackupSelectionBackupSelectionConditionsOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BackupBackupSelectionBackupSelectionConditionsOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BackupBackupSelectionBackupSelectionConditionsOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BackupBackupSelectionBackupSelectionConditionsOutputReference) InternalValue() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BackupBackupSelectionBackupSelectionConditionsOutputReference) StringEquals() BackupBackupSelectionBackupSelectionConditionsStringEqualsList {
	var returns BackupBackupSelectionBackupSelectionConditionsStringEqualsList
	_jsii_.Get(
		j,
		"stringEquals",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BackupBackupSelectionBackupSelectionConditionsOutputReference) StringEqualsInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"stringEqualsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BackupBackupSelectionBackupSelectionConditionsOutputReference) StringLike() BackupBackupSelectionBackupSelectionConditionsStringLikeList {
	var returns BackupBackupSelectionBackupSelectionConditionsStringLikeList
	_jsii_.Get(
		j,
		"stringLike",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BackupBackupSelectionBackupSelectionConditionsOutputReference) StringLikeInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"stringLikeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BackupBackupSelectionBackupSelectionConditionsOutputReference) StringNotEquals() BackupBackupSelectionBackupSelectionConditionsStringNotEqualsList {
	var returns BackupBackupSelectionBackupSelectionConditionsStringNotEqualsList
	_jsii_.Get(
		j,
		"stringNotEquals",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BackupBackupSelectionBackupSelectionConditionsOutputReference) StringNotEqualsInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"stringNotEqualsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BackupBackupSelectionBackupSelectionConditionsOutputReference) StringNotLike() BackupBackupSelectionBackupSelectionConditionsStringNotLikeList {
	var returns BackupBackupSelectionBackupSelectionConditionsStringNotLikeList
	_jsii_.Get(
		j,
		"stringNotLike",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BackupBackupSelectionBackupSelectionConditionsOutputReference) StringNotLikeInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"stringNotLikeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BackupBackupSelectionBackupSelectionConditionsOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BackupBackupSelectionBackupSelectionConditionsOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}


func NewBackupBackupSelectionBackupSelectionConditionsOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) BackupBackupSelectionBackupSelectionConditionsOutputReference {
	_init_.Initialize()

	if err := validateNewBackupBackupSelectionBackupSelectionConditionsOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_BackupBackupSelectionBackupSelectionConditionsOutputReference{}

	_jsii_.Create(
		"awscc.backupBackupSelection.BackupBackupSelectionBackupSelectionConditionsOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewBackupBackupSelectionBackupSelectionConditionsOutputReference_Override(b BackupBackupSelectionBackupSelectionConditionsOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"awscc.backupBackupSelection.BackupBackupSelectionBackupSelectionConditionsOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		b,
	)
}

func (j *jsiiProxy_BackupBackupSelectionBackupSelectionConditionsOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_BackupBackupSelectionBackupSelectionConditionsOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_BackupBackupSelectionBackupSelectionConditionsOutputReference)SetInternalValue(val interface{}) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_BackupBackupSelectionBackupSelectionConditionsOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_BackupBackupSelectionBackupSelectionConditionsOutputReference)SetTerraformResource(val cdktf.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (b *jsiiProxy_BackupBackupSelectionBackupSelectionConditionsOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		b,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (b *jsiiProxy_BackupBackupSelectionBackupSelectionConditionsOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
	if err := b.validateGetAnyMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]interface{}

	_jsii_.Invoke(
		b,
		"getAnyMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (b *jsiiProxy_BackupBackupSelectionBackupSelectionConditionsOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
	if err := b.validateGetBooleanAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		b,
		"getBooleanAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (b *jsiiProxy_BackupBackupSelectionBackupSelectionConditionsOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
	if err := b.validateGetBooleanMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*bool

	_jsii_.Invoke(
		b,
		"getBooleanMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (b *jsiiProxy_BackupBackupSelectionBackupSelectionConditionsOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
	if err := b.validateGetListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*string

	_jsii_.Invoke(
		b,
		"getListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (b *jsiiProxy_BackupBackupSelectionBackupSelectionConditionsOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
	if err := b.validateGetNumberAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *float64

	_jsii_.Invoke(
		b,
		"getNumberAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (b *jsiiProxy_BackupBackupSelectionBackupSelectionConditionsOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
	if err := b.validateGetNumberListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*float64

	_jsii_.Invoke(
		b,
		"getNumberListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (b *jsiiProxy_BackupBackupSelectionBackupSelectionConditionsOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
	if err := b.validateGetNumberMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*float64

	_jsii_.Invoke(
		b,
		"getNumberMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (b *jsiiProxy_BackupBackupSelectionBackupSelectionConditionsOutputReference) GetStringAttribute(terraformAttribute *string) *string {
	if err := b.validateGetStringAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *string

	_jsii_.Invoke(
		b,
		"getStringAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (b *jsiiProxy_BackupBackupSelectionBackupSelectionConditionsOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
	if err := b.validateGetStringMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*string

	_jsii_.Invoke(
		b,
		"getStringMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (b *jsiiProxy_BackupBackupSelectionBackupSelectionConditionsOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		b,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (b *jsiiProxy_BackupBackupSelectionBackupSelectionConditionsOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
	if err := b.validateInterpolationForAttributeParameters(property); err != nil {
		panic(err)
	}
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		b,
		"interpolationForAttribute",
		[]interface{}{property},
		&returns,
	)

	return returns
}

func (b *jsiiProxy_BackupBackupSelectionBackupSelectionConditionsOutputReference) PutStringEquals(value interface{}) {
	if err := b.validatePutStringEqualsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		b,
		"putStringEquals",
		[]interface{}{value},
	)
}

func (b *jsiiProxy_BackupBackupSelectionBackupSelectionConditionsOutputReference) PutStringLike(value interface{}) {
	if err := b.validatePutStringLikeParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		b,
		"putStringLike",
		[]interface{}{value},
	)
}

func (b *jsiiProxy_BackupBackupSelectionBackupSelectionConditionsOutputReference) PutStringNotEquals(value interface{}) {
	if err := b.validatePutStringNotEqualsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		b,
		"putStringNotEquals",
		[]interface{}{value},
	)
}

func (b *jsiiProxy_BackupBackupSelectionBackupSelectionConditionsOutputReference) PutStringNotLike(value interface{}) {
	if err := b.validatePutStringNotLikeParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		b,
		"putStringNotLike",
		[]interface{}{value},
	)
}

func (b *jsiiProxy_BackupBackupSelectionBackupSelectionConditionsOutputReference) ResetStringEquals() {
	_jsii_.InvokeVoid(
		b,
		"resetStringEquals",
		nil, // no parameters
	)
}

func (b *jsiiProxy_BackupBackupSelectionBackupSelectionConditionsOutputReference) ResetStringLike() {
	_jsii_.InvokeVoid(
		b,
		"resetStringLike",
		nil, // no parameters
	)
}

func (b *jsiiProxy_BackupBackupSelectionBackupSelectionConditionsOutputReference) ResetStringNotEquals() {
	_jsii_.InvokeVoid(
		b,
		"resetStringNotEquals",
		nil, // no parameters
	)
}

func (b *jsiiProxy_BackupBackupSelectionBackupSelectionConditionsOutputReference) ResetStringNotLike() {
	_jsii_.InvokeVoid(
		b,
		"resetStringNotLike",
		nil, // no parameters
	)
}

func (b *jsiiProxy_BackupBackupSelectionBackupSelectionConditionsOutputReference) Resolve(_context cdktf.IResolveContext) interface{} {
	if err := b.validateResolveParameters(_context); err != nil {
		panic(err)
	}
	var returns interface{}

	_jsii_.Invoke(
		b,
		"resolve",
		[]interface{}{_context},
		&returns,
	)

	return returns
}

func (b *jsiiProxy_BackupBackupSelectionBackupSelectionConditionsOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		b,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

