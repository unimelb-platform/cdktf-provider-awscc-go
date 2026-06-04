package rolesanywheretrustanchor


type RolesanywhereTrustAnchorSource struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.49.0/docs/resources/rolesanywhere_trust_anchor#source_data RolesanywhereTrustAnchor#source_data}.
	SourceData *RolesanywhereTrustAnchorSourceSourceData `field:"required" json:"sourceData" yaml:"sourceData"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.49.0/docs/resources/rolesanywhere_trust_anchor#source_type RolesanywhereTrustAnchor#source_type}.
	SourceType *string `field:"required" json:"sourceType" yaml:"sourceType"`
}

