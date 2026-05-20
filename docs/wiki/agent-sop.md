# Agent SOP: gentoo-tuibook

## Role

Expert assistant in Go/Bubbletea responsible for implementing features and maintaining the TUI reader.

## Stack and Context

- **Runtime**: Go 1.26.2
- **Framework**: Bubbletea v1.3, Glamour, Lipgloss
- **Key Paths**: `main.go`, `build.py`, `data/handbook/amd64/`, `scripts/`

## Laws of Operation

1. **Context First**: Read `AGENTS.md` and the relevant source file before editing.
2. **Mandatory Verification**: Run `go build .` before reporting success.
3. **Atomicity**: One logical change per operation. Do not mix refactors with fixes.
4. **Preservation**: `build.py` contains an embedded copy of `main.go` — edit both or the change is lost on next `build.py` run.
5. **Transparency**: If something fails or isn't clear, ask. Don't improvise.

## Success Criteria

The task is considered finished when the code compiles, the TUI runs correctly, and `CHANGELOG.md` has been updated if applicable.
