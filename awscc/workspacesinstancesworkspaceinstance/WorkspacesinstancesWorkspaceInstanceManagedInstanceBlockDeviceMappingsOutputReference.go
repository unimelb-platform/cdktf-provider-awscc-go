package workspacesinstancesworkspaceinstance

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/unimelb-platform/cdktf-provider-awscc-go/awscc/jsii"

	"github.com/hashicorp/terraform-cdk-go/cdktf"
	"github.com/unimelb-platform/cdktf-provider-awscc-go/awscc/workspacesinstancesworkspaceinstance/internal"
)

type WorkspacesinstancesWorkspaceInstanceManagedInstanceBlockDeviceMappingsOutputReference interface {
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
	DeviceName() *string
	SetDeviceName(val *string)
	DeviceNameInput() *string
	Ebs() WorkspacesinstancesWorkspaceInstanceManagedInstanceBlockDeviceMappingsEbsOutputReference
	EbsInput() interface{}
	// Experimental.
	Fqn() *string
	InternalValue() interface{}
	SetInternalValue(val interface{})
	NoDevice() *string
	SetNoDevice(val *string)
	NoDeviceInput() *string
	// Experimental.
	TerraformAttribute() *string
	// Experimental.
	SetTerraformAttribute(val *string)
	// Experimental.
	TerraformResource() cdktf.IInterpolatingParent
	// Experimental.
	SetTerraformResource(val cdktf.IInterpolatingParent)
	VirtualName() *string
	SetVirtualName(val *string)
	VirtualNameInput() *string
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
	PutEbs(value *WorkspacesinstancesWorkspaceInstanceManagedInstanceBlockDeviceMappingsEbs)
	ResetDeviceName()
	ResetEbs()
	ResetNoDevice()
	ResetVirtualName()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(_context cdktf.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for WorkspacesinstancesWorkspaceInstanceManagedInstanceBlockDeviceMappingsOutputReference
type jsiiProxy_WorkspacesinstancesWorkspaceInstanceManagedInstanceBlockDeviceMappingsOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_WorkspacesinstancesWorkspaceInstanceManagedInstanceBlockDeviceMappingsOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WorkspacesinstancesWorkspaceInstanceManagedInstanceBlockDeviceMappingsOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WorkspacesinstancesWorkspaceInstanceManagedInstanceBlockDeviceMappingsOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WorkspacesinstancesWorkspaceInstanceManagedInstanceBlockDeviceMappingsOutputReference) DeviceName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"deviceName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WorkspacesinstancesWorkspaceInstanceManagedInstanceBlockDeviceMappingsOutputReference) DeviceNameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"deviceNameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WorkspacesinstancesWorkspaceInstanceManagedInstanceBlockDeviceMappingsOutputReference) Ebs() WorkspacesinstancesWorkspaceInstanceManagedInstanceBlockDeviceMappingsEbsOutputReference {
	var returns WorkspacesinstancesWorkspaceInstanceManagedInstanceBlockDeviceMappingsEbsOutputReference
	_jsii_.Get(
		j,
		"ebs",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WorkspacesinstancesWorkspaceInstanceManagedInstanceBlockDeviceMappingsOutputReference) EbsInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"ebsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WorkspacesinstancesWorkspaceInstanceManagedInstanceBlockDeviceMappingsOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WorkspacesinstancesWorkspaceInstanceManagedInstanceBlockDeviceMappingsOutputReference) InternalValue() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WorkspacesinstancesWorkspaceInstanceManagedInstanceBlockDeviceMappingsOutputReference) NoDevice() *string {
	var returns *string
	_jsii_.Get(
		j,
		"noDevice",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WorkspacesinstancesWorkspaceInstanceManagedInstanceBlockDeviceMappingsOutputReference) NoDeviceInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"noDeviceInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WorkspacesinstancesWorkspaceInstanceManagedInstanceBlockDeviceMappingsOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WorkspacesinstancesWorkspaceInstanceManagedInstanceBlockDeviceMappingsOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WorkspacesinstancesWorkspaceInstanceManagedInstanceBlockDeviceMappingsOutputReference) VirtualName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"virtualName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WorkspacesinstancesWorkspaceInstanceManagedInstanceBlockDeviceMappingsOutputReference) VirtualNameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"virtualNameInput",
		&returns,
	)
	return returns
}


