package groundstationconfig


type GroundstationConfigConfigDataAntennaDownlinkDemodDecodeConfig struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.49.0/docs/resources/groundstation_config#decode_config GroundstationConfig#decode_config}.
	DecodeConfig *GroundstationConfigConfigDataAntennaDownlinkDemodDecodeConfigDecodeConfig `field:"optional" json:"decodeConfig" yaml:"decodeConfig"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.49.0/docs/resources/groundstation_config#demodulation_config GroundstationConfig#demodulation_config}.
	DemodulationConfig *GroundstationConfigConfigDataAntennaDownlinkDemodDecodeConfigDemodulationConfig `field:"optional" json:"demodulationConfig" yaml:"demodulationConfig"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.49.0/docs/resources/groundstation_config#spectrum_config GroundstationConfig#spectrum_config}.
	SpectrumConfig *GroundstationConfigConfigDataAntennaDownlinkDemodDecodeConfigSpectrumConfig `field:"optional" json:"spectrumConfig" yaml:"spectrumConfig"`
}

