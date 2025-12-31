---
title: "Ignores | The JavaScript Oxidation Compiler"
source_url: "https://oxc.rs/docs/guide/usage/linter/ignore-files"
fetched_at: "2025-12-31T10:44:14.234719+00:00"
---



Are you an LLM? You can read better optimized documentation at /docs/guide/usage/linter/ignore-files.md for this page in Markdown format

# Ignore files [​](https://oxc.rs/docs/guide/usage/linter/ignore-files.html#ignore-files)

Large repositories contain files that should not be linted, such as build output, vendored code, snapshots, or generated artifacts. Oxlint provides a predictable ignore model that works well in monorepos and CI.

## Default ignores [​](https://oxc.rs/docs/guide/usage/linter/ignore-files.html#default-ignores)

Oxlint automatically ignores:

* `.git` directories
* Minified files containing `.min.`, `-min.`, or `_min.` in the file name
* Files matched by `.gitignore` (global gitignore files are not respected)

Hidden files are not automatically ignored.

## `ignorePatterns` [​](https://oxc.rs/docs/guide/usage/linter/ignore-files.html#ignorepatterns)

The recommended approach is to define ignores in `.oxlintrc.json` using `ignorePatterns`. This keeps ignores close to the configuration they belong to and works naturally with nested configs.

Patterns are resolved relative to the configuration file.

jsonc

```
{
  "ignorePatterns": ["dist/**", "coverage/**", "vendor/**", "test/snapshots/**"],
}
```

In monorepos, nested configs can ignore package specific output without affecting the rest of the repository.

## .oxlintignore [​](https://oxc.rs/docs/guide/usage/linter/ignore-files.html#oxlintignore)

For repositories that prefer an ignore file, use `.oxlintignore`. The syntax is compatible with `.gitignore`, including comments and negation patterns.

Oxlint also supports `.eslintignore` for compatibility with existing ESLint setups. Existing `.eslintignore` files can remain in place during migration. New projects should prefer `ignorePatterns` in `.oxlintrc.json`.

## Ignore from the command line [​](https://oxc.rs/docs/guide/usage/linter/ignore-files.html#ignore-from-the-command-line)

CLI flags are useful for one off changes in CI or local debugging.

Use a custom ignore file:

bash

```
oxlint --ignore-path path/to/ignorefile
```

Add additional ignore patterns:

bash

```
oxlint --ignore-pattern 'dist/**' --ignore-pattern 'coverage/**'
```

Quote patterns to avoid shell glob expansion.

## Unignoring files [​](https://oxc.rs/docs/guide/usage/linter/ignore-files.html#unignoring-files)

Ignore files support negation patterns, which allow a directory to be ignored while keeping specific files.

To ignore everything under `build/` except one file, ignore the contents rather than the directory itself:

.oxlintignore

text

```
build/**/*
!build/keep.js
```

This keeps traversal possible while still ignoring almost everything.

## Disable ignoring [​](https://oxc.rs/docs/guide/usage/linter/ignore-files.html#disable-ignoring)

To disable all ignore behavior, including ignore files and CLI ignore options, use `--no-ignore`:

bash

```
oxlint --no-ignore
```
