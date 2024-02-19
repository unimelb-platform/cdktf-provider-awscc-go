package neptunedbcluster

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/unimelb-platform/cdktf-provider-awscc-go/awscc/jsii"

	"github.com/aws/constructs-go/constructs/v10"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
	"github.com/unimelb-platform/cdktf-provider-awscc-go/awscc/neptunedbcluster/internal"
)

// Represents a {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/neptune_db_cluster awscc_neptune_db_cluster}.
type NeptuneDbCluster interface {
	cdktf.TerraformResource
	AssociatedRoles() NeptuneDbClusterAssociatedRolesList
	AssociatedRolesInput() interface{}
	AvailabilityZones() *[]*string
	SetAvailabilityZones(val *[]*string)
	AvailabilityZonesInput() *[]*string
	BackupRetentionPeriod() *float64
	SetBackupRetentionPeriod(val *float64)
	BackupRetentionPeriodInput() *float64
	// Experimental.
	CdktfStack() cdktf.TerraformStack
	ClusterResourceId() *string
	// Experimental.
	Connection() interface{}
	// Experimental.
	SetConnection(val interface{})
	// Experimental.
	ConstructNodeMetadata() *map[string]interface{}
	CopyTagsToSnapshot() interface{}
	SetCopyTagsToSnapshot(val interface{})
	CopyTagsToSnapshotInput() interface{}
	// Experimental.
	Count() interface{}
	// Experimental.
	SetCount(val interface{})
	DbClusterIdentifier() *string
	SetDbClusterIdentifier(val *string)
	DbClusterIdentifierInput() *string
	DbClusterParameterGroupName() *string
	SetDbClusterParameterGroupName(val *string)
	DbClusterParameterGroupNameInput() *string
	DbInstanceParameterGroupName() *string
	SetDbInstanceParameterGroupName(val *string)
	DbInstanceParameterGroupNameInput() *string
	DbPort() *float64
	SetDbPort(val *float64)
	DbPortInput() *float64
	DbSubnetGroupName() *string
	SetDbSubnetGroupName(val *string)
	DbSubnetGroupNameInput() *string
	DeletionProtection() interface{}
	SetDeletionProtection(val interface{})
	DeletionProtectionInput() interface{}
	// Experimental.
	DependsOn() *[]*string
	// Experimental.
	SetDependsOn(val *[]*string)
	EnableCloudwatchLogsExports() *[]*string
	SetEnableCloudwatchLogsExports(val *[]*string)
	EnableCloudwatchLogsExportsInput() *[]*string
	Endpoint() *string
	EngineVersion() *string
	SetEngineVersion(val *string)
	EngineVersionInput() *string
	// Experimental.
	ForEach() cdktf.ITerraformIterator
	// Experimental.
	SetForEach(val cdktf.ITerraformIterator)
	// Experimental.
	Fqn() *string
	// Experimental.
	FriendlyUniqueId() *string
	IamAuthEnabled() interface{}
	SetIamAuthEnabled(val interface{})
	IamAuthEnabledInput() interface{}
	Id() *string
	KmsKeyId() *string
	SetKmsKeyId(val *string)
	KmsKeyIdInput() *string
	// Experimental.
	Lifecycle() *cdktf.TerraformResourceLifecycle
	// Experimental.
	SetLifecycle(val *cdktf.TerraformResourceLifecycle)
	// The tree node.
	Node() constructs.Node
	Port() *string
	PreferredBackupWindow() *string
	SetPreferredBackupWindow(val *string)
	PreferredBackupWindowInput() *string
	PreferredMaintenanceWindow() *string
	SetPreferredMaintenanceWindow(val *string)
	PreferredMaintenanceWindowInput() *string
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
	ReadEndpoint() *string
	RestoreToTime() *string
	SetRestoreToTime(val *string)
	RestoreToTimeInput() *string
	RestoreType() *string
	SetRestoreType(val *string)
	RestoreTypeInput() *string
	ServerlessScalingConfiguration() NeptuneDbClusterServerlessScalingConfigurationOutputReference
	ServerlessScalingConfigurationInput() interface{}
	SnapshotIdentifier() *string
	SetSnapshotIdentifier(val *string)
	SnapshotIdentifierInput() *string
	SourceDbClusterIdentifier() *string
	SetSourceDbClusterIdentifier(val *string)
	SourceDbClusterIdentifierInput() *string
	StorageEncrypted() interface{}
	SetStorageEncrypted(val interface{})
	StorageEncryptedInput() interface{}
	Tags() NeptuneDbClusterTagsList
	TagsInput() interface{}
	// Experimental.
	TerraformGeneratorMetadata() *cdktf.TerraformProviderGeneratorMetadata
	// Experimental.
	TerraformMetaArguments() *map[string]interface{}
	// Experimental.
	TerraformResourceType() *string
	UseLatestRestorableTime() interface{}
	SetUseLatestRestorableTime(val interface{})
	UseLatestRestorableTimeInput() interface{}
	VpcSecurityGroupIds() *[]*string
	SetVpcSecurityGroupIds(val *[]*string)
	VpcSecurityGroupIdsInput() *[]*string
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
	PutAssociatedRoles(value interface{})
	PutServerlessScalingConfiguration(value *NeptuneDbClusterServerlessScalingConfiguration)
	PutTags(value interface{})
	ResetAssociatedRoles()
	ResetAvailabilityZones()
	ResetBackupRetentionPeriod()
	ResetCopyTagsToSnapshot()
	ResetDbClusterIdentifier()
	ResetDbClusterParameterGroupName()
	ResetDbInstanceParameterGroupName()
	ResetDbPort()
	ResetDbSubnetGroupName()
	ResetDeletionProtection()
	ResetEnableCloudwatchLogsExports()
	ResetEngineVersion()
	ResetIamAuthEnabled()
	ResetKmsKeyId()
	// Resets a previously passed logical Id to use the auto-generated logical id again.
	// Experimental.
	ResetOverrideLogicalId()
	ResetPreferredBackupWindow()
	ResetPreferredMaintenanceWindow()
	ResetRestoreToTime()
	ResetRestoreType()
	ResetServerlessScalingConfiguration()
	ResetSnapshotIdentifier()
	ResetSourceDbClusterIdentifier()
	ResetStorageEncrypted()
	ResetTags()
	ResetUseLatestRestorableTime()
	ResetVpcSecurityGroupIds()
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

// The jsii proxy struct for NeptuneDbCluster
type jsiiProxy_NeptuneDbCluster struct {
	internal.Type__cdktfTerraformResource
}

func (j *jsiiProxy_NeptuneDbCluster) AssociatedRoles() NeptuneDbClusterAssociatedRolesList {
	var returns NeptuneDbClusterAssociatedRolesList
	_jsii_.Get(
		j,
		"associatedRoles",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NeptuneDbCluster) AssociatedRolesInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"associatedRolesInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NeptuneDbCluster) AvailabilityZones() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"availabilityZones",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NeptuneDbCluster) AvailabilityZonesInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"availabilityZonesInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NeptuneDbCluster) BackupRetentionPeriod() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"backupRetentionPeriod",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NeptuneDbCluster) BackupRetentionPeriodInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"backupRetentionPeriodInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NeptuneDbCluster) CdktfStack() cdktf.TerraformStack {
	var returns cdktf.TerraformStack
	_jsii_.Get(
		j,
		"cdktfStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NeptuneDbCluster) ClusterResourceId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"clusterResourceId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NeptuneDbCluster) Connection() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"connection",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NeptuneDbCluster) ConstructNodeMetadata() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"constructNodeMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NeptuneDbCluster) CopyTagsToSnapshot() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"copyTagsToSnapshot",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NeptuneDbCluster) CopyTagsToSnapshotInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"copyTagsToSnapshotInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NeptuneDbCluster) Count() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"count",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NeptuneDbCluster) DbClusterIdentifier() *string {
	var returns *string
	_jsii_.Get(
		j,
		"dbClusterIdentifier",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NeptuneDbCluster) DbClusterIdentifierInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"dbClusterIdentifierInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NeptuneDbCluster) DbClusterParameterGroupName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"dbClusterParameterGroupName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NeptuneDbCluster) DbClusterParameterGroupNameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"dbClusterParameterGroupNameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NeptuneDbCluster) DbInstanceParameterGroupName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"dbInstanceParameterGroupName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NeptuneDbCluster) DbInstanceParameterGroupNameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"dbInstanceParameterGroupNameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NeptuneDbCluster) DbPort() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"dbPort",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NeptuneDbCluster) DbPortInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"dbPortInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NeptuneDbCluster) DbSubnetGroupName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"dbSubnetGroupName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NeptuneDbCluster) DbSubnetGroupNameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"dbSubnetGroupNameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NeptuneDbCluster) DeletionProtection() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"deletionProtection",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NeptuneDbCluster) DeletionProtectionInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"deletionProtectionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NeptuneDbCluster) DependsOn() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"dependsOn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NeptuneDbCluster) EnableCloudwatchLogsExports() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"enableCloudwatchLogsExports",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NeptuneDbCluster) EnableCloudwatchLogsExportsInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"enableCloudwatchLogsExportsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NeptuneDbCluster) Endpoint() *string {
	var returns *string
	_jsii_.Get(
		j,
		"endpoint",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NeptuneDbCluster) EngineVersion() *string {
	var returns *string
	_jsii_.Get(
		j,
		"engineVersion",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NeptuneDbCluster) EngineVersionInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"engineVersionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NeptuneDbCluster) ForEach() cdktf.ITerraformIterator {
	var returns cdktf.ITerraformIterator
	_jsii_.Get(
		j,
		"forEach",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NeptuneDbCluster) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NeptuneDbCluster) FriendlyUniqueId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"friendlyUniqueId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NeptuneDbCluster) IamAuthEnabled() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"iamAuthEnabled",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NeptuneDbCluster) IamAuthEnabledInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"iamAuthEnabledInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NeptuneDbCluster) Id() *string {
	var returns *string
	_jsii_.Get(
		j,
		"id",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NeptuneDbCluster) KmsKeyId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"kmsKeyId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NeptuneDbCluster) KmsKeyIdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"kmsKeyIdInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NeptuneDbCluster) Lifecycle() *cdktf.TerraformResourceLifecycle {
	var returns *cdktf.TerraformResourceLifecycle
	_jsii_.Get(
		j,
		"lifecycle",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NeptuneDbCluster) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NeptuneDbCluster) Port() *string {
	var returns *string
	_jsii_.Get(
		j,
		"port",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NeptuneDbCluster) PreferredBackupWindow() *string {
	var returns *string
	_jsii_.Get(
		j,
		"preferredBackupWindow",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NeptuneDbCluster) PreferredBackupWindowInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"preferredBackupWindowInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NeptuneDbCluster) PreferredMaintenanceWindow() *string {
	var returns *string
	_jsii_.Get(
		j,
		"preferredMaintenanceWindow",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NeptuneDbCluster) PreferredMaintenanceWindowInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"preferredMaintenanceWindowInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NeptuneDbCluster) Provider() cdktf.TerraformProvider {
	var returns cdktf.TerraformProvider
	_jsii_.Get(
		j,
		"provider",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NeptuneDbCluster) Provisioners() *[]interface{} {
	var returns *[]interface{}
	_jsii_.Get(
		j,
		"provisioners",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NeptuneDbCluster) RawOverrides() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"rawOverrides",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NeptuneDbCluster) ReadEndpoint() *string {
	var returns *string
	_jsii_.Get(
		j,
		"readEndpoint",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NeptuneDbCluster) RestoreToTime() *string {
	var returns *string
	_jsii_.Get(
		j,
		"restoreToTime",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NeptuneDbCluster) RestoreToTimeInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"restoreToTimeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NeptuneDbCluster) RestoreType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"restoreType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NeptuneDbCluster) RestoreTypeInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"restoreTypeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NeptuneDbCluster) ServerlessScalingConfiguration() NeptuneDbClusterServerlessScalingConfigurationOutputReference {
	var returns NeptuneDbClusterServerlessScalingConfigurationOutputReference
	_jsii_.Get(
		j,
		"serverlessScalingConfiguration",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NeptuneDbCluster) ServerlessScalingConfigurationInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"serverlessScalingConfigurationInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NeptuneDbCluster) SnapshotIdentifier() *string {
	var returns *string
	_jsii_.Get(
		j,
		"snapshotIdentifier",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NeptuneDbCluster) SnapshotIdentifierInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"snapshotIdentifierInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NeptuneDbCluster) SourceDbClusterIdentifier() *string {
	var returns *string
	_jsii_.Get(
		j,
		"sourceDbClusterIdentifier",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NeptuneDbCluster) SourceDbClusterIdentifierInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"sourceDbClusterIdentifierInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NeptuneDbCluster) StorageEncrypted() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"storageEncrypted",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NeptuneDbCluster) StorageEncryptedInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"storageEncryptedInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NeptuneDbCluster) Tags() NeptuneDbClusterTagsList {
	var returns NeptuneDbClusterTagsList
	_jsii_.Get(
		j,
		"tags",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NeptuneDbCluster) TagsInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"tagsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NeptuneDbCluster) TerraformGeneratorMetadata() *cdktf.TerraformProviderGeneratorMetadata {
	var returns *cdktf.TerraformProviderGeneratorMetadata
	_jsii_.Get(
		j,
		"terraformGeneratorMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NeptuneDbCluster) TerraformMetaArguments() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"terraformMetaArguments",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NeptuneDbCluster) TerraformResourceType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformResourceType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NeptuneDbCluster) UseLatestRestorableTime() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"useLatestRestorableTime",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NeptuneDbCluster) UseLatestRestorableTimeInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"useLatestRestorableTimeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NeptuneDbCluster) VpcSecurityGroupIds() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"vpcSecurityGroupIds",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NeptuneDbCluster) VpcSecurityGroupIdsInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"vpcSecurityGroupIdsInput",
		&returns,
	)
	return returns
}


