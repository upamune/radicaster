---
title: "eslint/no-proto | The JavaScript Oxidation Compiler"
source_url: "https://oxc.rs/docs/guide/usage/linter/rules/eslint/no-proto"
fetched_at: "2025-12-31T10:44:14.234719+00:00"
---



Are you an LLM? You can read better optimized documentation at /docs/guide/usage/linter/rules/eslint/no-proto.md for this page in Markdown format

# eslint/no-proto Restriction [​](https://oxc.rs/docs/guide/usage/linter/rules/eslint/no-proto.html#eslint-no-proto)

🚧 An auto-fix is planned for this rule, but not implemented at this time.

### What it does [​](https://oxc.rs/docs/guide/usage/linter/rules/eslint/no-proto.html#what-it-does)

Disallow the use of the `__proto__` property.

### Why is this bad? [​](https://oxc.rs/docs/guide/usage/linter/rules/eslint/no-proto.html#why-is-this-bad)

The `__proto__` property has been deprecated as of ECMAScript 3.1 and shouldn’t be used in new code. Use `Object.getPrototypeOf` and `Object.setPrototypeOf` instead.

For more information, see [the MDN documentation](https://developer.mozilla.org/en-US/docs/Web/JavaScript/Reference/Global_Objects/Object/proto).

### Examples [​](https://oxc.rs/docs/guide/usage/linter/rules/eslint/no-proto.html#examples)

Examples of **incorrect** code for this rule:

javascript

```
/*eslint no-proto: "error"*/

var a = obj.__proto__;

var a = obj["__proto__"];

obj.__proto__ = b;

obj["__proto__"] = b;
```

## How to use [​](https://oxc.rs/docs/guide/usage/linter/rules/eslint/no-proto.html#how-to-use)

To **enable** this rule using the config file or in the CLI, you can use:

Config (.oxlintrc.json)CLI

json

```
{
  "rules": {
    "no-proto": "error"
  }
}
```

bash

```
oxlint --deny no-proto
```

## References [​](https://oxc.rs/docs/guide/usage/linter/rules/eslint/no-proto.html#references)

* [Rule Source](https://github.com/oxc-project/oxc/blob/1bf0ffc0f6859c90409a9701e62e8957ef1286cc/crates/oxc_linter/src/rules/eslint/no_proto.rs)
