package lightsailinstance

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/unimelb-platform/cdktf-provider-awscc-go/awscc/jsii"

	"github.com/hashicorp/terraform-cdk-go/cdktf"
	"github.com/unimelb-platform/cdktf-provider-awscc-go/awscc/lightsailinstance/internal"
)

type LightsailInstanceHardwareDisksOutputReference interface {
	cdktf.ComplexObject
	AttachedTo() *string
	SetAttachedTo(val *string)
	AttachedToInput() *string
	AttachmentState() *string
	SetAttachmentState(val *string)
	AttachmentStateInput() *string
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
	DiskName() *string
	SetDiskName(val *string)
	DiskNameInput() *string
	// Experimental.
	Fqn() *string
	InternalValue() interface{}
	SetInternalValue(val interface{})
	Iops() *float64
	SetIops(val *float64)
	IopsInput() *float64
	IsSystemDisk() interface{}
	SetIsSystemDisk(val interface{})
	IsSystemDiskInput() interface{}
	Path() *string
	SetPath(val *string)
	PathInput() *string
	SizeInGb() *string
	SetSizeInGb(val *string)
	SizeInGbInput() *string
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
	ResetAttachedTo()
	ResetAttachmentState()
	ResetIops()
	ResetIsSystemDisk()
	ResetSizeInGb()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(_context cdktf.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for LightsailInstanceHardwareDisksOutputReference
type jsiiProxy_LightsailInstanceHardwareDisksOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_LightsailInstanceHardwareDisksOutputReference) AttachedTo() *string {
	var returns *string
	_jsii_.Get(
		j,
		"attachedTo",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LightsailInstanceHardwareDisksOutputReference) AttachedToInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"attachedToInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LightsailInstanceHardwareDisksOutputReference) AttachmentState() *string {
	var returns *string
	_jsii_.Get(
		j,
		"attachmentState",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LightsailInstanceHardwareDisksOutputReference) AttachmentStateInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"attachmentStateInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LightsailInstanceHardwareDisksOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LightsailInstanceHardwareDisksOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LightsailInstanceHardwareDisksOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LightsailInstanceHardwareDisksOutputReference) DiskName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"diskName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LightsailInstanceHardwareDisksOutputReference) DiskNameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"diskNameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LightsailInstanceHardwareDisksOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LightsailInstanceHardwareDisksOutputReference) InternalValue() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LightsailInstanceHardwareDisksOutputReference) Iops() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"iops",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LightsailInstanceHardwareDisksOutputReference) IopsInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"iopsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LightsailInstanceHardwareDisksOutputReference) IsSystemDisk() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"isSystemDisk",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LightsailInstanceHardwareDisksOutputReference) IsSystemDiskInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"isSystemDiskInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LightsailInstanceHardwareDisksOutputReference) Path() *string {
	var returns *string
	_jsii_.Get(
		j,
		"path",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LightsailInstanceHardwareDisksOutputReference) PathInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"pathInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LightsailInstanceHardwareDisksOutputReference) SizeInGb() *string {
	var returns *string
	_jsii_.Get(
		j,
		"sizeInGb",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LightsailInstanceHardwareDisksOutputReference) SizeInGbInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"sizeInGbInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LightsailInstanceHardwareDisksOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LightsailInstanceHardwareDisksOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}


func NewLightsailInstanceHardwareDisksOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) LightsailInstanceHardwareDisksOutputReference {
	_init_.Initialize()

	if err := validateNewLightsailInstanceHardwareDisksOutputReferenceParameters(terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet); err != nil {
		panic(err)
	}
	j := jsiiProxy_LightsailInstanceHardwareDisksOutputReference{}

	_jsii_.Create(
		"awscc.lightsailInstance.LightsailInstanceHardwareDisksOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		&j,
	)

	return &j
}

func NewLightsailInstanceHardwareDisksOutputReference_Override(l LightsailInstanceHardwareDisksOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) {
	_init_.Initialize()

	_jsii_.Create(
		"awscc.lightsailInstance.LightsailInstanceHardwareDisksOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		l,
	)
}

func (j *jsiiProxy_LightsailInstanceHardwareDisksOutputReference)SetAttachedTo(val *string) {
	if err := j.validateSetAttachedToParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"attachedTo",
		val,
	)
}

func (j *jsiiProxy_LightsailInstanceHardwareDisksOutputReference)SetAttachmentState(val *string) {
	if err := j.validateSetAttachmentStateParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"attachmentState",
		val,
	)
}

func (j *jsiiProxy_LightsailInstanceHardwareDisksOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_LightsailInstanceHardwareDisksOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_LightsailInstanceHardwareDisksOutputReference)SetDiskName(val *string) {
	if err := j.validateSetDiskNameParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"diskName",
		val,
	)
}

