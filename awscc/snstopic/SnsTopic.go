package snstopic

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/unimelb-platform/cdktf-provider-awscc-go/awscc/jsii"

	"github.com/aws/constructs-go/constructs/v10"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
	"github.com/unimelb-platform/cdktf-provider-awscc-go/awscc/snstopic/internal"
)

// Represents a {@link https://registry.terraform.io/providers/hashicorp/awscc/1.49.0/docs/resources/sns_topic awscc_sns_topic}.
type SnsTopic interface {
	cdktf.TerraformResource
	ArchivePolicy() *string
	SetArchivePolicy(val *string)
	ArchivePolicyInput() *string
	// Experimental.
	CdktfStack() cdktf.TerraformStack
	// Experimental.
	Connection() interface{}
	// Experimental.
	SetConnection(val interface{})
	// Experimental.
	ConstructNodeMetadata() *map[string]interface{}
	ContentBasedDeduplication() interface{}
	SetContentBasedDeduplication(val interface{})
	ContentBasedDeduplicationInput() interface{}
	// Experimental.
	Count() interface{}
	// Experimental.
	SetCount(val interface{})
	DataProtectionPolicy() *string
	SetDataProtectionPolicy(val *string)
	DataProtectionPolicyInput() *string
	DeliveryStatusLogging() SnsTopicDeliveryStatusLoggingList
	DeliveryStatusLoggingInput() interface{}
	// Experimental.
	DependsOn() *[]*string
	// Experimental.
	SetDependsOn(val *[]*string)
	DisplayName() *string
	SetDisplayName(val *string)
	DisplayNameInput() *string
	FifoThroughputScope() *string
	SetFifoThroughputScope(val *string)
	FifoThroughputScopeInput() *string
	FifoTopic() interface{}
	SetFifoTopic(val interface{})
	FifoTopicInput() interface{}
	// Experimental.
	ForEach() cdktf.ITerraformIterator
	// Experimental.
	SetForEach(val cdktf.ITerraformIterator)
	// Experimental.
	Fqn() *string
	// Experimental.
	FriendlyUniqueId() *string
	Id() *string
	KmsMasterKeyId() *string
	SetKmsMasterKeyId(val *string)
	KmsMasterKeyIdInput() *string
	// Experimental.
	Lifecycle() *cdktf.TerraformResourceLifecycle
	// Experimental.
	SetLifecycle(val *cdktf.TerraformResourceLifecycle)
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
	SignatureVersion() *string
	SetSignatureVersion(val *string)
	SignatureVersionInput() *string
	Subscription() SnsTopicSubscriptionList
	SubscriptionInput() interface{}
	Tags() SnsTopicTagsList
	TagsInput() interface{}
	// Experimental.
	TerraformGeneratorMetadata() *cdktf.TerraformProviderGeneratorMetadata
	// Experimental.
	TerraformMetaArguments() *map[string]interface{}
	// Experimental.
	TerraformResourceType() *string
	TopicArn() *string
	TopicName() *string
	SetTopicName(val *string)
	TopicNameInput() *string
	TracingConfig() *string
	SetTracingConfig(val *string)
	TracingConfigInput() *string
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
	PutDeliveryStatusLogging(value interface{})
	PutSubscription(value interface{})
	PutTags(value interface{})
	ResetArchivePolicy()
	ResetContentBasedDeduplication()
	ResetDataProtectionPolicy()
	ResetDeliveryStatusLogging()
	ResetDisplayName()
	ResetFifoThroughputScope()
	ResetFifoTopic()
	ResetKmsMasterKeyId()
	// Resets a previously passed logical Id to use the auto-generated logical id again.
	// Experimental.
	ResetOverrideLogicalId()
	ResetSignatureVersion()
	ResetSubscription()
	ResetTags()
	ResetTopicName()
	ResetTracingConfig()
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

// The jsii proxy struct for SnsTopic
type jsiiProxy_SnsTopic struct {
	internal.Type__cdktfTerraformResource
}

func (j *jsiiProxy_SnsTopic) ArchivePolicy() *string {
	var returns *string
	_jsii_.Get(
		j,
		"archivePolicy",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SnsTopic) ArchivePolicyInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"archivePolicyInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SnsTopic) CdktfStack() cdktf.TerraformStack {
	var returns cdktf.TerraformStack
	_jsii_.Get(
		j,
		"cdktfStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SnsTopic) Connection() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"connection",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SnsTopic) ConstructNodeMetadata() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"constructNodeMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SnsTopic) ContentBasedDeduplication() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"contentBasedDeduplication",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SnsTopic) ContentBasedDeduplicationInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"contentBasedDeduplicationInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SnsTopic) Count() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"count",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SnsTopic) DataProtectionPolicy() *string {
	var returns *string
	_jsii_.Get(
		j,
		"dataProtectionPolicy",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SnsTopic) DataProtectionPolicyInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"dataProtectionPolicyInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SnsTopic) DeliveryStatusLogging() SnsTopicDeliveryStatusLoggingList {
	var returns SnsTopicDeliveryStatusLoggingList
	_jsii_.Get(
		j,
		"deliveryStatusLogging",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SnsTopic) DeliveryStatusLoggingInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"deliveryStatusLoggingInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SnsTopic) DependsOn() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"dependsOn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SnsTopic) DisplayName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"displayName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SnsTopic) DisplayNameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"displayNameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SnsTopic) FifoThroughputScope() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fifoThroughputScope",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SnsTopic) FifoThroughputScopeInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fifoThroughputScopeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SnsTopic) FifoTopic() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"fifoTopic",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SnsTopic) FifoTopicInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"fifoTopicInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SnsTopic) ForEach() cdktf.ITerraformIterator {
	var returns cdktf.ITerraformIterator
	_jsii_.Get(
		j,
		"forEach",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SnsTopic) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SnsTopic) FriendlyUniqueId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"friendlyUniqueId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SnsTopic) Id() *string {
	var returns *string
	_jsii_.Get(
		j,
		"id",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SnsTopic) KmsMasterKeyId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"kmsMasterKeyId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SnsTopic) KmsMasterKeyIdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"kmsMasterKeyIdInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SnsTopic) Lifecycle() *cdktf.TerraformResourceLifecycle {
	var returns *cdktf.TerraformResourceLifecycle
	_jsii_.Get(
		j,
		"lifecycle",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SnsTopic) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SnsTopic) Provider() cdktf.TerraformProvider {
	var returns cdktf.TerraformProvider
	_jsii_.Get(
		j,
		"provider",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SnsTopic) Provisioners() *[]interface{} {
	var returns *[]interface{}
	_jsii_.Get(
		j,
		"provisioners",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SnsTopic) RawOverrides() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"rawOverrides",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SnsTopic) SignatureVersion() *string {
	var returns *string
	_jsii_.Get(
		j,
		"signatureVersion",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SnsTopic) SignatureVersionInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"signatureVersionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SnsTopic) Subscription() SnsTopicSubscriptionList {
	var returns SnsTopicSubscriptionList
	_jsii_.Get(
		j,
		"subscription",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SnsTopic) SubscriptionInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"subscriptionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SnsTopic) Tags() SnsTopicTagsList {
	var returns SnsTopicTagsList
	_jsii_.Get(
		j,
		"tags",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SnsTopic) TagsInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"tagsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SnsTopic) TerraformGeneratorMetadata() *cdktf.TerraformProviderGeneratorMetadata {
	var returns *cdktf.TerraformProviderGeneratorMetadata
	_jsii_.Get(
		j,
		"terraformGeneratorMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SnsTopic) TerraformMetaArguments() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"terraformMetaArguments",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SnsTopic) TerraformResourceType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformResourceType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SnsTopic) TopicArn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"topicArn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SnsTopic) TopicName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"topicName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SnsTopic) TopicNameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"topicNameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SnsTopic) TracingConfig() *string {
	var returns *string
	_jsii_.Get(
		j,
		"tracingConfig",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SnsTopic) TracingConfigInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"tracingConfigInput",
		&returns,
	)
	return returns
}


