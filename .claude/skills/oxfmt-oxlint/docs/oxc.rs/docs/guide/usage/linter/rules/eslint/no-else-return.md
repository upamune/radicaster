---
title: "eslint/no-else-return | The JavaScript Oxidation Compiler"
source_url: "https://oxc.rs/docs/guide/usage/linter/rules/eslint/no-else-return"
fetched_at: "2025-12-31T10:44:14.234719+00:00"
---



Are you an LLM? You can read better optimized documentation at /docs/guide/usage/linter/rules/eslint/no-else-return.md for this page in Markdown format

# eslint/no-else-return Pedantic [​](https://oxc.rs/docs/guide/usage/linter/rules/eslint/no-else-return.html#eslint-no-else-return)

🛠️ An auto-fix is available for this rule for some violations.

### What it does [​](https://oxc.rs/docs/guide/usage/linter/rules/eslint/no-else-return.html#what-it-does)

Disallow `else` blocks after `return` statements in `if` statements

### Why is this bad? [​](https://oxc.rs/docs/guide/usage/linter/rules/eslint/no-else-return.html#why-is-this-bad)

If an `if` block contains a `return` statement, the `else` block becomes unnecessary. Its contents can be placed outside of the block.

javascript

```
function foo() {
  if (x) {
    return y;
  } else {
    return z;
  }
}
```

This rule is aimed at highlighting an unnecessary block of code following an `if` containing a return statement. As such, it will warn when it encounters an `else` following a chain of `if`s, all of them containing a `return` statement.

### Examples [​](https://oxc.rs/docs/guide/usage/linter/rules/eslint/no-else-return.html#examples)

#### `allowElseIf: true` [​](https://oxc.rs/docs/guide/usage/linter/rules/eslint/no-else-return.html#allowelseif-true)

Examples of **incorrect** code for this rule:

javascript

```
function foo1() {
  if (x) {
    return y;
  } else {
    return z;
  }
}

function foo2() {
  if (x) {
    return y;
  } else if (z) {
    return w;
  } else {
    return t;
  }
}

function foo3() {
  if (x) {
    return y;
  } else {
    var t = "foo";
  }

  return t;
}

function foo4() {
  if (error) {
    return "It failed";
  } else {
    if (loading) {
      return "It's still loading";
    }
  }
}

// Two warnings for nested occurrences
function foo5() {
  if (x) {
    if (y) {
      return y;
    } else {
      return x;
    }
  } else {
    return z;
  }
}
```

Examples of **correct** code for this rule:

javascript

```
function foo1() {
  if (x) {
    return y;
  }

  return z;
}

function foo2() {
  if (x) {
    return y;
  } else if (z) {
    var t = "foo";
  } else {
    return w;
  }
}

function foo3() {
  if (x) {
    if (z) {
      return y;
    }
  } else {
    return z;
  }
}

function foo4() {
  if (error) {
    return "It failed";
  } else if (loading) {
    return "It's still loading";
  }
}
```

## Configuration [​](https://oxc.rs/docs/guide/usage/linter/rules/eslint/no-else-return.html#configuration)

This rule accepts a configuration object with the following properties:

### allowElseIf [​](https://oxc.rs/docs/guide/usage/linter/rules/eslint/no-else-return.html#allowelseif)

type: `boolean`

default: `true`

Whether to allow `else if` blocks after a return statement.

Examples of **incorrect** code for this rule with `allowElseIf: false`:

javascript

```
function foo() {
  if (error) {
    return "It failed";
  } else if (loading) {
    return "It's still loading";
  }
}
```

Examples of **correct** code for this rule with `allowElseIf: false`:

javascript

```
function foo() {
  if (error) {
    return "It failed";
  }

  if (loading) {
    return "It's still loading";
  }
}
```

## How to use [​](https://oxc.rs/docs/guide/usage/linter/rules/eslint/no-else-return.html#how-to-use)

To **enable** this rule using the config file or in the CLI, you can use:

Config (.oxlintrc.json)CLI

json

```
{
  "rules": {
    "no-else-return": "error"
  }
}
```

bash

```
oxlint --deny no-else-return
```

## References [​](https://oxc.rs/docs/guide/usage/linter/rules/eslint/no-else-return.html#references)

* [Rule Source](https://github.com/oxc-project/oxc/blob/1bf0ffc0f6859c90409a9701e62e8957ef1286cc/crates/oxc_linter/src/rules/eslint/no_else_return.rs)
