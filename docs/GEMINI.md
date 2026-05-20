# Gemini CLI Rules

Specific instructions for the **Gemini CLI** agent.

## Working Context

- This project follows the Jules Dev Standard.
- Go 1.26.2 single-module project using Bubbletea, Glamour, Lipgloss.
- Content embedded via `//go:embed`. Fetch scripts at `scripts/`.

## Workflow

1. Research -> 2. Strategy -> 3. Execution (Plan-Act-Validate).

## Restrictions

- Do not modify files in `docs/` unless documenting a feature.
- Keep `CHANGELOG.md` updated with every `feat` or `fix`.
- `build.py` regenerates `main.go` — edits to one must be mirrored in the other.
