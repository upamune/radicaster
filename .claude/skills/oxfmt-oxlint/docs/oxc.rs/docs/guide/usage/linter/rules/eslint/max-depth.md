---
title: "eslint/max-depth | The JavaScript Oxidation Compiler"
source_url: "https://oxc.rs/docs/guide/usage/linter/rules/eslint/max-depth"
fetched_at: "2025-12-31T10:44:14.234719+00:00"
---



Are you an LLM? You can read better optimized documentation at /docs/guide/usage/linter/rules/eslint/max-depth.md for this page in Markdown format

# eslint/max-depth Pedantic [​](https://oxc.rs/docs/guide/usage/linter/rules/eslint/max-depth.html#eslint-max-depth)

### What it does [​](https://oxc.rs/docs/guide/usage/linter/rules/eslint/max-depth.html#what-it-does)

Enforce a maximum depth that blocks can be nested. This rule helps to limit the complexity of nested blocks, improving readability and maintainability by ensuring that code does not become too deeply nested.

### Why is this bad? [​](https://oxc.rs/docs/guide/usage/linter/rules/eslint/max-depth.html#why-is-this-bad)

Many developers consider code difficult to read if blocks are nested beyond a certain depth. Excessive nesting can make it harder to follow the flow of the code, increasing cognitive load and making maintenance more error-prone. By enforcing a maximum block depth, this rule encourages cleaner, more readable code.

### Examples [​](https://oxc.rs/docs/guide/usage/linter/rules/eslint/max-depth.html#examples)

Examples of **incorrect** code for this rule with the default `{ "max": 3 }` option:

js

```
function foo() {
  for (;;) { // Nested 1 deep
    while (true) { // Nested 2 deep
      if (true) { // Nested 3 deep
        if (true) { // Nested 4 deep }
      }
    }
  }
}
```

Examples of **correct** code for this rule with the default `{ "max": 3 }` option:

js

```
function foo() {
  for (;;) { // Nested 1 deep
    while (true) { // Nested 2 deep
      if (true) { // Nested 3 deep }
    }
  }
}
```

Note that class static blocks do not count as nested blocks, and that the depth in them is calculated separately from the enclosing context.

Example:

js

```
function foo() {
  if (true) { // Nested 1 deep
    class C {
      static {
        if (true) { // Nested 1 deep
          if (true) { // Nested 2 deep }
        }
      }
    }
  }
}
```

## Configuration [​](https://oxc.rs/docs/guide/usage/linter/rules/eslint/max-depth.html#configuration)

This rule accepts a configuration object with the following properties:

### max [​](https://oxc.rs/docs/guide/usage/linter/rules/eslint/max-depth.html#max)

type: `integer`

default: `4`

The `max` enforces a maximum depth that blocks can be nested

## How to use [​](https://oxc.rs/docs/guide/usage/linter/rules/eslint/max-depth.html#how-to-use)

To **enable** this rule using the config file or in the CLI, you can use:

Config (.oxlintrc.json)CLI

json

```
{
  "rules": {
    "max-depth": "error"
  }
}
```

bash

```
oxlint --deny max-depth
```

## References [​](https://oxc.rs/docs/guide/usage/linter/rules/eslint/max-depth.html#references)

* [Rule Source](https://github.com/oxc-project/oxc/blob/1bf0ffc0f6859c90409a9701e62e8957ef1286cc/crates/oxc_linter/src/rules/eslint/max_depth.rs)
