---
title: "typescript/no-deprecated | The JavaScript Oxidation Compiler"
source_url: "https://oxc.rs/docs/guide/usage/linter/rules/typescript/no-deprecated"
fetched_at: "2025-12-31T10:44:14.234719+00:00"
---



Are you an LLM? You can read better optimized documentation at /docs/guide/usage/linter/rules/typescript/no-deprecated.md for this page in Markdown format

# typescript/no-deprecated Pedantic [​](https://oxc.rs/docs/guide/usage/linter/rules/typescript/no-deprecated.html#typescript-no-deprecated)

💭 This rule requires [type information](https://oxc.rs/docs/guide/usage/linter/type-aware.html).

### What it does [​](https://oxc.rs/docs/guide/usage/linter/rules/typescript/no-deprecated.html#what-it-does)

Disallow using code marked as `@deprecated`.

### Why is this bad? [​](https://oxc.rs/docs/guide/usage/linter/rules/typescript/no-deprecated.html#why-is-this-bad)

The JSDoc `@deprecated` tag can be used to document some piece of code being deprecated. It's best to avoid using code marked as deprecated. This rule reports on any references to code marked as `@deprecated`.

TypeScript recognizes the `@deprecated` tag, allowing editors to visually indicate deprecated code — usually with a strikethrough. However, TypeScript doesn't report type errors for deprecated code on its own.

### Examples [​](https://oxc.rs/docs/guide/usage/linter/rules/typescript/no-deprecated.html#examples)

Examples of **incorrect** code for this rule:

ts

```
/** @deprecated Use apiV2 instead. */
declare function apiV1(): Promise<string>;
declare function apiV2(): Promise<string>;

await apiV1(); // Using deprecated function

import { parse } from "node:url";
// 'parse' is deprecated. Use the WHATWG URL API instead.
const url = parse("/foo");
```

Examples of **correct** code for this rule:

ts

```
/** @deprecated Use apiV2 instead. */
declare function apiV1(): Promise<string>;
declare function apiV2(): Promise<string>;

await apiV2(); // Using non-deprecated function

// Modern Node.js API, uses `new URL()`
const url2 = new URL("/foo", "http://www.example.com");
```

## Configuration [​](https://oxc.rs/docs/guide/usage/linter/rules/typescript/no-deprecated.html#configuration)

This rule accepts a configuration object with the following properties:

### allow [​](https://oxc.rs/docs/guide/usage/linter/rules/typescript/no-deprecated.html#allow)

type: `array`

default: `[]`

An array of type or value specifiers that are allowed to be used even if deprecated. Use this to allow specific deprecated APIs that you intentionally want to continue using.

#### allow[n] [​](https://oxc.rs/docs/guide/usage/linter/rules/typescript/no-deprecated.html#allow-n)

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

## How to use [​](https://oxc.rs/docs/guide/usage/linter/rules/typescript/no-deprecated.html#how-to-use)

To **enable** this rule using the config file or in the CLI, you can use:

Config (.oxlintrc.json)CLI

json

```
{
  "rules": {
    "typescript/no-deprecated": "error"
  }
}
```

bash

```
oxlint --type-aware --deny typescript/no-deprecated
```

## References [​](https://oxc.rs/docs/guide/usage/linter/rules/typescript/no-deprecated.html#references)

* [Rule Source](https://github.com/oxc-project/oxc/blob/1bf0ffc0f6859c90409a9701e62e8957ef1286cc/crates/oxc_linter/src/rules/typescript/no_deprecated.rs)
* [Rule Source (tsgolint)](https://github.com/oxc-project/tsgolint/blob/main/internal/rules/no_deprecated/no_deprecated.go)