func NewWorkspacesinstancesWorkspaceInstanceManagedInstanceBlockDeviceMappingsOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) WorkspacesinstancesWorkspaceInstanceManagedInstanceBlockDeviceMappingsOutputReference {
	_init_.Initialize()

	if err := validateNewWorkspacesinstancesWorkspaceInstanceManagedInstanceBlockDeviceMappingsOutputReferenceParameters(terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet); err != nil {
		panic(err)
	}
	j := jsiiProxy_WorkspacesinstancesWorkspaceInstanceManagedInstanceBlockDeviceMappingsOutputReference{}

	_jsii_.Create(
		"awscc.workspacesinstancesWorkspaceInstance.WorkspacesinstancesWorkspaceInstanceManagedInstanceBlockDeviceMappingsOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		&j,
	)

	return &j
}

func NewWorkspacesinstancesWorkspaceInstanceManagedInstanceBlockDeviceMappingsOutputReference_Override(w WorkspacesinstancesWorkspaceInstanceManagedInstanceBlockDeviceMappingsOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) {
	_init_.Initialize()

	_jsii_.Create(
		"awscc.workspacesinstancesWorkspaceInstance.WorkspacesinstancesWorkspaceInstanceManagedInstanceBlockDeviceMappingsOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		w,
	)
}

func (j *jsiiProxy_WorkspacesinstancesWorkspaceInstanceManagedInstanceBlockDeviceMappingsOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_WorkspacesinstancesWorkspaceInstanceManagedInstanceBlockDeviceMappingsOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_WorkspacesinstancesWorkspaceInstanceManagedInstanceBlockDeviceMappingsOutputReference)SetDeviceName(val *string) {
	if err := j.validateSetDeviceNameParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"deviceName",
		val,
	)
}

func (j *jsiiProxy_WorkspacesinstancesWorkspaceInstanceManagedInstanceBlockDeviceMappingsOutputReference)SetInternalValue(val interface{}) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_WorkspacesinstancesWorkspaceInstanceManagedInstanceBlockDeviceMappingsOutputReference)SetNoDevice(val *string) {
	if err := j.validateSetNoDeviceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"noDevice",
		val,
	)
}

func (j *jsiiProxy_WorkspacesinstancesWorkspaceInstanceManagedInstanceBlockDeviceMappingsOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_WorkspacesinstancesWorkspaceInstanceManagedInstanceBlockDeviceMappingsOutputReference)SetTerraformResource(val cdktf.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (j *jsiiProxy_WorkspacesinstancesWorkspaceInstanceManagedInstanceBlockDeviceMappingsOutputReference)SetVirtualName(val *string) {
	if err := j.validateSetVirtualNameParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"virtualName",
		val,
	)
}

