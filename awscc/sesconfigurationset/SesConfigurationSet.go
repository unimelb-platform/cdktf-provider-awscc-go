package sesconfigurationset

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/unimelb-platform/cdktf-provider-awscc-go/awscc/jsii"

	"github.com/aws/constructs-go/constructs/v10"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
	"github.com/unimelb-platform/cdktf-provider-awscc-go/awscc/sesconfigurationset/internal"
)

// Represents a {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/ses_configuration_set awscc_ses_configuration_set}.
type SesConfigurationSet interface {
	cdktf.TerraformResource
	// Experimental.
	CdktfStack() cdktf.TerraformStack
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
	DeliveryOptions() SesConfigurationSetDeliveryOptionsOutputReference
	DeliveryOptionsInput() interface{}
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
	ReputationOptions() SesConfigurationSetReputationOptionsOutputReference
	ReputationOptionsInput() interface{}
	SendingOptions() SesConfigurationSetSendingOptionsOutputReference
	SendingOptionsInput() interface{}
	SuppressionOptions() SesConfigurationSetSuppressionOptionsOutputReference
	SuppressionOptionsInput() interface{}
	// Experimental.
	TerraformGeneratorMetadata() *cdktf.TerraformProviderGeneratorMetadata
	// Experimental.
	TerraformMetaArguments() *map[string]interface{}
	// Experimental.
	TerraformResourceType() *string
	TrackingOptions() SesConfigurationSetTrackingOptionsOutputReference
	TrackingOptionsInput() interface{}
	VdmOptions() SesConfigurationSetVdmOptionsOutputReference
	VdmOptionsInput() interface{}
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
	PutDeliveryOptions(value *SesConfigurationSetDeliveryOptions)
	PutReputationOptions(value *SesConfigurationSetReputationOptions)
	PutSendingOptions(value *SesConfigurationSetSendingOptions)
	PutSuppressionOptions(value *SesConfigurationSetSuppressionOptions)
	PutTrackingOptions(value *SesConfigurationSetTrackingOptions)
	PutVdmOptions(value *SesConfigurationSetVdmOptions)
	ResetDeliveryOptions()
	ResetName()
	// Resets a previously passed logical Id to use the auto-generated logical id again.
	// Experimental.
	ResetOverrideLogicalId()
	ResetReputationOptions()
	ResetSendingOptions()
	ResetSuppressionOptions()
	ResetTrackingOptions()
	ResetVdmOptions()
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

// The jsii proxy struct for SesConfigurationSet
type jsiiProxy_SesConfigurationSet struct {
	internal.Type__cdktfTerraformResource
}

func (j *jsiiProxy_SesConfigurationSet) CdktfStack() cdktf.TerraformStack {
	var returns cdktf.TerraformStack
	_jsii_.Get(
		j,
		"cdktfStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SesConfigurationSet) Connection() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"connection",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SesConfigurationSet) ConstructNodeMetadata() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"constructNodeMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SesConfigurationSet) Count() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"count",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SesConfigurationSet) DeliveryOptions() SesConfigurationSetDeliveryOptionsOutputReference {
	var returns SesConfigurationSetDeliveryOptionsOutputReference
	_jsii_.Get(
		j,
		"deliveryOptions",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SesConfigurationSet) DeliveryOptionsInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"deliveryOptionsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SesConfigurationSet) DependsOn() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"dependsOn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SesConfigurationSet) ForEach() cdktf.ITerraformIterator {
	var returns cdktf.ITerraformIterator
	_jsii_.Get(
		j,
		"forEach",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SesConfigurationSet) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SesConfigurationSet) FriendlyUniqueId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"friendlyUniqueId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SesConfigurationSet) Id() *string {
	var returns *string
	_jsii_.Get(
		j,
		"id",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SesConfigurationSet) Lifecycle() *cdktf.TerraformResourceLifecycle {
	var returns *cdktf.TerraformResourceLifecycle
	_jsii_.Get(
		j,
		"lifecycle",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SesConfigurationSet) Name() *string {
	var returns *string
	_jsii_.Get(
		j,
		"name",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SesConfigurationSet) NameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"nameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SesConfigurationSet) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SesConfigurationSet) Provider() cdktf.TerraformProvider {
	var returns cdktf.TerraformProvider
	_jsii_.Get(
		j,
		"provider",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SesConfigurationSet) Provisioners() *[]interface{} {
	var returns *[]interface{}
	_jsii_.Get(
		j,
		"provisioners",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SesConfigurationSet) RawOverrides() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"rawOverrides",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SesConfigurationSet) ReputationOptions() SesConfigurationSetReputationOptionsOutputReference {
	var returns SesConfigurationSetReputationOptionsOutputReference
	_jsii_.Get(
		j,
		"reputationOptions",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SesConfigurationSet) ReputationOptionsInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"reputationOptionsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SesConfigurationSet) SendingOptions() SesConfigurationSetSendingOptionsOutputReference {
	var returns SesConfigurationSetSendingOptionsOutputReference
	_jsii_.Get(
		j,
		"sendingOptions",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SesConfigurationSet) SendingOptionsInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"sendingOptionsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SesConfigurationSet) SuppressionOptions() SesConfigurationSetSuppressionOptionsOutputReference {
	var returns SesConfigurationSetSuppressionOptionsOutputReference
	_jsii_.Get(
		j,
		"suppressionOptions",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SesConfigurationSet) SuppressionOptionsInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"suppressionOptionsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SesConfigurationSet) TerraformGeneratorMetadata() *cdktf.TerraformProviderGeneratorMetadata {
	var returns *cdktf.TerraformProviderGeneratorMetadata
	_jsii_.Get(
		j,
		"terraformGeneratorMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SesConfigurationSet) TerraformMetaArguments() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"terraformMetaArguments",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SesConfigurationSet) TerraformResourceType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformResourceType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SesConfigurationSet) TrackingOptions() SesConfigurationSetTrackingOptionsOutputReference {
	var returns SesConfigurationSetTrackingOptionsOutputReference
	_jsii_.Get(
		j,
		"trackingOptions",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SesConfigurationSet) TrackingOptionsInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"trackingOptionsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SesConfigurationSet) VdmOptions() SesConfigurationSetVdmOptionsOutputReference {
	var returns SesConfigurationSetVdmOptionsOutputReference
	_jsii_.Get(
		j,
		"vdmOptions",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SesConfigurationSet) VdmOptionsInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"vdmOptionsInput",
		&returns,
	)
	return returns
}


// Create a new {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/ses_configuration_set awscc_ses_configuration_set} Resource.
func NewSesConfigurationSet(scope constructs.Construct, id *string, config *SesConfigurationSetConfig) SesConfigurationSet {
	_init_.Initialize()

	if err := validateNewSesConfigurationSetParameters(scope, id, config); err != nil {
		panic(err)
	}
	j := jsiiProxy_SesConfigurationSet{}

	_jsii_.Create(
		"awscc.sesConfigurationSet.SesConfigurationSet",
		[]interface{}{scope, id, config},
		&j,
	)

	return &j
}

// Create a new {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/ses_configuration_set awscc_ses_configuration_set} Resource.
func NewSesConfigurationSet_Override(s SesConfigurationSet, scope constructs.Construct, id *string, config *SesConfigurationSetConfig) {
	_init_.Initialize()

	_jsii_.Create(
		"awscc.sesConfigurationSet.SesConfigurationSet",
		[]interface{}{scope, id, config},
		s,
	)
}

func (j *jsiiProxy_SesConfigurationSet)SetConnection(val interface{}) {
	if err := j.validateSetConnectionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"connection",
		val,
	)
}

func (j *jsiiProxy_SesConfigurationSet)SetCount(val interface{}) {
	if err := j.validateSetCountParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"count",
		val,
	)
}

func (j *jsiiProxy_SesConfigurationSet)SetDependsOn(val *[]*string) {
	_jsii_.Set(
		j,
		"dependsOn",
		val,
	)
}

func (j *jsiiProxy_SesConfigurationSet)SetForEach(val cdktf.ITerraformIterator) {
	_jsii_.Set(
		j,
		"forEach",
		val,
	)
}

func (j *jsiiProxy_SesConfigurationSet)SetLifecycle(val *cdktf.TerraformResourceLifecycle) {
	if err := j.validateSetLifecycleParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"lifecycle",
		val,
	)
}

