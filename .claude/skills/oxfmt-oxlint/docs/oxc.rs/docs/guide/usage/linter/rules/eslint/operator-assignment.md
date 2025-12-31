---
title: "eslint/operator-assignment | The JavaScript Oxidation Compiler"
source_url: "https://oxc.rs/docs/guide/usage/linter/rules/eslint/operator-assignment"
fetched_at: "2025-12-31T10:44:14.234719+00:00"
---



Are you an LLM? You can read better optimized documentation at /docs/guide/usage/linter/rules/eslint/operator-assignment.md for this page in Markdown format

# eslint/operator-assignment Style [​](https://oxc.rs/docs/guide/usage/linter/rules/eslint/operator-assignment.html#eslint-operator-assignment)

⚠️🛠️️ A dangerous auto-fix is available for this rule.

### What it does [​](https://oxc.rs/docs/guide/usage/linter/rules/eslint/operator-assignment.html#what-it-does)

This rule requires or disallows assignment operator shorthand where possible. It encourages the use of shorthand assignment operators like `+=`, `-=`, `*=`, `/=`, etc. to make the code more concise and readable.

### Why is this bad? [​](https://oxc.rs/docs/guide/usage/linter/rules/eslint/operator-assignment.html#why-is-this-bad)

JavaScript provides shorthand operators that combine variable assignment and simple mathematical operations. Failing to use these shorthand operators can lead to unnecessarily verbose code and can be seen as a missed opportunity for clarity and simplicity.

### Examples [​](https://oxc.rs/docs/guide/usage/linter/rules/eslint/operator-assignment.html#examples)

Examples of **incorrect** code for this rule with the default `always` option:

js

```
x = x + y;
x = y * x;
x[0] = x[0] / y;
x.y = x.y << z;
```

Examples of **correct** code for this rule with the default `always` option:

js

```
x = y;
x += y;
x = y * z;
x = x * y * z;
x[0] /= y;
x[foo()] = x[foo()] % 2;
x = y + x; // `+` is not always commutative (e.g. x = "abc")
```

Examples of **incorrect** code for this rule with the `never` option:

js

```
x *= y;
x ^= (y + z) / foo();
```

Examples of **correct** code for this rule with the `never` option:

js

```
x = x + y;
x.y = x.y / a.b;
```

## Configuration [​](https://oxc.rs/docs/guide/usage/linter/rules/eslint/operator-assignment.html#configuration)

This rule accepts one of the following string values:

### `"always"` [​](https://oxc.rs/docs/guide/usage/linter/rules/eslint/operator-assignment.html#always)

Requires assignment operator shorthand where possible.

### `"never"` [​](https://oxc.rs/docs/guide/usage/linter/rules/eslint/operator-assignment.html#never)

Disallows assignment operator shorthand.

## How to use [​](https://oxc.rs/docs/guide/usage/linter/rules/eslint/operator-assignment.html#how-to-use)

To **enable** this rule using the config file or in the CLI, you can use:

Config (.oxlintrc.json)CLI

json

```
{
  "rules": {
    "operator-assignment": "error"
  }
}
```

bash

```
oxlint --deny operator-assignment
```

## References [​](https://oxc.rs/docs/guide/usage/linter/rules/eslint/operator-assignment.html#references)

* [Rule Source](https://github.com/oxc-project/oxc/blob/1bf0ffc0f6859c90409a9701e62e8957ef1286cc/crates/oxc_linter/src/rules/eslint/operator_assignment.rs)
