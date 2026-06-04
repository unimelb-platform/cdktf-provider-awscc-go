package route53recordset

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/unimelb-platform/cdktf-provider-awscc-go/awscc/jsii"

	"github.com/aws/constructs-go/constructs/v10"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
	"github.com/unimelb-platform/cdktf-provider-awscc-go/awscc/route53recordset/internal"
)

// Represents a {@link https://registry.terraform.io/providers/hashicorp/awscc/1.49.0/docs/resources/route53_record_set awscc_route53_record_set}.
type Route53RecordSet interface {
	cdktf.TerraformResource
	AliasTarget() Route53RecordSetAliasTargetOutputReference
	AliasTargetInput() interface{}
	// Experimental.
	CdktfStack() cdktf.TerraformStack
	CidrRoutingConfig() Route53RecordSetCidrRoutingConfigOutputReference
	CidrRoutingConfigInput() interface{}
	Comment() *string
	SetComment(val *string)
	CommentInput() *string
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
	Failover() *string
	SetFailover(val *string)
	FailoverInput() *string
	// Experimental.
	ForEach() cdktf.ITerraformIterator
	// Experimental.
	SetForEach(val cdktf.ITerraformIterator)
	// Experimental.
	Fqn() *string
	// Experimental.
	FriendlyUniqueId() *string
	GeoLocation() Route53RecordSetGeoLocationOutputReference
	GeoLocationInput() interface{}
	GeoProximityLocation() Route53RecordSetGeoProximityLocationOutputReference
	GeoProximityLocationInput() interface{}
	HealthCheckId() *string
	SetHealthCheckId(val *string)
	HealthCheckIdInput() *string
	HostedZoneId() *string
	SetHostedZoneId(val *string)
	HostedZoneIdInput() *string
	HostedZoneName() *string
	SetHostedZoneName(val *string)
	HostedZoneNameInput() *string
	Id() *string
	// Experimental.
	Lifecycle() *cdktf.TerraformResourceLifecycle
	// Experimental.
	SetLifecycle(val *cdktf.TerraformResourceLifecycle)
	MultiValueAnswer() interface{}
	SetMultiValueAnswer(val interface{})
	MultiValueAnswerInput() interface{}
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
	RecordSetId() *string
	Region() *string
	SetRegion(val *string)
	RegionInput() *string
	ResourceRecords() *[]*string
	SetResourceRecords(val *[]*string)
	ResourceRecordsInput() *[]*string
	SetIdentifier() *string
	SetSetIdentifier(val *string)
	SetIdentifierInput() *string
	// Experimental.
	TerraformGeneratorMetadata() *cdktf.TerraformProviderGeneratorMetadata
	// Experimental.
	TerraformMetaArguments() *map[string]interface{}
	// Experimental.
	TerraformResourceType() *string
	Ttl() *string
	SetTtl(val *string)
	TtlInput() *string
	Type() *string
	SetType(val *string)
	TypeInput() *string
	Weight() *float64
	SetWeight(val *float64)
	WeightInput() *float64
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
	PutAliasTarget(value *Route53RecordSetAliasTarget)
	PutCidrRoutingConfig(value *Route53RecordSetCidrRoutingConfig)
	PutGeoLocation(value *Route53RecordSetGeoLocation)
	PutGeoProximityLocation(value *Route53RecordSetGeoProximityLocation)
	ResetAliasTarget()
	ResetCidrRoutingConfig()
	ResetComment()
	ResetFailover()
	ResetGeoLocation()
	ResetGeoProximityLocation()
	ResetHealthCheckId()
	ResetHostedZoneId()
	ResetHostedZoneName()
	ResetMultiValueAnswer()
	// Resets a previously passed logical Id to use the auto-generated logical id again.
	// Experimental.
	ResetOverrideLogicalId()
	ResetRegion()
	ResetResourceRecords()
	ResetSetIdentifier()
	ResetTtl()
	ResetWeight()
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

// The jsii proxy struct for Route53RecordSet
type jsiiProxy_Route53RecordSet struct {
	internal.Type__cdktfTerraformResource
}

func (j *jsiiProxy_Route53RecordSet) AliasTarget() Route53RecordSetAliasTargetOutputReference {
	var returns Route53RecordSetAliasTargetOutputReference
	_jsii_.Get(
		j,
		"aliasTarget",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Route53RecordSet) AliasTargetInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"aliasTargetInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Route53RecordSet) CdktfStack() cdktf.TerraformStack {
	var returns cdktf.TerraformStack
	_jsii_.Get(
		j,
		"cdktfStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Route53RecordSet) CidrRoutingConfig() Route53RecordSetCidrRoutingConfigOutputReference {
	var returns Route53RecordSetCidrRoutingConfigOutputReference
	_jsii_.Get(
		j,
		"cidrRoutingConfig",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Route53RecordSet) CidrRoutingConfigInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"cidrRoutingConfigInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Route53RecordSet) Comment() *string {
	var returns *string
	_jsii_.Get(
		j,
		"comment",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Route53RecordSet) CommentInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"commentInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Route53RecordSet) Connection() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"connection",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Route53RecordSet) ConstructNodeMetadata() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"constructNodeMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Route53RecordSet) Count() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"count",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Route53RecordSet) DependsOn() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"dependsOn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Route53RecordSet) Failover() *string {
	var returns *string
	_jsii_.Get(
		j,
		"failover",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Route53RecordSet) FailoverInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"failoverInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Route53RecordSet) ForEach() cdktf.ITerraformIterator {
	var returns cdktf.ITerraformIterator
	_jsii_.Get(
		j,
		"forEach",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Route53RecordSet) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Route53RecordSet) FriendlyUniqueId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"friendlyUniqueId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Route53RecordSet) GeoLocation() Route53RecordSetGeoLocationOutputReference {
	var returns Route53RecordSetGeoLocationOutputReference
	_jsii_.Get(
		j,
		"geoLocation",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Route53RecordSet) GeoLocationInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"geoLocationInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Route53RecordSet) GeoProximityLocation() Route53RecordSetGeoProximityLocationOutputReference {
	var returns Route53RecordSetGeoProximityLocationOutputReference
	_jsii_.Get(
		j,
		"geoProximityLocation",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Route53RecordSet) GeoProximityLocationInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"geoProximityLocationInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Route53RecordSet) HealthCheckId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"healthCheckId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Route53RecordSet) HealthCheckIdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"healthCheckIdInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Route53RecordSet) HostedZoneId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"hostedZoneId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Route53RecordSet) HostedZoneIdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"hostedZoneIdInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Route53RecordSet) HostedZoneName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"hostedZoneName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Route53RecordSet) HostedZoneNameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"hostedZoneNameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Route53RecordSet) Id() *string {
	var returns *string
	_jsii_.Get(
		j,
		"id",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Route53RecordSet) Lifecycle() *cdktf.TerraformResourceLifecycle {
	var returns *cdktf.TerraformResourceLifecycle
	_jsii_.Get(
		j,
		"lifecycle",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Route53RecordSet) MultiValueAnswer() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"multiValueAnswer",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Route53RecordSet) MultiValueAnswerInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"multiValueAnswerInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Route53RecordSet) Name() *string {
	var returns *string
	_jsii_.Get(
		j,
		"name",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Route53RecordSet) NameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"nameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Route53RecordSet) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Route53RecordSet) Provider() cdktf.TerraformProvider {
	var returns cdktf.TerraformProvider
	_jsii_.Get(
		j,
		"provider",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Route53RecordSet) Provisioners() *[]interface{} {
	var returns *[]interface{}
	_jsii_.Get(
		j,
		"provisioners",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Route53RecordSet) RawOverrides() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"rawOverrides",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Route53RecordSet) RecordSetId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"recordSetId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Route53RecordSet) Region() *string {
	var returns *string
	_jsii_.Get(
		j,
		"region",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Route53RecordSet) RegionInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"regionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Route53RecordSet) ResourceRecords() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"resourceRecords",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Route53RecordSet) ResourceRecordsInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"resourceRecordsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Route53RecordSet) SetIdentifier() *string {
	var returns *string
	_jsii_.Get(
		j,
		"setIdentifier",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Route53RecordSet) SetIdentifierInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"setIdentifierInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Route53RecordSet) TerraformGeneratorMetadata() *cdktf.TerraformProviderGeneratorMetadata {
	var returns *cdktf.TerraformProviderGeneratorMetadata
	_jsii_.Get(
		j,
		"terraformGeneratorMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Route53RecordSet) TerraformMetaArguments() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"terraformMetaArguments",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Route53RecordSet) TerraformResourceType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformResourceType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Route53RecordSet) Ttl() *string {
	var returns *string
	_jsii_.Get(
		j,
		"ttl",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Route53RecordSet) TtlInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"ttlInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Route53RecordSet) Type() *string {
	var returns *string
	_jsii_.Get(
		j,
		"type",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Route53RecordSet) TypeInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"typeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Route53RecordSet) Weight() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"weight",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Route53RecordSet) WeightInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"weightInput",
		&returns,
	)
	return returns
}


