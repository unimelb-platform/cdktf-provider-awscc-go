package redshiftserverlessnamespace

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/unimelb-platform/cdktf-provider-awscc-go/awscc/jsii"

	"github.com/aws/constructs-go/constructs/v10"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
	"github.com/unimelb-platform/cdktf-provider-awscc-go/awscc/redshiftserverlessnamespace/internal"
)

// Represents a {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/redshiftserverless_namespace awscc_redshiftserverless_namespace}.
type RedshiftserverlessNamespace interface {
	cdktf.TerraformResource
	AdminPasswordSecretKmsKeyId() *string
	SetAdminPasswordSecretKmsKeyId(val *string)
	AdminPasswordSecretKmsKeyIdInput() *string
	AdminUsername() *string
	SetAdminUsername(val *string)
	AdminUsernameInput() *string
	AdminUserPassword() *string
	SetAdminUserPassword(val *string)
	AdminUserPasswordInput() *string
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
	DbName() *string
	SetDbName(val *string)
	DbNameInput() *string
	DefaultIamRoleArn() *string
	SetDefaultIamRoleArn(val *string)
	DefaultIamRoleArnInput() *string
	// Experimental.
	DependsOn() *[]*string
	// Experimental.
	SetDependsOn(val *[]*string)
	FinalSnapshotName() *string
	SetFinalSnapshotName(val *string)
	FinalSnapshotNameInput() *string
	FinalSnapshotRetentionPeriod() *float64
	SetFinalSnapshotRetentionPeriod(val *float64)
	FinalSnapshotRetentionPeriodInput() *float64
	// Experimental.
	ForEach() cdktf.ITerraformIterator
	// Experimental.
	SetForEach(val cdktf.ITerraformIterator)
	// Experimental.
	Fqn() *string
	// Experimental.
	FriendlyUniqueId() *string
	IamRoles() *[]*string
	SetIamRoles(val *[]*string)
	IamRolesInput() *[]*string
	Id() *string
	KmsKeyId() *string
	SetKmsKeyId(val *string)
	KmsKeyIdInput() *string
	// Experimental.
	Lifecycle() *cdktf.TerraformResourceLifecycle
	// Experimental.
	SetLifecycle(val *cdktf.TerraformResourceLifecycle)
	LogExports() *[]*string
	SetLogExports(val *[]*string)
	LogExportsInput() *[]*string
	ManageAdminPassword() interface{}
	SetManageAdminPassword(val interface{})
	ManageAdminPasswordInput() interface{}
	Namespace() RedshiftserverlessNamespaceNamespaceOutputReference
	NamespaceName() *string
	SetNamespaceName(val *string)
	NamespaceNameInput() *string
	NamespaceResourcePolicy() *string
	SetNamespaceResourcePolicy(val *string)
	NamespaceResourcePolicyInput() *string
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
	RedshiftIdcApplicationArn() *string
	SetRedshiftIdcApplicationArn(val *string)
	RedshiftIdcApplicationArnInput() *string
	Tags() RedshiftserverlessNamespaceTagsList
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
	ImportFrom(id *string, provider cdktf.TerraformProvider)
	// Experimental.
	InterpolationForAttribute(terraformAttribute *string) cdktf.IResolvable
	// Moves this resource to the target resource given by moveTarget.
	// Experimental.
	MoveTo(moveTarget *string, index interface{})
	// Overrides the auto-generated logical ID with a specific ID.
	// Experimental.
	OverrideLogicalId(newLogicalId *string)
	PutTags(value interface{})
	ResetAdminPasswordSecretKmsKeyId()
	ResetAdminUsername()
	ResetAdminUserPassword()
	ResetDbName()
	ResetDefaultIamRoleArn()
	ResetFinalSnapshotName()
	ResetFinalSnapshotRetentionPeriod()
	ResetIamRoles()
	ResetKmsKeyId()
	ResetLogExports()
	ResetManageAdminPassword()
	ResetNamespaceResourcePolicy()
	// Resets a previously passed logical Id to use the auto-generated logical id again.
	// Experimental.
	ResetOverrideLogicalId()
	ResetRedshiftIdcApplicationArn()
	ResetTags()
	SynthesizeAttributes() *map[string]interface{}
	// Experimental.
	ToMetadata() interface{}
	// Returns a string representation of this construct.
	ToString() *string
	// Adds this resource to the terraform JSON output.
	// Experimental.
	ToTerraform() interface{}
}

