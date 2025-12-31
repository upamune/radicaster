---
title: "eslint/no-lonely-if | The JavaScript Oxidation Compiler"
source_url: "https://oxc.rs/docs/guide/usage/linter/rules/eslint/no-lonely-if"
fetched_at: "2025-12-31T10:44:14.234719+00:00"
---



Are you an LLM? You can read better optimized documentation at /docs/guide/usage/linter/rules/eslint/no-lonely-if.md for this page in Markdown format

# eslint/no-lonely-if Pedantic [​](https://oxc.rs/docs/guide/usage/linter/rules/eslint/no-lonely-if.html#eslint-no-lonely-if)

🚧 An auto-fix is planned for this rule, but not implemented at this time.

### What it does [​](https://oxc.rs/docs/guide/usage/linter/rules/eslint/no-lonely-if.html#what-it-does)

Disallow `if` statements as the only statement in `else` blocks

### Why is this bad? [​](https://oxc.rs/docs/guide/usage/linter/rules/eslint/no-lonely-if.html#why-is-this-bad)

When an `if` statement is the only statement in an `else` block, it is often clearer to use an `else if` instead.

### Examples [​](https://oxc.rs/docs/guide/usage/linter/rules/eslint/no-lonely-if.html#examples)

Examples of **incorrect** code for this rule:

js

```
if (condition) {
  // ...
} else {
  if (anotherCondition) {
    // ...
  }
}
```

js

```
if (condition) {
  // ...
} else {
  if (anotherCondition) {
    // ...
  } else {
    // ...
  }
}
```

Examples of **correct** code for this rule:

js

```
if (condition) {
  // ...
} else if (anotherCondition) {
  // ...
}
```

js

```
if (condition) {
  // ...
} else if (anotherCondition) {
  // ...
} else {
  // ...
}
```

js

```
if (condition) {
  // ...
} else {
  if (anotherCondition) {
    // ...
  }
  doSomething();
}
```

## How to use [​](https://oxc.rs/docs/guide/usage/linter/rules/eslint/no-lonely-if.html#how-to-use)

To **enable** this rule using the config file or in the CLI, you can use:

Config (.oxlintrc.json)CLI

json

```
{
  "rules": {
    "no-lonely-if": "error"
  }
}
```

bash

```
oxlint --deny no-lonely-if
```

## References [​](https://oxc.rs/docs/guide/usage/linter/rules/eslint/no-lonely-if.html#references)

* [Rule Source](https://github.com/oxc-project/oxc/blob/1bf0ffc0f6859c90409a9701e62e8957ef1286cc/crates/oxc_linter/src/rules/eslint/no_lonely_if.rs)
