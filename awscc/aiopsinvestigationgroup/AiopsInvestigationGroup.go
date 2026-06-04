package aiopsinvestigationgroup

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/unimelb-platform/cdktf-provider-awscc-go/awscc/jsii"

	"github.com/aws/constructs-go/constructs/v10"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
	"github.com/unimelb-platform/cdktf-provider-awscc-go/awscc/aiopsinvestigationgroup/internal"
)

// Represents a {@link https://registry.terraform.io/providers/hashicorp/awscc/1.49.0/docs/resources/aiops_investigation_group awscc_aiops_investigation_group}.
type AiopsInvestigationGroup interface {
	cdktf.TerraformResource
	Arn() *string
	// Experimental.
	CdktfStack() cdktf.TerraformStack
	ChatbotNotificationChannels() AiopsInvestigationGroupChatbotNotificationChannelsList
	ChatbotNotificationChannelsInput() interface{}
	// Experimental.
	Connection() interface{}
	// Experimental.
	SetConnection(val interface{})
	// Experimental.
	ConstructNodeMetadata() *map[string]interface{}
	// Experimental.
	Count() interface{}
	// Experimental.
	SetCount(val interface{})
	CreatedAt() *string
	CreatedBy() *string
	CrossAccountConfigurations() AiopsInvestigationGroupCrossAccountConfigurationsList
	CrossAccountConfigurationsInput() interface{}
	// Experimental.
	DependsOn() *[]*string
	// Experimental.
	SetDependsOn(val *[]*string)
	EncryptionConfig() AiopsInvestigationGroupEncryptionConfigOutputReference
	EncryptionConfigInput() interface{}
	// Experimental.
	ForEach() cdktf.ITerraformIterator
	// Experimental.
	SetForEach(val cdktf.ITerraformIterator)
	// Experimental.
	Fqn() *string
	// Experimental.
	FriendlyUniqueId() *string
	Id() *string
	InvestigationGroupPolicy() *string
	SetInvestigationGroupPolicy(val *string)
	InvestigationGroupPolicyInput() *string
	IsCloudTrailEventHistoryEnabled() interface{}
	SetIsCloudTrailEventHistoryEnabled(val interface{})
	IsCloudTrailEventHistoryEnabledInput() interface{}
	LastModifiedAt() *string
	LastModifiedBy() *string
	// Experimental.
	Lifecycle() *cdktf.TerraformResourceLifecycle
	// Experimental.
	SetLifecycle(val *cdktf.TerraformResourceLifecycle)
	Name() *string
	SetName(val *string)
	NameInput() *string
	// The tree node.
	Node() constructs.Node
	// Experimental.
	Provider() cdktf.TerraformProvider
	// Experimental.
	SetProvider(val cdktf.TerraformProvider)
	// Experimental.
	Provisioners() *[]interface{}
	// Experimental.
	SetProvisioners(val *[]interface{})
	// Experimental.
	RawOverrides() interface{}
	RetentionInDays() *float64
	SetRetentionInDays(val *float64)
	RetentionInDaysInput() *float64
	RoleArn() *string
	SetRoleArn(val *string)
	RoleArnInput() *string
	TagKeyBoundaries() *[]*string
	SetTagKeyBoundaries(val *[]*string)
	TagKeyBoundariesInput() *[]*string
	Tags() AiopsInvestigationGroupTagsList
	TagsInput() interface{}
	// Experimental.
	TerraformGeneratorMetadata() *cdktf.TerraformProviderGeneratorMetadata
	// Experimental.
	TerraformMetaArguments() *map[string]interface{}
	// Experimental.
	TerraformResourceType() *string
	// Adds a user defined moveTarget string to this resource to be later used in .moveTo(moveTarget) to resolve the location of the move.
	// Experimental.
	AddMoveTarget(moveTarget *string)
	// Experimental.
	AddOverride(path *string, value interface{})
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
	HasResourceMove() interface{}
	// Experimental.
	ImportFrom(id *string, provider cdktf.TerraformProvider)
	// Experimental.
	InterpolationForAttribute(terraformAttribute *string) cdktf.IResolvable
	// Move the resource corresponding to "id" to this resource.
	//
	// Note that the resource being moved from must be marked as moved using it's instance function.
	// Experimental.
	MoveFromId(id *string)
	// Moves this resource to the target resource given by moveTarget.
	// Experimental.
	MoveTo(moveTarget *string, index interface{})
	// Moves this resource to the resource corresponding to "id".
	// Experimental.
	MoveToId(id *string)
	// Overrides the auto-generated logical ID with a specific ID.
	// Experimental.
	OverrideLogicalId(newLogicalId *string)
	PutChatbotNotificationChannels(value interface{})
	PutCrossAccountConfigurations(value interface{})
	PutEncryptionConfig(value *AiopsInvestigationGroupEncryptionConfig)
	PutTags(value interface{})
	ResetChatbotNotificationChannels()
	ResetCrossAccountConfigurations()
	ResetEncryptionConfig()
	ResetInvestigationGroupPolicy()
	ResetIsCloudTrailEventHistoryEnabled()
	// Resets a previously passed logical Id to use the auto-generated logical id again.
	// Experimental.
	ResetOverrideLogicalId()
	ResetRetentionInDays()
	ResetRoleArn()
	ResetTagKeyBoundaries()
	ResetTags()
	SynthesizeAttributes() *map[string]interface{}
	SynthesizeHclAttributes() *map[string]interface{}
	// Experimental.
	ToHclTerraform() interface{}
	// Experimental.
	ToMetadata() interface{}
	// Returns a string representation of this construct.
	ToString() *string
	// Adds this resource to the terraform JSON output.
	// Experimental.
	ToTerraform() interface{}
}