// Create a new {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/neptune_db_cluster awscc_neptune_db_cluster} Resource.
func NewNeptuneDbCluster(scope constructs.Construct, id *string, config *NeptuneDbClusterConfig) NeptuneDbCluster {
	_init_.Initialize()

	if err := validateNewNeptuneDbClusterParameters(scope, id, config); err != nil {
		panic(err)
	}
	j := jsiiProxy_NeptuneDbCluster{}

	_jsii_.Create(
		"awscc.neptuneDbCluster.NeptuneDbCluster",
		[]interface{}{scope, id, config},
		&j,
	)

	return &j
}

// Create a new {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/neptune_db_cluster awscc_neptune_db_cluster} Resource.
func NewNeptuneDbCluster_Override(n NeptuneDbCluster, scope constructs.Construct, id *string, config *NeptuneDbClusterConfig) {
	_init_.Initialize()

	_jsii_.Create(
		"awscc.neptuneDbCluster.NeptuneDbCluster",
		[]interface{}{scope, id, config},
		n,
	)
}

func (j *jsiiProxy_NeptuneDbCluster)SetAvailabilityZones(val *[]*string) {
	if err := j.validateSetAvailabilityZonesParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"availabilityZones",
		val,
	)
}

func (j *jsiiProxy_NeptuneDbCluster)SetBackupRetentionPeriod(val *float64) {
	if err := j.validateSetBackupRetentionPeriodParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"backupRetentionPeriod",
		val,
	)
}

