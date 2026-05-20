<table border="0">
  <tr>
    <td width="200" align="center" valign="top">
      <img src="docs/assets/logo.svg" width="180" alt="gentoo-tuibook logo">
    </td>
    <td valign="top">
      <h1>gentoo-tuibook</h1>
      <p><strong>Terminal UI for the Gentoo Linux installation handbook</strong><br/>
      <em>Keyboard-driven reader that embeds the Gentoo handbook as Markdown, rendered via Bubbletea and Glamour.</em></p>
      <p>
        <a href="LICENSE"><img src="https://img.shields.io/badge/license-MIT-blue?style=plastic" alt="License MIT"></a>
        <img src="https://img.shields.io/badge/Built%20With-Go-00ADD8?style=plastic" alt="Built with Go">
      </p>
    </td>
  </tr>
</table>

---

<p align="center">
  <img src="docs/assets/screenshot.png" alt="gentoo-tuibook in action" width="900" />
</p>

---

<!--toc:start-->
- [gentoo-tuibook](#gentoo-tuibook)
  - [Overview](#overview)
  - [Using as a CLI Tool](#using-as-a-cli-tool)
  - [Installation](#installation)
  - [Key Features ](#key-features)
  - [Technical Architecture](#technical-architecture)
  - [Command Reference](#command-reference)
  - [Roadmap & Milestones](#roadmap--milestones)
  - [Acknowledgments](#acknowledgments)
  - [Contributing](#contributing)
  - [License](#license)
<!--toc:end-->

## Overview

**gentoo-tuibook** is a terminal-native reader for the Gentoo Linux installation handbook. No browser, no GUI — just the handbook rendered where Gentoo lives: in the terminal.

For command-line users, this project provides a powerful TUI interface.

At its heart lies **`github.com/gentoo-tuibook`**, a single-binary reader that embeds the full handbook content via `//go:embed`. Built upon **Bubbletea** and **Glamour**, it exposes a keyboard-driven split-pane reader with chapter selection and content rendering.

For end-users, gentoo-tuibook ships with **`gentoo-tuibook`**, a standalone binary with no external runtime dependencies.

---

<p align="center">
  <img src="docs/assets/demo.gif" alt="gentoo-tuibook demo" width="680" />
</p>

---

## Using as a CLI Tool

The CLI provides a straightforward interface to interact with gentoo-tuibook.

The CLI can be installed globally or used locally in your project.

```bash
go install github.com/gentoo-tuibook@latest
```

**Why use `gentoo-tuibook`?**
* **No browser required:** The handbook renders directly in your terminal, alongside your build session.
* **Offline-first:** All content is embedded in a single binary — no network calls after download.
* **Keyboard-driven:** Full navigation with familiar keybindings. No mouse needed.
* **Single binary:** Zero runtime dependencies. One file, everything included.

---

## Installation

### Via Go Install (Recommended)

```bash
go install github.com/gentoo-tuibook@latest
```

### From Source

```bash
git clone https://github.com/user/gentoo-tuibook.git
cd gentoo-tuibook
go build -o gentoo-tuibook.exe .
# Binary at: ./gentoo-tuibook.exe
```

---

## Key Features

*   **Handbook Browser**: Chapter list with split-pane content viewer — navigate chapters on the left, read on the right.
*   **Markdown Rendering**: Full Markdown rendering via Glamour with automatic terminal theme detection.
*   **Link Extraction**: Extracts and cycles through URLs in handbook content for quick browser access.
*   **Mode Switching**: Toggle between navigation mode (LISTA) and reading mode (READ) with link focus.
*   **Embedded Content**: Eleven handbook chapters embedded directly into the binary via `//go:embed`.
*   **Scriptable Content Fetch**: Two Go scripts in `scripts/` to pull fresh content from the Gentoo wiki API.
*   **Dual App Generations**: Alternate Green-themed version available via `build.py` codegen.

---

## Technical Architecture

A single-binary Go application — no server, no database, no runtime dependencies beyond the OS terminal.

```mermaid
graph TD
    User([User]) -->|Keyboard input| Cli(gentoo-tuibook TUI)

    subgraph gentoo-tuibook Application
        Cli -->|Model updates| Bubbletea(Bubbletea Engine)
        
        subgraph Core Loop
            Bubbletea --> Model[Application Model]
            Model --> View[Terminal View]
            View --> Bubbletea
        end

        Bubbletea -->|Renders| Content[Handbook Content]
        Bubbletea -.->|Optional| Browser(Browser for links)
    end

    Content --> FS[(Embedded MD Files)]
```

### Core Components

- **`github.com/gentoo-tuibook`**: Main `package main` with Bubbletea model/view/update loop, chapter list, viewport, link extraction, and Glamour-based Markdown rendering.
- **`build.py`**: Codegen script that regenerates `main.go` with an alternate (green-themed) version of the app.
- **`scripts/fetch.go` and `scripts/fetch_full.go`**: Standalone Go scripts that pull HTML from the Gentoo wiki API and convert to Markdown.

---

### Key Engineering Decisions

- **Flat `main` package over multi-package layout:** No server, no API, no multi-package logic justifies a single `package main`. Keeps compilation trivial and dependency graph minimal.
- **Embedded content over HTTP fetch at runtime:** The handbook changes slowly. Embedding makes the binary self-contained and offline-ready.
- **`build.py` codegen over feature flags:** Two UI generations coexist; the codegen script replaces `main.go` wholesale rather than adding conditional branches.

---

## Command Reference

For a comprehensive breakdown, see the **[Official Docs](docs/wiki/index.md)**.

*   **[Installation Guide](docs/wiki/development.md)**
*   **[CLI Command Reference](docs/wiki/index.md)**
*   **[Technical Architecture](docs/wiki/architecture.md)**

### Keyboard Shortcuts

| Key | Mode | Action |
| :--- | :--- | :--- |
| `↑/↓` | Navigation | Move chapter selection |
| `Enter` | Navigation | Open selected chapter |
| `r` | Reading | Enter read mode (link focus) |
| `Tab` | Reading | Cycle through extracted links |
| `Esc` | Both | Return to navigation mode |
| `q` / `Ctrl+C` | Both | Quit application |

---

## Roadmap & Milestones

| Version | Status | Milestone |
| --- | --- | --- |
| **v0.1.0** | ✅ | Initial TUI with handbook browser and Markdown rendering |
| **v0.2.0** | ⏳ | Search, bookmarks, configurable theme |
| **v0.3.0** | ⏳ | Multi-architecture support (ARM64, multi-page display) |

---

## Acknowledgments

- **[Charmbracelet](https://github.com/charmbracelet)** — Bubbletea, Glamour, Lipgloss, and the entire Charm ecosystem.
- **[Gentoo Wiki](https://wiki.gentoo.org)** — Source of the handbook content.

## Contributing

Pull requests are welcome. For major changes, open an RFC issue first. See `CONTRIBUTING.md` for guidelines.

## License

<p align="center">
  Engineered by the gentoo-tuibook contributors.<br>
  Released under the terms of the MIT License.
</p>