// The jsii proxy struct for RedshiftserverlessNamespace
type jsiiProxy_RedshiftserverlessNamespace struct {
	internal.Type__cdktfTerraformResource
}

func (j *jsiiProxy_RedshiftserverlessNamespace) AdminPasswordSecretKmsKeyId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"adminPasswordSecretKmsKeyId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_RedshiftserverlessNamespace) AdminPasswordSecretKmsKeyIdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"adminPasswordSecretKmsKeyIdInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_RedshiftserverlessNamespace) AdminUsername() *string {
	var returns *string
	_jsii_.Get(
		j,
		"adminUsername",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_RedshiftserverlessNamespace) AdminUsernameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"adminUsernameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_RedshiftserverlessNamespace) AdminUserPassword() *string {
	var returns *string
	_jsii_.Get(
		j,
		"adminUserPassword",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_RedshiftserverlessNamespace) AdminUserPasswordInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"adminUserPasswordInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_RedshiftserverlessNamespace) CdktfStack() cdktf.TerraformStack {
	var returns cdktf.TerraformStack
	_jsii_.Get(
		j,
		"cdktfStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_RedshiftserverlessNamespace) Connection() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"connection",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_RedshiftserverlessNamespace) ConstructNodeMetadata() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"constructNodeMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_RedshiftserverlessNamespace) Count() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"count",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_RedshiftserverlessNamespace) DbName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"dbName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_RedshiftserverlessNamespace) DbNameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"dbNameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_RedshiftserverlessNamespace) DefaultIamRoleArn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"defaultIamRoleArn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_RedshiftserverlessNamespace) DefaultIamRoleArnInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"defaultIamRoleArnInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_RedshiftserverlessNamespace) DependsOn() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"dependsOn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_RedshiftserverlessNamespace) FinalSnapshotName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"finalSnapshotName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_RedshiftserverlessNamespace) FinalSnapshotNameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"finalSnapshotNameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_RedshiftserverlessNamespace) FinalSnapshotRetentionPeriod() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"finalSnapshotRetentionPeriod",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_RedshiftserverlessNamespace) FinalSnapshotRetentionPeriodInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"finalSnapshotRetentionPeriodInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_RedshiftserverlessNamespace) ForEach() cdktf.ITerraformIterator {
	var returns cdktf.ITerraformIterator
	_jsii_.Get(
		j,
		"forEach",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_RedshiftserverlessNamespace) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_RedshiftserverlessNamespace) FriendlyUniqueId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"friendlyUniqueId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_RedshiftserverlessNamespace) IamRoles() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"iamRoles",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_RedshiftserverlessNamespace) IamRolesInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"iamRolesInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_RedshiftserverlessNamespace) Id() *string {
	var returns *string
	_jsii_.Get(
		j,
		"id",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_RedshiftserverlessNamespace) KmsKeyId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"kmsKeyId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_RedshiftserverlessNamespace) KmsKeyIdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"kmsKeyIdInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_RedshiftserverlessNamespace) Lifecycle() *cdktf.TerraformResourceLifecycle {
	var returns *cdktf.TerraformResourceLifecycle
	_jsii_.Get(
		j,
		"lifecycle",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_RedshiftserverlessNamespace) LogExports() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"logExports",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_RedshiftserverlessNamespace) LogExportsInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"logExportsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_RedshiftserverlessNamespace) ManageAdminPassword() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"manageAdminPassword",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_RedshiftserverlessNamespace) ManageAdminPasswordInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"manageAdminPasswordInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_RedshiftserverlessNamespace) Namespace() RedshiftserverlessNamespaceNamespaceOutputReference {
	var returns RedshiftserverlessNamespaceNamespaceOutputReference
	_jsii_.Get(
		j,
		"namespace",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_RedshiftserverlessNamespace) NamespaceName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"namespaceName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_RedshiftserverlessNamespace) NamespaceNameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"namespaceNameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_RedshiftserverlessNamespace) NamespaceResourcePolicy() *string {
	var returns *string
	_jsii_.Get(
		j,
		"namespaceResourcePolicy",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_RedshiftserverlessNamespace) NamespaceResourcePolicyInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"namespaceResourcePolicyInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_RedshiftserverlessNamespace) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_RedshiftserverlessNamespace) Provider() cdktf.TerraformProvider {
	var returns cdktf.TerraformProvider
	_jsii_.Get(
		j,
		"provider",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_RedshiftserverlessNamespace) Provisioners() *[]interface{} {
	var returns *[]interface{}
	_jsii_.Get(
		j,
		"provisioners",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_RedshiftserverlessNamespace) RawOverrides() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"rawOverrides",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_RedshiftserverlessNamespace) RedshiftIdcApplicationArn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"redshiftIdcApplicationArn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_RedshiftserverlessNamespace) RedshiftIdcApplicationArnInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"redshiftIdcApplicationArnInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_RedshiftserverlessNamespace) Tags() RedshiftserverlessNamespaceTagsList {
	var returns RedshiftserverlessNamespaceTagsList
	_jsii_.Get(
		j,
		"tags",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_RedshiftserverlessNamespace) TagsInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"tagsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_RedshiftserverlessNamespace) TerraformGeneratorMetadata() *cdktf.TerraformProviderGeneratorMetadata {
	var returns *cdktf.TerraformProviderGeneratorMetadata
	_jsii_.Get(
		j,
		"terraformGeneratorMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_RedshiftserverlessNamespace) TerraformMetaArguments() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"terraformMetaArguments",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_RedshiftserverlessNamespace) TerraformResourceType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformResourceType",
		&returns,
	)
	return returns
}


