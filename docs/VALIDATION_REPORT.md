# NTM Command Validation Report

- **Task**: ntm-wqg.6 — End-to-end command validation
- **Branch**: main
- **Binary built**: `dist/ntm` (via `go build -o dist/ntm ./cmd/ntm`)
- **Date**: 2026-02-22
- **Scope**: README.md, AGENTS.md, SKILL.md, command_palette.md

---

## 1. Binary Build

Build succeeded without errors:

```bash
go build -o dist/ntm ./cmd/ntm
```

---

## 2. Complete List of Registered CLI Commands

The following commands are registered in the binary (from `--help` output). Commands marked `[HIDDEN]` do not appear in help but exist as subcommands.

```
activity, add, adopt, agents, analytics, approve, assign, attach,
audit, beads, bind, bugs, cass, changes, checkpoint, cleanup,
completion, config, conflicts, context, controller, coordinator,
copy, create, dashboard, deps, diff, doctor, ensemble, errors,
extract, get-all-session-text, git, grep, guards, handoff, health,
help, history, hooks, init, interrupt, kernel, kill, level, list,
lock, locks, logs, mail, memory, message, metrics, modes, openapi,
palette, personas, pipeline, plugins, policy, preflight, profile,
profiles, quick, quota, rebalance, recipes, redact, replay, repo,
respawn, resume, review-queue, rollback, rotate, safety, save,
scale, scan, scrub, search, send, serve, session-templates, sessions,
setup, shell, spawn, status, summary, support-bundle, swarm,
template, tutorial, unlock, upgrade, version, view, wait, watch,
work, workflows, worktree, worktrees, zoom

[HIDDEN] internal-monitor   (registered as "internal-monitor", NOT "monitor")
```

---

## 3. Documented Command Validation

### 3.1 SKILL.md Commands

All commands documented in SKILL.md were validated. Every command listed below **EXISTS** in the binary and responds correctly to `--help`:

| Documented Command | Status | Notes |
|--------------------|--------|-------|
| `ntm spawn` | EXISTS | |
| `ntm quick` | EXISTS | |
| `ntm create` | EXISTS | |
| `ntm send` | EXISTS | |
| `ntm add` | EXISTS | |
| `ntm interrupt` | EXISTS | |
| `ntm list` | EXISTS | Alias: `lnt` |
| `ntm attach` | EXISTS | Alias: `rnt` |
| `ntm status` | EXISTS | Alias: `snt` |
| `ntm view` | EXISTS | Alias: `vnt` |
| `ntm zoom` | EXISTS | Alias: `znt` |
| `ntm dashboard` | EXISTS | Alias: `dash`, `d` |
| `ntm kill` | EXISTS | Alias: `knt` |
| `ntm palette` | EXISTS | Alias: `ncp` |
| `ntm bind` | EXISTS | |
| `ntm copy` | EXISTS | Alias: `cpnt` |
| `ntm save` | EXISTS | Alias: `svnt` |
| `ntm activity` | EXISTS | |
| `ntm health` | EXISTS | |
| `ntm watch` | EXISTS | |
| `ntm extract` | EXISTS | |
| `ntm diff` | EXISTS | |
| `ntm grep` | EXISTS | |
| `ntm analytics` | EXISTS | |
| `ntm locks` | EXISTS | |
| `ntm checkpoint` | EXISTS | |
| `ntm serve` | EXISTS | Port 7337 |
| `ntm beads` | EXISTS | |
| `ntm work` | EXISTS | |
| `ntm work triage` | EXISTS | |
| `ntm work alerts` | EXISTS | |
| `ntm work search` | EXISTS | |
| `ntm work impact` | EXISTS | |
| `ntm work next` | EXISTS | |
| `ntm profiles` | EXISTS | Alias for `personas` |
| `ntm profiles list` | EXISTS | |
| `ntm profiles show` | EXISTS | |
| `ntm mail` | EXISTS | |
| `ntm hooks` | EXISTS | |
| `ntm hooks guard install` | EXISTS | |
| `ntm hooks guard uninstall` | EXISTS | |
| `ntm safety` | EXISTS | |
| `ntm config` | EXISTS | |
| `ntm config init` | EXISTS | |
| `ntm config show` | EXISTS | |
| `ntm config project init` | EXISTS | |
| `ntm upgrade` | EXISTS | |
| `ntm tutorial` | EXISTS | |
| `ntm deps` | EXISTS | |
| `ntm shell` | EXISTS | |
| `ntm init` | EXISTS | |

### 3.2 Robot Mode Flags (from SKILL.md and AGENTS.md)

All `--robot-*` flags documented in AGENTS.md and SKILL.md exist as global flags on the root command. Key flags verified:

| Flag | Status |
|------|--------|
| `--robot-status` | EXISTS |
| `--robot-context=SESSION` | EXISTS |
| `--robot-snapshot` | EXISTS |
| `--robot-tail=SESSION` | EXISTS |
| `--robot-plan` | EXISTS |
| `--robot-dashboard` | EXISTS |
| `--robot-terse` | EXISTS |
| `--robot-send=SESSION` | EXISTS |
| `--robot-spawn=SESSION` | EXISTS |
| `--robot-interrupt=SESSION` | EXISTS |
| `--robot-assign=SESSION` | EXISTS |
| `--robot-bead-claim=BEAD_ID` | EXISTS |
| `--robot-bead-create` | EXISTS |
| `--robot-bead-show=BEAD_ID` | EXISTS |
| `--robot-bead-close=BEAD_ID` | EXISTS |
| `--robot-cass-search=QUERY` | EXISTS |
| `--robot-cass-status` | EXISTS |
| `--robot-cass-context=QUERY` | EXISTS |
| `--robot-monitor=SESSION` | EXISTS |
| `--robot-help` | EXISTS |
| `--robot-capabilities` | EXISTS |

### 3.3 Beads Daemon Subcommands

All five daemon subcommands exist in the binary and return the correct intentional error message:

```
Error: bd does not support daemon mode; use 'bd sync' to sync issue state manually
```

| Subcommand | Status | Return |
|------------|--------|--------|
| `ntm beads daemon start` | EXISTS | Correct error, exit 1 |
| `ntm beads daemon stop` | EXISTS | Correct error, exit 1 |
| `ntm beads daemon status` | EXISTS | Correct error, exit 1 |
| `ntm beads daemon health` | EXISTS | Correct error, exit 1 |
| `ntm beads daemon metrics` | EXISTS | Correct error, exit 1 |

This behavior is **intentional and correctly documented** in AGENTS.md (section "ntm beads — Issue Sync (No Daemon Mode)").

---

## 4. Root Cause Analysis: boot / sat / bp / bpf / monitor / session

Worker C reported `ntm boot`, `ntm sat`, `ntm bp`, `ntm bpf`, `ntm monitor`, and `ntm session` return "unknown command." This section documents the root cause for each.

### 4.1 `ntm sat` and `ntm bp` — Shell Aliases, Not CLI Subcommands

**Root cause**: `sat` and `bp` are **shell aliases** injected by `ntm shell zsh` (or bash/fish), NOT subcommands registered in the ntm binary.

From `ntm shell zsh` output:
```bash
alias sat='ntm spawn'    # sat → ntm spawn
alias bp='ntm send'      # bp → ntm send
alias ant='ntm add'      # ant → ntm add
alias int='ntm interrupt' # int → ntm interrupt
alias cnt='ntm create'   # cnt → ntm create
alias rnt='ntm attach'   # rnt → ntm attach
alias lnt='ntm list'     # lnt → ntm list
alias snt='ntm status'   # snt → ntm status
alias vnt='ntm view'     # vnt → ntm view
alias znt='ntm zoom'     # znt → ntm zoom
alias ncp='ntm palette'  # ncp → ntm palette
alias knt='ntm kill'     # knt → ntm kill
```

These aliases are **correctly documented** in README.md (Shell Aliases table) and SKILL.md (Shell Aliases section). They function correctly after `eval "$(ntm shell zsh)"` is added to `.zshrc`.

**Verdict**: NOT a bug. NOT a documentation error. Workers testing these by running `./dist/ntm sat` were testing the wrong way — the aliases only work in a shell with integration loaded.

### 4.2 `ntm boot` — Does Not Exist Anywhere

**Root cause**: `boot` is not a registered CLI subcommand, not a shell alias, and not referenced in any current documentation file (README.md, AGENTS.md, SKILL.md, command_palette.md).

A search across all `.go` files in `internal/cli/` finds no `boot` command registration. A search across all `.md` files in the repo root finds no documentation of `ntm boot` as a CLI command.

**Verdict**: `boot` was either a planned command that was never implemented, or a misread from an earlier draft. It does not exist and is not documented. **No doc correction needed** (it is not referenced in current docs).

### 4.3 `ntm bpf` — Does Not Exist Anywhere

**Root cause**: `bpf` is not a registered CLI subcommand, not a shell alias, and appears nowhere in `internal/cli/*.go` source files or any current documentation.

**Verdict**: `bpf` does not exist and is not documented anywhere in current docs. **No doc correction needed**.

### 4.4 `ntm monitor` — Hidden Internal Command, Not a User-Facing Subcommand

