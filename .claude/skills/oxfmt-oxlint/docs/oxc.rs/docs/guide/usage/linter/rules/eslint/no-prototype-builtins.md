---
title: "eslint/no-prototype-builtins | The JavaScript Oxidation Compiler"
source_url: "https://oxc.rs/docs/guide/usage/linter/rules/eslint/no-prototype-builtins"
fetched_at: "2025-12-31T10:44:14.234719+00:00"
---



Are you an LLM? You can read better optimized documentation at /docs/guide/usage/linter/rules/eslint/no-prototype-builtins.md for this page in Markdown format

# eslint/no-prototype-builtins Pedantic [​](https://oxc.rs/docs/guide/usage/linter/rules/eslint/no-prototype-builtins.html#eslint-no-prototype-builtins)

### What it does [​](https://oxc.rs/docs/guide/usage/linter/rules/eslint/no-prototype-builtins.html#what-it-does)

Disallow calling some Object.prototype methods directly on objects

### Why is this bad? [​](https://oxc.rs/docs/guide/usage/linter/rules/eslint/no-prototype-builtins.html#why-is-this-bad)

In ECMAScript 5.1, Object.create was added, which enables the creation of objects with a specified [[Prototype]]. Object.create(null) is a common pattern used to create objects that will be used as a Map. This can lead to errors when it is assumed that objects will have properties from Object.prototype. This rule prevents calling some Object.prototype methods directly from an object. Additionally, objects can have properties that shadow the builtins on Object.prototype, potentially causing unintended behavior or denial-of-service security vulnerabilities. For example, it would be unsafe for a webserver to parse JSON input from a client and call hasOwnProperty directly on the resulting object, because a malicious client could send a JSON value like {"hasOwnProperty": 1} and cause the server to crash.

To avoid subtle bugs like this, it’s better to always call these methods from Object.prototype. For example, foo.hasOwnProperty("bar") should be replaced with Object.prototype.hasOwnProperty.call(foo, "bar").

### Examples [​](https://oxc.rs/docs/guide/usage/linter/rules/eslint/no-prototype-builtins.html#examples)

Examples of **incorrect** code for this rule:

javascript

```
var hasBarProperty = foo.hasOwnProperty("bar");
var isPrototypeOfBar = foo.isPrototypeOf(bar);
var barIsEnumerable = foo.propertyIsEnumerable("bar");
```

## How to use [​](https://oxc.rs/docs/guide/usage/linter/rules/eslint/no-prototype-builtins.html#how-to-use)

To **enable** this rule using the config file or in the CLI, you can use:

Config (.oxlintrc.json)CLI

json

```
{
  "rules": {
    "no-prototype-builtins": "error"
  }
}
```

bash

```
oxlint --deny no-prototype-builtins
```

## References [​](https://oxc.rs/docs/guide/usage/linter/rules/eslint/no-prototype-builtins.html#references)

* [Rule Source](https://github.com/oxc-project/oxc/blob/1bf0ffc0f6859c90409a9701e62e8957ef1286cc/crates/oxc_linter/src/rules/eslint/no_prototype_builtins.rs)
