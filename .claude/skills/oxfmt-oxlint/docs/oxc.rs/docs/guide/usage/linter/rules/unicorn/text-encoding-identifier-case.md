---
title: "unicorn/text-encoding-identifier-case | The JavaScript Oxidation Compiler"
source_url: "https://oxc.rs/docs/guide/usage/linter/rules/unicorn/text-encoding-identifier-case"
fetched_at: "2025-12-31T10:44:14.234719+00:00"
---



Are you an LLM? You can read better optimized documentation at /docs/guide/usage/linter/rules/unicorn/text-encoding-identifier-case.md for this page in Markdown format

# unicorn/text-encoding-identifier-case Style [​](https://oxc.rs/docs/guide/usage/linter/rules/unicorn/text-encoding-identifier-case.html#unicorn-text-encoding-identifier-case)

🛠️ An auto-fix is available for this rule.

### What it does [​](https://oxc.rs/docs/guide/usage/linter/rules/unicorn/text-encoding-identifier-case.html#what-it-does)

This rule enforces consistent casing for text encoding identifiers, specifically:

* `'utf8'` instead of `'UTF-8'` or `'utf-8'`
* `'ascii'` instead of `'ASCII'`

### Why is this bad? [​](https://oxc.rs/docs/guide/usage/linter/rules/unicorn/text-encoding-identifier-case.html#why-is-this-bad)

Inconsistent casing of encoding identifiers reduces code readability and can lead to subtle confusion across a codebase. Although casing is not strictly enforced by ECMAScript or Node.js, using lowercase is the conventional and widely recognized style.

### Examples [​](https://oxc.rs/docs/guide/usage/linter/rules/unicorn/text-encoding-identifier-case.html#examples)

Examples of **incorrect** code for this rule:

javascript

```
import fs from "node:fs/promises";
async function bad() {
  await fs.readFile(file, "UTF-8");
  await fs.readFile(file, "ASCII");
  const string = buffer.toString("utf-8");
}
```

Examples of **correct** code for this rule:

javascript

```
import fs from "node:fs/promises";
async function good() {
  await fs.readFile(file, "utf8");
  await fs.readFile(file, "ascii");
  const string = buffer.toString("utf8");
}
```

## How to use [​](https://oxc.rs/docs/guide/usage/linter/rules/unicorn/text-encoding-identifier-case.html#how-to-use)

To **enable** this rule using the config file or in the CLI, you can use:

Config (.oxlintrc.json)CLI

json

```
{
  "rules": {
    "unicorn/text-encoding-identifier-case": "error"
  }
}
```

bash

```
oxlint --deny unicorn/text-encoding-identifier-case
```

## References [​](https://oxc.rs/docs/guide/usage/linter/rules/unicorn/text-encoding-identifier-case.html#references)

* [Rule Source](https://github.com/oxc-project/oxc/blob/1bf0ffc0f6859c90409a9701e62e8957ef1286cc/crates/oxc_linter/src/rules/unicorn/text_encoding_identifier_case.rs)