func (j *jsiiProxy_SesConfigurationSet)SetName(val *string) {
	if err := j.validateSetNameParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"name",
		val,
	)
}

func (j *jsiiProxy_SesConfigurationSet)SetProvider(val cdktf.TerraformProvider) {
	_jsii_.Set(
		j,
		"provider",
		val,
	)
}

func (j *jsiiProxy_SesConfigurationSet)SetProvisioners(val *[]interface{}) {
	if err := j.validateSetProvisionersParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"provisioners",
		val,
	)
}

// Generates CDKTF code for importing a SesConfigurationSet resource upon running "cdktf plan <stack-name>".
func SesConfigurationSet_GenerateConfigForImport(scope constructs.Construct, importToId *string, importFromId *string, provider cdktf.TerraformProvider) cdktf.ImportableResource {
	_init_.Initialize()

	if err := validateSesConfigurationSet_GenerateConfigForImportParameters(scope, importToId, importFromId); err != nil {
		panic(err)
	}
	var returns cdktf.ImportableResource

	_jsii_.StaticInvoke(
		"awscc.sesConfigurationSet.SesConfigurationSet",
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
func SesConfigurationSet_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	if err := validateSesConfigurationSet_IsConstructParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"awscc.sesConfigurationSet.SesConfigurationSet",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Experimental.
func SesConfigurationSet_IsTerraformElement(x interface{}) *bool {
	_init_.Initialize()

	if err := validateSesConfigurationSet_IsTerraformElementParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"awscc.sesConfigurationSet.SesConfigurationSet",
		"isTerraformElement",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Experimental.
func SesConfigurationSet_IsTerraformResource(x interface{}) *bool {
	_init_.Initialize()

	if err := validateSesConfigurationSet_IsTerraformResourceParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"awscc.sesConfigurationSet.SesConfigurationSet",
		"isTerraformResource",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func SesConfigurationSet_TfResourceType() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"awscc.sesConfigurationSet.SesConfigurationSet",
		"tfResourceType",
		&returns,
	)
	return returns
}

func (s *jsiiProxy_SesConfigurationSet) AddMoveTarget(moveTarget *string) {
	if err := s.validateAddMoveTargetParameters(moveTarget); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		s,
		"addMoveTarget",
		[]interface{}{moveTarget},
	)
}

func (s *jsiiProxy_SesConfigurationSet) AddOverride(path *string, value interface{}) {
	if err := s.validateAddOverrideParameters(path, value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		s,
		"addOverride",
		[]interface{}{path, value},
	)
}

func (s *jsiiProxy_SesConfigurationSet) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (s *jsiiProxy_SesConfigurationSet) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
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

func (s *jsiiProxy_SesConfigurationSet) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (s *jsiiProxy_SesConfigurationSet) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (s *jsiiProxy_SesConfigurationSet) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (s *jsiiProxy_SesConfigurationSet) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (s *jsiiProxy_SesConfigurationSet) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (s *jsiiProxy_SesConfigurationSet) GetStringAttribute(terraformAttribute *string) *string {
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

func (s *jsiiProxy_SesConfigurationSet) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (s *jsiiProxy_SesConfigurationSet) HasResourceMove() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		s,
		"hasResourceMove",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (s *jsiiProxy_SesConfigurationSet) ImportFrom(id *string, provider cdktf.TerraformProvider) {
	if err := s.validateImportFromParameters(id); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		s,
		"importFrom",
		[]interface{}{id, provider},
	)
}

func (s *jsiiProxy_SesConfigurationSet) InterpolationForAttribute(terraformAttribute *string) cdktf.IResolvable {
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

func (s *jsiiProxy_SesConfigurationSet) MoveFromId(id *string) {
	if err := s.validateMoveFromIdParameters(id); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		s,
		"moveFromId",
		[]interface{}{id},
	)
}

func (s *jsiiProxy_SesConfigurationSet) MoveTo(moveTarget *string, index interface{}) {
	if err := s.validateMoveToParameters(moveTarget, index); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		s,
		"moveTo",
		[]interface{}{moveTarget, index},
	)
}

