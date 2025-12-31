---
title: "eslint/no-unassigned-vars | The JavaScript Oxidation Compiler"
source_url: "https://oxc.rs/docs/guide/usage/linter/rules/eslint/no-unassigned-vars"
fetched_at: "2025-12-31T10:44:14.234719+00:00"
---



Are you an LLM? You can read better optimized documentation at /docs/guide/usage/linter/rules/eslint/no-unassigned-vars.md for this page in Markdown format

# eslint/no-unassigned-vars Correctness [​](https://oxc.rs/docs/guide/usage/linter/rules/eslint/no-unassigned-vars.html#eslint-no-unassigned-vars)

✅ This rule is turned on by default.

### What it does [​](https://oxc.rs/docs/guide/usage/linter/rules/eslint/no-unassigned-vars.html#what-it-does)

Disallow let or var variables that are read but never assigned

### Why is this bad? [​](https://oxc.rs/docs/guide/usage/linter/rules/eslint/no-unassigned-vars.html#why-is-this-bad)

This rule flags let or var declarations that are never assigned a value but are still read or used in the code. Since these variables will always be undefined, their usage is likely a programming mistake.

### Examples [​](https://oxc.rs/docs/guide/usage/linter/rules/eslint/no-unassigned-vars.html#examples)

Examples of **incorrect** code for this rule:

js

```
let status;
if (status === "ready") {
  console.log("Ready!");
}
```

Examples of **correct** code for this rule:

js

```
let message = "hello";
console.log(message);

let user;
user = getUser();
console.log(user.name);
```

## How to use [​](https://oxc.rs/docs/guide/usage/linter/rules/eslint/no-unassigned-vars.html#how-to-use)

To **enable** this rule using the config file or in the CLI, you can use:

Config (.oxlintrc.json)CLI

json

```
{
  "rules": {
    "no-unassigned-vars": "error"
  }
}
```

bash

```
oxlint --deny no-unassigned-vars
```

## References [​](https://oxc.rs/docs/guide/usage/linter/rules/eslint/no-unassigned-vars.html#references)

* [Rule Source](https://github.com/oxc-project/oxc/blob/1bf0ffc0f6859c90409a9701e62e8957ef1286cc/crates/oxc_linter/src/rules/eslint/no_unassigned_vars.rs)
