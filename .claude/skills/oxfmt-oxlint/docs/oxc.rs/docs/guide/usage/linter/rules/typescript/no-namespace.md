---
title: "typescript/no-namespace | The JavaScript Oxidation Compiler"
source_url: "https://oxc.rs/docs/guide/usage/linter/rules/typescript/no-namespace"
fetched_at: "2025-12-31T10:44:14.234719+00:00"
---



Are you an LLM? You can read better optimized documentation at /docs/guide/usage/linter/rules/typescript/no-namespace.md for this page in Markdown format

# typescript/no-namespace Restriction [​](https://oxc.rs/docs/guide/usage/linter/rules/typescript/no-namespace.html#typescript-no-namespace)

### What it does [​](https://oxc.rs/docs/guide/usage/linter/rules/typescript/no-namespace.html#what-it-does)

Disallow TypeScript namespaces.

### Why is this bad? [​](https://oxc.rs/docs/guide/usage/linter/rules/typescript/no-namespace.html#why-is-this-bad)

TypeScript historically allowed a form of code organization called "custom modules" (module Example {}), later renamed to "namespaces" (namespace Example). Namespaces are an outdated way to organize TypeScript code. ES2015 module syntax is now preferred (import/export).

### Examples [​](https://oxc.rs/docs/guide/usage/linter/rules/typescript/no-namespace.html#examples)

Examples of **incorrect** code for this rule:

typescript

```
module foo {}
namespace foo {}
declare module foo {}
declare namespace foo {}
```

Examples of **correct** code for this rule:

typescript

```
declare module "foo" {}
// anything inside a d.ts file
```

## Configuration [​](https://oxc.rs/docs/guide/usage/linter/rules/typescript/no-namespace.html#configuration)

This rule accepts a configuration object with the following properties:

### allowDeclarations [​](https://oxc.rs/docs/guide/usage/linter/rules/typescript/no-namespace.html#allowdeclarations)

type: `boolean`

default: `false`

Whether to allow declare with custom TypeScript namespaces.

Examples of **incorrect** code for this rule when `{ "allowDeclarations": true }`

typescript

```
module foo {}
namespace foo {}
```

Examples of **correct** code for this rule when `{ "allowDeclarations": true }`

typescript

```
declare module "foo" {}
declare module foo {}
declare namespace foo {}

declare global {
  namespace foo {}
}

declare module foo {
  namespace foo {}
}
```

Examples of **incorrect** code for this rule when `{ "allowDeclarations": false }`

typescript

```
module foo {}
namespace foo {}
declare module foo {}
declare namespace foo {}
```

Examples of **correct** code for this rule when `{ "allowDeclarations": false }`

typescript

```
declare module "foo" {}
```

### allowDefinitionFiles [​](https://oxc.rs/docs/guide/usage/linter/rules/typescript/no-namespace.html#allowdefinitionfiles)

type: `boolean`

default: `true`

Examples of **incorrect** code for this rule when `{ "allowDefinitionFiles": true }`

typescript

```
// if outside a d.ts file
module foo {}
namespace foo {}

// if outside a d.ts file
module foo {}
namespace foo {}
declare module foo {}
declare namespace foo {}
```

Examples of **correct** code for this rule when `{ "allowDefinitionFiles": true }`

typescript

```
declare module "foo" {}
// anything inside a d.ts file
```

## How to use [​](https://oxc.rs/docs/guide/usage/linter/rules/typescript/no-namespace.html#how-to-use)

To **enable** this rule using the config file or in the CLI, you can use:

Config (.oxlintrc.json)CLI

json

```
{
  "rules": {
    "typescript/no-namespace": "error"
  }
}
```

bash

```
oxlint --deny typescript/no-namespace
```

## References [​](https://oxc.rs/docs/guide/usage/linter/rules/typescript/no-namespace.html#references)

* [Rule Source](https://github.com/oxc-project/oxc/blob/1bf0ffc0f6859c90409a9701e62e8957ef1286cc/crates/oxc_linter/src/rules/typescript/no_namespace.rs)
