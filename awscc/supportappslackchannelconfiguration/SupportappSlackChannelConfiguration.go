package supportappslackchannelconfiguration

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/unimelb-platform/cdktf-provider-awscc-go/awscc/jsii"

	"github.com/aws/constructs-go/constructs/v10"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
	"github.com/unimelb-platform/cdktf-provider-awscc-go/awscc/supportappslackchannelconfiguration/internal"
)

// Represents a {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/supportapp_slack_channel_configuration awscc_supportapp_slack_channel_configuration}.
type SupportappSlackChannelConfiguration interface {
	cdktf.TerraformResource
	// Experimental.
	CdktfStack() cdktf.TerraformStack
	ChannelId() *string
	SetChannelId(val *string)
	ChannelIdInput() *string
	ChannelName() *string
	SetChannelName(val *string)
	ChannelNameInput() *string
	ChannelRoleArn() *string
	SetChannelRoleArn(val *string)
	ChannelRoleArnInput() *string
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
	// Experimental.
	DependsOn() *[]*string
	// Experimental.
	SetDependsOn(val *[]*string)
	// Experimental.
	ForEach() cdktf.ITerraformIterator
	// Experimental.
	SetForEach(val cdktf.ITerraformIterator)
	// Experimental.
	Fqn() *string
	// Experimental.
	FriendlyUniqueId() *string
	Id() *string
	// Experimental.
	Lifecycle() *cdktf.TerraformResourceLifecycle
	// Experimental.
	SetLifecycle(val *cdktf.TerraformResourceLifecycle)
	// The tree node.
	Node() constructs.Node
	NotifyOnAddCorrespondenceToCase() interface{}
	SetNotifyOnAddCorrespondenceToCase(val interface{})
	NotifyOnAddCorrespondenceToCaseInput() interface{}
	NotifyOnCaseSeverity() *string
	SetNotifyOnCaseSeverity(val *string)
	NotifyOnCaseSeverityInput() *string
	NotifyOnCreateOrReopenCase() interface{}
	SetNotifyOnCreateOrReopenCase(val interface{})
	NotifyOnCreateOrReopenCaseInput() interface{}
	NotifyOnResolveCase() interface{}
	SetNotifyOnResolveCase(val interface{})
	NotifyOnResolveCaseInput() interface{}
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
	TeamId() *string
	SetTeamId(val *string)
	TeamIdInput() *string
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
	ResetChannelName()
	ResetNotifyOnAddCorrespondenceToCase()
	ResetNotifyOnCreateOrReopenCase()
	ResetNotifyOnResolveCase()
	// Resets a previously passed logical Id to use the auto-generated logical id again.
	// Experimental.
	ResetOverrideLogicalId()
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

// The jsii proxy struct for SupportappSlackChannelConfiguration
type jsiiProxy_SupportappSlackChannelConfiguration struct {
	internal.Type__cdktfTerraformResource
}

func (j *jsiiProxy_SupportappSlackChannelConfiguration) CdktfStack() cdktf.TerraformStack {
	var returns cdktf.TerraformStack
	_jsii_.Get(
		j,
		"cdktfStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SupportappSlackChannelConfiguration) ChannelId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"channelId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SupportappSlackChannelConfiguration) ChannelIdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"channelIdInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SupportappSlackChannelConfiguration) ChannelName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"channelName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SupportappSlackChannelConfiguration) ChannelNameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"channelNameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SupportappSlackChannelConfiguration) ChannelRoleArn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"channelRoleArn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SupportappSlackChannelConfiguration) ChannelRoleArnInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"channelRoleArnInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SupportappSlackChannelConfiguration) Connection() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"connection",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SupportappSlackChannelConfiguration) ConstructNodeMetadata() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"constructNodeMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SupportappSlackChannelConfiguration) Count() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"count",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SupportappSlackChannelConfiguration) DependsOn() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"dependsOn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SupportappSlackChannelConfiguration) ForEach() cdktf.ITerraformIterator {
	var returns cdktf.ITerraformIterator
	_jsii_.Get(
		j,
		"forEach",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SupportappSlackChannelConfiguration) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SupportappSlackChannelConfiguration) FriendlyUniqueId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"friendlyUniqueId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SupportappSlackChannelConfiguration) Id() *string {
	var returns *string
	_jsii_.Get(
		j,
		"id",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SupportappSlackChannelConfiguration) Lifecycle() *cdktf.TerraformResourceLifecycle {
	var returns *cdktf.TerraformResourceLifecycle
	_jsii_.Get(
		j,
		"lifecycle",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SupportappSlackChannelConfiguration) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SupportappSlackChannelConfiguration) NotifyOnAddCorrespondenceToCase() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"notifyOnAddCorrespondenceToCase",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SupportappSlackChannelConfiguration) NotifyOnAddCorrespondenceToCaseInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"notifyOnAddCorrespondenceToCaseInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SupportappSlackChannelConfiguration) NotifyOnCaseSeverity() *string {
	var returns *string
	_jsii_.Get(
		j,
		"notifyOnCaseSeverity",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SupportappSlackChannelConfiguration) NotifyOnCaseSeverityInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"notifyOnCaseSeverityInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SupportappSlackChannelConfiguration) NotifyOnCreateOrReopenCase() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"notifyOnCreateOrReopenCase",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SupportappSlackChannelConfiguration) NotifyOnCreateOrReopenCaseInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"notifyOnCreateOrReopenCaseInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SupportappSlackChannelConfiguration) NotifyOnResolveCase() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"notifyOnResolveCase",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SupportappSlackChannelConfiguration) NotifyOnResolveCaseInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"notifyOnResolveCaseInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SupportappSlackChannelConfiguration) Provider() cdktf.TerraformProvider {
	var returns cdktf.TerraformProvider
	_jsii_.Get(
		j,
		"provider",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SupportappSlackChannelConfiguration) Provisioners() *[]interface{} {
	var returns *[]interface{}
	_jsii_.Get(
		j,
		"provisioners",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SupportappSlackChannelConfiguration) RawOverrides() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"rawOverrides",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SupportappSlackChannelConfiguration) TeamId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"teamId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SupportappSlackChannelConfiguration) TeamIdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"teamIdInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SupportappSlackChannelConfiguration) TerraformGeneratorMetadata() *cdktf.TerraformProviderGeneratorMetadata {
	var returns *cdktf.TerraformProviderGeneratorMetadata
	_jsii_.Get(
		j,
		"terraformGeneratorMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SupportappSlackChannelConfiguration) TerraformMetaArguments() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"terraformMetaArguments",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SupportappSlackChannelConfiguration) TerraformResourceType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformResourceType",
		&returns,
	)
	return returns
}


