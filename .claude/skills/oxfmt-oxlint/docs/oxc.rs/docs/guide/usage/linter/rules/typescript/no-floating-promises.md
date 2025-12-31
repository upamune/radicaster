---
title: "typescript/no-floating-promises | The JavaScript Oxidation Compiler"
source_url: "https://oxc.rs/docs/guide/usage/linter/rules/typescript/no-floating-promises"
fetched_at: "2025-12-31T10:44:14.234719+00:00"
---



Are you an LLM? You can read better optimized documentation at /docs/guide/usage/linter/rules/typescript/no-floating-promises.md for this page in Markdown format

# typescript/no-floating-promises Correctness [​](https://oxc.rs/docs/guide/usage/linter/rules/typescript/no-floating-promises.html#typescript-no-floating-promises)

✅ This rule is turned on by default when type-aware linting is enabled.

💭 This rule requires [type information](https://oxc.rs/docs/guide/usage/linter/type-aware.html).

🚧 An auto-fix is planned for this rule, but not implemented at this time.

### What it does [​](https://oxc.rs/docs/guide/usage/linter/rules/typescript/no-floating-promises.html#what-it-does)

This rule disallows "floating" Promises in TypeScript code, which is a Promise that is created without any code to handle its resolution or rejection.

This rule will report Promise-valued statements that are not treated in one of the following ways:

* Calling its `.then()` with two arguments
* Calling its `.catch()` with one argument
* `await`ing it
* `return`ing it
* `void`ing it

This rule also reports when an Array containing Promises is created and not properly handled. The main way to resolve this is by using one of the Promise concurrency methods to create a single Promise, then handling that according to the procedure above. These methods include:

* `Promise.all()`
* `Promise.allSettled()`
* `Promise.any()`
* `Promise.race()`

### Why is this bad? [​](https://oxc.rs/docs/guide/usage/linter/rules/typescript/no-floating-promises.html#why-is-this-bad)

Floating Promises can cause several issues, such as improperly sequenced operations, ignored Promise rejections, and more.

### Examples [​](https://oxc.rs/docs/guide/usage/linter/rules/typescript/no-floating-promises.html#examples)

Examples of **incorrect** code for this rule:

ts

```
const promise = new Promise((resolve, reject) => resolve("value"));
promise;

async function returnsPromise() {
  return "value";
}
returnsPromise().then(() => {});

Promise.reject("value").catch();

Promise.reject("value").finally();

[1, 2, 3].map(async (x) => x + 1);
```

Examples of **correct** code for this rule:

ts

```
const promise = new Promise((resolve, reject) => resolve("value"));
await promise;

async function returnsPromise() {
  return "value";
}

void returnsPromise();

returnsPromise().then(
  () => {},
  () => {},
);

Promise.reject("value").catch(() => {});

await Promise.reject("value").finally(() => {});

await Promise.all([1, 2, 3].map(async (x) => x + 1));
```

## Configuration [​](https://oxc.rs/docs/guide/usage/linter/rules/typescript/no-floating-promises.html#configuration)

This rule accepts a configuration object with the following properties:

### allowForKnownSafeCalls [​](https://oxc.rs/docs/guide/usage/linter/rules/typescript/no-floating-promises.html#allowforknownsafecalls)

type: `array`

default: `[]`

Allows specific calls to be ignored, specified as type or value specifiers.

#### allowForKnownSafeCalls[n] [​](https://oxc.rs/docs/guide/usage/linter/rules/typescript/no-floating-promises.html#allowforknownsafecalls-n)

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

### allowForKnownSafePromises [​](https://oxc.rs/docs/guide/usage/linter/rules/typescript/no-floating-promises.html#allowforknownsafepromises)

type: `array`

default: `[]`

Allows specific Promise types to be ignored, specified as type or value specifiers.

#### allowForKnownSafePromises[n] [​](https://oxc.rs/docs/guide/usage/linter/rules/typescript/no-floating-promises.html#allowforknownsafepromises-n)

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

### checkThenables [​](https://oxc.rs/docs/guide/usage/linter/rules/typescript/no-floating-promises.html#checkthenables)

type: `boolean`

default: `false`

Check for thenable objects that are not necessarily Promises.

### ignoreIIFE [​](https://oxc.rs/docs/guide/usage/linter/rules/typescript/no-floating-promises.html#ignoreiife)

type: `boolean`

default: `false`

Ignore immediately invoked function expressions (IIFEs).

### ignoreVoid [​](https://oxc.rs/docs/guide/usage/linter/rules/typescript/no-floating-promises.html#ignorevoid)

type: `boolean`

default: `true`

Ignore Promises that are void expressions.

## How to use [​](https://oxc.rs/docs/guide/usage/linter/rules/typescript/no-floating-promises.html#how-to-use)

To **enable** this rule using the config file or in the CLI, you can use:

Config (.oxlintrc.json)CLI

json

```
{
  "rules": {
    "typescript/no-floating-promises": "error"
  }
}
```

bash

```
oxlint --type-aware --deny typescript/no-floating-promises
```

## References [​](https://oxc.rs/docs/guide/usage/linter/rules/typescript/no-floating-promises.html#references)

* [Rule Source](https://github.com/oxc-project/oxc/blob/1bf0ffc0f6859c90409a9701e62e8957ef1286cc/crates/oxc_linter/src/rules/typescript/no_floating_promises.rs)
* [Rule Source (tsgolint)](https://github.com/oxc-project/tsgolint/blob/main/internal/rules/no_floating_promises/no_floating_promises.go)
