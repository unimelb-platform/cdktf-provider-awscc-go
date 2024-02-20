package backupbackupplan

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/unimelb-platform/cdktf-provider-awscc-go/awscc/jsii"

	"github.com/hashicorp/terraform-cdk-go/cdktf"
	"github.com/unimelb-platform/cdktf-provider-awscc-go/awscc/backupbackupplan/internal"
)

type BackupBackupPlanBackupPlanBackupPlanRuleOutputReference interface {
	cdktf.ComplexObject
	CompletionWindowMinutes() *float64
	SetCompletionWindowMinutes(val *float64)
	CompletionWindowMinutesInput() *float64
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
	CopyActions() BackupBackupPlanBackupPlanBackupPlanRuleCopyActionsList
	CopyActionsInput() interface{}
	// The creation stack of this resolvable which will be appended to errors thrown during resolution.
	//
	// If this returns an empty array the stack will not be attached.
	// Experimental.
	CreationStack() *[]*string
	EnableContinuousBackup() interface{}
	SetEnableContinuousBackup(val interface{})
	EnableContinuousBackupInput() interface{}
	// Experimental.
	Fqn() *string
	InternalValue() interface{}
	SetInternalValue(val interface{})
	Lifecycle() BackupBackupPlanBackupPlanBackupPlanRuleLifecycleOutputReference
	LifecycleInput() interface{}
	RecoveryPointTags() *map[string]*string
	SetRecoveryPointTags(val *map[string]*string)
	RecoveryPointTagsInput() *map[string]*string
	RuleName() *string
	SetRuleName(val *string)
	RuleNameInput() *string
	ScheduleExpression() *string
	SetScheduleExpression(val *string)
	ScheduleExpressionInput() *string
	ScheduleExpressionTimezone() *string
	SetScheduleExpressionTimezone(val *string)
	ScheduleExpressionTimezoneInput() *string
	StartWindowMinutes() *float64
	SetStartWindowMinutes(val *float64)
	StartWindowMinutesInput() *float64
	TargetBackupVault() *string
	SetTargetBackupVault(val *string)
	TargetBackupVaultInput() *string
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
	PutCopyActions(value interface{})
	PutLifecycle(value *BackupBackupPlanBackupPlanBackupPlanRuleLifecycle)
	ResetCompletionWindowMinutes()
	ResetCopyActions()
	ResetEnableContinuousBackup()
	ResetLifecycle()
	ResetRecoveryPointTags()
	ResetScheduleExpression()
	ResetScheduleExpressionTimezone()
	ResetStartWindowMinutes()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(_context cdktf.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for BackupBackupPlanBackupPlanBackupPlanRuleOutputReference
type jsiiProxy_BackupBackupPlanBackupPlanBackupPlanRuleOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_BackupBackupPlanBackupPlanBackupPlanRuleOutputReference) CompletionWindowMinutes() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"completionWindowMinutes",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BackupBackupPlanBackupPlanBackupPlanRuleOutputReference) CompletionWindowMinutesInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"completionWindowMinutesInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BackupBackupPlanBackupPlanBackupPlanRuleOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BackupBackupPlanBackupPlanBackupPlanRuleOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BackupBackupPlanBackupPlanBackupPlanRuleOutputReference) CopyActions() BackupBackupPlanBackupPlanBackupPlanRuleCopyActionsList {
	var returns BackupBackupPlanBackupPlanBackupPlanRuleCopyActionsList
	_jsii_.Get(
		j,
		"copyActions",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BackupBackupPlanBackupPlanBackupPlanRuleOutputReference) CopyActionsInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"copyActionsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BackupBackupPlanBackupPlanBackupPlanRuleOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BackupBackupPlanBackupPlanBackupPlanRuleOutputReference) EnableContinuousBackup() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"enableContinuousBackup",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BackupBackupPlanBackupPlanBackupPlanRuleOutputReference) EnableContinuousBackupInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"enableContinuousBackupInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BackupBackupPlanBackupPlanBackupPlanRuleOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BackupBackupPlanBackupPlanBackupPlanRuleOutputReference) InternalValue() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BackupBackupPlanBackupPlanBackupPlanRuleOutputReference) Lifecycle() BackupBackupPlanBackupPlanBackupPlanRuleLifecycleOutputReference {
	var returns BackupBackupPlanBackupPlanBackupPlanRuleLifecycleOutputReference
	_jsii_.Get(
		j,
		"lifecycle",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BackupBackupPlanBackupPlanBackupPlanRuleOutputReference) LifecycleInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"lifecycleInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BackupBackupPlanBackupPlanBackupPlanRuleOutputReference) RecoveryPointTags() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"recoveryPointTags",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BackupBackupPlanBackupPlanBackupPlanRuleOutputReference) RecoveryPointTagsInput() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"recoveryPointTagsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BackupBackupPlanBackupPlanBackupPlanRuleOutputReference) RuleName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"ruleName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BackupBackupPlanBackupPlanBackupPlanRuleOutputReference) RuleNameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"ruleNameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BackupBackupPlanBackupPlanBackupPlanRuleOutputReference) ScheduleExpression() *string {
	var returns *string
	_jsii_.Get(
		j,
		"scheduleExpression",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BackupBackupPlanBackupPlanBackupPlanRuleOutputReference) ScheduleExpressionInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"scheduleExpressionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BackupBackupPlanBackupPlanBackupPlanRuleOutputReference) ScheduleExpressionTimezone() *string {
	var returns *string
	_jsii_.Get(
		j,
		"scheduleExpressionTimezone",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BackupBackupPlanBackupPlanBackupPlanRuleOutputReference) ScheduleExpressionTimezoneInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"scheduleExpressionTimezoneInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BackupBackupPlanBackupPlanBackupPlanRuleOutputReference) StartWindowMinutes() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"startWindowMinutes",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BackupBackupPlanBackupPlanBackupPlanRuleOutputReference) StartWindowMinutesInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"startWindowMinutesInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BackupBackupPlanBackupPlanBackupPlanRuleOutputReference) TargetBackupVault() *string {
	var returns *string
	_jsii_.Get(
		j,
		"targetBackupVault",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BackupBackupPlanBackupPlanBackupPlanRuleOutputReference) TargetBackupVaultInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"targetBackupVaultInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BackupBackupPlanBackupPlanBackupPlanRuleOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BackupBackupPlanBackupPlanBackupPlanRuleOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}


func NewBackupBackupPlanBackupPlanBackupPlanRuleOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) BackupBackupPlanBackupPlanBackupPlanRuleOutputReference {
	_init_.Initialize()

	if err := validateNewBackupBackupPlanBackupPlanBackupPlanRuleOutputReferenceParameters(terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet); err != nil {
		panic(err)
	}
	j := jsiiProxy_BackupBackupPlanBackupPlanBackupPlanRuleOutputReference{}

	_jsii_.Create(
		"awscc.backupBackupPlan.BackupBackupPlanBackupPlanBackupPlanRuleOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		&j,
	)

	return &j
}

