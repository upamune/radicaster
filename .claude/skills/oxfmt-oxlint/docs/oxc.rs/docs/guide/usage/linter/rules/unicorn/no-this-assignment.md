---
title: "unicorn/no-this-assignment | The JavaScript Oxidation Compiler"
source_url: "https://oxc.rs/docs/guide/usage/linter/rules/unicorn/no-this-assignment"
fetched_at: "2025-12-31T10:44:14.234719+00:00"
---



Are you an LLM? You can read better optimized documentation at /docs/guide/usage/linter/rules/unicorn/no-this-assignment.md for this page in Markdown format

# unicorn/no-this-assignment Pedantic [​](https://oxc.rs/docs/guide/usage/linter/rules/unicorn/no-this-assignment.html#unicorn-no-this-assignment)

### What it does [​](https://oxc.rs/docs/guide/usage/linter/rules/unicorn/no-this-assignment.html#what-it-does)

Disallow assigning `this` to a variable.

### Why is this bad? [​](https://oxc.rs/docs/guide/usage/linter/rules/unicorn/no-this-assignment.html#why-is-this-bad)

Assigning `this` to a variable is unnecessary and confusing.

### Examples [​](https://oxc.rs/docs/guide/usage/linter/rules/unicorn/no-this-assignment.html#examples)

Examples of **incorrect** code for this rule:

javascript

```
const foo = this;
class Bar {
  method() {
    foo.baz();
  }
}

new Bar().method();
```

Examples of **correct** code for this rule:

javascript

```
class Bar {
  constructor(fooInstance) {
    this.fooInstance = fooInstance;
  }
  method() {
    this.fooInstance.baz();
  }
}

new Bar(this).method();
```

## How to use [​](https://oxc.rs/docs/guide/usage/linter/rules/unicorn/no-this-assignment.html#how-to-use)

To **enable** this rule using the config file or in the CLI, you can use:

Config (.oxlintrc.json)CLI

json

```
{
  "rules": {
    "unicorn/no-this-assignment": "error"
  }
}
```

bash

```
oxlint --deny unicorn/no-this-assignment
```

## References [​](https://oxc.rs/docs/guide/usage/linter/rules/unicorn/no-this-assignment.html#references)

* [Rule Source](https://github.com/oxc-project/oxc/blob/1bf0ffc0f6859c90409a9701e62e8957ef1286cc/crates/oxc_linter/src/rules/unicorn/no_this_assignment.rs)
