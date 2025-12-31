---
title: "typescript/only-throw-error | The JavaScript Oxidation Compiler"
source_url: "https://oxc.rs/docs/guide/usage/linter/rules/typescript/only-throw-error"
fetched_at: "2025-12-31T10:44:14.234719+00:00"
---



Are you an LLM? You can read better optimized documentation at /docs/guide/usage/linter/rules/typescript/only-throw-error.md for this page in Markdown format

# typescript/only-throw-error Pedantic [​](https://oxc.rs/docs/guide/usage/linter/rules/typescript/only-throw-error.html#typescript-only-throw-error)

💭 This rule requires [type information](https://oxc.rs/docs/guide/usage/linter/type-aware.html).

🚧 An auto-fix is planned for this rule, but not implemented at this time.

### What it does [​](https://oxc.rs/docs/guide/usage/linter/rules/typescript/only-throw-error.html#what-it-does)

This rule disallows throwing non-Error values.

### Why is this bad? [​](https://oxc.rs/docs/guide/usage/linter/rules/typescript/only-throw-error.html#why-is-this-bad)

It's considered good practice to only throw Error objects (or subclasses of Error). This is because Error objects automatically capture a stack trace, which is useful for debugging. Additionally, some tools and environments expect thrown values to be Error objects.

### Examples [​](https://oxc.rs/docs/guide/usage/linter/rules/typescript/only-throw-error.html#examples)

Examples of **incorrect** code for this rule:

ts

```
throw "error"; // throwing string

throw 42; // throwing number

throw true; // throwing boolean

throw { message: "error" }; // throwing plain object

throw null; // throwing null

throw undefined; // throwing undefined

const error = "Something went wrong";
throw error; // throwing non-Error variable
```

Examples of **correct** code for this rule:

ts

```
throw new Error("Something went wrong");

throw new TypeError("Invalid type");

throw new RangeError("Value out of range");

// Custom Error subclasses
class CustomError extends Error {
  constructor(message: string) {
    super(message);
    this.name = "CustomError";
  }
}
throw new CustomError("Custom error occurred");

// Variables that are Error objects
const error = new Error("Error message");
throw error;
```

## Configuration [​](https://oxc.rs/docs/guide/usage/linter/rules/typescript/only-throw-error.html#configuration)

This rule accepts a configuration object with the following properties:

### allow [​](https://oxc.rs/docs/guide/usage/linter/rules/typescript/only-throw-error.html#allow)

type: `array`

default: `[]`

An array of type or value specifiers for additional types that are allowed to be thrown. Use this to allow throwing custom error types.

#### allow[n] [​](https://oxc.rs/docs/guide/usage/linter/rules/typescript/only-throw-error.html#allow-n)

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

### allowThrowingAny [​](https://oxc.rs/docs/guide/usage/linter/rules/typescript/only-throw-error.html#allowthrowingany)

type: `boolean`

default: `true`

Whether to allow throwing values typed as `any`.

### allowThrowingUnknown [​](https://oxc.rs/docs/guide/usage/linter/rules/typescript/only-throw-error.html#allowthrowingunknown)

type: `boolean`

default: `true`

Whether to allow throwing values typed as `unknown`.

## How to use [​](https://oxc.rs/docs/guide/usage/linter/rules/typescript/only-throw-error.html#how-to-use)

To **enable** this rule using the config file or in the CLI, you can use:

Config (.oxlintrc.json)CLI

json

```
{
  "rules": {
    "typescript/only-throw-error": "error"
  }
}
```

bash

```
oxlint --type-aware --deny typescript/only-throw-error
```

## References [​](https://oxc.rs/docs/guide/usage/linter/rules/typescript/only-throw-error.html#references)

* [Rule Source](https://github.com/oxc-project/oxc/blob/1bf0ffc0f6859c90409a9701e62e8957ef1286cc/crates/oxc_linter/src/rules/typescript/only_throw_error.rs)
* [Rule Source (tsgolint)](https://github.com/oxc-project/tsgolint/blob/main/internal/rules/only_throw_error/only_throw_error.go)
