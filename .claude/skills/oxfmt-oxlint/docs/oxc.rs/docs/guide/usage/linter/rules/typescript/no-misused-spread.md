---
title: "typescript/no-misused-spread | The JavaScript Oxidation Compiler"
source_url: "https://oxc.rs/docs/guide/usage/linter/rules/typescript/no-misused-spread"
fetched_at: "2025-12-31T10:44:14.234719+00:00"
---



Are you an LLM? You can read better optimized documentation at /docs/guide/usage/linter/rules/typescript/no-misused-spread.md for this page in Markdown format

# typescript/no-misused-spread Correctness [​](https://oxc.rs/docs/guide/usage/linter/rules/typescript/no-misused-spread.html#typescript-no-misused-spread)

✅ This rule is turned on by default when type-aware linting is enabled.

💭 This rule requires [type information](https://oxc.rs/docs/guide/usage/linter/type-aware.html).

🚧 An auto-fix is planned for this rule, but not implemented at this time.

### What it does [​](https://oxc.rs/docs/guide/usage/linter/rules/typescript/no-misused-spread.html#what-it-does)

This rule disallows spreading syntax in places where it doesn't make sense or could cause runtime errors.

### Why is this bad? [​](https://oxc.rs/docs/guide/usage/linter/rules/typescript/no-misused-spread.html#why-is-this-bad)

The spread operator can be misused in ways that might not be immediately obvious but can cause runtime errors or unexpected behavior. This rule helps catch common misuses.

### Examples [​](https://oxc.rs/docs/guide/usage/linter/rules/typescript/no-misused-spread.html#examples)

Examples of **incorrect** code for this rule:

ts

```
// Spreading a non-iterable value in an array
const num = 42;
const arr = [...num]; // Runtime error: num is not iterable

// Spreading a Promise in an array
const promise = Promise.resolve([1, 2, 3]);
const arr2 = [...promise]; // Runtime error: Promise is not iterable

// Spreading non-object in object literal
const str = "hello";
const obj = { ...str }; // Creates { '0': 'h', '1': 'e', ... } which might be unexpected
```

Examples of **correct** code for this rule:

ts

```
// Spreading arrays
const arr1 = [1, 2, 3];
const arr2 = [...arr1];

// Spreading objects
const obj1 = { a: 1, b: 2 };
const obj2 = { ...obj1 };

// Spreading resolved Promise
const promise = Promise.resolve([1, 2, 3]);
const arr3 = [...(await promise)];

// Using Array.from for non-iterables if needed
const str = "hello";
const arr4 = Array.from(str); // ['h', 'e', 'l', 'l', 'o']
```

## Configuration [​](https://oxc.rs/docs/guide/usage/linter/rules/typescript/no-misused-spread.html#configuration)

This rule accepts a configuration object with the following properties:

### allow [​](https://oxc.rs/docs/guide/usage/linter/rules/typescript/no-misused-spread.html#allow)

type: `array`

default: `[]`

An array of type or value specifiers that are allowed to be spread even if they would normally be flagged as misused.

#### allow[n] [​](https://oxc.rs/docs/guide/usage/linter/rules/typescript/no-misused-spread.html#allow-n)

type: `string`

Type or value specifier for matching specific declarations

Supports four types of specifiers:

1. **String specifier** (deprecated): Universal match by name

json

```
"Promise"
```

2. **File specifier**: Match types/values declared in local files

json

```
{ "from": "file", "name": "MyType" }
{ "from": "file", "name": ["Type1", "Type2"] }
{ "from": "file", "name": "MyType", "path": "./types.ts" }
```

3. **Lib specifier**: Match TypeScript built-in lib types

json

```
{ "from": "lib", "name": "Promise" }
{ "from": "lib", "name": ["Promise", "PromiseLike"] }
```

4. **Package specifier**: Match types/values from npm packages

json

```
{ "from": "package", "name": "Observable", "package": "rxjs" }
{ "from": "package", "name": ["Observable", "Subject"], "package": "rxjs" }
```

## How to use [​](https://oxc.rs/docs/guide/usage/linter/rules/typescript/no-misused-spread.html#how-to-use)

To **enable** this rule using the config file or in the CLI, you can use:

Config (.oxlintrc.json)CLI

json

```
{
  "rules": {
    "typescript/no-misused-spread": "error"
  }
}
```

bash

```
oxlint --type-aware --deny typescript/no-misused-spread
```

## References [​](https://oxc.rs/docs/guide/usage/linter/rules/typescript/no-misused-spread.html#references)

* [Rule Source](https://github.com/oxc-project/oxc/blob/1bf0ffc0f6859c90409a9701e62e8957ef1286cc/crates/oxc_linter/src/rules/typescript/no_misused_spread.rs)
* [Rule Source (tsgolint)](https://github.com/oxc-project/tsgolint/blob/main/internal/rules/no_misused_spread/no_misused_spread.go)
