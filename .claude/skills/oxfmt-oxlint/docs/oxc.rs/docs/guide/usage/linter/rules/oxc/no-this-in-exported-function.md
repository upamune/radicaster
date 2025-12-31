---
title: "oxc/no-this-in-exported-function | The JavaScript Oxidation Compiler"
source_url: "https://oxc.rs/docs/guide/usage/linter/rules/oxc/no-this-in-exported-function"
fetched_at: "2025-12-31T10:44:14.234719+00:00"
---



Are you an LLM? You can read better optimized documentation at /docs/guide/usage/linter/rules/oxc/no-this-in-exported-function.md for this page in Markdown format

# oxc/no-this-in-exported-function Suspicious [​](https://oxc.rs/docs/guide/usage/linter/rules/oxc/no-this-in-exported-function.html#oxc-no-this-in-exported-function)

### What it does [​](https://oxc.rs/docs/guide/usage/linter/rules/oxc/no-this-in-exported-function.html#what-it-does)

Disallows the use of `this` in exported functions.

### Why is this bad? [​](https://oxc.rs/docs/guide/usage/linter/rules/oxc/no-this-in-exported-function.html#why-is-this-bad)

In most bundlers, the value of `this` is not preserved for exported functions. When a function is exported and imported in another module, `this` typically becomes `undefined` instead of the module namespace object. This can lead to unexpected runtime errors or incorrect behavior.

### Examples [​](https://oxc.rs/docs/guide/usage/linter/rules/oxc/no-this-in-exported-function.html#examples)

Examples of **incorrect** code for this rule:

javascript

```
export function foo() {
  console.log(this);
}

export default function bar() {
  this.something();
}

function baz() {
  const self = this;
}
export { baz };
```

Examples of **correct** code for this rule:

javascript

```
function foo() {
  console.log(this);
}

export const bar = () => {
  console.log(this);
};
```

## How to use [​](https://oxc.rs/docs/guide/usage/linter/rules/oxc/no-this-in-exported-function.html#how-to-use)

To **enable** this rule using the config file or in the CLI, you can use:

Config (.oxlintrc.json)CLI

json

```
{
  "rules": {
    "oxc/no-this-in-exported-function": "error"
  }
}
```

bash

```
oxlint --deny oxc/no-this-in-exported-function
```

## References [​](https://oxc.rs/docs/guide/usage/linter/rules/oxc/no-this-in-exported-function.html#references)

* [Rule Source](https://github.com/oxc-project/oxc/blob/1bf0ffc0f6859c90409a9701e62e8957ef1286cc/crates/oxc_linter/src/rules/oxc/no_this_in_exported_function.rs)