// Create a new {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/redshiftserverless_namespace awscc_redshiftserverless_namespace} Resource.
func NewRedshiftserverlessNamespace(scope constructs.Construct, id *string, config *RedshiftserverlessNamespaceConfig) RedshiftserverlessNamespace {
	_init_.Initialize()

	if err := validateNewRedshiftserverlessNamespaceParameters(scope, id, config); err != nil {
		panic(err)
	}
	j := jsiiProxy_RedshiftserverlessNamespace{}

	_jsii_.Create(
		"awscc.redshiftserverlessNamespace.RedshiftserverlessNamespace",
		[]interface{}{scope, id, config},
		&j,
	)

	return &j
}

// Create a new {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/redshiftserverless_namespace awscc_redshiftserverless_namespace} Resource.
func NewRedshiftserverlessNamespace_Override(r RedshiftserverlessNamespace, scope constructs.Construct, id *string, config *RedshiftserverlessNamespaceConfig) {
	_init_.Initialize()

	_jsii_.Create(
		"awscc.redshiftserverlessNamespace.RedshiftserverlessNamespace",
		[]interface{}{scope, id, config},
		r,
	)
}

func (j *jsiiProxy_RedshiftserverlessNamespace)SetAdminPasswordSecretKmsKeyId(val *string) {
	if err := j.validateSetAdminPasswordSecretKmsKeyIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"adminPasswordSecretKmsKeyId",
		val,
	)
}

func (j *jsiiProxy_RedshiftserverlessNamespace)SetAdminUsername(val *string) {
	if err := j.validateSetAdminUsernameParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"adminUsername",
		val,
	)
}

func (j *jsiiProxy_RedshiftserverlessNamespace)SetAdminUserPassword(val *string) {
	if err := j.validateSetAdminUserPasswordParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"adminUserPassword",
		val,
	)
}

