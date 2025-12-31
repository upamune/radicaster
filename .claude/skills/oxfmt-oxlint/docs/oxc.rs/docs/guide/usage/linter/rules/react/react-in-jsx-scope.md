---
title: "react/react-in-jsx-scope | The JavaScript Oxidation Compiler"
source_url: "https://oxc.rs/docs/guide/usage/linter/rules/react/react-in-jsx-scope"
fetched_at: "2025-12-31T10:44:14.234719+00:00"
---



Are you an LLM? You can read better optimized documentation at /docs/guide/usage/linter/rules/react/react-in-jsx-scope.md for this page in Markdown format

# react/react-in-jsx-scope Suspicious [​](https://oxc.rs/docs/guide/usage/linter/rules/react/react-in-jsx-scope.html#react-react-in-jsx-scope)

### What it does [​](https://oxc.rs/docs/guide/usage/linter/rules/react/react-in-jsx-scope.html#what-it-does)

Enforces that React is imported and in-scope when using JSX syntax.

Note that this rule is **not necessary** on React 17+ if you are using the new JSX Transform, and you can disable this rule and skip importing `React` in files with JSX syntax.

If your `tsconfig.json` has `jsx` set to `react-jsx` or `react-jsxdev`, you are using the new JSX Transform. For JavaScript projects using Babel, you are using the new JSX Transform if your React preset configuration (in `.babelrc` or `babel.config.js`) has `runtime: "automatic"`.

For more information, see [the React blog post on JSX Transform](https://legacy.reactjs.org/blog/2020/09/22/introducing-the-new-jsx-transform.html#eslint).

### Why is this bad? [​](https://oxc.rs/docs/guide/usage/linter/rules/react/react-in-jsx-scope.html#why-is-this-bad)

When using JSX, `<a />` expands to `React.createElement("a")`. Therefore the `React` variable must be in scope.

### Examples [​](https://oxc.rs/docs/guide/usage/linter/rules/react/react-in-jsx-scope.html#examples)

Examples of **incorrect** code for this rule:

jsx

```
const a = <a />;
```

Examples of **correct** code for this rule:

jsx

```
import React from "react";
const a = <a />;
```

## How to use [​](https://oxc.rs/docs/guide/usage/linter/rules/react/react-in-jsx-scope.html#how-to-use)

To **enable** this rule using the config file or in the CLI, you can use:

Config (.oxlintrc.json)CLI

json

```
{
  "plugins": ["react"],
  "rules": {
    "react/react-in-jsx-scope": "error"
  }
}
```

bash

```
oxlint --deny react/react-in-jsx-scope --react-plugin
```

## References [​](https://oxc.rs/docs/guide/usage/linter/rules/react/react-in-jsx-scope.html#references)

* [Rule Source](https://github.com/oxc-project/oxc/blob/1bf0ffc0f6859c90409a9701e62e8957ef1286cc/crates/oxc_linter/src/rules/react/react_in_jsx_scope.rs)
