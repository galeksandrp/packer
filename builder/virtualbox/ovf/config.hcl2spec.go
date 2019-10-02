// Code generated by "hcl2-schema -type Config"; DO NOT EDIT.\n

package ovf

import (
	"github.com/hashicorp/hcl2/hcldec"
	"github.com/zclconf/go-cty/cty"
)

func (*Config) HCL2Spec() map[string]hcldec.Spec {
	s := map[string]hcldec.Spec{
		"Checksum":                &hcldec.AttrSpec{Name: "checksum", Type: cty.String, Required: false},
		"ChecksumType":            &hcldec.AttrSpec{Name: "checksum_type", Type: cty.String, Required: false},
		"GuestAdditionsMode":      &hcldec.AttrSpec{Name: "guest_additions_mode", Type: cty.String, Required: false},
		"GuestAdditionsPath":      &hcldec.AttrSpec{Name: "guest_additions_path", Type: cty.String, Required: false},
		"GuestAdditionsInterface": &hcldec.AttrSpec{Name: "guest_additions_interface", Type: cty.String, Required: false},
		"GuestAdditionsSHA256":    &hcldec.AttrSpec{Name: "guest_additions_sha256", Type: cty.String, Required: false},
		"GuestAdditionsURL":       &hcldec.AttrSpec{Name: "guest_additions_url", Type: cty.String, Required: false},
		"ImportFlags":             &hcldec.AttrSpec{Name: "import_flags", Type: cty.List(cty.String), Required: false},
		"ImportOpts":              &hcldec.AttrSpec{Name: "import_opts", Type: cty.String, Required: false},
		"SourcePath":              &hcldec.AttrSpec{Name: "source_path", Type: cty.String, Required: false},
		"TargetPath":              &hcldec.AttrSpec{Name: "target_path", Type: cty.String, Required: false},
		"VMName":                  &hcldec.AttrSpec{Name: "vm_name", Type: cty.String, Required: false},
		"KeepRegistered":          &hcldec.AttrSpec{Name: "keep_registered", Type: cty.Bool, Required: false},
		"SkipExport":              &hcldec.AttrSpec{Name: "skip_export", Type: cty.Bool, Required: false},
	}
	for k, v := range (&Config{}).HTTPConfig.HCL2Spec() {
		s[k] = v
	}
	for k, v := range (&Config{}).FloppyConfig.HCL2Spec() {
		s[k] = v
	}
	for k, v := range (&Config{}).BootConfig.HCL2Spec() {
		s[k] = v
	}
	for k, v := range (&Config{}).ExportConfig.HCL2Spec() {
		s[k] = v
	}
	for k, v := range (&Config{}).OutputConfig.HCL2Spec() {
		s[k] = v
	}
	for k, v := range (&Config{}).RunConfig.HCL2Spec() {
		s[k] = v
	}
	for k, v := range (&Config{}).SSHConfig.HCL2Spec() {
		s[k] = v
	}
	for k, v := range (&Config{}).ShutdownConfig.HCL2Spec() {
		s[k] = v
	}
	for k, v := range (&Config{}).VBoxManageConfig.HCL2Spec() {
		s[k] = v
	}
	for k, v := range (&Config{}).VBoxVersionConfig.HCL2Spec() {
		s[k] = v
	}
	for k, v := range (&Config{}).GuestAdditionsConfig.HCL2Spec() {
		s[k] = v
	}
	return s
}
