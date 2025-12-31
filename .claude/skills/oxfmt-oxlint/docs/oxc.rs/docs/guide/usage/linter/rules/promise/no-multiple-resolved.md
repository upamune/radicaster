---
title: "promise/no-multiple-resolved | The JavaScript Oxidation Compiler"
source_url: "https://oxc.rs/docs/guide/usage/linter/rules/promise/no-multiple-resolved"
fetched_at: "2025-12-31T10:44:14.234719+00:00"
---



Are you an LLM? You can read better optimized documentation at /docs/guide/usage/linter/rules/promise/no-multiple-resolved.md for this page in Markdown format

# promise/no-multiple-resolved Suspicious [​](https://oxc.rs/docs/guide/usage/linter/rules/promise/no-multiple-resolved.html#promise-no-multiple-resolved)

### What it does [​](https://oxc.rs/docs/guide/usage/linter/rules/promise/no-multiple-resolved.html#what-it-does)

This rule warns of paths that resolve multiple times in executor functions that Promise constructors.

### Why is this bad? [​](https://oxc.rs/docs/guide/usage/linter/rules/promise/no-multiple-resolved.html#why-is-this-bad)

Multiple resolve/reject calls:

* Violate the Promise/A+ specification
* Have no effect on the Promise's behavior
* Make the code's intent unclear
* May indicate logical errors in the implementation

### Examples [​](https://oxc.rs/docs/guide/usage/linter/rules/promise/no-multiple-resolved.html#examples)

Examples of **incorrect** code for this rule:

javascript

```
new Promise((resolve, reject) => {
  fn((error, value) => {
    if (error) {
      reject(error);
    }

    resolve(value); // Both `reject` and `resolve` may be called.
  });
});
```

Examples of **correct** code for this rule:

javascript

```
new Promise((resolve, reject) => {
  fn((error, value) => {
    if (error) {
      reject(error);
    } else {
      resolve(value);
    }
  });
});
```

## How to use [​](https://oxc.rs/docs/guide/usage/linter/rules/promise/no-multiple-resolved.html#how-to-use)

To **enable** this rule using the config file or in the CLI, you can use:

Config (.oxlintrc.json)CLI

json

```
{
  "plugins": ["promise"],
  "rules": {
    "promise/no-multiple-resolved": "error"
  }
}
```

bash

```
oxlint --deny promise/no-multiple-resolved --promise-plugin
```

## References [​](https://oxc.rs/docs/guide/usage/linter/rules/promise/no-multiple-resolved.html#references)

* [Rule Source](https://github.com/oxc-project/oxc/blob/1bf0ffc0f6859c90409a9701e62e8957ef1286cc/crates/oxc_linter/src/rules/promise/no_multiple_resolved.rs)
