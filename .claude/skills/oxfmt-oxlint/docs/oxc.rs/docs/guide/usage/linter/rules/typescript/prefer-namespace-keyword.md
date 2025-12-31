---
title: "typescript/prefer-namespace-keyword | The JavaScript Oxidation Compiler"
source_url: "https://oxc.rs/docs/guide/usage/linter/rules/typescript/prefer-namespace-keyword"
fetched_at: "2025-12-31T10:44:14.234719+00:00"
---



Are you an LLM? You can read better optimized documentation at /docs/guide/usage/linter/rules/typescript/prefer-namespace-keyword.md for this page in Markdown format

# typescript/prefer-namespace-keyword Style [​](https://oxc.rs/docs/guide/usage/linter/rules/typescript/prefer-namespace-keyword.html#typescript-prefer-namespace-keyword)

🛠️ An auto-fix is available for this rule.

### What it does [​](https://oxc.rs/docs/guide/usage/linter/rules/typescript/prefer-namespace-keyword.html#what-it-does)

This rule reports when the module keyword is used instead of namespace. This rule does not report on the use of TypeScript module declarations to describe external APIs (declare module 'foo' {}).

### Why is this bad? [​](https://oxc.rs/docs/guide/usage/linter/rules/typescript/prefer-namespace-keyword.html#why-is-this-bad)

Namespaces are an outdated way to organize TypeScript code. ES2015 module syntax is now preferred (import/export). For projects still using custom modules / namespaces, it's preferred to refer to them as namespaces.

### Examples [​](https://oxc.rs/docs/guide/usage/linter/rules/typescript/prefer-namespace-keyword.html#examples)

Examples of **incorrect** code for this rule:

typescript

```
module Example {}
```

Examples of **correct** code for this rule:

typescript

```
namespace Example {}
```

## How to use [​](https://oxc.rs/docs/guide/usage/linter/rules/typescript/prefer-namespace-keyword.html#how-to-use)

To **enable** this rule using the config file or in the CLI, you can use:

Config (.oxlintrc.json)CLI

json

```
{
  "rules": {
    "typescript/prefer-namespace-keyword": "error"
  }
}
```

bash

```
oxlint --deny typescript/prefer-namespace-keyword
```

## References [​](https://oxc.rs/docs/guide/usage/linter/rules/typescript/prefer-namespace-keyword.html#references)

* [Rule Source](https://github.com/oxc-project/oxc/blob/1bf0ffc0f6859c90409a9701e62e8957ef1286cc/crates/oxc_linter/src/rules/typescript/prefer_namespace_keyword.rs)