// Create a new {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/supportapp_slack_channel_configuration awscc_supportapp_slack_channel_configuration} Resource.
func NewSupportappSlackChannelConfiguration(scope constructs.Construct, id *string, config *SupportappSlackChannelConfigurationConfig) SupportappSlackChannelConfiguration {
	_init_.Initialize()

	if err := validateNewSupportappSlackChannelConfigurationParameters(scope, id, config); err != nil {
		panic(err)
	}
	j := jsiiProxy_SupportappSlackChannelConfiguration{}

	_jsii_.Create(
		"awscc.supportappSlackChannelConfiguration.SupportappSlackChannelConfiguration",
		[]interface{}{scope, id, config},
		&j,
	)

	return &j
}

// Create a new {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/supportapp_slack_channel_configuration awscc_supportapp_slack_channel_configuration} Resource.
func NewSupportappSlackChannelConfiguration_Override(s SupportappSlackChannelConfiguration, scope constructs.Construct, id *string, config *SupportappSlackChannelConfigurationConfig) {
	_init_.Initialize()

	_jsii_.Create(
		"awscc.supportappSlackChannelConfiguration.SupportappSlackChannelConfiguration",
		[]interface{}{scope, id, config},
		s,
	)
}

func (j *jsiiProxy_SupportappSlackChannelConfiguration)SetChannelId(val *string) {
	if err := j.validateSetChannelIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"channelId",
		val,
	)
}

func (j *jsiiProxy_SupportappSlackChannelConfiguration)SetChannelName(val *string) {
	if err := j.validateSetChannelNameParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"channelName",
		val,
	)
}

func (j *jsiiProxy_SupportappSlackChannelConfiguration)SetChannelRoleArn(val *string) {
	if err := j.validateSetChannelRoleArnParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"channelRoleArn",
		val,
	)
}

func (j *jsiiProxy_SupportappSlackChannelConfiguration)SetConnection(val interface{}) {
	if err := j.validateSetConnectionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"connection",
		val,
	)
}

func (j *jsiiProxy_SupportappSlackChannelConfiguration)SetCount(val interface{}) {
	if err := j.validateSetCountParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"count",
		val,
	)
}

func (j *jsiiProxy_SupportappSlackChannelConfiguration)SetDependsOn(val *[]*string) {
	_jsii_.Set(
		j,
		"dependsOn",
		val,
	)
}

func (j *jsiiProxy_SupportappSlackChannelConfiguration)SetForEach(val cdktf.ITerraformIterator) {
	_jsii_.Set(
		j,
		"forEach",
		val,
	)
}

func (j *jsiiProxy_SupportappSlackChannelConfiguration)SetLifecycle(val *cdktf.TerraformResourceLifecycle) {
	if err := j.validateSetLifecycleParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"lifecycle",
		val,
	)
}