// Create a new {@link https://registry.terraform.io/providers/hashicorp/awscc/1.49.0/docs/resources/route53_record_set awscc_route53_record_set} Resource.
func NewRoute53RecordSet(scope constructs.Construct, id *string, config *Route53RecordSetConfig) Route53RecordSet {
	_init_.Initialize()

	if err := validateNewRoute53RecordSetParameters(scope, id, config); err != nil {
		panic(err)
	}
	j := jsiiProxy_Route53RecordSet{}

	_jsii_.Create(
		"awscc.route53RecordSet.Route53RecordSet",
		[]interface{}{scope, id, config},
		&j,
	)

	return &j
}

// Create a new {@link https://registry.terraform.io/providers/hashicorp/awscc/1.49.0/docs/resources/route53_record_set awscc_route53_record_set} Resource.
func NewRoute53RecordSet_Override(r Route53RecordSet, scope constructs.Construct, id *string, config *Route53RecordSetConfig) {
	_init_.Initialize()

	_jsii_.Create(
		"awscc.route53RecordSet.Route53RecordSet",
		[]interface{}{scope, id, config},
		r,
	)
}

func (j *jsiiProxy_Route53RecordSet)SetComment(val *string) {
	if err := j.validateSetCommentParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"comment",
		val,
	)
}