// The jsii proxy struct for AiopsInvestigationGroup
type jsiiProxy_AiopsInvestigationGroup struct {
	internal.Type__cdktfTerraformResource
}

func (j *jsiiProxy_AiopsInvestigationGroup) Arn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"arn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AiopsInvestigationGroup) CdktfStack() cdktf.TerraformStack {
	var returns cdktf.TerraformStack
	_jsii_.Get(
		j,
		"cdktfStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AiopsInvestigationGroup) ChatbotNotificationChannels() AiopsInvestigationGroupChatbotNotificationChannelsList {
	var returns AiopsInvestigationGroupChatbotNotificationChannelsList
	_jsii_.Get(
		j,
		"chatbotNotificationChannels",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AiopsInvestigationGroup) ChatbotNotificationChannelsInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"chatbotNotificationChannelsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AiopsInvestigationGroup) Connection() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"connection",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AiopsInvestigationGroup) ConstructNodeMetadata() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"constructNodeMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AiopsInvestigationGroup) Count() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"count",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AiopsInvestigationGroup) CreatedAt() *string {
	var returns *string
	_jsii_.Get(
		j,
		"createdAt",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AiopsInvestigationGroup) CreatedBy() *string {
	var returns *string
	_jsii_.Get(
		j,
		"createdBy",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AiopsInvestigationGroup) CrossAccountConfigurations() AiopsInvestigationGroupCrossAccountConfigurationsList {
	var returns AiopsInvestigationGroupCrossAccountConfigurationsList
	_jsii_.Get(
		j,
		"crossAccountConfigurations",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AiopsInvestigationGroup) CrossAccountConfigurationsInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"crossAccountConfigurationsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AiopsInvestigationGroup) DependsOn() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"dependsOn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AiopsInvestigationGroup) EncryptionConfig() AiopsInvestigationGroupEncryptionConfigOutputReference {
	var returns AiopsInvestigationGroupEncryptionConfigOutputReference
	_jsii_.Get(
		j,
		"encryptionConfig",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AiopsInvestigationGroup) EncryptionConfigInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"encryptionConfigInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AiopsInvestigationGroup) ForEach() cdktf.ITerraformIterator {
	var returns cdktf.ITerraformIterator
	_jsii_.Get(
		j,
		"forEach",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AiopsInvestigationGroup) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AiopsInvestigationGroup) FriendlyUniqueId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"friendlyUniqueId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AiopsInvestigationGroup) Id() *string {
	var returns *string
	_jsii_.Get(
		j,
		"id",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AiopsInvestigationGroup) InvestigationGroupPolicy() *string {
	var returns *string
	_jsii_.Get(
		j,
		"investigationGroupPolicy",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AiopsInvestigationGroup) InvestigationGroupPolicyInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"investigationGroupPolicyInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AiopsInvestigationGroup) IsCloudTrailEventHistoryEnabled() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"isCloudTrailEventHistoryEnabled",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AiopsInvestigationGroup) IsCloudTrailEventHistoryEnabledInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"isCloudTrailEventHistoryEnabledInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AiopsInvestigationGroup) LastModifiedAt() *string {
	var returns *string
	_jsii_.Get(
		j,
		"lastModifiedAt",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AiopsInvestigationGroup) LastModifiedBy() *string {
	var returns *string
	_jsii_.Get(
		j,
		"lastModifiedBy",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AiopsInvestigationGroup) Lifecycle() *cdktf.TerraformResourceLifecycle {
	var returns *cdktf.TerraformResourceLifecycle
	_jsii_.Get(
		j,
		"lifecycle",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AiopsInvestigationGroup) Name() *string {
	var returns *string
	_jsii_.Get(
		j,
		"name",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AiopsInvestigationGroup) NameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"nameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AiopsInvestigationGroup) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AiopsInvestigationGroup) Provider() cdktf.TerraformProvider {
	var returns cdktf.TerraformProvider
	_jsii_.Get(
		j,
		"provider",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AiopsInvestigationGroup) Provisioners() *[]interface{} {
	var returns *[]interface{}
	_jsii_.Get(
		j,
		"provisioners",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AiopsInvestigationGroup) RawOverrides() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"rawOverrides",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AiopsInvestigationGroup) RetentionInDays() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"retentionInDays",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AiopsInvestigationGroup) RetentionInDaysInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"retentionInDaysInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AiopsInvestigationGroup) RoleArn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"roleArn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AiopsInvestigationGroup) RoleArnInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"roleArnInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AiopsInvestigationGroup) TagKeyBoundaries() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"tagKeyBoundaries",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AiopsInvestigationGroup) TagKeyBoundariesInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"tagKeyBoundariesInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AiopsInvestigationGroup) Tags() AiopsInvestigationGroupTagsList {
	var returns AiopsInvestigationGroupTagsList
	_jsii_.Get(
		j,
		"tags",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AiopsInvestigationGroup) TagsInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"tagsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AiopsInvestigationGroup) TerraformGeneratorMetadata() *cdktf.TerraformProviderGeneratorMetadata {
	var returns *cdktf.TerraformProviderGeneratorMetadata
	_jsii_.Get(
		j,
		"terraformGeneratorMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AiopsInvestigationGroup) TerraformMetaArguments() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"terraformMetaArguments",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AiopsInvestigationGroup) TerraformResourceType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformResourceType",
		&returns,
	)
	return returns
}


// Create a new {@link https://registry.terraform.io/providers/hashicorp/awscc/1.49.0/docs/resources/aiops_investigation_group awscc_aiops_investigation_group} Resource.
func NewAiopsInvestigationGroup(scope constructs.Construct, id *string, config *AiopsInvestigationGroupConfig) AiopsInvestigationGroup {
	_init_.Initialize()

	if err := validateNewAiopsInvestigationGroupParameters(scope, id, config); err != nil {
		panic(err)
	}
	j := jsiiProxy_AiopsInvestigationGroup{}

	_jsii_.Create(
		"awscc.aiopsInvestigationGroup.AiopsInvestigationGroup",
		[]interface{}{scope, id, config},
		&j,
	)

	return &j
}

// Create a new {@link https://registry.terraform.io/providers/hashicorp/awscc/1.49.0/docs/resources/aiops_investigation_group awscc_aiops_investigation_group} Resource.
func NewAiopsInvestigationGroup_Override(a AiopsInvestigationGroup, scope constructs.Construct, id *string, config *AiopsInvestigationGroupConfig) {
	_init_.Initialize()

	_jsii_.Create(
		"awscc.aiopsInvestigationGroup.AiopsInvestigationGroup",
		[]interface{}{scope, id, config},
		a,
	)
}

func (j *jsiiProxy_AiopsInvestigationGroup)SetConnection(val interface{}) {
	if err := j.validateSetConnectionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"connection",
		val,
	)
}

func (j *jsiiProxy_AiopsInvestigationGroup)SetCount(val interface{}) {
	if err := j.validateSetCountParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"count",
		val,
	)
}

func (j *jsiiProxy_AiopsInvestigationGroup)SetDependsOn(val *[]*string) {
	_jsii_.Set(
		j,
		"dependsOn",
		val,
	)
}

func (j *jsiiProxy_AiopsInvestigationGroup)SetForEach(val cdktf.ITerraformIterator) {
	_jsii_.Set(
		j,
		"forEach",
		val,
	)
}

func (j *jsiiProxy_AiopsInvestigationGroup)SetInvestigationGroupPolicy(val *string) {
	if err := j.validateSetInvestigationGroupPolicyParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"investigationGroupPolicy",
		val,
	)
}

