---
title: "eslint/sort-imports | The JavaScript Oxidation Compiler"
source_url: "https://oxc.rs/docs/guide/usage/linter/rules/eslint/sort-imports"
fetched_at: "2025-12-31T10:44:14.234719+00:00"
---



Are you an LLM? You can read better optimized documentation at /docs/guide/usage/linter/rules/eslint/sort-imports.md for this page in Markdown format

# eslint/sort-imports Style [​](https://oxc.rs/docs/guide/usage/linter/rules/eslint/sort-imports.html#eslint-sort-imports)

🛠️ An auto-fix is available for this rule for some violations.

### What it does [​](https://oxc.rs/docs/guide/usage/linter/rules/eslint/sort-imports.html#what-it-does)

This rule checks all import declarations and verifies that all imports are first sorted by the used member syntax and then alphabetically by the first member or alias name.

When declaring multiple imports, a sorted list of import declarations make it easier for developers to read the code and find necessary imports later.

### Why is this bad? [​](https://oxc.rs/docs/guide/usage/linter/rules/eslint/sort-imports.html#why-is-this-bad)

Consistent import sorting can be useful for readability and maintainability of code.

### Examples [​](https://oxc.rs/docs/guide/usage/linter/rules/eslint/sort-imports.html#examples)

Examples of **incorrect** code for this rule:

javascript

```
import { b, a, c } from "foo.js";

import d from "foo.js";
import e from "bar.js";
```

## Configuration [​](https://oxc.rs/docs/guide/usage/linter/rules/eslint/sort-imports.html#configuration)

This rule accepts a configuration object with the following properties:

### allowSeparatedGroups [​](https://oxc.rs/docs/guide/usage/linter/rules/eslint/sort-imports.html#allowseparatedgroups)

type: `boolean`

default: `false`

When `true`, the rule allows import groups separated by blank lines to be treated independently.

### ignoreCase [​](https://oxc.rs/docs/guide/usage/linter/rules/eslint/sort-imports.html#ignorecase)

type: `boolean`

default: `false`

When `true`, the rule ignores case-sensitivity when sorting import names.

### ignoreDeclarationSort [​](https://oxc.rs/docs/guide/usage/linter/rules/eslint/sort-imports.html#ignoredeclarationsort)

type: `boolean`

default: `false`

When `true`, the rule ignores the sorting of import declarations (the order of `import` statements).

### ignoreMemberSort [​](https://oxc.rs/docs/guide/usage/linter/rules/eslint/sort-imports.html#ignoremembersort)

type: `boolean`

default: `false`

When `true`, the rule ignores the sorting of import members within a single import declaration.

### memberSyntaxSortOrder [​](https://oxc.rs/docs/guide/usage/linter/rules/eslint/sort-imports.html#membersyntaxsortorder)

type: `array`

default: `["none", "all", "multiple", "single"]`

Specifies the sort order of different import syntaxes. Must include all 4 kinds or else this will fall back to default.

#### memberSyntaxSortOrder[n] [​](https://oxc.rs/docs/guide/usage/linter/rules/eslint/sort-imports.html#membersyntaxsortorder-n)

type: `"none" | "all" | "multiple" | "single"`

## How to use [​](https://oxc.rs/docs/guide/usage/linter/rules/eslint/sort-imports.html#how-to-use)

To **enable** this rule using the config file or in the CLI, you can use:

Config (.oxlintrc.json)CLI

json

```
{
  "rules": {
    "sort-imports": "error"
  }
}
```

bash

```
oxlint --deny sort-imports
```

## References [​](https://oxc.rs/docs/guide/usage/linter/rules/eslint/sort-imports.html#references)

* [Rule Source](https://github.com/oxc-project/oxc/blob/1bf0ffc0f6859c90409a9701e62e8957ef1286cc/crates/oxc_linter/src/rules/eslint/sort_imports.rs)
