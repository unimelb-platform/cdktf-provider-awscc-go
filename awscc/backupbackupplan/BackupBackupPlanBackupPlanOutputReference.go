package backupbackupplan

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/unimelb-platform/cdktf-provider-awscc-go/awscc/jsii"

	"github.com/hashicorp/terraform-cdk-go/cdktf"
	"github.com/unimelb-platform/cdktf-provider-awscc-go/awscc/backupbackupplan/internal"
)

type BackupBackupPlanBackupPlanOutputReference interface {
	cdktf.ComplexObject
	AdvancedBackupSettings() BackupBackupPlanBackupPlanAdvancedBackupSettingsList
	AdvancedBackupSettingsInput() interface{}
	BackupPlanName() *string
	SetBackupPlanName(val *string)
	BackupPlanNameInput() *string
	BackupPlanRule() BackupBackupPlanBackupPlanBackupPlanRuleList
	BackupPlanRuleInput() interface{}
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
	PutAdvancedBackupSettings(value interface{})
	PutBackupPlanRule(value interface{})
	ResetAdvancedBackupSettings()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(_context cdktf.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for BackupBackupPlanBackupPlanOutputReference
type jsiiProxy_BackupBackupPlanBackupPlanOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_BackupBackupPlanBackupPlanOutputReference) AdvancedBackupSettings() BackupBackupPlanBackupPlanAdvancedBackupSettingsList {
	var returns BackupBackupPlanBackupPlanAdvancedBackupSettingsList
	_jsii_.Get(
		j,
		"advancedBackupSettings",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BackupBackupPlanBackupPlanOutputReference) AdvancedBackupSettingsInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"advancedBackupSettingsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BackupBackupPlanBackupPlanOutputReference) BackupPlanName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"backupPlanName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BackupBackupPlanBackupPlanOutputReference) BackupPlanNameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"backupPlanNameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BackupBackupPlanBackupPlanOutputReference) BackupPlanRule() BackupBackupPlanBackupPlanBackupPlanRuleList {
	var returns BackupBackupPlanBackupPlanBackupPlanRuleList
	_jsii_.Get(
		j,
		"backupPlanRule",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BackupBackupPlanBackupPlanOutputReference) BackupPlanRuleInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"backupPlanRuleInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BackupBackupPlanBackupPlanOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BackupBackupPlanBackupPlanOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BackupBackupPlanBackupPlanOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BackupBackupPlanBackupPlanOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BackupBackupPlanBackupPlanOutputReference) InternalValue() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BackupBackupPlanBackupPlanOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BackupBackupPlanBackupPlanOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}


func NewBackupBackupPlanBackupPlanOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) BackupBackupPlanBackupPlanOutputReference {
	_init_.Initialize()

	if err := validateNewBackupBackupPlanBackupPlanOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_BackupBackupPlanBackupPlanOutputReference{}

	_jsii_.Create(
		"awscc.backupBackupPlan.BackupBackupPlanBackupPlanOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewBackupBackupPlanBackupPlanOutputReference_Override(b BackupBackupPlanBackupPlanOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"awscc.backupBackupPlan.BackupBackupPlanBackupPlanOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		b,
	)
}

func (j *jsiiProxy_BackupBackupPlanBackupPlanOutputReference)SetBackupPlanName(val *string) {
	if err := j.validateSetBackupPlanNameParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"backupPlanName",
		val,
	)
}

func (j *jsiiProxy_BackupBackupPlanBackupPlanOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_BackupBackupPlanBackupPlanOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_BackupBackupPlanBackupPlanOutputReference)SetInternalValue(val interface{}) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_BackupBackupPlanBackupPlanOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_BackupBackupPlanBackupPlanOutputReference)SetTerraformResource(val cdktf.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (b *jsiiProxy_BackupBackupPlanBackupPlanOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		b,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (b *jsiiProxy_BackupBackupPlanBackupPlanOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (b *jsiiProxy_BackupBackupPlanBackupPlanOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
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

func (b *jsiiProxy_BackupBackupPlanBackupPlanOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (b *jsiiProxy_BackupBackupPlanBackupPlanOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (b *jsiiProxy_BackupBackupPlanBackupPlanOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (b *jsiiProxy_BackupBackupPlanBackupPlanOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (b *jsiiProxy_BackupBackupPlanBackupPlanOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (b *jsiiProxy_BackupBackupPlanBackupPlanOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (b *jsiiProxy_BackupBackupPlanBackupPlanOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (b *jsiiProxy_BackupBackupPlanBackupPlanOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		b,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (b *jsiiProxy_BackupBackupPlanBackupPlanOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
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

func (b *jsiiProxy_BackupBackupPlanBackupPlanOutputReference) PutAdvancedBackupSettings(value interface{}) {
	if err := b.validatePutAdvancedBackupSettingsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		b,
		"putAdvancedBackupSettings",
		[]interface{}{value},
	)
}

func (b *jsiiProxy_BackupBackupPlanBackupPlanOutputReference) PutBackupPlanRule(value interface{}) {
	if err := b.validatePutBackupPlanRuleParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		b,
		"putBackupPlanRule",
		[]interface{}{value},
	)
}

func (b *jsiiProxy_BackupBackupPlanBackupPlanOutputReference) ResetAdvancedBackupSettings() {
	_jsii_.InvokeVoid(
		b,
		"resetAdvancedBackupSettings",
		nil, // no parameters
	)
}

func (b *jsiiProxy_BackupBackupPlanBackupPlanOutputReference) Resolve(_context cdktf.IResolveContext) interface{} {
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

func (b *jsiiProxy_BackupBackupPlanBackupPlanOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		b,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

