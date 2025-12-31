---
title: "typescript/no-unnecessary-template-expression | The JavaScript Oxidation Compiler"
source_url: "https://oxc.rs/docs/guide/usage/linter/rules/typescript/no-unnecessary-template-expression"
fetched_at: "2025-12-31T10:44:14.234719+00:00"
---



Are you an LLM? You can read better optimized documentation at /docs/guide/usage/linter/rules/typescript/no-unnecessary-template-expression.md for this page in Markdown format

# typescript/no-unnecessary-template-expression Suspicious [​](https://oxc.rs/docs/guide/usage/linter/rules/typescript/no-unnecessary-template-expression.html#typescript-no-unnecessary-template-expression)

💭 This rule requires [type information](https://oxc.rs/docs/guide/usage/linter/type-aware.html).

🚧 An auto-fix is planned for this rule, but not implemented at this time.

### What it does [​](https://oxc.rs/docs/guide/usage/linter/rules/typescript/no-unnecessary-template-expression.html#what-it-does)

This rule disallows unnecessary template literals.

### Why is this bad? [​](https://oxc.rs/docs/guide/usage/linter/rules/typescript/no-unnecessary-template-expression.html#why-is-this-bad)

Template literals should only be used when they are needed for string interpolation or multi-line strings. Using template literals when a simple string would suffice adds unnecessary complexity.

### Examples [​](https://oxc.rs/docs/guide/usage/linter/rules/typescript/no-unnecessary-template-expression.html#examples)

Examples of **incorrect** code for this rule:

ts

```
const str1 = `Hello world`;

const str2 = `42`;

const str3 = `true`;

// Template with only literal expressions
const str4 = `${"Hello"} ${"world"}`;
```

Examples of **correct** code for this rule:

ts

```
const str1 = "Hello world";

const str2 = "42";

const str3 = "true";

// Template with variable interpolation
const name = "world";
const str4 = `Hello ${name}`;

// Multi-line string
const multiline = `
  Hello
  world
`;

// Template with expression
const str5 = `Result: ${1 + 2}`;
```

## How to use [​](https://oxc.rs/docs/guide/usage/linter/rules/typescript/no-unnecessary-template-expression.html#how-to-use)

To **enable** this rule using the config file or in the CLI, you can use:

Config (.oxlintrc.json)CLI

json

```
{
  "rules": {
    "typescript/no-unnecessary-template-expression": "error"
  }
}
```

bash

```
oxlint --type-aware --deny typescript/no-unnecessary-template-expression
```

## References [​](https://oxc.rs/docs/guide/usage/linter/rules/typescript/no-unnecessary-template-expression.html#references)

* [Rule Source](https://github.com/oxc-project/oxc/blob/1bf0ffc0f6859c90409a9701e62e8957ef1286cc/crates/oxc_linter/src/rules/typescript/no_unnecessary_template_expression.rs)
* [Rule Source (tsgolint)](https://github.com/oxc-project/tsgolint/blob/main/internal/rules/no_unnecessary_template_expression/no_unnecessary_template_expression.go)
