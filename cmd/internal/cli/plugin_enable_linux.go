package cli

import (
	"os"

	"github.com/spf13/cobra"
	"github.com/sylabs/singularity/docs"
	"github.com/sylabs/singularity/internal/app/singularity"
	"github.com/sylabs/singularity/internal/pkg/buildcfg"
	"github.com/sylabs/singularity/internal/pkg/sylog"
)

// PluginEnableCmd enables the named plugin
//
// singularity plugin enable <name>
var PluginEnableCmd = &cobra.Command{
	Run: func(cmd *cobra.Command, args []string) {
		err := singularity.EnablePlugin(args[0], buildcfg.LIBEXECDIR)
		if err != nil {
			if os.IsNotExist(err) {
				sylog.Errorf("Failed to enable plugin %q: plugin not found.", args[0])
			} else {
				sylog.Errorf("Failed to enable plugin %q: %s.", args[0], err)
			}
		}
	},
	DisableFlagsInUseLine: true,
	Args:                  cobra.ExactArgs(1),

	Use:     docs.PluginEnableUse,
	Short:   docs.PluginEnableShort,
	Long:    docs.PluginEnableLong,
	Example: docs.PluginEnableExample,
}