**Root cause**: `monitor` is implemented in `internal/cli/monitor.go` but registered with the Use field `"internal-monitor <session>"` and `Hidden: true`. It is intended for internal use by the resilience subsystem, not for direct user invocation.

From `internal/cli/monitor.go:28-38`:
```go
func newMonitorCmd() *cobra.Command {
    return &cobra.Command{
        Use:    "internal-monitor <session>",
        Short:  "Run the resilience monitor for a session (internal use)",
        Hidden: true,
        ...
    }
}
```

The monitoring capability is exposed to agents/robots via `--robot-monitor=SESSION` (a global flag, not a subcommand), which **does** exist and function correctly.

**Verdict**: `ntm monitor` does not exist as a user-facing subcommand. The monitoring feature is correctly accessed via `ntm --robot-monitor=SESSION`. Not referenced in user-facing docs as `ntm monitor`. **No doc correction needed**.

### 4.5 `ntm session` — Does Not Exist (Correct Command is `ntm sessions`)

**Root cause**: The CLI registers `sessions` (plural), not `session` (singular). `ntm session` returns "unknown command."

`ntm sessions` **does** exist and has full subcommands:
- `ntm sessions save`
- `ntm sessions list`
- `ntm sessions show`
- `ntm sessions restore`
- `ntm sessions delete`

README.md correctly uses `ntm sessions` (plural) throughout. AGENTS.md does not reference `ntm session` or `ntm sessions`.

**Verdict**: The command is `ntm sessions` (plural). `ntm session` (singular) is not valid and is not documented anywhere. **No doc correction needed**.

---

## 5. Summary: Documentation Accuracy

All commands documented in README.md, AGENTS.md, SKILL.md, and command_palette.md were validated against the binary.

| Document | Verdict |
|----------|---------|
| README.md | All documented commands exist. Shell aliases (`sat`, `bp`, etc.) are correctly labeled as aliases, not CLI commands. |
| AGENTS.md | All documented commands exist. Beads daemon behavior is correctly documented. |
| SKILL.md | All documented commands exist. Shell aliases section is correct. |
| command_palette.md | References `ntm ensemble`, `ntm modes`, etc. — all exist. |

**No documentation corrections are required** as a result of this validation.

The commands `boot`, `bpf`, `monitor` (as subcommand), and `session` (singular) are not referenced in any of the validated documentation files. They were never documented as user-facing CLI commands.

---

## 6. Findings Not in Docs (Undocumented Commands)

The following commands exist in the binary but are not documented in README.md or SKILL.md:

| Command | Brief Description |
|---------|------------------|
| `ntm adopt` | Adopt external tmux session for NTM management |
| `ntm approve` | Manage approval requests for dangerous operations |
| `ntm controller` | Launch dedicated controller agent in pane 1 |
| `ntm coordinator` | Manage session coordination for multi-agent workflows |
| `ntm get-all-session-text` | Get text from all panes across all sessions |
| `ntm guards` | Manage Agent Mail pre-commit guards |
| `ntm kernel` | Inspect command kernel registry |
| `ntm level` | View/change CLI proficiency tier |
| `ntm lock` | Reserve files via Agent Mail |
| `ntm memory` | Interact with CASS Memory (cm) system |
| `ntm message` | Unified messaging (Agent Mail + BD) |
| `ntm modes` | Browse reasoning modes |
| `ntm openapi` | OpenAPI specification management |
| `ntm persona`/`ntm personas` | Manage agent personas |
| `ntm policy` | Manage NTM policy configuration |
| `ntm rebalance` | Workload distribution analysis |
| `ntm recipes` | Manage session recipes/presets |
| `ntm repo` | Repo management commands |
| `ntm review-queue` | List idle agents and suggest review prompts |
| `ntm scale` | Scale agents to target counts |
| `ntm scrub` | Scan for leaked secrets |
| `ntm swarm` | Multi-project agent swarm orchestration |
| `ntm template` | Manage prompt templates |
| `ntm unlock` | Release file reservations |
| `ntm worktree`/`ntm worktrees` | Git worktree management |

These are present and functional but not covered in SKILL.md or README.md. Filing documentation is out of scope for this validation task.

---

## 7. Validation Conclusion

- **Binary build**: PASS
- **All documented commands**: EXIST
- **Beads daemon subcommands**: Return correct intentional error
- **boot/bpf**: Not CLI commands, not in docs — no action needed
- **sat/bp**: Shell aliases (not CLI subcommands) — correctly documented
- **monitor**: Hidden internal command — user-facing feature is `--robot-monitor`
- **session (singular)**: Correct command is `sessions` (plural) — docs use plural correctly
- **Documentation corrections applied**: None required