func (j *jsiiProxy_RedshiftserverlessNamespace)SetConnection(val interface{}) {
	if err := j.validateSetConnectionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"connection",
		val,
	)
}

func (j *jsiiProxy_RedshiftserverlessNamespace)SetCount(val interface{}) {
	if err := j.validateSetCountParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"count",
		val,
	)
}

func (j *jsiiProxy_RedshiftserverlessNamespace)SetDbName(val *string) {
	if err := j.validateSetDbNameParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"dbName",
		val,
	)
}

func (j *jsiiProxy_RedshiftserverlessNamespace)SetDefaultIamRoleArn(val *string) {
	if err := j.validateSetDefaultIamRoleArnParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"defaultIamRoleArn",
		val,
	)
}

func (j *jsiiProxy_RedshiftserverlessNamespace)SetDependsOn(val *[]*string) {
	_jsii_.Set(
		j,
		"dependsOn",
		val,
	)
}

func (j *jsiiProxy_RedshiftserverlessNamespace)SetFinalSnapshotName(val *string) {
	if err := j.validateSetFinalSnapshotNameParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"finalSnapshotName",
		val,
	)
}

func (j *jsiiProxy_RedshiftserverlessNamespace)SetFinalSnapshotRetentionPeriod(val *float64) {
	if err := j.validateSetFinalSnapshotRetentionPeriodParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"finalSnapshotRetentionPeriod",
		val,
	)
}

func (j *jsiiProxy_RedshiftserverlessNamespace)SetForEach(val cdktf.ITerraformIterator) {
	_jsii_.Set(
		j,
		"forEach",
		val,
	)
}

func (j *jsiiProxy_RedshiftserverlessNamespace)SetIamRoles(val *[]*string) {
	if err := j.validateSetIamRolesParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"iamRoles",
		val,
	)
}

func (j *jsiiProxy_RedshiftserverlessNamespace)SetKmsKeyId(val *string) {
	if err := j.validateSetKmsKeyIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"kmsKeyId",
		val,
	)
}

func (j *jsiiProxy_RedshiftserverlessNamespace)SetLifecycle(val *cdktf.TerraformResourceLifecycle) {
	if err := j.validateSetLifecycleParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"lifecycle",
		val,
	)
}

func (j *jsiiProxy_RedshiftserverlessNamespace)SetLogExports(val *[]*string) {
	if err := j.validateSetLogExportsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"logExports",
		val,
	)
}

func (j *jsiiProxy_RedshiftserverlessNamespace)SetManageAdminPassword(val interface{}) {
	if err := j.validateSetManageAdminPasswordParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"manageAdminPassword",
		val,
	)
}

func (j *jsiiProxy_RedshiftserverlessNamespace)SetNamespaceName(val *string) {
	if err := j.validateSetNamespaceNameParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"namespaceName",
		val,
	)
}

func (j *jsiiProxy_RedshiftserverlessNamespace)SetNamespaceResourcePolicy(val *string) {
	if err := j.validateSetNamespaceResourcePolicyParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"namespaceResourcePolicy",
		val,
	)
}

func (j *jsiiProxy_RedshiftserverlessNamespace)SetProvider(val cdktf.TerraformProvider) {
	_jsii_.Set(
		j,
		"provider",
		val,
	)
}

func (j *jsiiProxy_RedshiftserverlessNamespace)SetProvisioners(val *[]interface{}) {
	if err := j.validateSetProvisionersParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"provisioners",
		val,
	)
}

func (j *jsiiProxy_RedshiftserverlessNamespace)SetRedshiftIdcApplicationArn(val *string) {
	if err := j.validateSetRedshiftIdcApplicationArnParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"redshiftIdcApplicationArn",
		val,
	)
}

