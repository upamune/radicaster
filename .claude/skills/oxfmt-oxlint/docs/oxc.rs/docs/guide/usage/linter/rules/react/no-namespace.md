---
title: "react/no-namespace | The JavaScript Oxidation Compiler"
source_url: "https://oxc.rs/docs/guide/usage/linter/rules/react/no-namespace"
fetched_at: "2025-12-31T10:44:14.234719+00:00"
---



Are you an LLM? You can read better optimized documentation at /docs/guide/usage/linter/rules/react/no-namespace.md for this page in Markdown format

# react/no-namespace Suspicious [​](https://oxc.rs/docs/guide/usage/linter/rules/react/no-namespace.html#react-no-namespace)

### What it does [​](https://oxc.rs/docs/guide/usage/linter/rules/react/no-namespace.html#what-it-does)

Enforce that namespaces are not used in React elements.

### Why is this bad? [​](https://oxc.rs/docs/guide/usage/linter/rules/react/no-namespace.html#why-is-this-bad)

Namespaces in React elements, such as svg:circle, are not supported by React.

### Examples [​](https://oxc.rs/docs/guide/usage/linter/rules/react/no-namespace.html#examples)

Examples of **incorrect** code for this rule:

jsx

```
<ns:TestComponent />
<Ns:TestComponent />
```

Examples of **correct** code for this rule:

jsx

```
<TestComponent />
<testComponent />
```

## How to use [​](https://oxc.rs/docs/guide/usage/linter/rules/react/no-namespace.html#how-to-use)

To **enable** this rule using the config file or in the CLI, you can use:

Config (.oxlintrc.json)CLI

json

```
{
  "plugins": ["react"],
  "rules": {
    "react/no-namespace": "error"
  }
}
```

bash

```
oxlint --deny react/no-namespace --react-plugin
```

## References [​](https://oxc.rs/docs/guide/usage/linter/rules/react/no-namespace.html#references)

* [Rule Source](https://github.com/oxc-project/oxc/blob/1bf0ffc0f6859c90409a9701e62e8957ef1286cc/crates/oxc_linter/src/rules/react/no_namespace.rs)
