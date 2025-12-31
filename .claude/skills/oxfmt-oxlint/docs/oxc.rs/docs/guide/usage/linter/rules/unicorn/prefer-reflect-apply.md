---
title: "unicorn/prefer-reflect-apply | The JavaScript Oxidation Compiler"
source_url: "https://oxc.rs/docs/guide/usage/linter/rules/unicorn/prefer-reflect-apply"
fetched_at: "2025-12-31T10:44:14.234719+00:00"
---



Are you an LLM? You can read better optimized documentation at /docs/guide/usage/linter/rules/unicorn/prefer-reflect-apply.md for this page in Markdown format

# unicorn/prefer-reflect-apply Style [​](https://oxc.rs/docs/guide/usage/linter/rules/unicorn/prefer-reflect-apply.html#unicorn-prefer-reflect-apply)

### What it does [​](https://oxc.rs/docs/guide/usage/linter/rules/unicorn/prefer-reflect-apply.html#what-it-does)

### Why is this bad? [​](https://oxc.rs/docs/guide/usage/linter/rules/unicorn/prefer-reflect-apply.html#why-is-this-bad)

`Reflect.apply()` is arguably less verbose and easier to understand. In addition, when you accept arbitrary methods, it's not safe to assume `.apply()` exists or is not overridden.

### Examples [​](https://oxc.rs/docs/guide/usage/linter/rules/unicorn/prefer-reflect-apply.html#examples)

Examples of **incorrect** code for this rule:

javascript

```
foo.apply(null, [42]);
```

Examples of **correct** code for this rule:

javascript

```
Reflect.apply(foo, null);
```

## How to use [​](https://oxc.rs/docs/guide/usage/linter/rules/unicorn/prefer-reflect-apply.html#how-to-use)

To **enable** this rule using the config file or in the CLI, you can use:

Config (.oxlintrc.json)CLI

json

```
{
  "rules": {
    "unicorn/prefer-reflect-apply": "error"
  }
}
```

bash

```
oxlint --deny unicorn/prefer-reflect-apply
```

## References [​](https://oxc.rs/docs/guide/usage/linter/rules/unicorn/prefer-reflect-apply.html#references)

* [Rule Source](https://github.com/oxc-project/oxc/blob/1bf0ffc0f6859c90409a9701e62e8957ef1286cc/crates/oxc_linter/src/rules/unicorn/prefer_reflect_apply.rs)
