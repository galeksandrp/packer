// Code generated by "hcl2-schema -type AccessConfig,VaultAWSEngineOptions"; DO NOT EDIT.\n

package common

import (
	"github.com/hashicorp/hcl2/hcldec"
	"github.com/zclconf/go-cty/cty"
)

func (*AccessConfig) HCL2Spec() map[string]hcldec.Spec {
	s := map[string]hcldec.Spec{
		"AccessKey":             &hcldec.AttrSpec{Name: "access_key", Type: cty.String, Required: false},
		"CustomEndpointEc2":     &hcldec.AttrSpec{Name: "custom_endpoint_ec2", Type: cty.String, Required: false},
		"DecodeAuthZMessages":   &hcldec.AttrSpec{Name: "decode_authorization_messages", Type: cty.Bool, Required: false},
		"InsecureSkipTLSVerify": &hcldec.AttrSpec{Name: "insecure_skip_tls_verify", Type: cty.Bool, Required: false},
		"MFACode":               &hcldec.AttrSpec{Name: "mfa_code", Type: cty.String, Required: false},
		"ProfileName":           &hcldec.AttrSpec{Name: "profile", Type: cty.String, Required: false},
		"RawRegion":             &hcldec.AttrSpec{Name: "region", Type: cty.String, Required: false},
		"SecretKey":             &hcldec.AttrSpec{Name: "secret_key", Type: cty.String, Required: false},
		"SkipValidation":        &hcldec.AttrSpec{Name: "skip_region_validation", Type: cty.Bool, Required: false},
		"SkipMetadataApiCheck":  &hcldec.AttrSpec{Name: "skip_metadata_api_check", Type: cty.Bool, Required: false},
		"Token":                 &hcldec.AttrSpec{Name: "token", Type: cty.String, Required: false},
		"vault_aws_engine":      &hcldec.BlockObjectSpec{TypeName: "VaultAWSEngineOptions", LabelNames: []string(nil), Nested: hcldec.ObjectSpec((&AccessConfig{}).VaultAWSEngine.HCL2Spec())},
	}
	return s
}

func (*VaultAWSEngineOptions) HCL2Spec() map[string]hcldec.Spec {
	s := map[string]hcldec.Spec{
		"Name":       &hcldec.AttrSpec{Name: "name", Type: cty.String, Required: false},
		"RoleARN":    &hcldec.AttrSpec{Name: "role_arn", Type: cty.String, Required: false},
		"TTL":        &hcldec.AttrSpec{Name: "ttl", Type: cty.String, Required: false},
		"EngineName": &hcldec.AttrSpec{Name: "engine_name", Type: cty.String, Required: false},
	}
	return s
}
