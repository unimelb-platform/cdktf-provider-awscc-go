package groundstationconfig


type GroundstationConfigConfigDataAntennaUplinkConfigSpectrumConfig struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.49.0/docs/resources/groundstation_config#center_frequency GroundstationConfig#center_frequency}.
	CenterFrequency *GroundstationConfigConfigDataAntennaUplinkConfigSpectrumConfigCenterFrequency `field:"optional" json:"centerFrequency" yaml:"centerFrequency"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.49.0/docs/resources/groundstation_config#polarization GroundstationConfig#polarization}.
	Polarization *string `field:"optional" json:"polarization" yaml:"polarization"`
}