func (j *jsiiProxy_AiopsInvestigationGroup)SetIsCloudTrailEventHistoryEnabled(val interface{}) {
	if err := j.validateSetIsCloudTrailEventHistoryEnabledParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"isCloudTrailEventHistoryEnabled",
		val,
	)
}

func (j *jsiiProxy_AiopsInvestigationGroup)SetLifecycle(val *cdktf.TerraformResourceLifecycle) {
	if err := j.validateSetLifecycleParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"lifecycle",
		val,
	)
}

func (j *jsiiProxy_AiopsInvestigationGroup)SetName(val *string) {
	if err := j.validateSetNameParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"name",
		val,
	)
}

func (j *jsiiProxy_AiopsInvestigationGroup)SetProvider(val cdktf.TerraformProvider) {
	_jsii_.Set(
		j,
		"provider",
		val,
	)
}

func (j *jsiiProxy_AiopsInvestigationGroup)SetProvisioners(val *[]interface{}) {
	if err := j.validateSetProvisionersParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"provisioners",
		val,
	)
}

func (j *jsiiProxy_AiopsInvestigationGroup)SetRetentionInDays(val *float64) {
	if err := j.validateSetRetentionInDaysParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"retentionInDays",
		val,
	)
}

func (j *jsiiProxy_AiopsInvestigationGroup)SetRoleArn(val *string) {
	if err := j.validateSetRoleArnParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"roleArn",
		val,
	)
}