// Create a new {@link https://registry.terraform.io/providers/hashicorp/awscc/1.49.0/docs/resources/sns_topic awscc_sns_topic} Resource.
func NewSnsTopic(scope constructs.Construct, id *string, config *SnsTopicConfig) SnsTopic {
	_init_.Initialize()

	if err := validateNewSnsTopicParameters(scope, id, config); err != nil {
		panic(err)
	}
	j := jsiiProxy_SnsTopic{}

	_jsii_.Create(
		"awscc.snsTopic.SnsTopic",
		[]interface{}{scope, id, config},
		&j,
	)

	return &j
}

// Create a new {@link https://registry.terraform.io/providers/hashicorp/awscc/1.49.0/docs/resources/sns_topic awscc_sns_topic} Resource.
func NewSnsTopic_Override(s SnsTopic, scope constructs.Construct, id *string, config *SnsTopicConfig) {
	_init_.Initialize()

	_jsii_.Create(
		"awscc.snsTopic.SnsTopic",
		[]interface{}{scope, id, config},
		s,
	)
}

func (j *jsiiProxy_SnsTopic)SetArchivePolicy(val *string) {
	if err := j.validateSetArchivePolicyParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"archivePolicy",
		val,
	)
}

func (j *jsiiProxy_SnsTopic)SetConnection(val interface{}) {
	if err := j.validateSetConnectionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"connection",
		val,
	)
}

