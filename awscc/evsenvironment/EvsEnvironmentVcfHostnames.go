package evsenvironment


type EvsEnvironmentVcfHostnames struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.49.0/docs/resources/evs_environment#cloud_builder EvsEnvironment#cloud_builder}.
	CloudBuilder *string `field:"required" json:"cloudBuilder" yaml:"cloudBuilder"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.49.0/docs/resources/evs_environment#nsx EvsEnvironment#nsx}.
	Nsx *string `field:"required" json:"nsx" yaml:"nsx"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.49.0/docs/resources/evs_environment#nsx_edge_1 EvsEnvironment#nsx_edge_1}.
	NsxEdge1 *string `field:"required" json:"nsxEdge1" yaml:"nsxEdge1"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.49.0/docs/resources/evs_environment#nsx_edge_2 EvsEnvironment#nsx_edge_2}.
	NsxEdge2 *string `field:"required" json:"nsxEdge2" yaml:"nsxEdge2"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.49.0/docs/resources/evs_environment#nsx_manager_1 EvsEnvironment#nsx_manager_1}.
	NsxManager1 *string `field:"required" json:"nsxManager1" yaml:"nsxManager1"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.49.0/docs/resources/evs_environment#nsx_manager_2 EvsEnvironment#nsx_manager_2}.
	NsxManager2 *string `field:"required" json:"nsxManager2" yaml:"nsxManager2"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.49.0/docs/resources/evs_environment#nsx_manager_3 EvsEnvironment#nsx_manager_3}.
	NsxManager3 *string `field:"required" json:"nsxManager3" yaml:"nsxManager3"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.49.0/docs/resources/evs_environment#sddc_manager EvsEnvironment#sddc_manager}.
	SddcManager *string `field:"required" json:"sddcManager" yaml:"sddcManager"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.49.0/docs/resources/evs_environment#v_center EvsEnvironment#v_center}.
	VCenter *string `field:"required" json:"vCenter" yaml:"vCenter"`
}

