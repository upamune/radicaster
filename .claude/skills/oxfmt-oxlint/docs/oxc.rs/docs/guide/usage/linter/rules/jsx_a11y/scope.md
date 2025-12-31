---
title: "jsx_a11y/scope | The JavaScript Oxidation Compiler"
source_url: "https://oxc.rs/docs/guide/usage/linter/rules/jsx_a11y/scope"
fetched_at: "2025-12-31T10:44:14.234719+00:00"
---



Are you an LLM? You can read better optimized documentation at /docs/guide/usage/linter/rules/jsx\_a11y/scope.md for this page in Markdown format

# jsx\_a11y/scope Correctness [​](https://oxc.rs/docs/guide/usage/linter/rules/jsx_a11y/scope.html#jsx-a11y-scope)

🛠️ An auto-fix is available for this rule.

### What it does [​](https://oxc.rs/docs/guide/usage/linter/rules/jsx_a11y/scope.html#what-it-does)

The scope prop should be used only on `<th>` elements.

### Why is this bad? [​](https://oxc.rs/docs/guide/usage/linter/rules/jsx_a11y/scope.html#why-is-this-bad)

The scope attribute makes table navigation much easier for screen reader users, provided that it is used correctly. Incorrectly used, scope can make table navigation much harder and less efficient. A screen reader operates under the assumption that a table has a header and that this header specifies a scope. Because of the way screen readers function, having an accurate header makes viewing a table far more accessible and more efficient for people who use the device.

### Examples [​](https://oxc.rs/docs/guide/usage/linter/rules/jsx_a11y/scope.html#examples)

Examples of **incorrect** code for this rule:

jsx

```
<div scope />
```

Examples of **correct** code for this rule:

jsx

```
<th scope="col" />
<th scope={scope} />
```

## How to use [​](https://oxc.rs/docs/guide/usage/linter/rules/jsx_a11y/scope.html#how-to-use)

To **enable** this rule using the config file or in the CLI, you can use:

Config (.oxlintrc.json)CLI

json

```
{
  "plugins": ["jsx-a11y"],
  "rules": {
    "jsx-a11y/scope": "error"
  }
}
```

bash

```
oxlint --deny jsx-a11y/scope --jsx-a11y-plugin
```

## References [​](https://oxc.rs/docs/guide/usage/linter/rules/jsx_a11y/scope.html#references)

* [Rule Source](https://github.com/oxc-project/oxc/blob/1bf0ffc0f6859c90409a9701e62e8957ef1286cc/crates/oxc_linter/src/rules/jsx_a11y/scope.rs)