func (j *jsiiProxy_SnsTopic)SetContentBasedDeduplication(val interface{}) {
	if err := j.validateSetContentBasedDeduplicationParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"contentBasedDeduplication",
		val,
	)
}

func (j *jsiiProxy_SnsTopic)SetCount(val interface{}) {
	if err := j.validateSetCountParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"count",
		val,
	)
}

func (j *jsiiProxy_SnsTopic)SetDataProtectionPolicy(val *string) {
	if err := j.validateSetDataProtectionPolicyParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"dataProtectionPolicy",
		val,
	)
}

func (j *jsiiProxy_SnsTopic)SetDependsOn(val *[]*string) {
	_jsii_.Set(
		j,
		"dependsOn",
		val,
	)
}

func (j *jsiiProxy_SnsTopic)SetDisplayName(val *string) {
	if err := j.validateSetDisplayNameParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"displayName",
		val,
	)
}

func (j *jsiiProxy_SnsTopic)SetFifoThroughputScope(val *string) {
	if err := j.validateSetFifoThroughputScopeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"fifoThroughputScope",
		val,
	)
}

func (j *jsiiProxy_SnsTopic)SetFifoTopic(val interface{}) {
	if err := j.validateSetFifoTopicParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"fifoTopic",
		val,
	)
}

func (j *jsiiProxy_SnsTopic)SetForEach(val cdktf.ITerraformIterator) {
	_jsii_.Set(
		j,
		"forEach",
		val,
	)
}

func (j *jsiiProxy_SnsTopic)SetKmsMasterKeyId(val *string) {
	if err := j.validateSetKmsMasterKeyIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"kmsMasterKeyId",
		val,
	)
}

func (j *jsiiProxy_SnsTopic)SetLifecycle(val *cdktf.TerraformResourceLifecycle) {
	if err := j.validateSetLifecycleParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"lifecycle",
		val,
	)
}

func (j *jsiiProxy_SnsTopic)SetProvider(val cdktf.TerraformProvider) {
	_jsii_.Set(
		j,
		"provider",
		val,
	)
}

func (j *jsiiProxy_SnsTopic)SetProvisioners(val *[]interface{}) {
	if err := j.validateSetProvisionersParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"provisioners",
		val,
	)
}

func (j *jsiiProxy_SnsTopic)SetSignatureVersion(val *string) {
	if err := j.validateSetSignatureVersionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"signatureVersion",
		val,
	)
}

func (j *jsiiProxy_SnsTopic)SetTopicName(val *string) {
	if err := j.validateSetTopicNameParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"topicName",
		val,
	)
}

