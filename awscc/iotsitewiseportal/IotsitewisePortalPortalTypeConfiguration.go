package iotsitewiseportal


type IotsitewisePortalPortalTypeConfiguration struct {
	// List of enabled Tools for a certain portal.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.49.0/docs/resources/iotsitewise_portal#portal_tools IotsitewisePortal#portal_tools}
	PortalTools *[]*string `field:"optional" json:"portalTools" yaml:"portalTools"`
}