// Generates CDKTF code for importing a RedshiftserverlessNamespace resource upon running "cdktf plan <stack-name>".
func RedshiftserverlessNamespace_GenerateConfigForImport(scope constructs.Construct, importToId *string, importFromId *string, provider cdktf.TerraformProvider) cdktf.ImportableResource {
	_init_.Initialize()

	if err := validateRedshiftserverlessNamespace_GenerateConfigForImportParameters(scope, importToId, importFromId); err != nil {
		panic(err)
	}
	var returns cdktf.ImportableResource

	_jsii_.StaticInvoke(
		"awscc.redshiftserverlessNamespace.RedshiftserverlessNamespace",
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
func RedshiftserverlessNamespace_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	if err := validateRedshiftserverlessNamespace_IsConstructParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"awscc.redshiftserverlessNamespace.RedshiftserverlessNamespace",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Experimental.
func RedshiftserverlessNamespace_IsTerraformElement(x interface{}) *bool {
	_init_.Initialize()

	if err := validateRedshiftserverlessNamespace_IsTerraformElementParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"awscc.redshiftserverlessNamespace.RedshiftserverlessNamespace",
		"isTerraformElement",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Experimental.
func RedshiftserverlessNamespace_IsTerraformResource(x interface{}) *bool {
	_init_.Initialize()

	if err := validateRedshiftserverlessNamespace_IsTerraformResourceParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"awscc.redshiftserverlessNamespace.RedshiftserverlessNamespace",
		"isTerraformResource",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func RedshiftserverlessNamespace_TfResourceType() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"awscc.redshiftserverlessNamespace.RedshiftserverlessNamespace",
		"tfResourceType",
		&returns,
	)
	return returns
}

func (r *jsiiProxy_RedshiftserverlessNamespace) AddMoveTarget(moveTarget *string) {
	if err := r.validateAddMoveTargetParameters(moveTarget); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		r,
		"addMoveTarget",
		[]interface{}{moveTarget},
	)
}

func (r *jsiiProxy_RedshiftserverlessNamespace) AddOverride(path *string, value interface{}) {
	if err := r.validateAddOverrideParameters(path, value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		r,
		"addOverride",
		[]interface{}{path, value},
	)
}

func (r *jsiiProxy_RedshiftserverlessNamespace) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (r *jsiiProxy_RedshiftserverlessNamespace) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
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

func (r *jsiiProxy_RedshiftserverlessNamespace) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (r *jsiiProxy_RedshiftserverlessNamespace) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (r *jsiiProxy_RedshiftserverlessNamespace) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (r *jsiiProxy_RedshiftserverlessNamespace) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (r *jsiiProxy_RedshiftserverlessNamespace) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (r *jsiiProxy_RedshiftserverlessNamespace) GetStringAttribute(terraformAttribute *string) *string {
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

func (r *jsiiProxy_RedshiftserverlessNamespace) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (r *jsiiProxy_RedshiftserverlessNamespace) ImportFrom(id *string, provider cdktf.TerraformProvider) {
	if err := r.validateImportFromParameters(id); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		r,
		"importFrom",
		[]interface{}{id, provider},
	)
}

func (r *jsiiProxy_RedshiftserverlessNamespace) InterpolationForAttribute(terraformAttribute *string) cdktf.IResolvable {
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

func (r *jsiiProxy_RedshiftserverlessNamespace) MoveTo(moveTarget *string, index interface{}) {
	if err := r.validateMoveToParameters(moveTarget, index); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		r,
		"moveTo",
		[]interface{}{moveTarget, index},
	)
}

func (r *jsiiProxy_RedshiftserverlessNamespace) OverrideLogicalId(newLogicalId *string) {
	if err := r.validateOverrideLogicalIdParameters(newLogicalId); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		r,
		"overrideLogicalId",
		[]interface{}{newLogicalId},
	)
}

func (r *jsiiProxy_RedshiftserverlessNamespace) PutTags(value interface{}) {
	if err := r.validatePutTagsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		r,
		"putTags",
		[]interface{}{value},
	)
}

