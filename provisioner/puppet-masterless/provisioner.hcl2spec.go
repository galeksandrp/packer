// Code generated by "hcl2-schema -type Config"; DO NOT EDIT.\n

package puppetmasterless

import (
	"github.com/hashicorp/hcl2/hcldec"
	"github.com/zclconf/go-cty/cty"
)

func (*Config) HCL2Spec() map[string]hcldec.Spec {
	s := map[string]hcldec.Spec{
		"CleanStagingDir":  &hcldec.AttrSpec{Name: "clean_staging_directory", Type: cty.Bool, Required: false},
		"GuestOSType":      &hcldec.AttrSpec{Name: "guest_os_type", Type: cty.String, Required: false},
		"ExecuteCommand":   &hcldec.AttrSpec{Name: "execute_command", Type: cty.String, Required: false},
		"ExtraArguments":   &hcldec.AttrSpec{Name: "extra_arguments", Type: cty.List(cty.String), Required: false},
		"Facter":           &hcldec.BlockAttrsSpec{TypeName: "facter", ElementType: cty.String, Required: false},
		"HieraConfigPath":  &hcldec.AttrSpec{Name: "hiera_config_path", Type: cty.String, Required: false},
		"IgnoreExitCodes":  &hcldec.AttrSpec{Name: "ignore_exit_codes", Type: cty.Bool, Required: false},
		"ModulePaths":      &hcldec.AttrSpec{Name: "module_paths", Type: cty.List(cty.String), Required: false},
		"ManifestFile":     &hcldec.AttrSpec{Name: "manifest_file", Type: cty.String, Required: false},
		"ManifestDir":      &hcldec.AttrSpec{Name: "manifest_dir", Type: cty.String, Required: false},
		"PreventSudo":      &hcldec.AttrSpec{Name: "prevent_sudo", Type: cty.Bool, Required: false},
		"PuppetBinDir":     &hcldec.AttrSpec{Name: "puppet_bin_dir", Type: cty.String, Required: false},
		"StagingDir":       &hcldec.AttrSpec{Name: "staging_directory", Type: cty.String, Required: false},
		"WorkingDir":       &hcldec.AttrSpec{Name: "working_directory", Type: cty.String, Required: false},
		"ElevatedUser":     &hcldec.AttrSpec{Name: "elevated_user", Type: cty.String, Required: false},
		"ElevatedPassword": &hcldec.AttrSpec{Name: "elevated_password", Type: cty.String, Required: false},
	}
	return s
}