func (j *jsiiProxy_NeptuneDbCluster)SetConnection(val interface{}) {
	if err := j.validateSetConnectionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"connection",
		val,
	)
}

func (j *jsiiProxy_NeptuneDbCluster)SetCopyTagsToSnapshot(val interface{}) {
	if err := j.validateSetCopyTagsToSnapshotParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"copyTagsToSnapshot",
		val,
	)
}

func (j *jsiiProxy_NeptuneDbCluster)SetCount(val interface{}) {
	if err := j.validateSetCountParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"count",
		val,
	)
}

func (j *jsiiProxy_NeptuneDbCluster)SetDbClusterIdentifier(val *string) {
	if err := j.validateSetDbClusterIdentifierParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"dbClusterIdentifier",
		val,
	)
}

func (j *jsiiProxy_NeptuneDbCluster)SetDbClusterParameterGroupName(val *string) {
	if err := j.validateSetDbClusterParameterGroupNameParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"dbClusterParameterGroupName",
		val,
	)
}

func (j *jsiiProxy_NeptuneDbCluster)SetDbInstanceParameterGroupName(val *string) {
	if err := j.validateSetDbInstanceParameterGroupNameParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"dbInstanceParameterGroupName",
		val,
	)
}

func (j *jsiiProxy_NeptuneDbCluster)SetDbPort(val *float64) {
	if err := j.validateSetDbPortParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"dbPort",
		val,
	)
}

