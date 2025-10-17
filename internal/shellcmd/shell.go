package shellcmd

import (
	"fmt"
	"os/exec"
	"strings"

	"github.com/spf13/cobra"
)

var shellCmd = &cobra.Command{
	Use:   ".shell",
	Short: `Run a shell command in the terminal`,
	Args:  cobra.ArbitraryArgs,
	RunE: func(cmd *cobra.Command, args []string) error {
		_, ok := cmd.Context().Value(dbCtx{}).(*DbCmdConfig)
		if !ok {
			return fmt.Errorf("missing db connection")
		}

		cmd_input := strings.Join(args, " ")
		shell_cmd := exec.Command(cmd_input)

		if err := shell_cmd.Run(); err != nil {
			return nil
		}

		return nil
	},
}
