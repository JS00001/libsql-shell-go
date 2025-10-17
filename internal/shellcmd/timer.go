package shellcmd

import (
	"fmt"
	"strings"

	"github.com/JS00001/libsql-shell-go/pkg/shell/enums"
	"github.com/spf13/cobra"
)

var timerCmd = &cobra.Command{
	Use:   ".timer",
	Short: "Toggle query timing on or off",
	Args:  cobra.MaximumNArgs(1),
	ValidArgs: []string{
		string(enums.ON),
		string(enums.OFF),
	},
	RunE: func(cmd *cobra.Command, args []string) error {
		validTimerArgs := strings.Join(cmd.ValidArgs, ", ")
		config, ok := cmd.Context().Value(dbCtx{}).(*DbCmdConfig)

		if !ok {
			return fmt.Errorf("missing db connection")
		}

		currentTimer := config.GetTimer()
		if len(args) == 0 {
			return fmt.Errorf("No args provided. Current timer is %s. Valid args are %s", currentTimer, validTimerArgs)
		}

		mode := args[0]
		switch mode {
		case string(enums.OFF):
			config.SetTimer(enums.OFF)
		case string(enums.ON):
			config.SetTimer(enums.ON)
		default:
			return fmt.Errorf("Invalid args. Current timer is %s. Valid args are %s", currentTimer, validTimerArgs)
		}

		return nil
	},
}