func (j *jsiiProxy_NeptuneDbCluster)SetDbSubnetGroupName(val *string) {
	if err := j.validateSetDbSubnetGroupNameParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"dbSubnetGroupName",
		val,
	)
}

func (j *jsiiProxy_NeptuneDbCluster)SetDeletionProtection(val interface{}) {
	if err := j.validateSetDeletionProtectionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"deletionProtection",
		val,
	)
}

func (j *jsiiProxy_NeptuneDbCluster)SetDependsOn(val *[]*string) {
	_jsii_.Set(
		j,
		"dependsOn",
		val,
	)
}

func (j *jsiiProxy_NeptuneDbCluster)SetEnableCloudwatchLogsExports(val *[]*string) {
	if err := j.validateSetEnableCloudwatchLogsExportsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"enableCloudwatchLogsExports",
		val,
	)
}

func (j *jsiiProxy_NeptuneDbCluster)SetEngineVersion(val *string) {
	if err := j.validateSetEngineVersionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"engineVersion",
		val,
	)
}

func (j *jsiiProxy_NeptuneDbCluster)SetForEach(val cdktf.ITerraformIterator) {
	_jsii_.Set(
		j,
		"forEach",
		val,
	)
}

func (j *jsiiProxy_NeptuneDbCluster)SetIamAuthEnabled(val interface{}) {
	if err := j.validateSetIamAuthEnabledParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"iamAuthEnabled",
		val,
	)
}