func (j *jsiiProxy_SnsTopic)SetTracingConfig(val *string) {
	if err := j.validateSetTracingConfigParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"tracingConfig",
		val,
	)
}

// Generates CDKTF code for importing a SnsTopic resource upon running "cdktf plan <stack-name>".
func SnsTopic_GenerateConfigForImport(scope constructs.Construct, importToId *string, importFromId *string, provider cdktf.TerraformProvider) cdktf.ImportableResource {
	_init_.Initialize()

	if err := validateSnsTopic_GenerateConfigForImportParameters(scope, importToId, importFromId); err != nil {
		panic(err)
	}
	var returns cdktf.ImportableResource

	_jsii_.StaticInvoke(
		"awscc.snsTopic.SnsTopic",
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
func SnsTopic_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	if err := validateSnsTopic_IsConstructParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"awscc.snsTopic.SnsTopic",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Experimental.
func SnsTopic_IsTerraformElement(x interface{}) *bool {
	_init_.Initialize()

	if err := validateSnsTopic_IsTerraformElementParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"awscc.snsTopic.SnsTopic",
		"isTerraformElement",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Experimental.
func SnsTopic_IsTerraformResource(x interface{}) *bool {
	_init_.Initialize()

	if err := validateSnsTopic_IsTerraformResourceParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"awscc.snsTopic.SnsTopic",
		"isTerraformResource",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func SnsTopic_TfResourceType() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"awscc.snsTopic.SnsTopic",
		"tfResourceType",
		&returns,
	)
	return returns
}

func (s *jsiiProxy_SnsTopic) AddMoveTarget(moveTarget *string) {
	if err := s.validateAddMoveTargetParameters(moveTarget); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		s,
		"addMoveTarget",
		[]interface{}{moveTarget},
	)
}

func (s *jsiiProxy_SnsTopic) AddOverride(path *string, value interface{}) {
	if err := s.validateAddOverrideParameters(path, value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		s,
		"addOverride",
		[]interface{}{path, value},
	)
}

func (s *jsiiProxy_SnsTopic) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (s *jsiiProxy_SnsTopic) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
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

func (s *jsiiProxy_SnsTopic) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (s *jsiiProxy_SnsTopic) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (s *jsiiProxy_SnsTopic) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (s *jsiiProxy_SnsTopic) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (s *jsiiProxy_SnsTopic) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (s *jsiiProxy_SnsTopic) GetStringAttribute(terraformAttribute *string) *string {
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

func (s *jsiiProxy_SnsTopic) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (s *jsiiProxy_SnsTopic) HasResourceMove() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		s,
		"hasResourceMove",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (s *jsiiProxy_SnsTopic) ImportFrom(id *string, provider cdktf.TerraformProvider) {
	if err := s.validateImportFromParameters(id); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		s,
		"importFrom",
		[]interface{}{id, provider},
	)
}

func (s *jsiiProxy_SnsTopic) InterpolationForAttribute(terraformAttribute *string) cdktf.IResolvable {
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

func (s *jsiiProxy_SnsTopic) MoveFromId(id *string) {
	if err := s.validateMoveFromIdParameters(id); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		s,
		"moveFromId",
		[]interface{}{id},
	)
}

func (s *jsiiProxy_SnsTopic) MoveTo(moveTarget *string, index interface{}) {
	if err := s.validateMoveToParameters(moveTarget, index); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		s,
		"moveTo",
		[]interface{}{moveTarget, index},
	)
}

func (s *jsiiProxy_SnsTopic) MoveToId(id *string) {
	if err := s.validateMoveToIdParameters(id); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		s,
		"moveToId",
		[]interface{}{id},
	)
}

func (s *jsiiProxy_SnsTopic) OverrideLogicalId(newLogicalId *string) {
	if err := s.validateOverrideLogicalIdParameters(newLogicalId); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		s,
		"overrideLogicalId",
		[]interface{}{newLogicalId},
	)
}