func (w *jsiiProxy_WorkspacesinstancesWorkspaceInstanceManagedInstanceBlockDeviceMappingsOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		w,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (w *jsiiProxy_WorkspacesinstancesWorkspaceInstanceManagedInstanceBlockDeviceMappingsOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
	if err := w.validateGetAnyMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]interface{}

	_jsii_.Invoke(
		w,
		"getAnyMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (w *jsiiProxy_WorkspacesinstancesWorkspaceInstanceManagedInstanceBlockDeviceMappingsOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
	if err := w.validateGetBooleanAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		w,
		"getBooleanAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (w *jsiiProxy_WorkspacesinstancesWorkspaceInstanceManagedInstanceBlockDeviceMappingsOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
	if err := w.validateGetBooleanMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*bool

	_jsii_.Invoke(
		w,
		"getBooleanMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (w *jsiiProxy_WorkspacesinstancesWorkspaceInstanceManagedInstanceBlockDeviceMappingsOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
	if err := w.validateGetListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*string

	_jsii_.Invoke(
		w,
		"getListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (w *jsiiProxy_WorkspacesinstancesWorkspaceInstanceManagedInstanceBlockDeviceMappingsOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
	if err := w.validateGetNumberAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *float64

	_jsii_.Invoke(
		w,
		"getNumberAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (w *jsiiProxy_WorkspacesinstancesWorkspaceInstanceManagedInstanceBlockDeviceMappingsOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
	if err := w.validateGetNumberListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*float64

	_jsii_.Invoke(
		w,
		"getNumberListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (w *jsiiProxy_WorkspacesinstancesWorkspaceInstanceManagedInstanceBlockDeviceMappingsOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
	if err := w.validateGetNumberMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*float64

	_jsii_.Invoke(
		w,
		"getNumberMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (w *jsiiProxy_WorkspacesinstancesWorkspaceInstanceManagedInstanceBlockDeviceMappingsOutputReference) GetStringAttribute(terraformAttribute *string) *string {
	if err := w.validateGetStringAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *string

	_jsii_.Invoke(
		w,
		"getStringAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (w *jsiiProxy_WorkspacesinstancesWorkspaceInstanceManagedInstanceBlockDeviceMappingsOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
	if err := w.validateGetStringMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*string

	_jsii_.Invoke(
		w,
		"getStringMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (w *jsiiProxy_WorkspacesinstancesWorkspaceInstanceManagedInstanceBlockDeviceMappingsOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		w,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (w *jsiiProxy_WorkspacesinstancesWorkspaceInstanceManagedInstanceBlockDeviceMappingsOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
	if err := w.validateInterpolationForAttributeParameters(property); err != nil {
		panic(err)
	}
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		w,
		"interpolationForAttribute",
		[]interface{}{property},
		&returns,
	)

	return returns
}

func (w *jsiiProxy_WorkspacesinstancesWorkspaceInstanceManagedInstanceBlockDeviceMappingsOutputReference) PutEbs(value *WorkspacesinstancesWorkspaceInstanceManagedInstanceBlockDeviceMappingsEbs) {
	if err := w.validatePutEbsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		w,
		"putEbs",
		[]interface{}{value},
	)
}

func (w *jsiiProxy_WorkspacesinstancesWorkspaceInstanceManagedInstanceBlockDeviceMappingsOutputReference) ResetDeviceName() {
	_jsii_.InvokeVoid(
		w,
		"resetDeviceName",
		nil, // no parameters
	)
}

func (w *jsiiProxy_WorkspacesinstancesWorkspaceInstanceManagedInstanceBlockDeviceMappingsOutputReference) ResetEbs() {
	_jsii_.InvokeVoid(
		w,
		"resetEbs",
		nil, // no parameters
	)
}

func (w *jsiiProxy_WorkspacesinstancesWorkspaceInstanceManagedInstanceBlockDeviceMappingsOutputReference) ResetNoDevice() {
	_jsii_.InvokeVoid(
		w,
		"resetNoDevice",
		nil, // no parameters
	)
}

func (w *jsiiProxy_WorkspacesinstancesWorkspaceInstanceManagedInstanceBlockDeviceMappingsOutputReference) ResetVirtualName() {
	_jsii_.InvokeVoid(
		w,
		"resetVirtualName",
		nil, // no parameters
	)
}

func (w *jsiiProxy_WorkspacesinstancesWorkspaceInstanceManagedInstanceBlockDeviceMappingsOutputReference) Resolve(_context cdktf.IResolveContext) interface{} {
	if err := w.validateResolveParameters(_context); err != nil {
		panic(err)
	}
	var returns interface{}

	_jsii_.Invoke(
		w,
		"resolve",
		[]interface{}{_context},
		&returns,
	)

	return returns
}

func (w *jsiiProxy_WorkspacesinstancesWorkspaceInstanceManagedInstanceBlockDeviceMappingsOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		w,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

