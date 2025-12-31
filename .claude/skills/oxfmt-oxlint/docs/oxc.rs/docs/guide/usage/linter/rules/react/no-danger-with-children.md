---
title: "react/no-danger-with-children | The JavaScript Oxidation Compiler"
source_url: "https://oxc.rs/docs/guide/usage/linter/rules/react/no-danger-with-children"
fetched_at: "2025-12-31T10:44:14.234719+00:00"
---



Are you an LLM? You can read better optimized documentation at /docs/guide/usage/linter/rules/react/no-danger-with-children.md for this page in Markdown format

# react/no-danger-with-children Correctness [​](https://oxc.rs/docs/guide/usage/linter/rules/react/no-danger-with-children.html#react-no-danger-with-children)

### What it does [​](https://oxc.rs/docs/guide/usage/linter/rules/react/no-danger-with-children.html#what-it-does)

Disallows when a DOM element is using both `children` and `dangerouslySetInnerHTML` properties.

### Why is this bad? [​](https://oxc.rs/docs/guide/usage/linter/rules/react/no-danger-with-children.html#why-is-this-bad)

React will throw a warning if this rule is ignored and both `children` and `dangerouslySetInnerHTML` are used.

### Examples [​](https://oxc.rs/docs/guide/usage/linter/rules/react/no-danger-with-children.html#examples)

Examples of **incorrect** code for this rule:

jsx

```
<div dangerouslySetInnerHTML={{ __html: "HTML" }}>Children</div>;
React.createElement("div", { dangerouslySetInnerHTML: { __html: "HTML" } }, "Children");
```

Examples of **correct** code for this rule:

jsx

```
<div>Children</div>
<div dangerouslySetInnerHTML={{ __html: "HTML" }} />
```

## How to use [​](https://oxc.rs/docs/guide/usage/linter/rules/react/no-danger-with-children.html#how-to-use)

To **enable** this rule using the config file or in the CLI, you can use:

Config (.oxlintrc.json)CLI

json

```
{
  "plugins": ["react"],
  "rules": {
    "react/no-danger-with-children": "error"
  }
}
```

bash

```
oxlint --deny react/no-danger-with-children --react-plugin
```

## References [​](https://oxc.rs/docs/guide/usage/linter/rules/react/no-danger-with-children.html#references)

* [Rule Source](https://github.com/oxc-project/oxc/blob/1bf0ffc0f6859c90409a9701e62e8957ef1286cc/crates/oxc_linter/src/rules/react/no_danger_with_children.rs)
