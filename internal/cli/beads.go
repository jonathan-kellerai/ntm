package cli

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

func newBeadsCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "beads",
		Short: "Manage the beads (bd) issue tracking integration",
		Long: `Manage the beads (bd) issue tracking integration.

NOTE: bd does not support a daemon mode. All subcommands under 'ntm beads daemon'
return an informative error. To sync issue state with the git remote, run:

  bd sync

Examples:
  ntm beads daemon status          # Returns error explaining bd has no daemon mode
  ntm beads daemon health          # Returns error explaining bd has no daemon mode`,
	}

	cmd.AddCommand(newBeadsDaemonCmd())

	return cmd
}

func newBeadsDaemonCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "daemon",
		Short: "Manage the BD daemon lifecycle",
	}

	cmd.AddCommand(
		newBeadsDaemonStartCmd(),
		newBeadsDaemonStopCmd(),
		newBeadsDaemonStatusCmd(),
		newBeadsDaemonHealthCmd(),
		newBeadsDaemonMetricsCmd(),
	)

	return cmd
}

func newBeadsDaemonStartCmd() *cobra.Command {
	var (
		sessionID  string
		autoCommit bool
		autoPush   bool
		interval   string
		foreground bool
	)

	cmd := &cobra.Command{
		Use:   "start",
		Short: "Start the BD daemon",
		Long: `Attempt to start the BD daemon — this command always returns an error.

bd does not support daemon mode. There is no 'bd daemon --start' command.
To sync issue state with the git remote, run:

  bd sync`,
		RunE: func(cmd *cobra.Command, args []string) error {
			projectDir, err := os.Getwd()
			if err != nil {
				return fmt.Errorf("get working directory: %w", err)
			}

			// If session specified, use supervisor
			if sessionID != "" {
				return startBDWithSupervisor(projectDir, sessionID, autoCommit, autoPush, interval)
			}

			// Otherwise, run bd daemon directly
			return startBDDirect(projectDir, autoCommit, autoPush, interval, foreground)
		},
	}

	cmd.Flags().StringVar(&sessionID, "session", "", "NTM session ID (uses supervisor)")
	cmd.Flags().BoolVar(&autoCommit, "auto-commit", true, "Automatically commit changes")
	cmd.Flags().BoolVar(&autoPush, "auto-push", false, "Automatically push commits (requires policy approval)")
	cmd.Flags().StringVar(&interval, "interval", "5s", "Sync check interval")
	cmd.Flags().BoolVar(&foreground, "foreground", false, "Run in foreground (standalone mode only)")

	return cmd
}

func newBeadsDaemonStopCmd() *cobra.Command {
	var sessionID string

	cmd := &cobra.Command{
		Use:   "stop",
		Short: "Stop the BD daemon",
		RunE: func(cmd *cobra.Command, args []string) error {
			projectDir, err := os.Getwd()
			if err != nil {
				return fmt.Errorf("get working directory: %w", err)
			}

			// If session specified, use supervisor
			if sessionID != "" {
				return stopBDWithSupervisor(projectDir, sessionID)
			}

			// Otherwise, run bd daemon --stop directly
			return runBDCommand(projectDir, "--stop")
		},
	}

	cmd.Flags().StringVar(&sessionID, "session", "", "NTM session ID (uses supervisor)")

	return cmd
}

func newBeadsDaemonStatusCmd() *cobra.Command {
	var sessionID string

	cmd := &cobra.Command{
		Use:   "status",
		Short: "Show BD daemon status",
		RunE: func(cmd *cobra.Command, args []string) error {
			projectDir, err := os.Getwd()
			if err != nil {
				return fmt.Errorf("get working directory: %w", err)
			}

			// If session specified, check supervisor status
			if sessionID != "" {
				return showBDSupervisorStatus(projectDir, sessionID)
			}

			// Otherwise, run bd daemon --status
			return runBDCommand(projectDir, "--status", "--json")
		},
	}

	cmd.Flags().StringVar(&sessionID, "session", "", "NTM session ID (uses supervisor)")

	return cmd
}

func newBeadsDaemonHealthCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "health",
		Short: "Check BD daemon health",
		RunE: func(cmd *cobra.Command, args []string) error {
			projectDir, err := os.Getwd()
			if err != nil {
				return fmt.Errorf("get working directory: %w", err)
			}

			// Run bd daemon health
			return runBDCommand(projectDir, "health", "--json")
		},
	}

	return cmd
}

func newBeadsDaemonMetricsCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "metrics",
		Short: "Show BD daemon metrics",
		RunE: func(cmd *cobra.Command, args []string) error {
			projectDir, err := os.Getwd()
			if err != nil {
				return fmt.Errorf("get working directory: %w", err)
			}

			// Run bd daemon --metrics
			return runBDCommand(projectDir, "--metrics", "--json")
		},
	}

	return cmd
}

// startBDWithSupervisor is a stub that surfaces a clear error when a caller
// attempts to start bd via the NTM supervisor. bd has no daemon mode, so
// rather than silently failing or panicking, this function returns an
// actionable error directing the user to 'bd sync'.
func startBDWithSupervisor(projectDir, sessionID string, autoCommit, autoPush bool, interval string) error {
	return fmt.Errorf("bd does not support daemon mode; use 'bd sync' to sync issue state manually")
}

// stopBDWithSupervisor is a stub that surfaces a clear error when a caller
// attempts to stop a bd daemon via the NTM supervisor. bd has no daemon mode,
// so no process to stop; the error guides users to 'bd sync'.
func stopBDWithSupervisor(projectDir, sessionID string) error {
	return fmt.Errorf("bd does not support daemon mode; use 'bd sync' to sync issue state manually")
}

// showBDSupervisorStatus is a stub that surfaces a clear error when a caller
// requests bd daemon status from the NTM supervisor. bd has no daemon mode,
// so there is no status to report; the error guides users to 'bd sync'.
func showBDSupervisorStatus(projectDir, sessionID string) error {
	return fmt.Errorf("bd does not support daemon mode; use 'bd sync' to sync issue state manually")
}

// startBDDirect is a stub that surfaces a clear error when a caller
// attempts to start bd directly (standalone mode). bd has no daemon mode,
// so the error guides users to 'bd sync'.
func startBDDirect(projectDir string, autoCommit, autoPush bool, interval string, foreground bool) error {
	return fmt.Errorf("bd does not support daemon mode; use 'bd sync' to sync issue state manually")
}

// runBDCommand is a stub that surfaces a clear error for any bd daemon
// subcommand (stop, status, health, metrics). bd has no daemon mode; these
// subcommands are wired up in the CLI for discoverability but always return
// an informative error directing users to 'bd sync'.
func runBDCommand(projectDir string, args ...string) error {
	return fmt.Errorf("bd does not support daemon mode; use 'bd sync' to sync issue state manually")
}