func (s *jsiiProxy_SnsTopic) PutDeliveryStatusLogging(value interface{}) {
	if err := s.validatePutDeliveryStatusLoggingParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		s,
		"putDeliveryStatusLogging",
		[]interface{}{value},
	)
}

func (s *jsiiProxy_SnsTopic) PutSubscription(value interface{}) {
	if err := s.validatePutSubscriptionParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		s,
		"putSubscription",
		[]interface{}{value},
	)
}

func (s *jsiiProxy_SnsTopic) PutTags(value interface{}) {
	if err := s.validatePutTagsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		s,
		"putTags",
		[]interface{}{value},
	)
}

func (s *jsiiProxy_SnsTopic) ResetArchivePolicy() {
	_jsii_.InvokeVoid(
		s,
		"resetArchivePolicy",
		nil, // no parameters
	)
}

func (s *jsiiProxy_SnsTopic) ResetContentBasedDeduplication() {
	_jsii_.InvokeVoid(
		s,
		"resetContentBasedDeduplication",
		nil, // no parameters
	)
}

func (s *jsiiProxy_SnsTopic) ResetDataProtectionPolicy() {
	_jsii_.InvokeVoid(
		s,
		"resetDataProtectionPolicy",
		nil, // no parameters
	)
}

func (s *jsiiProxy_SnsTopic) ResetDeliveryStatusLogging() {
	_jsii_.InvokeVoid(
		s,
		"resetDeliveryStatusLogging",
		nil, // no parameters
	)
}

func (s *jsiiProxy_SnsTopic) ResetDisplayName() {
	_jsii_.InvokeVoid(
		s,
		"resetDisplayName",
		nil, // no parameters
	)
}

func (s *jsiiProxy_SnsTopic) ResetFifoThroughputScope() {
	_jsii_.InvokeVoid(
		s,
		"resetFifoThroughputScope",
		nil, // no parameters
	)
}

func (s *jsiiProxy_SnsTopic) ResetFifoTopic() {
	_jsii_.InvokeVoid(
		s,
		"resetFifoTopic",
		nil, // no parameters
	)
}

func (s *jsiiProxy_SnsTopic) ResetKmsMasterKeyId() {
	_jsii_.InvokeVoid(
		s,
		"resetKmsMasterKeyId",
		nil, // no parameters
	)
}

func (s *jsiiProxy_SnsTopic) ResetOverrideLogicalId() {
	_jsii_.InvokeVoid(
		s,
		"resetOverrideLogicalId",
		nil, // no parameters
	)
}

func (s *jsiiProxy_SnsTopic) ResetSignatureVersion() {
	_jsii_.InvokeVoid(
		s,
		"resetSignatureVersion",
		nil, // no parameters
	)
}

func (s *jsiiProxy_SnsTopic) ResetSubscription() {
	_jsii_.InvokeVoid(
		s,
		"resetSubscription",
		nil, // no parameters
	)
}

func (s *jsiiProxy_SnsTopic) ResetTags() {
	_jsii_.InvokeVoid(
		s,
		"resetTags",
		nil, // no parameters
	)
}

func (s *jsiiProxy_SnsTopic) ResetTopicName() {
	_jsii_.InvokeVoid(
		s,
		"resetTopicName",
		nil, // no parameters
	)
}

func (s *jsiiProxy_SnsTopic) ResetTracingConfig() {
	_jsii_.InvokeVoid(
		s,
		"resetTracingConfig",
		nil, // no parameters
	)
}

func (s *jsiiProxy_SnsTopic) SynthesizeAttributes() *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		s,
		"synthesizeAttributes",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (s *jsiiProxy_SnsTopic) SynthesizeHclAttributes() *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		s,
		"synthesizeHclAttributes",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (s *jsiiProxy_SnsTopic) ToHclTerraform() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		s,
		"toHclTerraform",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (s *jsiiProxy_SnsTopic) ToMetadata() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		s,
		"toMetadata",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (s *jsiiProxy_SnsTopic) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		s,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (s *jsiiProxy_SnsTopic) ToTerraform() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		s,
		"toTerraform",
		nil, // no parameters
		&returns,
	)

	return returns
}