func (r *jsiiProxy_RedshiftserverlessNamespace) ResetAdminPasswordSecretKmsKeyId() {
	_jsii_.InvokeVoid(
		r,
		"resetAdminPasswordSecretKmsKeyId",
		nil, // no parameters
	)
}

func (r *jsiiProxy_RedshiftserverlessNamespace) ResetAdminUsername() {
	_jsii_.InvokeVoid(
		r,
		"resetAdminUsername",
		nil, // no parameters
	)
}

func (r *jsiiProxy_RedshiftserverlessNamespace) ResetAdminUserPassword() {
	_jsii_.InvokeVoid(
		r,
		"resetAdminUserPassword",
		nil, // no parameters
	)
}

func (r *jsiiProxy_RedshiftserverlessNamespace) ResetDbName() {
	_jsii_.InvokeVoid(
		r,
		"resetDbName",
		nil, // no parameters
	)
}

func (r *jsiiProxy_RedshiftserverlessNamespace) ResetDefaultIamRoleArn() {
	_jsii_.InvokeVoid(
		r,
		"resetDefaultIamRoleArn",
		nil, // no parameters
	)
}

func (r *jsiiProxy_RedshiftserverlessNamespace) ResetFinalSnapshotName() {
	_jsii_.InvokeVoid(
		r,
		"resetFinalSnapshotName",
		nil, // no parameters
	)
}

func (r *jsiiProxy_RedshiftserverlessNamespace) ResetFinalSnapshotRetentionPeriod() {
	_jsii_.InvokeVoid(
		r,
		"resetFinalSnapshotRetentionPeriod",
		nil, // no parameters
	)
}

func (r *jsiiProxy_RedshiftserverlessNamespace) ResetIamRoles() {
	_jsii_.InvokeVoid(
		r,
		"resetIamRoles",
		nil, // no parameters
	)
}

func (r *jsiiProxy_RedshiftserverlessNamespace) ResetKmsKeyId() {
	_jsii_.InvokeVoid(
		r,
		"resetKmsKeyId",
		nil, // no parameters
	)
}

func (r *jsiiProxy_RedshiftserverlessNamespace) ResetLogExports() {
	_jsii_.InvokeVoid(
		r,
		"resetLogExports",
		nil, // no parameters
	)
}

func (r *jsiiProxy_RedshiftserverlessNamespace) ResetManageAdminPassword() {
	_jsii_.InvokeVoid(
		r,
		"resetManageAdminPassword",
		nil, // no parameters
	)
}

func (r *jsiiProxy_RedshiftserverlessNamespace) ResetNamespaceResourcePolicy() {
	_jsii_.InvokeVoid(
		r,
		"resetNamespaceResourcePolicy",
		nil, // no parameters
	)
}

func (r *jsiiProxy_RedshiftserverlessNamespace) ResetOverrideLogicalId() {
	_jsii_.InvokeVoid(
		r,
		"resetOverrideLogicalId",
		nil, // no parameters
	)
}

func (r *jsiiProxy_RedshiftserverlessNamespace) ResetRedshiftIdcApplicationArn() {
	_jsii_.InvokeVoid(
		r,
		"resetRedshiftIdcApplicationArn",
		nil, // no parameters
	)
}

func (r *jsiiProxy_RedshiftserverlessNamespace) ResetTags() {
	_jsii_.InvokeVoid(
		r,
		"resetTags",
		nil, // no parameters
	)
}

func (r *jsiiProxy_RedshiftserverlessNamespace) SynthesizeAttributes() *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		r,
		"synthesizeAttributes",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (r *jsiiProxy_RedshiftserverlessNamespace) ToMetadata() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		r,
		"toMetadata",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (r *jsiiProxy_RedshiftserverlessNamespace) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		r,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (r *jsiiProxy_RedshiftserverlessNamespace) ToTerraform() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		r,
		"toTerraform",
		nil, // no parameters
		&returns,
	)

	return returns
}