func (j *jsiiProxy_AiopsInvestigationGroup)SetTagKeyBoundaries(val *[]*string) {
	if err := j.validateSetTagKeyBoundariesParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"tagKeyBoundaries",
		val,
	)
}

// Generates CDKTF code for importing a AiopsInvestigationGroup resource upon running "cdktf plan <stack-name>".
func AiopsInvestigationGroup_GenerateConfigForImport(scope constructs.Construct, importToId *string, importFromId *string, provider cdktf.TerraformProvider) cdktf.ImportableResource {
	_init_.Initialize()

	if err := validateAiopsInvestigationGroup_GenerateConfigForImportParameters(scope, importToId, importFromId); err != nil {
		panic(err)
	}
	var returns cdktf.ImportableResource

	_jsii_.StaticInvoke(
		"awscc.aiopsInvestigationGroup.AiopsInvestigationGroup",
		"generateConfigForImport",
		[]interface{}{scope, importToId, importFromId, provider},
		&returns,
	)

	return returns
}

// Checks if `x` is a construct.
//
// Use this method instead of `instanceof` to properly detect `Construct`
// instances, even when the construct library is symlinked.
//
// Explanation: in JavaScript, multiple copies of the `constructs` library on
// disk are seen as independent, completely different libraries. As a
// consequence, the class `Construct` in each copy of the `constructs` library
// is seen as a different class, and an instance of one class will not test as
// `instanceof` the other class. `npm install` will not create installations
// like this, but users may manually symlink construct libraries together or
// use a monorepo tool: in those cases, multiple copies of the `constructs`
// library can be accidentally installed, and `instanceof` will behave
// unpredictably. It is safest to avoid using `instanceof`, and using
// this type-testing method instead.
//
// Returns: true if `x` is an object created from a class which extends `Construct`.
func AiopsInvestigationGroup_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	if err := validateAiopsInvestigationGroup_IsConstructParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"awscc.aiopsInvestigationGroup.AiopsInvestigationGroup",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Experimental.
func AiopsInvestigationGroup_IsTerraformElement(x interface{}) *bool {
	_init_.Initialize()

	if err := validateAiopsInvestigationGroup_IsTerraformElementParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"awscc.aiopsInvestigationGroup.AiopsInvestigationGroup",
		"isTerraformElement",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Experimental.
func AiopsInvestigationGroup_IsTerraformResource(x interface{}) *bool {
	_init_.Initialize()

	if err := validateAiopsInvestigationGroup_IsTerraformResourceParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"awscc.aiopsInvestigationGroup.AiopsInvestigationGroup",
		"isTerraformResource",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func AiopsInvestigationGroup_TfResourceType() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"awscc.aiopsInvestigationGroup.AiopsInvestigationGroup",
		"tfResourceType",
		&returns,
	)
	return returns
}

func (a *jsiiProxy_AiopsInvestigationGroup) AddMoveTarget(moveTarget *string) {
	if err := a.validateAddMoveTargetParameters(moveTarget); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"addMoveTarget",
		[]interface{}{moveTarget},
	)
}

func (a *jsiiProxy_AiopsInvestigationGroup) AddOverride(path *string, value interface{}) {
	if err := a.validateAddOverrideParameters(path, value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"addOverride",
		[]interface{}{path, value},
	)
}

func (a *jsiiProxy_AiopsInvestigationGroup) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (a *jsiiProxy_AiopsInvestigationGroup) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
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

