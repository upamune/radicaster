---
title: "jest/no-large-snapshots | The JavaScript Oxidation Compiler"
source_url: "https://oxc.rs/docs/guide/usage/linter/rules/jest/no-large-snapshots"
fetched_at: "2025-12-31T10:44:14.234719+00:00"
---



Are you an LLM? You can read better optimized documentation at /docs/guide/usage/linter/rules/jest/no-large-snapshots.md for this page in Markdown format

# jest/no-large-snapshots Style [​](https://oxc.rs/docs/guide/usage/linter/rules/jest/no-large-snapshots.html#jest-no-large-snapshots)

### What it does [​](https://oxc.rs/docs/guide/usage/linter/rules/jest/no-large-snapshots.html#what-it-does)

Disallow large snapshots.

### Why is this bad? [​](https://oxc.rs/docs/guide/usage/linter/rules/jest/no-large-snapshots.html#why-is-this-bad)

When using Jest's snapshot capability one should be mindful of the size of created snapshots. As a general best practice snapshots should be limited in size in order to be more manageable and reviewable. A stored snapshot is only as good as its review and as such keeping it short, sweet, and readable is important to allow for thorough reviews.

### Example [​](https://oxc.rs/docs/guide/usage/linter/rules/jest/no-large-snapshots.html#example)

Examples of **incorrect** code for this rule:

javascript

```
exports[`a large snapshot 1`] = `
line 1
line 2
line 3
line 4
line 5
line 6
line 7
line 8
line 9
line 10
line 11
line 12
line 13
line 14
line 15
line 16
line 17
line 18
line 19
line 20
line 21
line 22
line 23
line 24
line 25
line 26
line 27
line 28
line 29
line 30
line 31
line 32
line 33
line 34
line 35
line 36
line 37
line 38
line 39
line 40
line 41
line 42
line 43
line 44
line 45
line 46
line 47
line 48
line 49
line 50
line 51
`;
```

Examples of **incorrect** code for this rule:

js

```
exports[`a more manageable and readable snapshot 1`] = `
line 1
line 2
line 3
line 4
`;
```

This rule is compatible with [eslint-plugin-vitest](https://github.com/vitest-dev/eslint-plugin-vitest/blob/main/docs/rules/no-large-snapshots.md), to use it, add the following configuration to your `.oxlintrc.json`:

json

```
{
  "rules": {
    "vitest/no-large-snapshots": "error"
  }
}
```

## Configuration [​](https://oxc.rs/docs/guide/usage/linter/rules/jest/no-large-snapshots.html#configuration)

This rule accepts a configuration object with the following properties:

### allowedSnapshots [​](https://oxc.rs/docs/guide/usage/linter/rules/jest/no-large-snapshots.html#allowedsnapshots)

type: `Record<string, array>`

default: `{}`

A map of snapshot file paths to arrays of snapshot names that are allowed to exceed the size limit. Snapshot names can be specified as regular expressions.

### inlineMaxSize [​](https://oxc.rs/docs/guide/usage/linter/rules/jest/no-large-snapshots.html#inlinemaxsize)

type: `integer`

default: `50`

Maximum number of lines allowed for inline snapshots.

### maxSize [​](https://oxc.rs/docs/guide/usage/linter/rules/jest/no-large-snapshots.html#maxsize)

type: `integer`

default: `50`

Maximum number of lines allowed for external snapshot files.

## How to use [​](https://oxc.rs/docs/guide/usage/linter/rules/jest/no-large-snapshots.html#how-to-use)

To **enable** this rule using the config file or in the CLI, you can use:

Config (.oxlintrc.json)CLI

json

```
{
  "plugins": ["jest"],
  "rules": {
    "jest/no-large-snapshots": "error"
  }
}
```

bash

```
oxlint --deny jest/no-large-snapshots --jest-plugin
```

## References [​](https://oxc.rs/docs/guide/usage/linter/rules/jest/no-large-snapshots.html#references)

* [Rule Source](https://github.com/oxc-project/oxc/blob/1bf0ffc0f6859c90409a9701e62e8957ef1286cc/crates/oxc_linter/src/rules/jest/no_large_snapshots.rs)
