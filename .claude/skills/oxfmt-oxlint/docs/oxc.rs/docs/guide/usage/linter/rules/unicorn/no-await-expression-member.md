---
title: "unicorn/no-await-expression-member | The JavaScript Oxidation Compiler"
source_url: "https://oxc.rs/docs/guide/usage/linter/rules/unicorn/no-await-expression-member"
fetched_at: "2025-12-31T10:44:14.234719+00:00"
---



Are you an LLM? You can read better optimized documentation at /docs/guide/usage/linter/rules/unicorn/no-await-expression-member.md for this page in Markdown format

# unicorn/no-await-expression-member Style [​](https://oxc.rs/docs/guide/usage/linter/rules/unicorn/no-await-expression-member.html#unicorn-no-await-expression-member)

⚠️🛠️️ A dangerous auto-fix is available for this rule.

### What it does [​](https://oxc.rs/docs/guide/usage/linter/rules/unicorn/no-await-expression-member.html#what-it-does)

Disallows member access from `await` expressions.

### Why is this bad? [​](https://oxc.rs/docs/guide/usage/linter/rules/unicorn/no-await-expression-member.html#why-is-this-bad)

When accessing a member from an `await` expression, the `await` expression has to be parenthesized, which is not readable.

### Examples [​](https://oxc.rs/docs/guide/usage/linter/rules/unicorn/no-await-expression-member.html#examples)

Examples of **incorrect** code for this rule:

javascript

```
async function bad() {
  const secondElement = (await getArray())[1];
}
```

Examples of **correct** code for this rule:

javascript

```
async function good() {
  const [, secondElement] = await getArray();
}
```

## How to use [​](https://oxc.rs/docs/guide/usage/linter/rules/unicorn/no-await-expression-member.html#how-to-use)

To **enable** this rule using the config file or in the CLI, you can use:

Config (.oxlintrc.json)CLI

json

```
{
  "rules": {
    "unicorn/no-await-expression-member": "error"
  }
}
```

bash

```
oxlint --deny unicorn/no-await-expression-member
```

## References [​](https://oxc.rs/docs/guide/usage/linter/rules/unicorn/no-await-expression-member.html#references)

* [Rule Source](https://github.com/oxc-project/oxc/blob/1bf0ffc0f6859c90409a9701e62e8957ef1286cc/crates/oxc_linter/src/rules/unicorn/no_await_expression_member.rs)
