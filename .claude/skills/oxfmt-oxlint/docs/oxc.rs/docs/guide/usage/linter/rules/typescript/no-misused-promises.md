---
title: "typescript/no-misused-promises | The JavaScript Oxidation Compiler"
source_url: "https://oxc.rs/docs/guide/usage/linter/rules/typescript/no-misused-promises"
fetched_at: "2025-12-31T10:44:14.234719+00:00"
---



Are you an LLM? You can read better optimized documentation at /docs/guide/usage/linter/rules/typescript/no-misused-promises.md for this page in Markdown format

# typescript/no-misused-promises Pedantic [​](https://oxc.rs/docs/guide/usage/linter/rules/typescript/no-misused-promises.html#typescript-no-misused-promises)

💭 This rule requires [type information](https://oxc.rs/docs/guide/usage/linter/type-aware.html).

🚧 An auto-fix is planned for this rule, but not implemented at this time.

### What it does [​](https://oxc.rs/docs/guide/usage/linter/rules/typescript/no-misused-promises.html#what-it-does)

This rule forbids providing Promises to logical locations such as if statements in places where the TypeScript compiler allows them but they are not handled properly. These situations can often arise due to a missing `await` keyword or just a misunderstanding of the way async functions are handled/awaited.

### Why is this bad? [​](https://oxc.rs/docs/guide/usage/linter/rules/typescript/no-misused-promises.html#why-is-this-bad)

Misused promises can cause crashes or other unexpected behavior, unless there are possibly some global unhandled promise handlers registered.

### Examples [​](https://oxc.rs/docs/guide/usage/linter/rules/typescript/no-misused-promises.html#examples)

Examples of **incorrect** code for this rule:

ts

```
// Promises in conditionals:
const promise = Promise.resolve("value");
if (promise) {
  // Do something
}

// Promises where `void` return was expected:
[1, 2, 3].forEach(async (value) => {
  await fetch(`/${value}`);
});

// Spreading Promises:
const getData = () => fetch("/");
console.log({ foo: 42, ...getData() });
```

Examples of **correct** code for this rule:

ts

```
// Awaiting the Promise to get its value in a conditional:
const promise = Promise.resolve("value");
if (await promise) {
  // Do something
}

// Using a `for-of` with `await` inside (instead of `forEach`):
for (const value of [1, 2, 3]) {
  await fetch(`/${value}`);
}

// Spreading data returned from Promise, instead of the Promise itself:
const getData = () => fetch("/");
console.log({ foo: 42, ...(await getData()) });
```

## Configuration [​](https://oxc.rs/docs/guide/usage/linter/rules/typescript/no-misused-promises.html#configuration)

This rule accepts a configuration object with the following properties:

### checksConditionals [​](https://oxc.rs/docs/guide/usage/linter/rules/typescript/no-misused-promises.html#checksconditionals)

type: `boolean`

default: `true`

Whether to check if Promises are used in conditionals. When true, disallows using Promises in conditions where a boolean is expected.

### checksSpreads [​](https://oxc.rs/docs/guide/usage/linter/rules/typescript/no-misused-promises.html#checksspreads)

type: `boolean`

default: `true`

Whether to check if Promises are used in spread syntax. When true, disallows spreading Promise values.

### checksVoidReturn [​](https://oxc.rs/docs/guide/usage/linter/rules/typescript/no-misused-promises.html#checksvoidreturn)

type: `object | boolean`

#### checksVoidReturn.arguments [​](https://oxc.rs/docs/guide/usage/linter/rules/typescript/no-misused-promises.html#checksvoidreturn-arguments)

type: `boolean`

default: `true`

Whether to check Promise-returning functions passed as arguments to void-returning functions.

#### checksVoidReturn.attributes [​](https://oxc.rs/docs/guide/usage/linter/rules/typescript/no-misused-promises.html#checksvoidreturn-attributes)

type: `boolean`

default: `true`

Whether to check Promise-returning functions in JSX attributes expecting void.

#### checksVoidReturn.inheritedMethods [​](https://oxc.rs/docs/guide/usage/linter/rules/typescript/no-misused-promises.html#checksvoidreturn-inheritedmethods)

type: `boolean`

default: `true`

Whether to check Promise-returning methods that override void-returning inherited methods.

#### checksVoidReturn.properties [​](https://oxc.rs/docs/guide/usage/linter/rules/typescript/no-misused-promises.html#checksvoidreturn-properties)

type: `boolean`

default: `true`

Whether to check Promise-returning functions assigned to object properties expecting void.

#### checksVoidReturn.returns [​](https://oxc.rs/docs/guide/usage/linter/rules/typescript/no-misused-promises.html#checksvoidreturn-returns)

type: `boolean`

default: `true`

Whether to check Promise values returned from void-returning functions.

#### checksVoidReturn.variables [​](https://oxc.rs/docs/guide/usage/linter/rules/typescript/no-misused-promises.html#checksvoidreturn-variables)

type: `boolean`

default: `true`

Whether to check Promise-returning functions assigned to variables typed as void-returning.

## How to use [​](https://oxc.rs/docs/guide/usage/linter/rules/typescript/no-misused-promises.html#how-to-use)

To **enable** this rule using the config file or in the CLI, you can use:

Config (.oxlintrc.json)CLI

json

```
{
  "rules": {
    "typescript/no-misused-promises": "error"
  }
}
```

bash

```
oxlint --type-aware --deny typescript/no-misused-promises
```

## References [​](https://oxc.rs/docs/guide/usage/linter/rules/typescript/no-misused-promises.html#references)

* [Rule Source](https://github.com/oxc-project/oxc/blob/1bf0ffc0f6859c90409a9701e62e8957ef1286cc/crates/oxc_linter/src/rules/typescript/no_misused_promises.rs)
* [Rule Source (tsgolint)](https://github.com/oxc-project/tsgolint/blob/main/internal/rules/no_misused_promises/no_misused_promises.go)
