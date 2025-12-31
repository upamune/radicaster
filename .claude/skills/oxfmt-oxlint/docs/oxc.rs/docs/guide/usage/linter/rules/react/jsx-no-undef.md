---
title: "react/jsx-no-undef | The JavaScript Oxidation Compiler"
source_url: "https://oxc.rs/docs/guide/usage/linter/rules/react/jsx-no-undef"
fetched_at: "2025-12-31T10:44:14.234719+00:00"
---



Are you an LLM? You can read better optimized documentation at /docs/guide/usage/linter/rules/react/jsx-no-undef.md for this page in Markdown format

# react/jsx-no-undef Correctness [​](https://oxc.rs/docs/guide/usage/linter/rules/react/jsx-no-undef.html#react-jsx-no-undef)

### What it does [​](https://oxc.rs/docs/guide/usage/linter/rules/react/jsx-no-undef.html#what-it-does)

Disallow undeclared variables in JSX

### Why is this bad? [​](https://oxc.rs/docs/guide/usage/linter/rules/react/jsx-no-undef.html#why-is-this-bad)

It is most likely a potential ReferenceError caused by a misspelling of a variable or parameter name.

### Examples [​](https://oxc.rs/docs/guide/usage/linter/rules/react/jsx-no-undef.html#examples)

Examples of **incorrect** code for this rule:

jsx

```
const A = () => <App />;
const C = <B />;
```

## How to use [​](https://oxc.rs/docs/guide/usage/linter/rules/react/jsx-no-undef.html#how-to-use)

To **enable** this rule using the config file or in the CLI, you can use:

Config (.oxlintrc.json)CLI

json

```
{
  "plugins": ["react"],
  "rules": {
    "react/jsx-no-undef": "error"
  }
}
```

bash

```
oxlint --deny react/jsx-no-undef --react-plugin
```

## References [​](https://oxc.rs/docs/guide/usage/linter/rules/react/jsx-no-undef.html#references)

* [Rule Source](https://github.com/oxc-project/oxc/blob/1bf0ffc0f6859c90409a9701e62e8957ef1286cc/crates/oxc_linter/src/rules/react/jsx_no_undef.rs)