func NewBackupBackupPlanBackupPlanBackupPlanRuleOutputReference_Override(b BackupBackupPlanBackupPlanBackupPlanRuleOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) {
	_init_.Initialize()

	_jsii_.Create(
		"awscc.backupBackupPlan.BackupBackupPlanBackupPlanBackupPlanRuleOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		b,
	)
}

func (j *jsiiProxy_BackupBackupPlanBackupPlanBackupPlanRuleOutputReference)SetCompletionWindowMinutes(val *float64) {
	if err := j.validateSetCompletionWindowMinutesParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"completionWindowMinutes",
		val,
	)
}

func (j *jsiiProxy_BackupBackupPlanBackupPlanBackupPlanRuleOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_BackupBackupPlanBackupPlanBackupPlanRuleOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_BackupBackupPlanBackupPlanBackupPlanRuleOutputReference)SetEnableContinuousBackup(val interface{}) {
	if err := j.validateSetEnableContinuousBackupParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"enableContinuousBackup",
		val,
	)
}

func (j *jsiiProxy_BackupBackupPlanBackupPlanBackupPlanRuleOutputReference)SetInternalValue(val interface{}) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_BackupBackupPlanBackupPlanBackupPlanRuleOutputReference)SetRecoveryPointTags(val *map[string]*string) {
	if err := j.validateSetRecoveryPointTagsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"recoveryPointTags",
		val,
	)
}

func (j *jsiiProxy_BackupBackupPlanBackupPlanBackupPlanRuleOutputReference)SetRuleName(val *string) {
	if err := j.validateSetRuleNameParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"ruleName",
		val,
	)
}

func (j *jsiiProxy_BackupBackupPlanBackupPlanBackupPlanRuleOutputReference)SetScheduleExpression(val *string) {
	if err := j.validateSetScheduleExpressionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"scheduleExpression",
		val,
	)
}

func (j *jsiiProxy_BackupBackupPlanBackupPlanBackupPlanRuleOutputReference)SetScheduleExpressionTimezone(val *string) {
	if err := j.validateSetScheduleExpressionTimezoneParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"scheduleExpressionTimezone",
		val,
	)
}

func (j *jsiiProxy_BackupBackupPlanBackupPlanBackupPlanRuleOutputReference)SetStartWindowMinutes(val *float64) {
	if err := j.validateSetStartWindowMinutesParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"startWindowMinutes",
		val,
	)
}

