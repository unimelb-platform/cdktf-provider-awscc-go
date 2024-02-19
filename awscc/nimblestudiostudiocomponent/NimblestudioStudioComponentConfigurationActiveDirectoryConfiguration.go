package nimblestudiostudiocomponent


type NimblestudioStudioComponentConfigurationActiveDirectoryConfiguration struct {
	// <p>A collection of custom attributes for an Active Directory computer.</p>.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/nimblestudio_studio_component#computer_attributes NimblestudioStudioComponent#computer_attributes}
	ComputerAttributes interface{} `field:"optional" json:"computerAttributes" yaml:"computerAttributes"`
	// <p>The directory ID of the Directory Service for Microsoft Active Directory to access             using this studio component.</p>.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/nimblestudio_studio_component#directory_id NimblestudioStudioComponent#directory_id}
	DirectoryId *string `field:"optional" json:"directoryId" yaml:"directoryId"`
	// <p>The distinguished name (DN) and organizational unit (OU) of an Active Directory             computer.</p>.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/nimblestudio_studio_component#organizational_unit_distinguished_name NimblestudioStudioComponent#organizational_unit_distinguished_name}
	OrganizationalUnitDistinguishedName *string `field:"optional" json:"organizationalUnitDistinguishedName" yaml:"organizationalUnitDistinguishedName"`
}