func (s *jsiiProxy_SesConfigurationSet) MoveToId(id *string) {
	if err := s.validateMoveToIdParameters(id); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		s,
		"moveToId",
		[]interface{}{id},
	)
}

func (s *jsiiProxy_SesConfigurationSet) OverrideLogicalId(newLogicalId *string) {
	if err := s.validateOverrideLogicalIdParameters(newLogicalId); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		s,
		"overrideLogicalId",
		[]interface{}{newLogicalId},
	)
}

func (s *jsiiProxy_SesConfigurationSet) PutDeliveryOptions(value *SesConfigurationSetDeliveryOptions) {
	if err := s.validatePutDeliveryOptionsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		s,
		"putDeliveryOptions",
		[]interface{}{value},
	)
}

func (s *jsiiProxy_SesConfigurationSet) PutReputationOptions(value *SesConfigurationSetReputationOptions) {
	if err := s.validatePutReputationOptionsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		s,
		"putReputationOptions",
		[]interface{}{value},
	)
}

func (s *jsiiProxy_SesConfigurationSet) PutSendingOptions(value *SesConfigurationSetSendingOptions) {
	if err := s.validatePutSendingOptionsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		s,
		"putSendingOptions",
		[]interface{}{value},
	)
}

func (s *jsiiProxy_SesConfigurationSet) PutSuppressionOptions(value *SesConfigurationSetSuppressionOptions) {
	if err := s.validatePutSuppressionOptionsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		s,
		"putSuppressionOptions",
		[]interface{}{value},
	)
}