func (j *jsiiProxy_BackupBackupPlanBackupPlanBackupPlanRuleOutputReference)SetTargetBackupVault(val *string) {
	if err := j.validateSetTargetBackupVaultParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"targetBackupVault",
		val,
	)
}

func (j *jsiiProxy_BackupBackupPlanBackupPlanBackupPlanRuleOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_BackupBackupPlanBackupPlanBackupPlanRuleOutputReference)SetTerraformResource(val cdktf.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (b *jsiiProxy_BackupBackupPlanBackupPlanBackupPlanRuleOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		b,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (b *jsiiProxy_BackupBackupPlanBackupPlanBackupPlanRuleOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (b *jsiiProxy_BackupBackupPlanBackupPlanBackupPlanRuleOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
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

func (b *jsiiProxy_BackupBackupPlanBackupPlanBackupPlanRuleOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (b *jsiiProxy_BackupBackupPlanBackupPlanBackupPlanRuleOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (b *jsiiProxy_BackupBackupPlanBackupPlanBackupPlanRuleOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (b *jsiiProxy_BackupBackupPlanBackupPlanBackupPlanRuleOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (b *jsiiProxy_BackupBackupPlanBackupPlanBackupPlanRuleOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (b *jsiiProxy_BackupBackupPlanBackupPlanBackupPlanRuleOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (b *jsiiProxy_BackupBackupPlanBackupPlanBackupPlanRuleOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (b *jsiiProxy_BackupBackupPlanBackupPlanBackupPlanRuleOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		b,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (b *jsiiProxy_BackupBackupPlanBackupPlanBackupPlanRuleOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
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

func (b *jsiiProxy_BackupBackupPlanBackupPlanBackupPlanRuleOutputReference) PutCopyActions(value interface{}) {
	if err := b.validatePutCopyActionsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		b,
		"putCopyActions",
		[]interface{}{value},
	)
}

func (b *jsiiProxy_BackupBackupPlanBackupPlanBackupPlanRuleOutputReference) PutLifecycle(value *BackupBackupPlanBackupPlanBackupPlanRuleLifecycle) {
	if err := b.validatePutLifecycleParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		b,
		"putLifecycle",
		[]interface{}{value},
	)
}

func (b *jsiiProxy_BackupBackupPlanBackupPlanBackupPlanRuleOutputReference) ResetCompletionWindowMinutes() {
	_jsii_.InvokeVoid(
		b,
		"resetCompletionWindowMinutes",
		nil, // no parameters
	)
}

func (b *jsiiProxy_BackupBackupPlanBackupPlanBackupPlanRuleOutputReference) ResetCopyActions() {
	_jsii_.InvokeVoid(
		b,
		"resetCopyActions",
		nil, // no parameters
	)
}

func (b *jsiiProxy_BackupBackupPlanBackupPlanBackupPlanRuleOutputReference) ResetEnableContinuousBackup() {
	_jsii_.InvokeVoid(
		b,
		"resetEnableContinuousBackup",
		nil, // no parameters
	)
}

func (b *jsiiProxy_BackupBackupPlanBackupPlanBackupPlanRuleOutputReference) ResetLifecycle() {
	_jsii_.InvokeVoid(
		b,
		"resetLifecycle",
		nil, // no parameters
	)
}

func (b *jsiiProxy_BackupBackupPlanBackupPlanBackupPlanRuleOutputReference) ResetRecoveryPointTags() {
	_jsii_.InvokeVoid(
		b,
		"resetRecoveryPointTags",
		nil, // no parameters
	)
}

func (b *jsiiProxy_BackupBackupPlanBackupPlanBackupPlanRuleOutputReference) ResetScheduleExpression() {
	_jsii_.InvokeVoid(
		b,
		"resetScheduleExpression",
		nil, // no parameters
	)
}

func (b *jsiiProxy_BackupBackupPlanBackupPlanBackupPlanRuleOutputReference) ResetScheduleExpressionTimezone() {
	_jsii_.InvokeVoid(
		b,
		"resetScheduleExpressionTimezone",
		nil, // no parameters
	)
}

func (b *jsiiProxy_BackupBackupPlanBackupPlanBackupPlanRuleOutputReference) ResetStartWindowMinutes() {
	_jsii_.InvokeVoid(
		b,
		"resetStartWindowMinutes",
		nil, // no parameters
	)
}

func (b *jsiiProxy_BackupBackupPlanBackupPlanBackupPlanRuleOutputReference) Resolve(_context cdktf.IResolveContext) interface{} {
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

func (b *jsiiProxy_BackupBackupPlanBackupPlanBackupPlanRuleOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		b,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

