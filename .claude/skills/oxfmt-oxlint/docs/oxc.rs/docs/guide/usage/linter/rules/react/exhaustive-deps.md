---
title: "react/exhaustive-deps | The JavaScript Oxidation Compiler"
source_url: "https://oxc.rs/docs/guide/usage/linter/rules/react/exhaustive-deps"
fetched_at: "2025-12-31T10:44:14.234719+00:00"
---



Are you an LLM? You can read better optimized documentation at /docs/guide/usage/linter/rules/react/exhaustive-deps.md for this page in Markdown format

# react/exhaustive-deps Correctness [​](https://oxc.rs/docs/guide/usage/linter/rules/react/exhaustive-deps.html#react-exhaustive-deps)

⚠️🛠️️💡 A dangerous auto-fix and a suggestion are available for this rule.

### What it does [​](https://oxc.rs/docs/guide/usage/linter/rules/react/exhaustive-deps.html#what-it-does)

Verifies the list of dependencies for Hooks like `useEffect` and similar.

### Why is this bad? [​](https://oxc.rs/docs/guide/usage/linter/rules/react/exhaustive-deps.html#why-is-this-bad)

React Hooks like `useEffect` and similar require a list of dependencies to be passed as an argument. This list is used to determine when the effect should be re-run. If the list is missing or incomplete, the effect may run more often than necessary, or not at all.

### Examples [​](https://oxc.rs/docs/guide/usage/linter/rules/react/exhaustive-deps.html#examples)

Examples of **incorrect** code for this rule:

javascript

```
function MyComponent(props) {
  useEffect(() => {
    console.log(props.foo);
  }, []);
  // `props` is missing from the dependencies array
  return <div />;
}
```

Examples of **correct** code for this rule:

javascript

```
function MyComponent(props) {
  useEffect(() => {
    console.log(props.foo);
  }, [props]);
  return <div />;
}
```

## Configuration [​](https://oxc.rs/docs/guide/usage/linter/rules/react/exhaustive-deps.html#configuration)

This rule accepts a configuration object with the following properties:

### additionalHooks [​](https://oxc.rs/docs/guide/usage/linter/rules/react/exhaustive-deps.html#additionalhooks)

type: `string | null`

default: `null`

Optionally provide a regex of additional hooks to check.

## How to use [​](https://oxc.rs/docs/guide/usage/linter/rules/react/exhaustive-deps.html#how-to-use)

To **enable** this rule using the config file or in the CLI, you can use:

Config (.oxlintrc.json)CLI

json

```
{
  "plugins": ["react"],
  "rules": {
    "react/exhaustive-deps": "error"
  }
}
```

bash

```
oxlint --deny react/exhaustive-deps --react-plugin
```

## References [​](https://oxc.rs/docs/guide/usage/linter/rules/react/exhaustive-deps.html#references)

* [Rule Source](https://github.com/oxc-project/oxc/blob/1bf0ffc0f6859c90409a9701e62e8957ef1286cc/crates/oxc_linter/src/rules/react/exhaustive_deps.rs)