func (a *jsiiProxy_AiopsInvestigationGroup) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (a *jsiiProxy_AiopsInvestigationGroup) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (a *jsiiProxy_AiopsInvestigationGroup) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (a *jsiiProxy_AiopsInvestigationGroup) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (a *jsiiProxy_AiopsInvestigationGroup) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (a *jsiiProxy_AiopsInvestigationGroup) GetStringAttribute(terraformAttribute *string) *string {
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

func (a *jsiiProxy_AiopsInvestigationGroup) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (a *jsiiProxy_AiopsInvestigationGroup) HasResourceMove() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		a,
		"hasResourceMove",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AiopsInvestigationGroup) ImportFrom(id *string, provider cdktf.TerraformProvider) {
	if err := a.validateImportFromParameters(id); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"importFrom",
		[]interface{}{id, provider},
	)
}

func (a *jsiiProxy_AiopsInvestigationGroup) InterpolationForAttribute(terraformAttribute *string) cdktf.IResolvable {
	if err := a.validateInterpolationForAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		a,
		"interpolationForAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AiopsInvestigationGroup) MoveFromId(id *string) {
	if err := a.validateMoveFromIdParameters(id); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"moveFromId",
		[]interface{}{id},
	)
}

func (a *jsiiProxy_AiopsInvestigationGroup) MoveTo(moveTarget *string, index interface{}) {
	if err := a.validateMoveToParameters(moveTarget, index); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"moveTo",
		[]interface{}{moveTarget, index},
	)
}

func (a *jsiiProxy_AiopsInvestigationGroup) MoveToId(id *string) {
	if err := a.validateMoveToIdParameters(id); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"moveToId",
		[]interface{}{id},
	)
}

func (a *jsiiProxy_AiopsInvestigationGroup) OverrideLogicalId(newLogicalId *string) {
	if err := a.validateOverrideLogicalIdParameters(newLogicalId); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"overrideLogicalId",
		[]interface{}{newLogicalId},
	)
}

func (a *jsiiProxy_AiopsInvestigationGroup) PutChatbotNotificationChannels(value interface{}) {
	if err := a.validatePutChatbotNotificationChannelsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"putChatbotNotificationChannels",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_AiopsInvestigationGroup) PutCrossAccountConfigurations(value interface{}) {
	if err := a.validatePutCrossAccountConfigurationsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"putCrossAccountConfigurations",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_AiopsInvestigationGroup) PutEncryptionConfig(value *AiopsInvestigationGroupEncryptionConfig) {
	if err := a.validatePutEncryptionConfigParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"putEncryptionConfig",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_AiopsInvestigationGroup) PutTags(value interface{}) {
	if err := a.validatePutTagsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"putTags",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_AiopsInvestigationGroup) ResetChatbotNotificationChannels() {
	_jsii_.InvokeVoid(
		a,
		"resetChatbotNotificationChannels",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AiopsInvestigationGroup) ResetCrossAccountConfigurations() {
	_jsii_.InvokeVoid(
		a,
		"resetCrossAccountConfigurations",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AiopsInvestigationGroup) ResetEncryptionConfig() {
	_jsii_.InvokeVoid(
		a,
		"resetEncryptionConfig",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AiopsInvestigationGroup) ResetInvestigationGroupPolicy() {
	_jsii_.InvokeVoid(
		a,
		"resetInvestigationGroupPolicy",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AiopsInvestigationGroup) ResetIsCloudTrailEventHistoryEnabled() {
	_jsii_.InvokeVoid(
		a,
		"resetIsCloudTrailEventHistoryEnabled",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AiopsInvestigationGroup) ResetOverrideLogicalId() {
	_jsii_.InvokeVoid(
		a,
		"resetOverrideLogicalId",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AiopsInvestigationGroup) ResetRetentionInDays() {
	_jsii_.InvokeVoid(
		a,
		"resetRetentionInDays",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AiopsInvestigationGroup) ResetRoleArn() {
	_jsii_.InvokeVoid(
		a,
		"resetRoleArn",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AiopsInvestigationGroup) ResetTagKeyBoundaries() {
	_jsii_.InvokeVoid(
		a,
		"resetTagKeyBoundaries",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AiopsInvestigationGroup) ResetTags() {
	_jsii_.InvokeVoid(
		a,
		"resetTags",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AiopsInvestigationGroup) SynthesizeAttributes() *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		a,
		"synthesizeAttributes",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AiopsInvestigationGroup) SynthesizeHclAttributes() *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		a,
		"synthesizeHclAttributes",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AiopsInvestigationGroup) ToHclTerraform() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		a,
		"toHclTerraform",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AiopsInvestigationGroup) ToMetadata() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		a,
		"toMetadata",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AiopsInvestigationGroup) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		a,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AiopsInvestigationGroup) ToTerraform() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		a,
		"toTerraform",
		nil, // no parameters
		&returns,
	)

	return returns
}

