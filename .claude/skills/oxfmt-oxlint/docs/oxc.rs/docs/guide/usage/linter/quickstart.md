---
title: "Quickstart | The JavaScript Oxidation Compiler"
source_url: "https://oxc.rs/docs/guide/usage/linter/quickstart"
fetched_at: "2025-12-31T10:44:14.234719+00:00"
---



Are you an LLM? You can read better optimized documentation at /docs/guide/usage/linter/quickstart.md for this page in Markdown format

# Quickstart [​](https://oxc.rs/docs/guide/usage/linter/quickstart.html#quickstart)

This page shows the recommended setup for Oxlint and the most common workflows, with copy-paste commands.

## Install [​](https://oxc.rs/docs/guide/usage/linter/quickstart.html#install)

Install Oxlint as a dev dependency:

sh

```
pnpm add -D oxlint
```

Add lint commands to `package.json`:

json

```
{
  "scripts": {
    "lint": "oxlint",
    "lint:fix": "oxlint --fix"
  }
}
```

Run it:

sh

```
pnpm run lint
```

Apply fixes:

sh

```
pnpm run lint:fix
```

## Usage [​](https://oxc.rs/docs/guide/usage/linter/quickstart.html#usage)

For the complete list of options, see the [CLI reference](https://oxc.rs/docs/guide/usage/linter/cli.html).

sh

```
oxlint [OPTIONS] [PATH]...
```

If `PATH` is omitted, Oxlint lints the current working directory.

## Common workflows [​](https://oxc.rs/docs/guide/usage/linter/quickstart.html#common-workflows)

### Pre-commit with [lint-staged](https://github.com/lint-staged/lint-staged) [​](https://oxc.rs/docs/guide/usage/linter/quickstart.html#pre-commit-with-lint-staged)

json

```
{
  "lint-staged": {
    "*.{js,jsx,ts,tsx,mjs,cjs}": "pnpm run lint"
  }
}
```

### Create a config file [​](https://oxc.rs/docs/guide/usage/linter/quickstart.html#create-a-config-file)

Initialize the `.oxlintrc.json` config with default values:

sh

```
oxlint --init
```

Use a config file explicitly:

sh

```
oxlint -c ./oxlintrc.json
# or
oxlint --config ./oxlintrc.json
```

### Fix problems [​](https://oxc.rs/docs/guide/usage/linter/quickstart.html#fix-problems)

Apply safe fixes:

sh

```
oxlint --fix
```

Apply suggestions (may change program behavior):

sh

```
oxlint --fix-suggestions
```

Apply dangerous fixes and suggestions:

sh

```
oxlint --fix-dangerously
```

See [Automatic fixes](https://oxc.rs/docs/guide/usage/linter/automatic-fixes.html) for guidance on when to use each mode.

### Ignore files [​](https://oxc.rs/docs/guide/usage/linter/quickstart.html#ignore-files)

Use an ignore file:

sh

```
oxlint --ignore-path .oxlintignore
```

Add ignore patterns from the command line:

sh

```
oxlint --ignore-pattern "dist/**" --ignore-pattern "*.min.js"
```

Disable ignore handling:

sh

```
oxlint --no-ignore
```

Follow symlinks (Oxlint ignores symlinks by default):

sh

```
oxlint --symlinks
```

See [Ignore files](https://oxc.rs/docs/guide/usage/linter/ignore-files.html).

### Fail CI reliably [​](https://oxc.rs/docs/guide/usage/linter/quickstart.html#fail-ci-reliably)

Only report errors:

sh

```
oxlint --quiet
```

Fail if any warnings are found:

sh

```
oxlint --deny-warnings
```

Fail if warnings exceed a threshold:

sh

```
oxlint --max-warnings 0
```

See [CI setup](https://oxc.rs/docs/guide/usage/linter/ci.html).

### Use machine-readable output [​](https://oxc.rs/docs/guide/usage/linter/quickstart.html#use-machine-readable-output)

Select an output format:

sh

```
oxlint -f json
```

Available formats include: `default`, `json`, `unix`, `checkstyle`, `github`, `stylish`.

### Inspect the effective configuration [​](https://oxc.rs/docs/guide/usage/linter/quickstart.html#inspect-the-effective-configuration)

Print the configuration that would be used for a file:

sh

```
oxlint --print-config path/to/file.ts
```

### List available rules [​](https://oxc.rs/docs/guide/usage/linter/quickstart.html#list-available-rules)

List registered rules:

sh

```
oxlint --rules
```

The full list is in the [Rules reference](https://oxc.rs/docs/guide/usage/linter/rules.html).

## Next steps [​](https://oxc.rs/docs/guide/usage/linter/quickstart.html#next-steps)

* Configure rules, plugins, and ignores: [Configuration](https://oxc.rs/docs/guide/usage/linter/config.html)
* [Setup editors](https://oxc.rs/docs/guide/usage/linter/editors.html)
* [Setup CI](https://oxc.rs/docs/guide/usage/linter/ci.html)
* Learn advanced features: [Multi-file analysis](https://oxc.rs/docs/guide/usage/linter/multi-file-analysis.html), [Type-aware linting](https://oxc.rs/docs/guide/usage/linter/type-aware.html), [JS plugins](https://oxc.rs/docs/guide/usage/linter/js-plugins.html)

* [CLI reference](https://oxc.rs/docs/guide/usage/linter/cli.html)
