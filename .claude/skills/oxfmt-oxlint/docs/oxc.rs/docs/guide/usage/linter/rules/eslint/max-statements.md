---
title: "eslint/max-statements | The JavaScript Oxidation Compiler"
source_url: "https://oxc.rs/docs/guide/usage/linter/rules/eslint/max-statements"
fetched_at: "2025-12-31T10:44:14.234719+00:00"
---



Are you an LLM? You can read better optimized documentation at /docs/guide/usage/linter/rules/eslint/max-statements.md for this page in Markdown format

# eslint/max-statements Style [​](https://oxc.rs/docs/guide/usage/linter/rules/eslint/max-statements.html#eslint-max-statements)

### What it does [​](https://oxc.rs/docs/guide/usage/linter/rules/eslint/max-statements.html#what-it-does)

Enforce a maximum number of statements in a function. This rule ensures that functions do not exceed a specified statements count, promoting smaller, more focused functions that are easier to maintain and understand.

### Why is this bad? [​](https://oxc.rs/docs/guide/usage/linter/rules/eslint/max-statements.html#why-is-this-bad)

Some people consider large functions a code smell. Large functions tend to do a lot of things and can make it hard to follow what's going on. This rule can help avoid large functions.

### Examples [​](https://oxc.rs/docs/guide/usage/linter/rules/eslint/max-statements.html#examples)

Examples of **incorrect** code for this rule with the default `{ "max": 10 }` option:

js

```
function foo() {
  const foo1 = 1;
  const foo2 = 2;
  const foo3 = 3;
  const foo4 = 4;
  const foo5 = 5;
  const foo6 = 6;
  const foo7 = 7;
  const foo8 = 8;
  const foo9 = 9;
  const foo10 = 10;

  const foo11 = 11; // Too many.
}

const bar = () => {
  const foo1 = 1;
  const foo2 = 2;
  const foo3 = 3;
  const foo4 = 4;
  const foo5 = 5;
  const foo6 = 6;
  const foo7 = 7;
  const foo8 = 8;
  const foo9 = 9;
  const foo10 = 10;

  const foo11 = 11; // Too many.
};
```

Examples of **correct** code for this rule with the default `{ "max": 10 }` option:

js

```
function foo() {
  const foo1 = 1;
  const foo2 = 2;
  const foo3 = 3;
  const foo4 = 4;
  const foo5 = 5;
  const foo6 = 6;
  const foo7 = 7;
  const foo8 = 8;
  const foo9 = 9;
  return function () {
    // 10

    // The number of statements in the inner function does not count toward the
    // statement maximum.

    let bar;
    let baz;
    return 42;
  };
}

const bar = () => {
  const foo1 = 1;
  const foo2 = 2;
  const foo3 = 3;
  const foo4 = 4;
  const foo5 = 5;
  const foo6 = 6;
  const foo7 = 7;
  const foo8 = 8;
  const foo9 = 9;
  return function () {
    // 10

    // The number of statements in the inner function does not count toward the
    // statement maximum.

    let bar;
    let baz;
    return 42;
  };
};
```

Note that this rule does not apply to class static blocks, and that statements in class static blocks do not count as statements in the enclosing function.

Examples of **correct** code for this rule with `{ "max": 2 }` option:

js

```
function foo() {
  let one;
  let two = class {
    static {
      let three;
      let four;
      let five;
      if (six) {
        let seven;
        let eight;
        let nine;
      }
    }
  };
}
```

Examples of additional **correct** code for this rule with the `{ "max": 10 }, { "ignoreTopLevelFunctions": true }` options:

js

```
function foo() {
  const foo1 = 1;
  const foo2 = 2;
  const foo3 = 3;
  const foo4 = 4;
  const foo5 = 5;
  const foo6 = 6;
  const foo7 = 7;
  const foo8 = 8;
  const foo9 = 9;
  const foo10 = 10;
  const foo11 = 11;
}
```

## Configuration [​](https://oxc.rs/docs/guide/usage/linter/rules/eslint/max-statements.html#configuration)

This rule accepts a configuration object with the following properties:

### ignoreTopLevelFunctions [​](https://oxc.rs/docs/guide/usage/linter/rules/eslint/max-statements.html#ignoretoplevelfunctions)

type: `boolean`

default: `false`

Whether to ignore top-level functions.

### max [​](https://oxc.rs/docs/guide/usage/linter/rules/eslint/max-statements.html#max)

type: `integer`

default: `10`

Maximum number of statements allowed per function.

## How to use [​](https://oxc.rs/docs/guide/usage/linter/rules/eslint/max-statements.html#how-to-use)

To **enable** this rule using the config file or in the CLI, you can use:

Config (.oxlintrc.json)CLI

json

```
{
  "rules": {
    "max-statements": "error"
  }
}
```

bash

```
oxlint --deny max-statements
```

## References [​](https://oxc.rs/docs/guide/usage/linter/rules/eslint/max-statements.html#references)

* [Rule Source](https://github.com/oxc-project/oxc/blob/1bf0ffc0f6859c90409a9701e62e8957ef1286cc/crates/oxc_linter/src/rules/eslint/max_statements.rs)
