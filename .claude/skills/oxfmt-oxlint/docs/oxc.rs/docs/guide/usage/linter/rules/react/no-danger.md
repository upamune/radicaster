---
title: "react/no-danger | The JavaScript Oxidation Compiler"
source_url: "https://oxc.rs/docs/guide/usage/linter/rules/react/no-danger"
fetched_at: "2025-12-31T10:44:14.234719+00:00"
---



Are you an LLM? You can read better optimized documentation at /docs/guide/usage/linter/rules/react/no-danger.md for this page in Markdown format

# react/no-danger Restriction [​](https://oxc.rs/docs/guide/usage/linter/rules/react/no-danger.html#react-no-danger)

### What it does [​](https://oxc.rs/docs/guide/usage/linter/rules/react/no-danger.html#what-it-does)

This rule prevents the use of `dangerouslySetInnerHTML` prop.

### Why is this bad? [​](https://oxc.rs/docs/guide/usage/linter/rules/react/no-danger.html#why-is-this-bad)

`dangerouslySetInnerHTML` is a way to inject HTML into your React component. This is dangerous because it can easily lead to XSS vulnerabilities.

### Examples [​](https://oxc.rs/docs/guide/usage/linter/rules/react/no-danger.html#examples)

Examples of **incorrect** code for this rule:

jsx

```
import React from "react";

const Hello = <div dangerouslySetInnerHTML={{ __html: "Hello World" }}></div>;
```

Examples of **correct** code for this rule:

jsx

```
import React from "react";

const Hello = <div>Hello World</div>;
```

## How to use [​](https://oxc.rs/docs/guide/usage/linter/rules/react/no-danger.html#how-to-use)

To **enable** this rule using the config file or in the CLI, you can use:

Config (.oxlintrc.json)CLI

json

```
{
  "plugins": ["react"],
  "rules": {
    "react/no-danger": "error"
  }
}
```

bash

```
oxlint --deny react/no-danger --react-plugin
```

## References [​](https://oxc.rs/docs/guide/usage/linter/rules/react/no-danger.html#references)

* [Rule Source](https://github.com/oxc-project/oxc/blob/1bf0ffc0f6859c90409a9701e62e8957ef1286cc/crates/oxc_linter/src/rules/react/no_danger.rs)
