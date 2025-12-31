---
title: "Command-line Interface | The JavaScript Oxidation Compiler"
source_url: "https://oxc.rs/docs/guide/usage/formatter/cli"
fetched_at: "2025-12-31T10:44:14.234719+00:00"
---



Are you an LLM? You can read better optimized documentation at /docs/guide/usage/formatter/cli.md for this page in Markdown format

# Command-line Interface [​](https://oxc.rs/docs/guide/usage/formatter/cli.html#command-line-interface)

## Usage [​](https://oxc.rs/docs/guide/usage/formatter/cli.html#usage)

**`oxfmt`** [**`-c`**=*`PATH`*] [*`PATH`*]...

## Mode Options: [​](https://oxc.rs/docs/guide/usage/formatter/cli.html#mode-options)

* **`--init`** — Initialize `.oxfmtrc.json` with default values
* **`--migrate`**=*`SOURCE`* — Migrate configuration to `.oxfmtrc.json` from specified source Available sources: prettier
* **`--lsp`** — Start language server protocol (LSP) server
* **`--stdin-filepath`**=*`PATH`* — Specify the file name to use to infer which parser to use

## Output Options: [​](https://oxc.rs/docs/guide/usage/formatter/cli.html#output-options)

* **`--write`** — Format and write files in place (default)
* **`--check`** — Check if files are formatted, also show statistics
* **`--list-different`** — List files that would be changed

## Config Options [​](https://oxc.rs/docs/guide/usage/formatter/cli.html#config-options)

* **`-c`**, **`--config`**=*`PATH`* — Path to the configuration file

## Ignore Options [​](https://oxc.rs/docs/guide/usage/formatter/cli.html#ignore-options)

* **`--ignore-path`**=*`PATH`* — Path to ignore file(s). Can be specified multiple times. If not specified, .gitignore and .prettierignore in the current directory are used.
* **`--with-node-modules`** — Format code in node\_modules directory (skipped by default)

## Runtime Options [​](https://oxc.rs/docs/guide/usage/formatter/cli.html#runtime-options)

* **`--no-error-on-unmatched-pattern`** — Do not exit with error when pattern is unmatched
* **`--threads`**=*`INT`* — Number of threads to use. Set to 1 for using only 1 CPU core.

## Available positional items: [​](https://oxc.rs/docs/guide/usage/formatter/cli.html#available-positional-items)

* *`PATH`* — Single file, single path or list of paths. If not provided, current working directory is used. Glob is supported only for exclude patterns like `'!**/fixtures/*.js'`.

## Available options: [​](https://oxc.rs/docs/guide/usage/formatter/cli.html#available-options)

* **`-h`**, **`--help`** — Prints help information
* **`-V`**, **`--version`** — Prints version information