func (j *jsiiProxy_SupportappSlackChannelConfiguration)SetNotifyOnAddCorrespondenceToCase(val interface{}) {
	if err := j.validateSetNotifyOnAddCorrespondenceToCaseParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"notifyOnAddCorrespondenceToCase",
		val,
	)
}

func (j *jsiiProxy_SupportappSlackChannelConfiguration)SetNotifyOnCaseSeverity(val *string) {
	if err := j.validateSetNotifyOnCaseSeverityParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"notifyOnCaseSeverity",
		val,
	)
}

func (j *jsiiProxy_SupportappSlackChannelConfiguration)SetNotifyOnCreateOrReopenCase(val interface{}) {
	if err := j.validateSetNotifyOnCreateOrReopenCaseParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"notifyOnCreateOrReopenCase",
		val,
	)
}

func (j *jsiiProxy_SupportappSlackChannelConfiguration)SetNotifyOnResolveCase(val interface{}) {
	if err := j.validateSetNotifyOnResolveCaseParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"notifyOnResolveCase",
		val,
	)
}

func (j *jsiiProxy_SupportappSlackChannelConfiguration)SetProvider(val cdktf.TerraformProvider) {
	_jsii_.Set(
		j,
		"provider",
		val,
	)
}

func (j *jsiiProxy_SupportappSlackChannelConfiguration)SetProvisioners(val *[]interface{}) {
	if err := j.validateSetProvisionersParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"provisioners",
		val,
	)
}

func (j *jsiiProxy_SupportappSlackChannelConfiguration)SetTeamId(val *string) {
	if err := j.validateSetTeamIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"teamId",
		val,
	)
}