func (j *jsiiProxy_NeptuneDbCluster)SetKmsKeyId(val *string) {
	if err := j.validateSetKmsKeyIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"kmsKeyId",
		val,
	)
}

func (j *jsiiProxy_NeptuneDbCluster)SetLifecycle(val *cdktf.TerraformResourceLifecycle) {
	if err := j.validateSetLifecycleParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"lifecycle",
		val,
	)
}

func (j *jsiiProxy_NeptuneDbCluster)SetPreferredBackupWindow(val *string) {
	if err := j.validateSetPreferredBackupWindowParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"preferredBackupWindow",
		val,
	)
}

func (j *jsiiProxy_NeptuneDbCluster)SetPreferredMaintenanceWindow(val *string) {
	if err := j.validateSetPreferredMaintenanceWindowParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"preferredMaintenanceWindow",
		val,
	)
}

func (j *jsiiProxy_NeptuneDbCluster)SetProvider(val cdktf.TerraformProvider) {
	_jsii_.Set(
		j,
		"provider",
		val,
	)
}

func (j *jsiiProxy_NeptuneDbCluster)SetProvisioners(val *[]interface{}) {
	if err := j.validateSetProvisionersParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"provisioners",
		val,
	)
}

func (j *jsiiProxy_NeptuneDbCluster)SetRestoreToTime(val *string) {
	if err := j.validateSetRestoreToTimeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"restoreToTime",
		val,
	)
}

func (j *jsiiProxy_NeptuneDbCluster)SetRestoreType(val *string) {
	if err := j.validateSetRestoreTypeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"restoreType",
		val,
	)
}

func (j *jsiiProxy_NeptuneDbCluster)SetSnapshotIdentifier(val *string) {
	if err := j.validateSetSnapshotIdentifierParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"snapshotIdentifier",
		val,
	)
}

func (j *jsiiProxy_NeptuneDbCluster)SetSourceDbClusterIdentifier(val *string) {
	if err := j.validateSetSourceDbClusterIdentifierParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"sourceDbClusterIdentifier",
		val,
	)
}

func (j *jsiiProxy_NeptuneDbCluster)SetStorageEncrypted(val interface{}) {
	if err := j.validateSetStorageEncryptedParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"storageEncrypted",
		val,
	)
}

func (j *jsiiProxy_NeptuneDbCluster)SetUseLatestRestorableTime(val interface{}) {
	if err := j.validateSetUseLatestRestorableTimeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"useLatestRestorableTime",
		val,
	)
}