func (j *jsiiProxy_Route53RecordSet)SetConnection(val interface{}) {
	if err := j.validateSetConnectionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"connection",
		val,
	)
}

func (j *jsiiProxy_Route53RecordSet)SetCount(val interface{}) {
	if err := j.validateSetCountParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"count",
		val,
	)
}

func (j *jsiiProxy_Route53RecordSet)SetDependsOn(val *[]*string) {
	_jsii_.Set(
		j,
		"dependsOn",
		val,
	)
}

func (j *jsiiProxy_Route53RecordSet)SetFailover(val *string) {
	if err := j.validateSetFailoverParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"failover",
		val,
	)
}

func (j *jsiiProxy_Route53RecordSet)SetForEach(val cdktf.ITerraformIterator) {
	_jsii_.Set(
		j,
		"forEach",
		val,
	)
}

func (j *jsiiProxy_Route53RecordSet)SetHealthCheckId(val *string) {
	if err := j.validateSetHealthCheckIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"healthCheckId",
		val,
	)
}

func (j *jsiiProxy_Route53RecordSet)SetHostedZoneId(val *string) {
	if err := j.validateSetHostedZoneIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"hostedZoneId",
		val,
	)
}

func (j *jsiiProxy_Route53RecordSet)SetHostedZoneName(val *string) {
	if err := j.validateSetHostedZoneNameParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"hostedZoneName",
		val,
	)
}

