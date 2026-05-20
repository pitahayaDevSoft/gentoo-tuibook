# Architecture and Decisions (ADRs)

## System Overview

A single-binary terminal user interface (TUI) that renders the Gentoo Linux installation handbook as embedded Markdown, using the Charm ecosystem (Bubbletea + Glamour + Lipgloss).

## ADRs

### ADR 0001: Flat `main` Package with Embedded Content

**Status**: Accepted
**Date**: 2026-05-18

#### Context

The app is a focused TUI reader — no server, no API, no multi-package logic. It needs to read Markdown files and display them in the terminal.

#### Decision

Keep a single `main` package with `//go:embed` for handbook content. No sub-packages, no internal module structure. The entire app is one file (`main.go`) with two supporting fetch scripts in `scripts/`.

#### Consequences

- **Positive**: Instant compilation, trivial dependency graph, single-binary deployment.
- **Negative**: Any other component that needs handbook content must be in the same package or duplicate the embed.