func (j *jsiiProxy_NeptuneDbCluster)SetVpcSecurityGroupIds(val *[]*string) {
	if err := j.validateSetVpcSecurityGroupIdsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"vpcSecurityGroupIds",
		val,
	)
}

// Generates CDKTF code for importing a NeptuneDbCluster resource upon running "cdktf plan <stack-name>".
func NeptuneDbCluster_GenerateConfigForImport(scope constructs.Construct, importToId *string, importFromId *string, provider cdktf.TerraformProvider) cdktf.ImportableResource {
	_init_.Initialize()

	if err := validateNeptuneDbCluster_GenerateConfigForImportParameters(scope, importToId, importFromId); err != nil {
		panic(err)
	}
	var returns cdktf.ImportableResource

	_jsii_.StaticInvoke(
		"awscc.neptuneDbCluster.NeptuneDbCluster",
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
func NeptuneDbCluster_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	if err := validateNeptuneDbCluster_IsConstructParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"awscc.neptuneDbCluster.NeptuneDbCluster",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Experimental.
func NeptuneDbCluster_IsTerraformElement(x interface{}) *bool {
	_init_.Initialize()

	if err := validateNeptuneDbCluster_IsTerraformElementParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"awscc.neptuneDbCluster.NeptuneDbCluster",
		"isTerraformElement",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Experimental.
func NeptuneDbCluster_IsTerraformResource(x interface{}) *bool {
	_init_.Initialize()

	if err := validateNeptuneDbCluster_IsTerraformResourceParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"awscc.neptuneDbCluster.NeptuneDbCluster",
		"isTerraformResource",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func NeptuneDbCluster_TfResourceType() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"awscc.neptuneDbCluster.NeptuneDbCluster",
		"tfResourceType",
		&returns,
	)
	return returns
}

func (n *jsiiProxy_NeptuneDbCluster) AddMoveTarget(moveTarget *string) {
	if err := n.validateAddMoveTargetParameters(moveTarget); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		n,
		"addMoveTarget",
		[]interface{}{moveTarget},
	)
}

func (n *jsiiProxy_NeptuneDbCluster) AddOverride(path *string, value interface{}) {
	if err := n.validateAddOverrideParameters(path, value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		n,
		"addOverride",
		[]interface{}{path, value},
	)
}

func (n *jsiiProxy_NeptuneDbCluster) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
	if err := n.validateGetAnyMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]interface{}

	_jsii_.Invoke(
		n,
		"getAnyMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (n *jsiiProxy_NeptuneDbCluster) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
	if err := n.validateGetBooleanAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		n,
		"getBooleanAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (n *jsiiProxy_NeptuneDbCluster) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
	if err := n.validateGetBooleanMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*bool

	_jsii_.Invoke(
		n,
		"getBooleanMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (n *jsiiProxy_NeptuneDbCluster) GetListAttribute(terraformAttribute *string) *[]*string {
	if err := n.validateGetListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*string

	_jsii_.Invoke(
		n,
		"getListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (n *jsiiProxy_NeptuneDbCluster) GetNumberAttribute(terraformAttribute *string) *float64 {
	if err := n.validateGetNumberAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *float64

	_jsii_.Invoke(
		n,
		"getNumberAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (n *jsiiProxy_NeptuneDbCluster) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
	if err := n.validateGetNumberListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*float64

	_jsii_.Invoke(
		n,
		"getNumberListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (n *jsiiProxy_NeptuneDbCluster) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
	if err := n.validateGetNumberMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*float64

	_jsii_.Invoke(
		n,
		"getNumberMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (n *jsiiProxy_NeptuneDbCluster) GetStringAttribute(terraformAttribute *string) *string {
	if err := n.validateGetStringAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *string

	_jsii_.Invoke(
		n,
		"getStringAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (n *jsiiProxy_NeptuneDbCluster) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
	if err := n.validateGetStringMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*string

	_jsii_.Invoke(
		n,
		"getStringMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (n *jsiiProxy_NeptuneDbCluster) HasResourceMove() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		n,
		"hasResourceMove",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (n *jsiiProxy_NeptuneDbCluster) ImportFrom(id *string, provider cdktf.TerraformProvider) {
	if err := n.validateImportFromParameters(id); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		n,
		"importFrom",
		[]interface{}{id, provider},
	)
}

func (n *jsiiProxy_NeptuneDbCluster) InterpolationForAttribute(terraformAttribute *string) cdktf.IResolvable {
	if err := n.validateInterpolationForAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		n,
		"interpolationForAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (n *jsiiProxy_NeptuneDbCluster) MoveFromId(id *string) {
	if err := n.validateMoveFromIdParameters(id); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		n,
		"moveFromId",
		[]interface{}{id},
	)
}

func (n *jsiiProxy_NeptuneDbCluster) MoveTo(moveTarget *string, index interface{}) {
	if err := n.validateMoveToParameters(moveTarget, index); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		n,
		"moveTo",
		[]interface{}{moveTarget, index},
	)
}

func (n *jsiiProxy_NeptuneDbCluster) MoveToId(id *string) {
	if err := n.validateMoveToIdParameters(id); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		n,
		"moveToId",
		[]interface{}{id},
	)
}

func (n *jsiiProxy_NeptuneDbCluster) OverrideLogicalId(newLogicalId *string) {
	if err := n.validateOverrideLogicalIdParameters(newLogicalId); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		n,
		"overrideLogicalId",
		[]interface{}{newLogicalId},
	)
}

func (n *jsiiProxy_NeptuneDbCluster) PutAssociatedRoles(value interface{}) {
	if err := n.validatePutAssociatedRolesParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		n,
		"putAssociatedRoles",
		[]interface{}{value},
	)
}

func (n *jsiiProxy_NeptuneDbCluster) PutServerlessScalingConfiguration(value *NeptuneDbClusterServerlessScalingConfiguration) {
	if err := n.validatePutServerlessScalingConfigurationParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		n,
		"putServerlessScalingConfiguration",
		[]interface{}{value},
	)
}

func (n *jsiiProxy_NeptuneDbCluster) PutTags(value interface{}) {
	if err := n.validatePutTagsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		n,
		"putTags",
		[]interface{}{value},
	)
}

