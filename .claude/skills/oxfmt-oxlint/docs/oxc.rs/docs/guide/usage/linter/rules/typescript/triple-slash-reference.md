---
title: "typescript/triple-slash-reference | The JavaScript Oxidation Compiler"
source_url: "https://oxc.rs/docs/guide/usage/linter/rules/typescript/triple-slash-reference"
fetched_at: "2025-12-31T10:44:14.234719+00:00"
---



Are you an LLM? You can read better optimized documentation at /docs/guide/usage/linter/rules/typescript/triple-slash-reference.md for this page in Markdown format

# typescript/triple-slash-reference Correctness [​](https://oxc.rs/docs/guide/usage/linter/rules/typescript/triple-slash-reference.html#typescript-triple-slash-reference)

✅ This rule is turned on by default.

### What it does [​](https://oxc.rs/docs/guide/usage/linter/rules/typescript/triple-slash-reference.html#what-it-does)

Disallow certain triple slash directives in favor of ES module import declarations.

### Why is this bad? [​](https://oxc.rs/docs/guide/usage/linter/rules/typescript/triple-slash-reference.html#why-is-this-bad)

Use of triple-slash reference type directives is generally discouraged in favor of ECMAScript Module imports.

### Examples [​](https://oxc.rs/docs/guide/usage/linter/rules/typescript/triple-slash-reference.html#examples)

Examples of **incorrect** code for this rule:

ts

```
/// <reference lib="code" />
globalThis.value;
```

## Configuration [​](https://oxc.rs/docs/guide/usage/linter/rules/typescript/triple-slash-reference.html#configuration)

This rule accepts a configuration object with the following properties:

### lib [​](https://oxc.rs/docs/guide/usage/linter/rules/typescript/triple-slash-reference.html#lib)

type: `"always" | "never"`

default: `"always"`

What to enforce for `/// <reference lib="..." />` references.

#### `"always"` [​](https://oxc.rs/docs/guide/usage/linter/rules/typescript/triple-slash-reference.html#always)

Allow triple-slash `lib` references.

#### `"never"` [​](https://oxc.rs/docs/guide/usage/linter/rules/typescript/triple-slash-reference.html#never)

Disallow triple-slash `lib` references.

### path [​](https://oxc.rs/docs/guide/usage/linter/rules/typescript/triple-slash-reference.html#path)

type: `"always" | "never"`

default: `"never"`

What to enforce for `/// <reference path="..." />` references.

#### `"always"` [​](https://oxc.rs/docs/guide/usage/linter/rules/typescript/triple-slash-reference.html#always-1)

Allow triple-slash `path` references.

#### `"never"` [​](https://oxc.rs/docs/guide/usage/linter/rules/typescript/triple-slash-reference.html#never-1)

Disallow triple-slash `path` references.

### types [​](https://oxc.rs/docs/guide/usage/linter/rules/typescript/triple-slash-reference.html#types)

type: `"always" | "never" | "prefer-import"`

default: `"prefer-import"`

What to enforce for `/// <reference types="..." />` references.

#### `"always"` [​](https://oxc.rs/docs/guide/usage/linter/rules/typescript/triple-slash-reference.html#always-2)

Allow triple-slash `types` references.

#### `"never"` [​](https://oxc.rs/docs/guide/usage/linter/rules/typescript/triple-slash-reference.html#never-2)

Disallow triple-slash `types` references.

#### `"prefer-import"` [​](https://oxc.rs/docs/guide/usage/linter/rules/typescript/triple-slash-reference.html#prefer-import)

Prefer ES module import declarations over triple-slash `types` references. This option only reports when there is an existing `import` declaration for the same module.

For example, this would be reported as a lint violation with `prefer-import`:

ts

```
/// <reference types="foo" />
import { bar } from "foo";
```

## How to use [​](https://oxc.rs/docs/guide/usage/linter/rules/typescript/triple-slash-reference.html#how-to-use)

To **enable** this rule using the config file or in the CLI, you can use:

Config (.oxlintrc.json)CLI

json

```
{
  "rules": {
    "typescript/triple-slash-reference": "error"
  }
}
```

bash

```
oxlint --deny typescript/triple-slash-reference
```

## References [​](https://oxc.rs/docs/guide/usage/linter/rules/typescript/triple-slash-reference.html#references)

* [Rule Source](https://github.com/oxc-project/oxc/blob/1bf0ffc0f6859c90409a9701e62e8957ef1286cc/crates/oxc_linter/src/rules/typescript/triple_slash_reference.rs)