func (j *jsiiProxy_LightsailInstanceHardwareDisksOutputReference)SetInternalValue(val interface{}) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_LightsailInstanceHardwareDisksOutputReference)SetIops(val *float64) {
	if err := j.validateSetIopsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"iops",
		val,
	)
}

func (j *jsiiProxy_LightsailInstanceHardwareDisksOutputReference)SetIsSystemDisk(val interface{}) {
	if err := j.validateSetIsSystemDiskParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"isSystemDisk",
		val,
	)
}

func (j *jsiiProxy_LightsailInstanceHardwareDisksOutputReference)SetPath(val *string) {
	if err := j.validateSetPathParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"path",
		val,
	)
}

func (j *jsiiProxy_LightsailInstanceHardwareDisksOutputReference)SetSizeInGb(val *string) {
	if err := j.validateSetSizeInGbParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"sizeInGb",
		val,
	)
}

func (j *jsiiProxy_LightsailInstanceHardwareDisksOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_LightsailInstanceHardwareDisksOutputReference)SetTerraformResource(val cdktf.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (l *jsiiProxy_LightsailInstanceHardwareDisksOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		l,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (l *jsiiProxy_LightsailInstanceHardwareDisksOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
	if err := l.validateGetAnyMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]interface{}

	_jsii_.Invoke(
		l,
		"getAnyMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (l *jsiiProxy_LightsailInstanceHardwareDisksOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
	if err := l.validateGetBooleanAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		l,
		"getBooleanAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (l *jsiiProxy_LightsailInstanceHardwareDisksOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
	if err := l.validateGetBooleanMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*bool

	_jsii_.Invoke(
		l,
		"getBooleanMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (l *jsiiProxy_LightsailInstanceHardwareDisksOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
	if err := l.validateGetListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*string

	_jsii_.Invoke(
		l,
		"getListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (l *jsiiProxy_LightsailInstanceHardwareDisksOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
	if err := l.validateGetNumberAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *float64

	_jsii_.Invoke(
		l,
		"getNumberAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (l *jsiiProxy_LightsailInstanceHardwareDisksOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
	if err := l.validateGetNumberListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*float64

	_jsii_.Invoke(
		l,
		"getNumberListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (l *jsiiProxy_LightsailInstanceHardwareDisksOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
	if err := l.validateGetNumberMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*float64

	_jsii_.Invoke(
		l,
		"getNumberMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (l *jsiiProxy_LightsailInstanceHardwareDisksOutputReference) GetStringAttribute(terraformAttribute *string) *string {
	if err := l.validateGetStringAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *string

	_jsii_.Invoke(
		l,
		"getStringAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (l *jsiiProxy_LightsailInstanceHardwareDisksOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
	if err := l.validateGetStringMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*string

	_jsii_.Invoke(
		l,
		"getStringMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (l *jsiiProxy_LightsailInstanceHardwareDisksOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		l,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (l *jsiiProxy_LightsailInstanceHardwareDisksOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
	if err := l.validateInterpolationForAttributeParameters(property); err != nil {
		panic(err)
	}
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		l,
		"interpolationForAttribute",
		[]interface{}{property},
		&returns,
	)

	return returns
}

func (l *jsiiProxy_LightsailInstanceHardwareDisksOutputReference) ResetAttachedTo() {
	_jsii_.InvokeVoid(
		l,
		"resetAttachedTo",
		nil, // no parameters
	)
}

func (l *jsiiProxy_LightsailInstanceHardwareDisksOutputReference) ResetAttachmentState() {
	_jsii_.InvokeVoid(
		l,
		"resetAttachmentState",
		nil, // no parameters
	)
}

func (l *jsiiProxy_LightsailInstanceHardwareDisksOutputReference) ResetIops() {
	_jsii_.InvokeVoid(
		l,
		"resetIops",
		nil, // no parameters
	)
}

func (l *jsiiProxy_LightsailInstanceHardwareDisksOutputReference) ResetIsSystemDisk() {
	_jsii_.InvokeVoid(
		l,
		"resetIsSystemDisk",
		nil, // no parameters
	)
}

func (l *jsiiProxy_LightsailInstanceHardwareDisksOutputReference) ResetSizeInGb() {
	_jsii_.InvokeVoid(
		l,
		"resetSizeInGb",
		nil, // no parameters
	)
}

func (l *jsiiProxy_LightsailInstanceHardwareDisksOutputReference) Resolve(_context cdktf.IResolveContext) interface{} {
	if err := l.validateResolveParameters(_context); err != nil {
		panic(err)
	}
	var returns interface{}

	_jsii_.Invoke(
		l,
		"resolve",
		[]interface{}{_context},
		&returns,
	)

	return returns
}

func (l *jsiiProxy_LightsailInstanceHardwareDisksOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		l,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

