---
title: "react/jsx-no-duplicate-props | The JavaScript Oxidation Compiler"
source_url: "https://oxc.rs/docs/guide/usage/linter/rules/react/jsx-no-duplicate-props"
fetched_at: "2025-12-31T10:44:14.234719+00:00"
---



Are you an LLM? You can read better optimized documentation at /docs/guide/usage/linter/rules/react/jsx-no-duplicate-props.md for this page in Markdown format

# react/jsx-no-duplicate-props Correctness [​](https://oxc.rs/docs/guide/usage/linter/rules/react/jsx-no-duplicate-props.html#react-jsx-no-duplicate-props)

### What it does [​](https://oxc.rs/docs/guide/usage/linter/rules/react/jsx-no-duplicate-props.html#what-it-does)

This rule prevents duplicate props in JSX elements.

### Why is this bad? [​](https://oxc.rs/docs/guide/usage/linter/rules/react/jsx-no-duplicate-props.html#why-is-this-bad)

Having duplicate props in a JSX element is most likely a mistake. Creating JSX elements with duplicate props can cause unexpected behavior in your application.

### Examples [​](https://oxc.rs/docs/guide/usage/linter/rules/react/jsx-no-duplicate-props.html#examples)

Examples of **incorrect** code for this rule:

jsx

```
<App a a />;
<App foo={2} bar baz foo={3} />;
```

Examples of **correct** code for this rule:

jsx

```
<App a />;
<App bar baz foo={3} />;
```

### Differences from eslint-plugin-react [​](https://oxc.rs/docs/guide/usage/linter/rules/react/jsx-no-duplicate-props.html#differences-from-eslint-plugin-react)

This rule does not support the `ignoreCase` option. Props with different cases are considered distinct and will not be flagged as duplicates (e.g., `<App foo Foo />` is allowed). This is intentional, as props are case-sensitive in JSX.

## How to use [​](https://oxc.rs/docs/guide/usage/linter/rules/react/jsx-no-duplicate-props.html#how-to-use)

To **enable** this rule using the config file or in the CLI, you can use:

Config (.oxlintrc.json)CLI

json

```
{
  "plugins": ["react"],
  "rules": {
    "react/jsx-no-duplicate-props": "error"
  }
}
```

bash

```
oxlint --deny react/jsx-no-duplicate-props --react-plugin
```

## References [​](https://oxc.rs/docs/guide/usage/linter/rules/react/jsx-no-duplicate-props.html#references)

* [Rule Source](https://github.com/oxc-project/oxc/blob/1bf0ffc0f6859c90409a9701e62e8957ef1286cc/crates/oxc_linter/src/rules/react/jsx_no_duplicate_props.rs)
