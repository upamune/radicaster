---
title: "react/jsx-props-no-spread-multi | The JavaScript Oxidation Compiler"
source_url: "https://oxc.rs/docs/guide/usage/linter/rules/react/jsx-props-no-spread-multi"
fetched_at: "2025-12-31T10:44:14.234719+00:00"
---



Are you an LLM? You can read better optimized documentation at /docs/guide/usage/linter/rules/react/jsx-props-no-spread-multi.md for this page in Markdown format

# react/jsx-props-no-spread-multi Correctness [​](https://oxc.rs/docs/guide/usage/linter/rules/react/jsx-props-no-spread-multi.html#react-jsx-props-no-spread-multi)

🛠️ An auto-fix is available for this rule.

### What it does [​](https://oxc.rs/docs/guide/usage/linter/rules/react/jsx-props-no-spread-multi.html#what-it-does)

Enforces that any unique expression is only spread once.

### Why is this bad? [​](https://oxc.rs/docs/guide/usage/linter/rules/react/jsx-props-no-spread-multi.html#why-is-this-bad)

Generally spreading the same expression twice is an indicator of a mistake since any attribute between the spreads may be overridden when the intent was not to. Even when that is not the case this will lead to unnecessary computations being performed.

### Examples [​](https://oxc.rs/docs/guide/usage/linter/rules/react/jsx-props-no-spread-multi.html#examples)

Examples of **incorrect** code for this rule:

jsx

```
<App {...props} myAttr="1" {...props} />
```

Examples of **correct** code for this rule:

jsx

```
<App myAttr="1" {...props} />
<App {...props} myAttr="1" />
```

## How to use [​](https://oxc.rs/docs/guide/usage/linter/rules/react/jsx-props-no-spread-multi.html#how-to-use)

To **enable** this rule using the config file or in the CLI, you can use:

Config (.oxlintrc.json)CLI

json

```
{
  "plugins": ["react"],
  "rules": {
    "react/jsx-props-no-spread-multi": "error"
  }
}
```

bash

```
oxlint --deny react/jsx-props-no-spread-multi --react-plugin
```

## References [​](https://oxc.rs/docs/guide/usage/linter/rules/react/jsx-props-no-spread-multi.html#references)

* [Rule Source](https://github.com/oxc-project/oxc/blob/1bf0ffc0f6859c90409a9701e62e8957ef1286cc/crates/oxc_linter/src/rules/react/jsx_props_no_spread_multi.rs)