func (n *jsiiProxy_NeptuneDbCluster) ResetAssociatedRoles() {
	_jsii_.InvokeVoid(
		n,
		"resetAssociatedRoles",
		nil, // no parameters
	)
}

func (n *jsiiProxy_NeptuneDbCluster) ResetAvailabilityZones() {
	_jsii_.InvokeVoid(
		n,
		"resetAvailabilityZones",
		nil, // no parameters
	)
}

func (n *jsiiProxy_NeptuneDbCluster) ResetBackupRetentionPeriod() {
	_jsii_.InvokeVoid(
		n,
		"resetBackupRetentionPeriod",
		nil, // no parameters
	)
}

func (n *jsiiProxy_NeptuneDbCluster) ResetCopyTagsToSnapshot() {
	_jsii_.InvokeVoid(
		n,
		"resetCopyTagsToSnapshot",
		nil, // no parameters
	)
}

func (n *jsiiProxy_NeptuneDbCluster) ResetDbClusterIdentifier() {
	_jsii_.InvokeVoid(
		n,
		"resetDbClusterIdentifier",
		nil, // no parameters
	)
}

func (n *jsiiProxy_NeptuneDbCluster) ResetDbClusterParameterGroupName() {
	_jsii_.InvokeVoid(
		n,
		"resetDbClusterParameterGroupName",
		nil, // no parameters
	)
}

func (n *jsiiProxy_NeptuneDbCluster) ResetDbInstanceParameterGroupName() {
	_jsii_.InvokeVoid(
		n,
		"resetDbInstanceParameterGroupName",
		nil, // no parameters
	)
}

func (n *jsiiProxy_NeptuneDbCluster) ResetDbPort() {
	_jsii_.InvokeVoid(
		n,
		"resetDbPort",
		nil, // no parameters
	)
}

func (n *jsiiProxy_NeptuneDbCluster) ResetDbSubnetGroupName() {
	_jsii_.InvokeVoid(
		n,
		"resetDbSubnetGroupName",
		nil, // no parameters
	)
}

func (n *jsiiProxy_NeptuneDbCluster) ResetDeletionProtection() {
	_jsii_.InvokeVoid(
		n,
		"resetDeletionProtection",
		nil, // no parameters
	)
}

