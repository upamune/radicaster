---
title: "typescript/no-non-null-asserted-optional-chain | The JavaScript Oxidation Compiler"
source_url: "https://oxc.rs/docs/guide/usage/linter/rules/typescript/no-non-null-asserted-optional-chain"
fetched_at: "2025-12-31T10:44:14.234719+00:00"
---



Are you an LLM? You can read better optimized documentation at /docs/guide/usage/linter/rules/typescript/no-non-null-asserted-optional-chain.md for this page in Markdown format

# typescript/no-non-null-asserted-optional-chain Correctness [​](https://oxc.rs/docs/guide/usage/linter/rules/typescript/no-non-null-asserted-optional-chain.html#typescript-no-non-null-asserted-optional-chain)

✅ This rule is turned on by default.

💡 A suggestion is available for this rule.

### What it does [​](https://oxc.rs/docs/guide/usage/linter/rules/typescript/no-non-null-asserted-optional-chain.html#what-it-does)

Disallow non-null assertions after an optional chain expression.

### Why is this bad? [​](https://oxc.rs/docs/guide/usage/linter/rules/typescript/no-non-null-asserted-optional-chain.html#why-is-this-bad)

By design, optional chain expressions (`?.`) provide `undefined` as the expression's value, if the object being accessed is `null` or `undefined`, instead of throwing an error. Using a non-null assertion (`!`) to assert the result of an optional chain expression is contradictory and likely wrong, as it indicates the code is both expecting the value to be potentially `null` or `undefined` and non-null at the same time.

In most cases, either:

1. The object is not nullable and did not need the `?.` for its property lookup
2. The non-null assertion is incorrect and introduces a type safety hole.

### Examples [​](https://oxc.rs/docs/guide/usage/linter/rules/typescript/no-non-null-asserted-optional-chain.html#examples)

Examples of **incorrect** code for this rule:

ts

```
foo?.bar!;
foo?.bar()!;
```

Examples of **correct** code for this rule:

ts

```
foo?.bar;
foo.bar!;
```

## How to use [​](https://oxc.rs/docs/guide/usage/linter/rules/typescript/no-non-null-asserted-optional-chain.html#how-to-use)

To **enable** this rule using the config file or in the CLI, you can use:

Config (.oxlintrc.json)CLI

json

```
{
  "rules": {
    "typescript/no-non-null-asserted-optional-chain": "error"
  }
}
```

bash

```
oxlint --deny typescript/no-non-null-asserted-optional-chain
```

## References [​](https://oxc.rs/docs/guide/usage/linter/rules/typescript/no-non-null-asserted-optional-chain.html#references)

* [Rule Source](https://github.com/oxc-project/oxc/blob/1bf0ffc0f6859c90409a9701e62e8957ef1286cc/crates/oxc_linter/src/rules/typescript/no_non_null_asserted_optional_chain.rs)
