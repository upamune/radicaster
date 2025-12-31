---
title: "eslint/no-useless-call | The JavaScript Oxidation Compiler"
source_url: "https://oxc.rs/docs/guide/usage/linter/rules/eslint/no-useless-call"
fetched_at: "2025-12-31T10:44:14.234719+00:00"
---



Are you an LLM? You can read better optimized documentation at /docs/guide/usage/linter/rules/eslint/no-useless-call.md for this page in Markdown format

# eslint/no-useless-call Perf [​](https://oxc.rs/docs/guide/usage/linter/rules/eslint/no-useless-call.html#eslint-no-useless-call)

### What it does [​](https://oxc.rs/docs/guide/usage/linter/rules/eslint/no-useless-call.html#what-it-does)

Disallow unnecessary calls to `.call()` and `.apply()`

### Why is this bad? [​](https://oxc.rs/docs/guide/usage/linter/rules/eslint/no-useless-call.html#why-is-this-bad)

`Function.prototype.call()` and `Function.prototype.apply()` are slower than the normal function invocation.

This rule compares code statically to check whether or not thisArg is changed. So if the code about thisArg is a dynamic expression, this rule cannot judge correctly.

### Examples [​](https://oxc.rs/docs/guide/usage/linter/rules/eslint/no-useless-call.html#examples)

Examples of **incorrect** code for this rule:

js

```
// These are the same as `foo(1, 2, 3);`
foo.call(undefined, 1, 2, 3);
foo.apply(undefined, [1, 2, 3]);
foo.call(null, 1, 2, 3);
foo.apply(null, [1, 2, 3]);

// These are the same as `obj.foo(1, 2, 3);`
obj.foo.call(obj, 1, 2, 3);
obj.foo.apply(obj, [1, 2, 3]);
```

Examples of **correct** code for this rule:

js

```
// The `this` binding is different.
foo.call(obj, 1, 2, 3);
foo.apply(obj, [1, 2, 3]);
obj.foo.call(null, 1, 2, 3);
obj.foo.apply(null, [1, 2, 3]);
obj.foo.call(otherObj, 1, 2, 3);
obj.foo.apply(otherObj, [1, 2, 3]);

// The argument list is variadic.
// Those are warned by the `prefer-spread` rule.
foo.apply(undefined, args);
foo.apply(null, args);
obj.foo.apply(obj, args);
```

## How to use [​](https://oxc.rs/docs/guide/usage/linter/rules/eslint/no-useless-call.html#how-to-use)

To **enable** this rule using the config file or in the CLI, you can use:

Config (.oxlintrc.json)CLI

json

```
{
  "rules": {
    "no-useless-call": "error"
  }
}
```

bash

```
oxlint --deny no-useless-call
```

## References [​](https://oxc.rs/docs/guide/usage/linter/rules/eslint/no-useless-call.html#references)

* [Rule Source](https://github.com/oxc-project/oxc/blob/1bf0ffc0f6859c90409a9701e62e8957ef1286cc/crates/oxc_linter/src/rules/eslint/no_useless_call.rs)
