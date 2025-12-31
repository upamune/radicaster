---
title: "typescript/no-base-to-string | The JavaScript Oxidation Compiler"
source_url: "https://oxc.rs/docs/guide/usage/linter/rules/typescript/no-base-to-string"
fetched_at: "2025-12-31T10:44:14.234719+00:00"
---



Are you an LLM? You can read better optimized documentation at /docs/guide/usage/linter/rules/typescript/no-base-to-string.md for this page in Markdown format

# typescript/no-base-to-string Correctness [​](https://oxc.rs/docs/guide/usage/linter/rules/typescript/no-base-to-string.html#typescript-no-base-to-string)

✅ This rule is turned on by default when type-aware linting is enabled.

💭 This rule requires [type information](https://oxc.rs/docs/guide/usage/linter/type-aware.html).

🚧 An auto-fix is planned for this rule, but not implemented at this time.

### What it does [​](https://oxc.rs/docs/guide/usage/linter/rules/typescript/no-base-to-string.html#what-it-does)

This rule requires toString() and toLocaleString() calls to only be called on objects which provide useful information when stringified.

### Why is this bad? [​](https://oxc.rs/docs/guide/usage/linter/rules/typescript/no-base-to-string.html#why-is-this-bad)

JavaScript's toString() method returns '[object Object]' on plain objects, which is not useful information. This rule prevents toString() and toLocaleString() from being called on objects that return less useful strings.

### Examples [​](https://oxc.rs/docs/guide/usage/linter/rules/typescript/no-base-to-string.html#examples)

Examples of **incorrect** code for this rule:

ts

```
// These will evaluate to '[object Object]'
({}).toString();
({ foo: "bar" }).toString();
({ foo: "bar" }).toLocaleString();

// This will evaluate to 'Symbol()'
Symbol("foo").toString();
```

Examples of **correct** code for this rule:

ts

```
const someString = "Hello world";
someString.toString();

const someNumber = 42;
someNumber.toString();

const someBoolean = true;
someBoolean.toString();

class CustomToString {
  toString() {
    return "CustomToString";
  }
}
new CustomToString().toString();
```

## Configuration [​](https://oxc.rs/docs/guide/usage/linter/rules/typescript/no-base-to-string.html#configuration)

This rule accepts a configuration object with the following properties:

### checkUnknown [​](https://oxc.rs/docs/guide/usage/linter/rules/typescript/no-base-to-string.html#checkunknown)

type: `boolean`

default: `false`

Whether to also check values of type `unknown`. When `true`, calling toString on `unknown` values will be flagged. Default is `false`.

### ignoredTypeNames [​](https://oxc.rs/docs/guide/usage/linter/rules/typescript/no-base-to-string.html#ignoredtypenames)

type: `string[]`

default: `["Error", "RegExp", "URL", "URLSearchParams"]`

A list of type names to ignore when checking for unsafe toString usage. These types are considered safe to call toString on even if they don't provide a custom implementation.

## How to use [​](https://oxc.rs/docs/guide/usage/linter/rules/typescript/no-base-to-string.html#how-to-use)

To **enable** this rule using the config file or in the CLI, you can use:

Config (.oxlintrc.json)CLI

json

```
{
  "rules": {
    "typescript/no-base-to-string": "error"
  }
}
```

bash

```
oxlint --type-aware --deny typescript/no-base-to-string
```

## References [​](https://oxc.rs/docs/guide/usage/linter/rules/typescript/no-base-to-string.html#references)

* [Rule Source](https://github.com/oxc-project/oxc/blob/1bf0ffc0f6859c90409a9701e62e8957ef1286cc/crates/oxc_linter/src/rules/typescript/no_base_to_string.rs)
* [Rule Source (tsgolint)](https://github.com/oxc-project/tsgolint/blob/main/internal/rules/no_base_to_string/no_base_to_string.go)
