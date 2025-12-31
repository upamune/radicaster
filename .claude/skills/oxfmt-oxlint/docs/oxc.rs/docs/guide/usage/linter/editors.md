---
title: "Setup editors | The JavaScript Oxidation Compiler"
source_url: "https://oxc.rs/docs/guide/usage/linter/editors"
fetched_at: "2025-12-31T10:44:14.234719+00:00"
---



Are you an LLM? You can read better optimized documentation at /docs/guide/usage/linter/editors.md for this page in Markdown format

# Setup editors [​](https://oxc.rs/docs/guide/usage/linter/editors.html#setup-editors)

Oxlint editor extensions start the language server with `oxlint --lsp` from your **project installation**, so `oxlint` must be installed in the project.

See [Quickstart](https://oxc.rs/docs/guide/usage/linter/quickstart.html) to install and run Oxlint from the command line.

## Supported editors [​](https://oxc.rs/docs/guide/usage/linter/editors.html#supported-editors)

* [VS Code](https://oxc.rs/docs/guide/usage/linter/editors.html#vs-code)
* [Zed](https://oxc.rs/docs/guide/usage/linter/editors.html#zed)
* [JetBrains](https://oxc.rs/docs/guide/usage/linter/editors.html#jetbrains)
* [coc.nvim](https://oxc.rs/docs/guide/usage/linter/editors.html#cocnvim)
* [Other editors](https://oxc.rs/docs/guide/usage/linter/editors.html#other-editors)

## VS Code [​](https://oxc.rs/docs/guide/usage/linter/editors.html#vs-code)

### Install [​](https://oxc.rs/docs/guide/usage/linter/editors.html#install)

Download the official Oxc VS Code extension from:

* [Visual Studio Marketplace](https://marketplace.visualstudio.com/items?itemName=oxc.oxc-vscode)
* [Open VSX Registry](https://open-vsx.org/extension/oxc/oxc-vscode)

The extension is also compatible with other VS Code-based editors, including Cursor.

### Use (recommended setup for teams) [​](https://oxc.rs/docs/guide/usage/linter/editors.html#use-recommended-setup-for-teams)

1. Recommend the extension in your repository so contributors install the same tooling.

Create `.vscode/extensions.json`:

json

```
{
  "recommendations": ["oxc.oxc-vscode"]
}
```

2. Enable fix on save (optional).

Add to `.vscode/settings.json`:

jsonc

```
{
  "editor.codeActionsOnSave": {
    "source.fixAll.oxc": true,
  },
}
```

### Usage and configuration reference [​](https://oxc.rs/docs/guide/usage/linter/editors.html#usage-and-configuration-reference)

* <https://github.com/oxc-project/oxc/tree/main/editors/vscode>

## Zed [​](https://oxc.rs/docs/guide/usage/linter/editors.html#zed)

### Install [​](https://oxc.rs/docs/guide/usage/linter/editors.html#install-1)

* [Oxc Zed Extension](https://zed.dev/extensions/oxc)

### Use [​](https://oxc.rs/docs/guide/usage/linter/editors.html#use)

Configure the extension in Zed’s `settings.json` (workspace or user settings), then open your project as a folder/workspace.

### Usage and configuration reference [​](https://oxc.rs/docs/guide/usage/linter/editors.html#usage-and-configuration-reference-1)

* <https://github.com/oxc-project/oxc-zed>

---

## JetBrains [​](https://oxc.rs/docs/guide/usage/linter/editors.html#jetbrains)

IntelliJ IDEA and WebStorm

### Install [​](https://oxc.rs/docs/guide/usage/linter/editors.html#install-2)

* [Oxc in JetBrains Marketplace](https://plugins.jetbrains.com/plugin/27061-oxc)

### Use [​](https://oxc.rs/docs/guide/usage/linter/editors.html#use-1)

Install the plugin, restart the IDE, and open your repository as a project.

### Usage and configuration reference [​](https://oxc.rs/docs/guide/usage/linter/editors.html#usage-and-configuration-reference-2)

* <https://github.com/oxc-project/oxc-intellij-plugin>

## coc.nvim [​](https://oxc.rs/docs/guide/usage/linter/editors.html#coc-nvim)

### Install [​](https://oxc.rs/docs/guide/usage/linter/editors.html#install-3)

vim

```
:CocInstall coc-oxc
```

### Usage and configuration reference [​](https://oxc.rs/docs/guide/usage/linter/editors.html#usage-and-configuration-reference-3)

* <https://github.com/oxc-project/coc-oxc>

## Other editors [​](https://oxc.rs/docs/guide/usage/linter/editors.html#other-editors)

If your editor supports custom LSP configuration (Neovim LSP, Emacs, Helix, Sublime LSP, etc.), configure it to launch:

bash

```
oxlint --lsp
```

### Reference (language server implementation and issues) [​](https://oxc.rs/docs/guide/usage/linter/editors.html#reference-language-server-implementation-and-issues)

* <https://github.com/oxc-project/oxc/tree/main/crates/oxc_language_server>