func (j *jsiiProxy_Route53RecordSet)SetLifecycle(val *cdktf.TerraformResourceLifecycle) {
	if err := j.validateSetLifecycleParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"lifecycle",
		val,
	)
}

func (j *jsiiProxy_Route53RecordSet)SetMultiValueAnswer(val interface{}) {
	if err := j.validateSetMultiValueAnswerParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"multiValueAnswer",
		val,
	)
}

func (j *jsiiProxy_Route53RecordSet)SetName(val *string) {
	if err := j.validateSetNameParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"name",
		val,
	)
}

func (j *jsiiProxy_Route53RecordSet)SetProvider(val cdktf.TerraformProvider) {
	_jsii_.Set(
		j,
		"provider",
		val,
	)
}

func (j *jsiiProxy_Route53RecordSet)SetProvisioners(val *[]interface{}) {
	if err := j.validateSetProvisionersParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"provisioners",
		val,
	)
}

func (j *jsiiProxy_Route53RecordSet)SetRegion(val *string) {
	if err := j.validateSetRegionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"region",
		val,
	)
}

func (j *jsiiProxy_Route53RecordSet)SetResourceRecords(val *[]*string) {
	if err := j.validateSetResourceRecordsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"resourceRecords",
		val,
	)
}

func (j *jsiiProxy_Route53RecordSet)SetSetIdentifier(val *string) {
	if err := j.validateSetSetIdentifierParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"setIdentifier",
		val,
	)
}

func (j *jsiiProxy_Route53RecordSet)SetTtl(val *string) {
	if err := j.validateSetTtlParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"ttl",
		val,
	)
}

func (j *jsiiProxy_Route53RecordSet)SetType(val *string) {
	if err := j.validateSetTypeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"type",
		val,
	)
}

func (j *jsiiProxy_Route53RecordSet)SetWeight(val *float64) {
	if err := j.validateSetWeightParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"weight",
		val,
	)
}

