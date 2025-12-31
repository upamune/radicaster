---
title: "eslint/no-debugger | The JavaScript Oxidation Compiler"
source_url: "https://oxc.rs/docs/guide/usage/linter/rules/eslint/no-debugger"
fetched_at: "2025-12-31T10:44:14.234719+00:00"
---



Are you an LLM? You can read better optimized documentation at /docs/guide/usage/linter/rules/eslint/no-debugger.md for this page in Markdown format

# eslint/no-debugger Correctness [​](https://oxc.rs/docs/guide/usage/linter/rules/eslint/no-debugger.html#eslint-no-debugger)

✅ This rule is turned on by default.

🛠️ An auto-fix is available for this rule.

### What it does [​](https://oxc.rs/docs/guide/usage/linter/rules/eslint/no-debugger.html#what-it-does)

Checks for usage of the `debugger` statement

### Why is this bad? [​](https://oxc.rs/docs/guide/usage/linter/rules/eslint/no-debugger.html#why-is-this-bad)

`debugger` statements do not affect functionality when a debugger isn't attached. They're most commonly an accidental debugging leftover.

### Examples [​](https://oxc.rs/docs/guide/usage/linter/rules/eslint/no-debugger.html#examples)

Examples of **incorrect** code for this rule:

javascript

```
async function main() {
  const data = await getData();
  const result = complexCalculation(data);
  debugger;
}
```

Examples of **correct** code for this rule:

javascript

```
async function main() {
  const data = await getData();
  const result = complexCalculation(data);
}
```

## How to use [​](https://oxc.rs/docs/guide/usage/linter/rules/eslint/no-debugger.html#how-to-use)

To **enable** this rule using the config file or in the CLI, you can use:

Config (.oxlintrc.json)CLI

json

```
{
  "rules": {
    "no-debugger": "error"
  }
}
```

bash

```
oxlint --deny no-debugger
```

## References [​](https://oxc.rs/docs/guide/usage/linter/rules/eslint/no-debugger.html#references)

* [Rule Source](https://github.com/oxc-project/oxc/blob/1bf0ffc0f6859c90409a9701e62e8957ef1286cc/crates/oxc_linter/src/rules/eslint/no_debugger.rs)