func (n *jsiiProxy_NeptuneDbCluster) ResetEnableCloudwatchLogsExports() {
	_jsii_.InvokeVoid(
		n,
		"resetEnableCloudwatchLogsExports",
		nil, // no parameters
	)
}

func (n *jsiiProxy_NeptuneDbCluster) ResetEngineVersion() {
	_jsii_.InvokeVoid(
		n,
		"resetEngineVersion",
		nil, // no parameters
	)
}

func (n *jsiiProxy_NeptuneDbCluster) ResetIamAuthEnabled() {
	_jsii_.InvokeVoid(
		n,
		"resetIamAuthEnabled",
		nil, // no parameters
	)
}

func (n *jsiiProxy_NeptuneDbCluster) ResetKmsKeyId() {
	_jsii_.InvokeVoid(
		n,
		"resetKmsKeyId",
		nil, // no parameters
	)
}

func (n *jsiiProxy_NeptuneDbCluster) ResetOverrideLogicalId() {
	_jsii_.InvokeVoid(
		n,
		"resetOverrideLogicalId",
		nil, // no parameters
	)
}

func (n *jsiiProxy_NeptuneDbCluster) ResetPreferredBackupWindow() {
	_jsii_.InvokeVoid(
		n,
		"resetPreferredBackupWindow",
		nil, // no parameters
	)
}

func (n *jsiiProxy_NeptuneDbCluster) ResetPreferredMaintenanceWindow() {
	_jsii_.InvokeVoid(
		n,
		"resetPreferredMaintenanceWindow",
		nil, // no parameters
	)
}

func (n *jsiiProxy_NeptuneDbCluster) ResetRestoreToTime() {
	_jsii_.InvokeVoid(
		n,
		"resetRestoreToTime",
		nil, // no parameters
	)
}

func (n *jsiiProxy_NeptuneDbCluster) ResetRestoreType() {
	_jsii_.InvokeVoid(
		n,
		"resetRestoreType",
		nil, // no parameters
	)
}

func (n *jsiiProxy_NeptuneDbCluster) ResetServerlessScalingConfiguration() {
	_jsii_.InvokeVoid(
		n,
		"resetServerlessScalingConfiguration",
		nil, // no parameters
	)
}

func (n *jsiiProxy_NeptuneDbCluster) ResetSnapshotIdentifier() {
	_jsii_.InvokeVoid(
		n,
		"resetSnapshotIdentifier",
		nil, // no parameters
	)
}

func (n *jsiiProxy_NeptuneDbCluster) ResetSourceDbClusterIdentifier() {
	_jsii_.InvokeVoid(
		n,
		"resetSourceDbClusterIdentifier",
		nil, // no parameters
	)
}

func (n *jsiiProxy_NeptuneDbCluster) ResetStorageEncrypted() {
	_jsii_.InvokeVoid(
		n,
		"resetStorageEncrypted",
		nil, // no parameters
	)
}

func (n *jsiiProxy_NeptuneDbCluster) ResetTags() {
	_jsii_.InvokeVoid(
		n,
		"resetTags",
		nil, // no parameters
	)
}

func (n *jsiiProxy_NeptuneDbCluster) ResetUseLatestRestorableTime() {
	_jsii_.InvokeVoid(
		n,
		"resetUseLatestRestorableTime",
		nil, // no parameters
	)
}

func (n *jsiiProxy_NeptuneDbCluster) ResetVpcSecurityGroupIds() {
	_jsii_.InvokeVoid(
		n,
		"resetVpcSecurityGroupIds",
		nil, // no parameters
	)
}

func (n *jsiiProxy_NeptuneDbCluster) SynthesizeAttributes() *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		n,
		"synthesizeAttributes",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (n *jsiiProxy_NeptuneDbCluster) SynthesizeHclAttributes() *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		n,
		"synthesizeHclAttributes",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (n *jsiiProxy_NeptuneDbCluster) ToHclTerraform() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		n,
		"toHclTerraform",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (n *jsiiProxy_NeptuneDbCluster) ToMetadata() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		n,
		"toMetadata",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (n *jsiiProxy_NeptuneDbCluster) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		n,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (n *jsiiProxy_NeptuneDbCluster) ToTerraform() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		n,
		"toTerraform",
		nil, // no parameters
		&returns,
	)

	return returns
}