// Generates CDKTF code for importing a Route53RecordSet resource upon running "cdktf plan <stack-name>".
func Route53RecordSet_GenerateConfigForImport(scope constructs.Construct, importToId *string, importFromId *string, provider cdktf.TerraformProvider) cdktf.ImportableResource {
	_init_.Initialize()

	if err := validateRoute53RecordSet_GenerateConfigForImportParameters(scope, importToId, importFromId); err != nil {
		panic(err)
	}
	var returns cdktf.ImportableResource

	_jsii_.StaticInvoke(
		"awscc.route53RecordSet.Route53RecordSet",
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
func Route53RecordSet_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	if err := validateRoute53RecordSet_IsConstructParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"awscc.route53RecordSet.Route53RecordSet",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Experimental.
func Route53RecordSet_IsTerraformElement(x interface{}) *bool {
	_init_.Initialize()

	if err := validateRoute53RecordSet_IsTerraformElementParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"awscc.route53RecordSet.Route53RecordSet",
		"isTerraformElement",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Experimental.
func Route53RecordSet_IsTerraformResource(x interface{}) *bool {
	_init_.Initialize()

	if err := validateRoute53RecordSet_IsTerraformResourceParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"awscc.route53RecordSet.Route53RecordSet",
		"isTerraformResource",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func Route53RecordSet_TfResourceType() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"awscc.route53RecordSet.Route53RecordSet",
		"tfResourceType",
		&returns,
	)
	return returns
}

func (r *jsiiProxy_Route53RecordSet) AddMoveTarget(moveTarget *string) {
	if err := r.validateAddMoveTargetParameters(moveTarget); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		r,
		"addMoveTarget",
		[]interface{}{moveTarget},
	)
}

func (r *jsiiProxy_Route53RecordSet) AddOverride(path *string, value interface{}) {
	if err := r.validateAddOverrideParameters(path, value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		r,
		"addOverride",
		[]interface{}{path, value},
	)
}

func (r *jsiiProxy_Route53RecordSet) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
	if err := r.validateGetAnyMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]interface{}

	_jsii_.Invoke(
		r,
		"getAnyMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (r *jsiiProxy_Route53RecordSet) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
	if err := r.validateGetBooleanAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		r,
		"getBooleanAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (r *jsiiProxy_Route53RecordSet) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
	if err := r.validateGetBooleanMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*bool

	_jsii_.Invoke(
		r,
		"getBooleanMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (r *jsiiProxy_Route53RecordSet) GetListAttribute(terraformAttribute *string) *[]*string {
	if err := r.validateGetListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*string

	_jsii_.Invoke(
		r,
		"getListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (r *jsiiProxy_Route53RecordSet) GetNumberAttribute(terraformAttribute *string) *float64 {
	if err := r.validateGetNumberAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *float64

	_jsii_.Invoke(
		r,
		"getNumberAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (r *jsiiProxy_Route53RecordSet) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
	if err := r.validateGetNumberListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*float64

	_jsii_.Invoke(
		r,
		"getNumberListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (r *jsiiProxy_Route53RecordSet) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
	if err := r.validateGetNumberMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*float64

	_jsii_.Invoke(
		r,
		"getNumberMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (r *jsiiProxy_Route53RecordSet) GetStringAttribute(terraformAttribute *string) *string {
	if err := r.validateGetStringAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *string

	_jsii_.Invoke(
		r,
		"getStringAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (r *jsiiProxy_Route53RecordSet) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
	if err := r.validateGetStringMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*string

	_jsii_.Invoke(
		r,
		"getStringMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (r *jsiiProxy_Route53RecordSet) HasResourceMove() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		r,
		"hasResourceMove",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (r *jsiiProxy_Route53RecordSet) ImportFrom(id *string, provider cdktf.TerraformProvider) {
	if err := r.validateImportFromParameters(id); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		r,
		"importFrom",
		[]interface{}{id, provider},
	)
}

func (r *jsiiProxy_Route53RecordSet) InterpolationForAttribute(terraformAttribute *string) cdktf.IResolvable {
	if err := r.validateInterpolationForAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		r,
		"interpolationForAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (r *jsiiProxy_Route53RecordSet) MoveFromId(id *string) {
	if err := r.validateMoveFromIdParameters(id); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		r,
		"moveFromId",
		[]interface{}{id},
	)
}

func (r *jsiiProxy_Route53RecordSet) MoveTo(moveTarget *string, index interface{}) {
	if err := r.validateMoveToParameters(moveTarget, index); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		r,
		"moveTo",
		[]interface{}{moveTarget, index},
	)
}

func (r *jsiiProxy_Route53RecordSet) MoveToId(id *string) {
	if err := r.validateMoveToIdParameters(id); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		r,
		"moveToId",
		[]interface{}{id},
	)
}

func (r *jsiiProxy_Route53RecordSet) OverrideLogicalId(newLogicalId *string) {
	if err := r.validateOverrideLogicalIdParameters(newLogicalId); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		r,
		"overrideLogicalId",
		[]interface{}{newLogicalId},
	)
}

func (r *jsiiProxy_Route53RecordSet) PutAliasTarget(value *Route53RecordSetAliasTarget) {
	if err := r.validatePutAliasTargetParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		r,
		"putAliasTarget",
		[]interface{}{value},
	)
}

func (r *jsiiProxy_Route53RecordSet) PutCidrRoutingConfig(value *Route53RecordSetCidrRoutingConfig) {
	if err := r.validatePutCidrRoutingConfigParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		r,
		"putCidrRoutingConfig",
		[]interface{}{value},
	)
}

func (r *jsiiProxy_Route53RecordSet) PutGeoLocation(value *Route53RecordSetGeoLocation) {
	if err := r.validatePutGeoLocationParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		r,
		"putGeoLocation",
		[]interface{}{value},
	)
}

func (r *jsiiProxy_Route53RecordSet) PutGeoProximityLocation(value *Route53RecordSetGeoProximityLocation) {
	if err := r.validatePutGeoProximityLocationParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		r,
		"putGeoProximityLocation",
		[]interface{}{value},
	)
}

