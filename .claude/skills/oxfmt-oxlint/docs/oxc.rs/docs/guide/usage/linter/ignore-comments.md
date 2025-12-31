---
title: "Inline ignore comments | The JavaScript Oxidation Compiler"
source_url: "https://oxc.rs/docs/guide/usage/linter/ignore-comments"
fetched_at: "2025-12-31T10:44:14.234719+00:00"
---



Are you an LLM? You can read better optimized documentation at /docs/guide/usage/linter/ignore-comments.md for this page in Markdown format

# Inline ignore comments [​](https://oxc.rs/docs/guide/usage/linter/ignore-comments.html#inline-ignore-comments)

Ignore comments provide an escape hatch for exceptional cases where a rule is correct in general but needs to be suppressed in a small, well scoped section of code. Inline comments override configuration files.

Oxlint supports line comments (`//`) and block comments (`/* */`). Comments must start with one of the keywords below.

## oxlint-disable [​](https://oxc.rs/docs/guide/usage/linter/ignore-comments.html#oxlint-disable)

Disable one or more rules until the end of the file, or until they are re enabled.

js

```
// Disable Oxlint for the rest of the file
/* oxlint-disable */

// Disable a single rule in this file
/* oxlint-disable no-console */

// Disable multiple rules in this file
/* oxlint-disable no-console, no-alert */
```

## oxlint-enable [​](https://oxc.rs/docs/guide/usage/linter/ignore-comments.html#oxlint-enable)

Enable one or more rules until the end of the file, or until they are disabled again.

js

```
/* oxlint-enable no-console */

/* oxlint-enable no-console, no-alert */
```

## oxlint-disable-line [​](https://oxc.rs/docs/guide/usage/linter/ignore-comments.html#oxlint-disable-line)

Disable one or more rules on the current line.

js

```
console.log("Hello, world!"); // oxlint-disable-line no-console

console.log(x++); // oxlint-disable-line no-console, no-plusplus
```

## oxlint-enable-line [​](https://oxc.rs/docs/guide/usage/linter/ignore-comments.html#oxlint-enable-line)

Enable one or more rules on the current line.

js

```
// oxlint-enable-line no-console
```

## oxlint-disable-next-line [​](https://oxc.rs/docs/guide/usage/linter/ignore-comments.html#oxlint-disable-next-line)

Disable one or more rules on the line following the comment, then re enable them automatically.

js

```
// oxlint-disable-next-line no-console
console.log("Hello, world!"); // allowed because of the previous comment
console.log(x++); // not allowed because the previous comment only applied to the previous line

// oxlint-disable-next-line no-console, no-plusplus
console.log("Hello, world!"); // allowed
```

## oxlint-enable-next-line [​](https://oxc.rs/docs/guide/usage/linter/ignore-comments.html#oxlint-enable-next-line)

Enable one or more rules on the line following the comment, then re enable automatically.

js

```
// oxlint-enable-next-line no-console
```

## ESLint compatibility [​](https://oxc.rs/docs/guide/usage/linter/ignore-comments.html#eslint-compatibility)

For compatibility with existing ESLint codebases, the same keywords are supported with `oxlint` replaced by `eslint`, such as `/* eslint-disable */` and `// eslint-disable-next-line`.

The `oxlint-*` form is recommended. The `eslint-*` form is useful during migration for rules that Oxlint does not support yet.

## Rule options cannot be changed inline [​](https://oxc.rs/docs/guide/usage/linter/ignore-comments.html#rule-options-cannot-be-changed-inline)

Ignore comments can enable or disable rules, but they cannot change rule options. Rule options belong in configuration files.

## Report unused ignore comments [​](https://oxc.rs/docs/guide/usage/linter/ignore-comments.html#report-unused-ignore-comments)

Reporting unused ignore comments is disabled by default. When enabled, Oxlint reports comments such as `// oxlint-disable-line` when no diagnostics would have been reported on that line.

Enable reporting:

bash

```
oxlint --report-unused-disable-directives
```

Specify the severity:

bash

```
oxlint --report-unused-disable-directives-severity error
```

Only one of these options can be used at a time.
