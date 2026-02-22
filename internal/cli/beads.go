package cli

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

func newBeadsCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "beads",
		Short: "Manage the beads (bd) daemon for issue tracking sync",
		Long: `Manage the beads daemon that automatically syncs issues with git remote.

The daemon handles:
  - Automatic export of database changes to JSONL
  - Auto-commit and push when configured
  - Pull and import of remote changes
  - Health monitoring and auto-restart

Examples:
  ntm beads daemon start           # Start BD daemon for current session
  ntm beads daemon stop            # Stop BD daemon
  ntm beads daemon status          # Show daemon status
  ntm beads daemon health          # Check daemon health
  ntm beads daemon metrics         # Show detailed metrics`,
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
		Long: `Start the BD daemon for automatic issue sync.

When running within an NTM session, the daemon is managed by the supervisor
with automatic health monitoring and restart capability.

For standalone use, run 'bd daemon --start' directly.`,
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

// startBDWithSupervisor is a no-op stub: bd does not support daemon mode.
func startBDWithSupervisor(projectDir, sessionID string, autoCommit, autoPush bool, interval string) error {
	return fmt.Errorf("bd does not support daemon mode; use 'bd sync' to sync issue state manually")
}

// stopBDWithSupervisor is a no-op stub: bd does not support daemon mode.
func stopBDWithSupervisor(projectDir, sessionID string) error {
	return fmt.Errorf("bd does not support daemon mode; use 'bd sync' to sync issue state manually")
}

// showBDSupervisorStatus is a no-op stub: bd does not support daemon mode.
func showBDSupervisorStatus(projectDir, sessionID string) error {
	return fmt.Errorf("bd does not support daemon mode; use 'bd sync' to sync issue state manually")
}

// startBDDirect is a no-op stub: bd does not support daemon mode.
func startBDDirect(projectDir string, autoCommit, autoPush bool, interval string, foreground bool) error {
	return fmt.Errorf("bd does not support daemon mode; use 'bd sync' to sync issue state manually")
}

// runBDCommand is a no-op stub: bd does not support daemon mode.
func runBDCommand(projectDir string, args ...string) error {
	return fmt.Errorf("bd does not support daemon mode; use 'bd sync' to sync issue state manually")
}