func (r *jsiiProxy_Route53RecordSet) ResetAliasTarget() {
	_jsii_.InvokeVoid(
		r,
		"resetAliasTarget",
		nil, // no parameters
	)
}

func (r *jsiiProxy_Route53RecordSet) ResetCidrRoutingConfig() {
	_jsii_.InvokeVoid(
		r,
		"resetCidrRoutingConfig",
		nil, // no parameters
	)
}

func (r *jsiiProxy_Route53RecordSet) ResetComment() {
	_jsii_.InvokeVoid(
		r,
		"resetComment",
		nil, // no parameters
	)
}

func (r *jsiiProxy_Route53RecordSet) ResetFailover() {
	_jsii_.InvokeVoid(
		r,
		"resetFailover",
		nil, // no parameters
	)
}

func (r *jsiiProxy_Route53RecordSet) ResetGeoLocation() {
	_jsii_.InvokeVoid(
		r,
		"resetGeoLocation",
		nil, // no parameters
	)
}

func (r *jsiiProxy_Route53RecordSet) ResetGeoProximityLocation() {
	_jsii_.InvokeVoid(
		r,
		"resetGeoProximityLocation",
		nil, // no parameters
	)
}

func (r *jsiiProxy_Route53RecordSet) ResetHealthCheckId() {
	_jsii_.InvokeVoid(
		r,
		"resetHealthCheckId",
		nil, // no parameters
	)
}

func (r *jsiiProxy_Route53RecordSet) ResetHostedZoneId() {
	_jsii_.InvokeVoid(
		r,
		"resetHostedZoneId",
		nil, // no parameters
	)
}

func (r *jsiiProxy_Route53RecordSet) ResetHostedZoneName() {
	_jsii_.InvokeVoid(
		r,
		"resetHostedZoneName",
		nil, // no parameters
	)
}

func (r *jsiiProxy_Route53RecordSet) ResetMultiValueAnswer() {
	_jsii_.InvokeVoid(
		r,
		"resetMultiValueAnswer",
		nil, // no parameters
	)
}

func (r *jsiiProxy_Route53RecordSet) ResetOverrideLogicalId() {
	_jsii_.InvokeVoid(
		r,
		"resetOverrideLogicalId",
		nil, // no parameters
	)
}

func (r *jsiiProxy_Route53RecordSet) ResetRegion() {
	_jsii_.InvokeVoid(
		r,
		"resetRegion",
		nil, // no parameters
	)
}

func (r *jsiiProxy_Route53RecordSet) ResetResourceRecords() {
	_jsii_.InvokeVoid(
		r,
		"resetResourceRecords",
		nil, // no parameters
	)
}

func (r *jsiiProxy_Route53RecordSet) ResetSetIdentifier() {
	_jsii_.InvokeVoid(
		r,
		"resetSetIdentifier",
		nil, // no parameters
	)
}

func (r *jsiiProxy_Route53RecordSet) ResetTtl() {
	_jsii_.InvokeVoid(
		r,
		"resetTtl",
		nil, // no parameters
	)
}

func (r *jsiiProxy_Route53RecordSet) ResetWeight() {
	_jsii_.InvokeVoid(
		r,
		"resetWeight",
		nil, // no parameters
	)
}

func (r *jsiiProxy_Route53RecordSet) SynthesizeAttributes() *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		r,
		"synthesizeAttributes",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (r *jsiiProxy_Route53RecordSet) SynthesizeHclAttributes() *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		r,
		"synthesizeHclAttributes",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (r *jsiiProxy_Route53RecordSet) ToHclTerraform() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		r,
		"toHclTerraform",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (r *jsiiProxy_Route53RecordSet) ToMetadata() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		r,
		"toMetadata",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (r *jsiiProxy_Route53RecordSet) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		r,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (r *jsiiProxy_Route53RecordSet) ToTerraform() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		r,
		"toTerraform",
		nil, // no parameters
		&returns,
	)

	return returns
}

