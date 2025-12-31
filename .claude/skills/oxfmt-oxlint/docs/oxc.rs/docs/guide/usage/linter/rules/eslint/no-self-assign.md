---
title: "eslint/no-self-assign | The JavaScript Oxidation Compiler"
source_url: "https://oxc.rs/docs/guide/usage/linter/rules/eslint/no-self-assign"
fetched_at: "2025-12-31T10:44:14.234719+00:00"
---



Are you an LLM? You can read better optimized documentation at /docs/guide/usage/linter/rules/eslint/no-self-assign.md for this page in Markdown format

# eslint/no-self-assign Correctness [​](https://oxc.rs/docs/guide/usage/linter/rules/eslint/no-self-assign.html#eslint-no-self-assign)

✅ This rule is turned on by default.

### What it does [​](https://oxc.rs/docs/guide/usage/linter/rules/eslint/no-self-assign.html#what-it-does)

Disallow assignments where both sides are exactly the same.

### Why is this bad? [​](https://oxc.rs/docs/guide/usage/linter/rules/eslint/no-self-assign.html#why-is-this-bad)

Self assignments have no effect, so probably those are an error due to incomplete refactoring. Those indicate that what you should do is still remaining.

### Examples [​](https://oxc.rs/docs/guide/usage/linter/rules/eslint/no-self-assign.html#examples)

Examples of **incorrect** code for this rule:

javascript

```
foo = foo;

[a, b] = [a, b];
[a, ...b] = [x, ...b];

({ a, b } = { a, x });

foo &&= foo;
foo ||= foo;
foo ??= foo;
```

javascript

```
obj.a = obj.a;
obj.a.b = obj.a.b;
obj["a"] = obj["a"];
obj[a] = obj[a];
```

Examples of **correct** code for this rule:

javascript

```
foo = bar;
[a, b] = [b, a];

// This pattern is warned by the `no-use-before-define` rule.
let foo = foo;

// The default values have an effect.
[foo = 1] = [foo];

// This ignores if there is a function call.
obj.a().b = obj.a().b;
a().b = a().b;

// `&=` and `|=` have an effect on non-integers.
foo &= foo;
foo |= foo;
```

## Configuration [​](https://oxc.rs/docs/guide/usage/linter/rules/eslint/no-self-assign.html#configuration)

This rule accepts a configuration object with the following properties:

### props [​](https://oxc.rs/docs/guide/usage/linter/rules/eslint/no-self-assign.html#props)

type: `boolean`

default: `true`

The `props` option when set to `false`, disables the checking of properties.

With `props` set to `false` the following are examples of correct code:

javascript

```
obj.a = obj.a;
obj.a.b = obj.a.b;
obj["a"] = obj["a"];
obj[a] = obj[a];
```

## How to use [​](https://oxc.rs/docs/guide/usage/linter/rules/eslint/no-self-assign.html#how-to-use)

To **enable** this rule using the config file or in the CLI, you can use:

Config (.oxlintrc.json)CLI

json

```
{
  "rules": {
    "no-self-assign": "error"
  }
}
```

bash

```
oxlint --deny no-self-assign
```

## References [​](https://oxc.rs/docs/guide/usage/linter/rules/eslint/no-self-assign.html#references)

* [Rule Source](https://github.com/oxc-project/oxc/blob/1bf0ffc0f6859c90409a9701e62e8957ef1286cc/crates/oxc_linter/src/rules/eslint/no_self_assign.rs)
