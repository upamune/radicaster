---
title: "eslint/no-template-curly-in-string | The JavaScript Oxidation Compiler"
source_url: "https://oxc.rs/docs/guide/usage/linter/rules/eslint/no-template-curly-in-string"
fetched_at: "2025-12-31T10:44:14.234719+00:00"
---



Are you an LLM? You can read better optimized documentation at /docs/guide/usage/linter/rules/eslint/no-template-curly-in-string.md for this page in Markdown format

# eslint/no-template-curly-in-string Style [​](https://oxc.rs/docs/guide/usage/linter/rules/eslint/no-template-curly-in-string.html#eslint-no-template-curly-in-string)

### What it does [​](https://oxc.rs/docs/guide/usage/linter/rules/eslint/no-template-curly-in-string.html#what-it-does)

Disallow template literal placeholder syntax in regular strings. This rule ensures that expressions like `${variable}` are only used within template literals, avoiding incorrect usage in regular strings.

### Why is this bad? [​](https://oxc.rs/docs/guide/usage/linter/rules/eslint/no-template-curly-in-string.html#why-is-this-bad)

ECMAScript 6 allows programmers to create strings containing variables or expressions using template literals. This is done by embedding expressions like `${variable}` between backticks. If regular quotes (`'` or `"`) are used with template literal syntax, it results in the literal string `"${variable}"` instead of evaluating the expression. This rule helps to avoid this mistake, ensuring that expressions are correctly evaluated inside template literals.

### Examples [​](https://oxc.rs/docs/guide/usage/linter/rules/eslint/no-template-curly-in-string.html#examples)

Examples of **incorrect** code for this rule:

javascript

```
"Hello ${name}!";
"Hello ${name}!";
"Time: ${12 * 60 * 60 * 1000}";
```

Examples of **correct** code for this rule:

javascript

```
`Hello ${name}!`;
`Time: ${12 * 60 * 60 * 1000}`;
templateFunction`Hello ${name}`;
```

## How to use [​](https://oxc.rs/docs/guide/usage/linter/rules/eslint/no-template-curly-in-string.html#how-to-use)

To **enable** this rule using the config file or in the CLI, you can use:

Config (.oxlintrc.json)CLI

json

```
{
  "rules": {
    "no-template-curly-in-string": "error"
  }
}
```

bash

```
oxlint --deny no-template-curly-in-string
```

## References [​](https://oxc.rs/docs/guide/usage/linter/rules/eslint/no-template-curly-in-string.html#references)

* [Rule Source](https://github.com/oxc-project/oxc/blob/1bf0ffc0f6859c90409a9701e62e8957ef1286cc/crates/oxc_linter/src/rules/eslint/no_template_curly_in_string.rs)
