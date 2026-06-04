package evsenvironment


type EvsEnvironmentInitialVlans struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.49.0/docs/resources/evs_environment#edge_v_tep EvsEnvironment#edge_v_tep}.
	EdgeVTep *EvsEnvironmentInitialVlansEdgeVTep `field:"optional" json:"edgeVTep" yaml:"edgeVTep"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.49.0/docs/resources/evs_environment#expansion_vlan_1 EvsEnvironment#expansion_vlan_1}.
	ExpansionVlan1 *EvsEnvironmentInitialVlansExpansionVlan1 `field:"optional" json:"expansionVlan1" yaml:"expansionVlan1"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.49.0/docs/resources/evs_environment#expansion_vlan_2 EvsEnvironment#expansion_vlan_2}.
	ExpansionVlan2 *EvsEnvironmentInitialVlansExpansionVlan2 `field:"optional" json:"expansionVlan2" yaml:"expansionVlan2"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.49.0/docs/resources/evs_environment#hcx EvsEnvironment#hcx}.
	Hcx *EvsEnvironmentInitialVlansHcx `field:"optional" json:"hcx" yaml:"hcx"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.49.0/docs/resources/evs_environment#nsx_up_link EvsEnvironment#nsx_up_link}.
	NsxUpLink *EvsEnvironmentInitialVlansNsxUpLink `field:"optional" json:"nsxUpLink" yaml:"nsxUpLink"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.49.0/docs/resources/evs_environment#vmk_management EvsEnvironment#vmk_management}.
	VmkManagement *EvsEnvironmentInitialVlansVmkManagement `field:"optional" json:"vmkManagement" yaml:"vmkManagement"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.49.0/docs/resources/evs_environment#vm_management EvsEnvironment#vm_management}.
	VmManagement *EvsEnvironmentInitialVlansVmManagement `field:"optional" json:"vmManagement" yaml:"vmManagement"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.49.0/docs/resources/evs_environment#v_motion EvsEnvironment#v_motion}.
	VMotion *EvsEnvironmentInitialVlansVMotion `field:"optional" json:"vMotion" yaml:"vMotion"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.49.0/docs/resources/evs_environment#v_san EvsEnvironment#v_san}.
	VSan *EvsEnvironmentInitialVlansVSan `field:"optional" json:"vSan" yaml:"vSan"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.49.0/docs/resources/evs_environment#v_tep EvsEnvironment#v_tep}.
	VTep *EvsEnvironmentInitialVlansVTep `field:"optional" json:"vTep" yaml:"vTep"`
}