// Generates CDKTF code for importing a SupportappSlackChannelConfiguration resource upon running "cdktf plan <stack-name>".
func SupportappSlackChannelConfiguration_GenerateConfigForImport(scope constructs.Construct, importToId *string, importFromId *string, provider cdktf.TerraformProvider) cdktf.ImportableResource {
	_init_.Initialize()

	if err := validateSupportappSlackChannelConfiguration_GenerateConfigForImportParameters(scope, importToId, importFromId); err != nil {
		panic(err)
	}
	var returns cdktf.ImportableResource

	_jsii_.StaticInvoke(
		"awscc.supportappSlackChannelConfiguration.SupportappSlackChannelConfiguration",
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
func SupportappSlackChannelConfiguration_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	if err := validateSupportappSlackChannelConfiguration_IsConstructParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"awscc.supportappSlackChannelConfiguration.SupportappSlackChannelConfiguration",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Experimental.
func SupportappSlackChannelConfiguration_IsTerraformElement(x interface{}) *bool {
	_init_.Initialize()

	if err := validateSupportappSlackChannelConfiguration_IsTerraformElementParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"awscc.supportappSlackChannelConfiguration.SupportappSlackChannelConfiguration",
		"isTerraformElement",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Experimental.
func SupportappSlackChannelConfiguration_IsTerraformResource(x interface{}) *bool {
	_init_.Initialize()

	if err := validateSupportappSlackChannelConfiguration_IsTerraformResourceParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"awscc.supportappSlackChannelConfiguration.SupportappSlackChannelConfiguration",
		"isTerraformResource",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func SupportappSlackChannelConfiguration_TfResourceType() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"awscc.supportappSlackChannelConfiguration.SupportappSlackChannelConfiguration",
		"tfResourceType",
		&returns,
	)
	return returns
}

func (s *jsiiProxy_SupportappSlackChannelConfiguration) AddMoveTarget(moveTarget *string) {
	if err := s.validateAddMoveTargetParameters(moveTarget); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		s,
		"addMoveTarget",
		[]interface{}{moveTarget},
	)
}

func (s *jsiiProxy_SupportappSlackChannelConfiguration) AddOverride(path *string, value interface{}) {
	if err := s.validateAddOverrideParameters(path, value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		s,
		"addOverride",
		[]interface{}{path, value},
	)
}

func (s *jsiiProxy_SupportappSlackChannelConfiguration) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
	if err := s.validateGetAnyMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]interface{}

	_jsii_.Invoke(
		s,
		"getAnyMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (s *jsiiProxy_SupportappSlackChannelConfiguration) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
	if err := s.validateGetBooleanAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		s,
		"getBooleanAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (s *jsiiProxy_SupportappSlackChannelConfiguration) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
	if err := s.validateGetBooleanMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*bool

	_jsii_.Invoke(
		s,
		"getBooleanMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (s *jsiiProxy_SupportappSlackChannelConfiguration) GetListAttribute(terraformAttribute *string) *[]*string {
	if err := s.validateGetListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*string

	_jsii_.Invoke(
		s,
		"getListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (s *jsiiProxy_SupportappSlackChannelConfiguration) GetNumberAttribute(terraformAttribute *string) *float64 {
	if err := s.validateGetNumberAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *float64

	_jsii_.Invoke(
		s,
		"getNumberAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (s *jsiiProxy_SupportappSlackChannelConfiguration) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
	if err := s.validateGetNumberListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*float64

	_jsii_.Invoke(
		s,
		"getNumberListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (s *jsiiProxy_SupportappSlackChannelConfiguration) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
	if err := s.validateGetNumberMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*float64

	_jsii_.Invoke(
		s,
		"getNumberMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (s *jsiiProxy_SupportappSlackChannelConfiguration) GetStringAttribute(terraformAttribute *string) *string {
	if err := s.validateGetStringAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *string

	_jsii_.Invoke(
		s,
		"getStringAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (s *jsiiProxy_SupportappSlackChannelConfiguration) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
	if err := s.validateGetStringMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*string

	_jsii_.Invoke(
		s,
		"getStringMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (s *jsiiProxy_SupportappSlackChannelConfiguration) HasResourceMove() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		s,
		"hasResourceMove",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (s *jsiiProxy_SupportappSlackChannelConfiguration) ImportFrom(id *string, provider cdktf.TerraformProvider) {
	if err := s.validateImportFromParameters(id); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		s,
		"importFrom",
		[]interface{}{id, provider},
	)
}

func (s *jsiiProxy_SupportappSlackChannelConfiguration) InterpolationForAttribute(terraformAttribute *string) cdktf.IResolvable {
	if err := s.validateInterpolationForAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		s,
		"interpolationForAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (s *jsiiProxy_SupportappSlackChannelConfiguration) MoveFromId(id *string) {
	if err := s.validateMoveFromIdParameters(id); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		s,
		"moveFromId",
		[]interface{}{id},
	)
}

func (s *jsiiProxy_SupportappSlackChannelConfiguration) MoveTo(moveTarget *string, index interface{}) {
	if err := s.validateMoveToParameters(moveTarget, index); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		s,
		"moveTo",
		[]interface{}{moveTarget, index},
	)
}

func (s *jsiiProxy_SupportappSlackChannelConfiguration) MoveToId(id *string) {
	if err := s.validateMoveToIdParameters(id); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		s,
		"moveToId",
		[]interface{}{id},
	)
}

func (s *jsiiProxy_SupportappSlackChannelConfiguration) OverrideLogicalId(newLogicalId *string) {
	if err := s.validateOverrideLogicalIdParameters(newLogicalId); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		s,
		"overrideLogicalId",
		[]interface{}{newLogicalId},
	)
}

func (s *jsiiProxy_SupportappSlackChannelConfiguration) ResetChannelName() {
	_jsii_.InvokeVoid(
		s,
		"resetChannelName",
		nil, // no parameters
	)
}

func (s *jsiiProxy_SupportappSlackChannelConfiguration) ResetNotifyOnAddCorrespondenceToCase() {
	_jsii_.InvokeVoid(
		s,
		"resetNotifyOnAddCorrespondenceToCase",
		nil, // no parameters
	)
}

func (s *jsiiProxy_SupportappSlackChannelConfiguration) ResetNotifyOnCreateOrReopenCase() {
	_jsii_.InvokeVoid(
		s,
		"resetNotifyOnCreateOrReopenCase",
		nil, // no parameters
	)
}

func (s *jsiiProxy_SupportappSlackChannelConfiguration) ResetNotifyOnResolveCase() {
	_jsii_.InvokeVoid(
		s,
		"resetNotifyOnResolveCase",
		nil, // no parameters
	)
}

func (s *jsiiProxy_SupportappSlackChannelConfiguration) ResetOverrideLogicalId() {
	_jsii_.InvokeVoid(
		s,
		"resetOverrideLogicalId",
		nil, // no parameters
	)
}

func (s *jsiiProxy_SupportappSlackChannelConfiguration) SynthesizeAttributes() *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		s,
		"synthesizeAttributes",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (s *jsiiProxy_SupportappSlackChannelConfiguration) SynthesizeHclAttributes() *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		s,
		"synthesizeHclAttributes",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (s *jsiiProxy_SupportappSlackChannelConfiguration) ToHclTerraform() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		s,
		"toHclTerraform",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (s *jsiiProxy_SupportappSlackChannelConfiguration) ToMetadata() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		s,
		"toMetadata",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (s *jsiiProxy_SupportappSlackChannelConfiguration) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		s,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (s *jsiiProxy_SupportappSlackChannelConfiguration) ToTerraform() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		s,
		"toTerraform",
		nil, // no parameters
		&returns,
	)

	return returns
}

