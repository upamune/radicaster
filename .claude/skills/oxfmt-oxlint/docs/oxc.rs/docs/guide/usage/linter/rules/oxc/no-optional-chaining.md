---
title: "oxc/no-optional-chaining | The JavaScript Oxidation Compiler"
source_url: "https://oxc.rs/docs/guide/usage/linter/rules/oxc/no-optional-chaining"
fetched_at: "2025-12-31T10:44:14.234719+00:00"
---



Are you an LLM? You can read better optimized documentation at /docs/guide/usage/linter/rules/oxc/no-optional-chaining.md for this page in Markdown format

# oxc/no-optional-chaining Restriction [​](https://oxc.rs/docs/guide/usage/linter/rules/oxc/no-optional-chaining.html#oxc-no-optional-chaining)

### What it does [​](https://oxc.rs/docs/guide/usage/linter/rules/oxc/no-optional-chaining.html#what-it-does)

Disallow [optional chaining](https://developer.mozilla.org/en-US/docs/Web/JavaScript/Reference/Operators/Optional_chaining).

### Why is this bad? [​](https://oxc.rs/docs/guide/usage/linter/rules/oxc/no-optional-chaining.html#why-is-this-bad)

You may want to use this rule if you need to maintain compatibility with older environments. However, optional chaining has been supported in all major browsers since 2020 and is generally safe to use today.

In some cases, transpiling optional chaining can result in verbose helper code that impacts bundle size and performance. This rule is useful when you need to avoid the overhead of transpiled optional chaining. This is only relevant if you are polyfilling to support browsers from pre-2020.

In most codebases at this point, you should not use this rule.

### Examples [​](https://oxc.rs/docs/guide/usage/linter/rules/oxc/no-optional-chaining.html#examples)

Examples of **incorrect** code for this rule:

javascript

```
const foo = obj?.foo;
obj.fn?.();
```

## Configuration [​](https://oxc.rs/docs/guide/usage/linter/rules/oxc/no-optional-chaining.html#configuration)

This rule accepts a configuration object with the following properties:

### message [​](https://oxc.rs/docs/guide/usage/linter/rules/oxc/no-optional-chaining.html#message)

type: `string`

default: `""`

A custom help message to display when optional chaining is found. For example, "Our output target is ES2016, and optional chaining results in verbose helpers and should be avoided."

## How to use [​](https://oxc.rs/docs/guide/usage/linter/rules/oxc/no-optional-chaining.html#how-to-use)

To **enable** this rule using the config file or in the CLI, you can use:

Config (.oxlintrc.json)CLI

json

```
{
  "rules": {
    "oxc/no-optional-chaining": "error"
  }
}
```

bash

```
oxlint --deny oxc/no-optional-chaining
```

## References [​](https://oxc.rs/docs/guide/usage/linter/rules/oxc/no-optional-chaining.html#references)

* [Rule Source](https://github.com/oxc-project/oxc/blob/1bf0ffc0f6859c90409a9701e62e8957ef1286cc/crates/oxc_linter/src/rules/oxc/no_optional_chaining.rs)
