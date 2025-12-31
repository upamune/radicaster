---
title: "eslint/max-nested-callbacks | The JavaScript Oxidation Compiler"
source_url: "https://oxc.rs/docs/guide/usage/linter/rules/eslint/max-nested-callbacks"
fetched_at: "2025-12-31T10:44:14.234719+00:00"
---



Are you an LLM? You can read better optimized documentation at /docs/guide/usage/linter/rules/eslint/max-nested-callbacks.md for this page in Markdown format

# eslint/max-nested-callbacks Pedantic [​](https://oxc.rs/docs/guide/usage/linter/rules/eslint/max-nested-callbacks.html#eslint-max-nested-callbacks)

### What it does [​](https://oxc.rs/docs/guide/usage/linter/rules/eslint/max-nested-callbacks.html#what-it-does)

Enforce a maximum depth that callbacks can be nested. This rule helps to limit the complexity of callback nesting, ensuring that callbacks do not become too deeply nested, improving code readability and maintainability.

### Why is this bad? [​](https://oxc.rs/docs/guide/usage/linter/rules/eslint/max-nested-callbacks.html#why-is-this-bad)

Many JavaScript libraries use the callback pattern to manage asynchronous operations. A program of any complexity will most likely need to manage several asynchronous operations at various levels of concurrency. A common pitfall is nesting callbacks excessively, making code harder to read and understand.

### Examples [​](https://oxc.rs/docs/guide/usage/linter/rules/eslint/max-nested-callbacks.html#examples)

Examples of **incorrect** code for this rule with the `{ "max": 3 }` option:

js

```
foo1(function () {
  foo2(function () {
    foo3(function () {
      foo4(function () {
        // ...
      });
    });
  });
});
```

Examples of **correct** code for this rule with the `{ "max": 3 }` option:

js

```
foo1(handleFoo1);

function handleFoo1() {
  foo2(handleFoo2);
}

function handleFoo2() {
  foo3(handleFoo3);
}

function handleFoo3() {
  foo4(handleFoo4);
}

function handleFoo4() {
  foo5();
}
```

## Configuration [​](https://oxc.rs/docs/guide/usage/linter/rules/eslint/max-nested-callbacks.html#configuration)

This rule accepts a configuration object with the following properties:

### max [​](https://oxc.rs/docs/guide/usage/linter/rules/eslint/max-nested-callbacks.html#max)

type: `integer`

default: `10`

The `max` enforces a maximum depth that callbacks can be nested.

## How to use [​](https://oxc.rs/docs/guide/usage/linter/rules/eslint/max-nested-callbacks.html#how-to-use)

To **enable** this rule using the config file or in the CLI, you can use:

Config (.oxlintrc.json)CLI

json

```
{
  "rules": {
    "max-nested-callbacks": "error"
  }
}
```

bash

```
oxlint --deny max-nested-callbacks
```

## References [​](https://oxc.rs/docs/guide/usage/linter/rules/eslint/max-nested-callbacks.html#references)

* [Rule Source](https://github.com/oxc-project/oxc/blob/1bf0ffc0f6859c90409a9701e62e8957ef1286cc/crates/oxc_linter/src/rules/eslint/max_nested_callbacks.rs)
