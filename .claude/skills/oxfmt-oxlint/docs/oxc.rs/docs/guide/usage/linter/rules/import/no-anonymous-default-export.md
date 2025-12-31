---
title: "import/no-anonymous-default-export | The JavaScript Oxidation Compiler"
source_url: "https://oxc.rs/docs/guide/usage/linter/rules/import/no-anonymous-default-export"
fetched_at: "2025-12-31T10:44:14.234719+00:00"
---



Are you an LLM? You can read better optimized documentation at /docs/guide/usage/linter/rules/import/no-anonymous-default-export.md for this page in Markdown format

# import/no-anonymous-default-export Style [​](https://oxc.rs/docs/guide/usage/linter/rules/import/no-anonymous-default-export.html#import-no-anonymous-default-export)

### What it does [​](https://oxc.rs/docs/guide/usage/linter/rules/import/no-anonymous-default-export.html#what-it-does)

Reports if a module's default export is unnamed. This includes several types of unnamed data types; literals, object expressions, arrays, anonymous functions, arrow functions, and anonymous class declarations.

### Why is this bad? [​](https://oxc.rs/docs/guide/usage/linter/rules/import/no-anonymous-default-export.html#why-is-this-bad)

Ensuring that default exports are named helps improve the grepability of the codebase by encouraging the re-use of the same identifier for the module's default export at its declaration site and at its import sites.

### Examples [​](https://oxc.rs/docs/guide/usage/linter/rules/import/no-anonymous-default-export.html#examples)

Examples of **incorrect** code for this rule:

js

```
export default [];
export default () => {};
export default class {};
export default function() {};
export default foo(bar);
export default 123;
export default {};
export default new Foo();
export default `foo`;
export default /^123/;
```

Examples of **correct** code for this rule:

js

```
const foo = 123;
export default foo;
export default function foo() {};
export default class MyClass {};
export default function foo() {};
export default foo(bar);
/* eslint import/no-anonymous-default-export: ['error', {"allowLiteral": true}] */
export default 123;
/* eslint import/no-anonymous-default-export: ['error, {"allowArray": true}] */
export default []
/* eslint import/no-anonymous-default-export: ['error, {"allowArrowFunction": true}] */
export default () => {};
/* eslint import/no-anonymous-default-export: ['error, {"allowAnonymousClass": true}] */
export default class {};
/* eslint import/no-anonymous-default-export: ['error, {"allowAnonymousFunction": true}] */
export default function() {};
/* eslint import/no-anonymous-default-export: ['error, {"allowObject": true}] */
export default {};
/* eslint import/no-anonymous-default-export: ['error, {"allowNew": true}] */
export default new Foo();
/* eslint import/no-anonymous-default-export: ['error, {"allowCallExpression": true}] */
export default foo(bar);
```

By default, all types of anonymous default exports are forbidden, but any types can be selectively allowed by toggling them on in the options.

## Configuration [​](https://oxc.rs/docs/guide/usage/linter/rules/import/no-anonymous-default-export.html#configuration)

This rule accepts a configuration object with the following properties:

### allowAnonymousClass [​](https://oxc.rs/docs/guide/usage/linter/rules/import/no-anonymous-default-export.html#allowanonymousclass)

type: `boolean`

default: `false`

Allow anonymous class as default export.

### allowAnonymousFunction [​](https://oxc.rs/docs/guide/usage/linter/rules/import/no-anonymous-default-export.html#allowanonymousfunction)

type: `boolean`

default: `false`

Allow anonymous function as default export.

### allowArray [​](https://oxc.rs/docs/guide/usage/linter/rules/import/no-anonymous-default-export.html#allowarray)

type: `boolean`

default: `false`

Allow anonymous array as default export.

### allowArrowFunction [​](https://oxc.rs/docs/guide/usage/linter/rules/import/no-anonymous-default-export.html#allowarrowfunction)

type: `boolean`

default: `false`

Allow anonymous arrow function as default export.

### allowCallExpression [​](https://oxc.rs/docs/guide/usage/linter/rules/import/no-anonymous-default-export.html#allowcallexpression)

type: `boolean`

default: `true`

Allow anonymous call expression as default export.

### allowLiteral [​](https://oxc.rs/docs/guide/usage/linter/rules/import/no-anonymous-default-export.html#allowliteral)

type: `boolean`

default: `false`

Allow anonymous literal as default export.

### allowNew [​](https://oxc.rs/docs/guide/usage/linter/rules/import/no-anonymous-default-export.html#allownew)

type: `boolean`

default: `false`

Allow anonymous new expression as default export.

### allowObject [​](https://oxc.rs/docs/guide/usage/linter/rules/import/no-anonymous-default-export.html#allowobject)

type: `boolean`

default: `false`

Allow anonymous object as default export.

## How to use [​](https://oxc.rs/docs/guide/usage/linter/rules/import/no-anonymous-default-export.html#how-to-use)

To **enable** this rule using the config file or in the CLI, you can use:

Config (.oxlintrc.json)CLI

json

```
{
  "plugins": ["import"],
  "rules": {
    "import/no-anonymous-default-export": "error"
  }
}
```

bash

```
oxlint --deny import/no-anonymous-default-export --import-plugin
```

## References [​](https://oxc.rs/docs/guide/usage/linter/rules/import/no-anonymous-default-export.html#references)

* [Rule Source](https://github.com/oxc-project/oxc/blob/1bf0ffc0f6859c90409a9701e62e8957ef1286cc/crates/oxc_linter/src/rules/import/no_anonymous_default_export.rs)
