---
title: "typescript/no-unsafe-type-assertion | The JavaScript Oxidation Compiler"
source_url: "https://oxc.rs/docs/guide/usage/linter/rules/typescript/no-unsafe-type-assertion"
fetched_at: "2025-12-31T10:44:14.234719+00:00"
---



Are you an LLM? You can read better optimized documentation at /docs/guide/usage/linter/rules/typescript/no-unsafe-type-assertion.md for this page in Markdown format

# typescript/no-unsafe-type-assertion Suspicious [​](https://oxc.rs/docs/guide/usage/linter/rules/typescript/no-unsafe-type-assertion.html#typescript-no-unsafe-type-assertion)

💭 This rule requires [type information](https://oxc.rs/docs/guide/usage/linter/type-aware.html).

🚧 An auto-fix is planned for this rule, but not implemented at this time.

### What it does [​](https://oxc.rs/docs/guide/usage/linter/rules/typescript/no-unsafe-type-assertion.html#what-it-does)

Disallows unsafe type assertions that narrow a type.

### Why is this bad? [​](https://oxc.rs/docs/guide/usage/linter/rules/typescript/no-unsafe-type-assertion.html#why-is-this-bad)

Type assertions that narrow a type bypass TypeScript's type-checking and can lead to runtime errors. Type assertions that broaden a type are safe because TypeScript essentially knows *less* about a type. Instead of using type assertions to narrow a type, it's better to rely on type guards, which help avoid potential runtime errors caused by unsafe type assertions.

### Examples [​](https://oxc.rs/docs/guide/usage/linter/rules/typescript/no-unsafe-type-assertion.html#examples)

Examples of **incorrect** code for this rule:

ts

```
function f() {
  return Math.random() < 0.5 ? 42 : "oops";
}
const z = f() as number;

const items = [1, "2", 3, "4"];
const number = items[0] as number;
```

Examples of **correct** code for this rule:

ts

```
function f() {
  return Math.random() < 0.5 ? 42 : "oops";
}
const z = f() as number | string | boolean;

const items = [1, "2", 3, "4"];
const number = items[0] as number | string | undefined;
```

## How to use [​](https://oxc.rs/docs/guide/usage/linter/rules/typescript/no-unsafe-type-assertion.html#how-to-use)

To **enable** this rule using the config file or in the CLI, you can use:

Config (.oxlintrc.json)CLI

json

```
{
  "rules": {
    "typescript/no-unsafe-type-assertion": "error"
  }
}
```

bash

```
oxlint --type-aware --deny typescript/no-unsafe-type-assertion
```

## References [​](https://oxc.rs/docs/guide/usage/linter/rules/typescript/no-unsafe-type-assertion.html#references)

* [Rule Source](https://github.com/oxc-project/oxc/blob/1bf0ffc0f6859c90409a9701e62e8957ef1286cc/crates/oxc_linter/src/rules/typescript/no_unsafe_type_assertion.rs)
* [Rule Source (tsgolint)](https://github.com/oxc-project/tsgolint/blob/main/internal/rules/no_unsafe_type_assertion/no_unsafe_type_assertion.go)