func (s *jsiiProxy_SesConfigurationSet) PutTrackingOptions(value *SesConfigurationSetTrackingOptions) {
	if err := s.validatePutTrackingOptionsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		s,
		"putTrackingOptions",
		[]interface{}{value},
	)
}

func (s *jsiiProxy_SesConfigurationSet) PutVdmOptions(value *SesConfigurationSetVdmOptions) {
	if err := s.validatePutVdmOptionsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		s,
		"putVdmOptions",
		[]interface{}{value},
	)
}

func (s *jsiiProxy_SesConfigurationSet) ResetDeliveryOptions() {
	_jsii_.InvokeVoid(
		s,
		"resetDeliveryOptions",
		nil, // no parameters
	)
}

func (s *jsiiProxy_SesConfigurationSet) ResetName() {
	_jsii_.InvokeVoid(
		s,
		"resetName",
		nil, // no parameters
	)
}

func (s *jsiiProxy_SesConfigurationSet) ResetOverrideLogicalId() {
	_jsii_.InvokeVoid(
		s,
		"resetOverrideLogicalId",
		nil, // no parameters
	)
}

func (s *jsiiProxy_SesConfigurationSet) ResetReputationOptions() {
	_jsii_.InvokeVoid(
		s,
		"resetReputationOptions",
		nil, // no parameters
	)
}

func (s *jsiiProxy_SesConfigurationSet) ResetSendingOptions() {
	_jsii_.InvokeVoid(
		s,
		"resetSendingOptions",
		nil, // no parameters
	)
}

func (s *jsiiProxy_SesConfigurationSet) ResetSuppressionOptions() {
	_jsii_.InvokeVoid(
		s,
		"resetSuppressionOptions",
		nil, // no parameters
	)
}

func (s *jsiiProxy_SesConfigurationSet) ResetTrackingOptions() {
	_jsii_.InvokeVoid(
		s,
		"resetTrackingOptions",
		nil, // no parameters
	)
}

func (s *jsiiProxy_SesConfigurationSet) ResetVdmOptions() {
	_jsii_.InvokeVoid(
		s,
		"resetVdmOptions",
		nil, // no parameters
	)
}

func (s *jsiiProxy_SesConfigurationSet) SynthesizeAttributes() *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		s,
		"synthesizeAttributes",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (s *jsiiProxy_SesConfigurationSet) SynthesizeHclAttributes() *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		s,
		"synthesizeHclAttributes",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (s *jsiiProxy_SesConfigurationSet) ToHclTerraform() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		s,
		"toHclTerraform",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (s *jsiiProxy_SesConfigurationSet) ToMetadata() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		s,
		"toMetadata",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (s *jsiiProxy_SesConfigurationSet) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		s,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (s *jsiiProxy_SesConfigurationSet) ToTerraform() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		s,
		"toTerraform",
		nil, // no parameters
		&returns,
	)

	return returns
}

