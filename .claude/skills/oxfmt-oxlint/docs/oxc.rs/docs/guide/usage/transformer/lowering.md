---
title: "Syntax Lowering | The JavaScript Oxidation Compiler"
source_url: "https://oxc.rs/docs/guide/usage/transformer/lowering"
fetched_at: "2025-12-31T10:44:14.234719+00:00"
---



Are you an LLM? You can read better optimized documentation at /docs/guide/usage/transformer/lowering.md for this page in Markdown format

# Syntax Lowering [​](https://oxc.rs/docs/guide/usage/transformer/lowering.html#syntax-lowering)

Oxc transformer supports lowering ESNext to ES2015 syntax.

## Target [​](https://oxc.rs/docs/guide/usage/transformer/lowering.html#target)

Oxc transformer receives a `target` option to specify the target runtime. This will determine which syntaxes are lowered and which warnings are emitted.

Each target environment is an environment name followed by a version number. The following environment names are currently supported:

* `chrome`
* `deno`
* `edge`
* `firefox`
* `hermes`
* `ie`
* `ios`
* `node`
* `opera`
* `rhino`
* `safari`
* `samsung`
* `es`

The values that are supported by [esbuild's target option](https://esbuild.github.io/api/#target) are supported.

js

```
import { transform } from "oxc-transform";

const result = await transform("lib.js", "const foo = a ?? b;", {
  target: ["chrome87", "es2022"],
});
```

## Transformations [​](https://oxc.rs/docs/guide/usage/transformer/lowering.html#transformations)

Oxc supports lowering the syntaxes below. Note that RegExp related transformations only transforms the RegExp literal (`/foo/v`) to use a RegExp constructor (`new RegExp('foo', 'v')`). You will need to use a polyfill together to support older browsers.

### ES2026 [​](https://oxc.rs/docs/guide/usage/transformer/lowering.html#es2026)

* Explicit Resource Management (`using a = foo()`)

### ES2024 [​](https://oxc.rs/docs/guide/usage/transformer/lowering.html#es2024)

* RegExp v flag with set notation + properties of strings (`/\p{Emoji}--\p{ASCII}/v`)

### ES2022 [​](https://oxc.rs/docs/guide/usage/transformer/lowering.html#es2022)

* Class Static Block (`class A { static { foo() } }`)
* Class Fields (`class A { foo = 1; #bar = 2; static baz = 3; static qux = 4; foobar(a) { return #bar in a } }`)
* RegExp Match Indices (`/foo/d`)

### ES2021 [​](https://oxc.rs/docs/guide/usage/transformer/lowering.html#es2021)

* Logical Assignment Operators (`foo ||= bar`)
* Numeric separators (Note: this is not implemented as a transform, but the codegen always removes the separators)

### ES2020 [​](https://oxc.rs/docs/guide/usage/transformer/lowering.html#es2020)

* Nullish coalescing operator (`foo ?? bar`)
* Optional Chaining (`foo?.bar`)
* Export namespace from (`export * as foo from "bar"`)

### ES2019 [​](https://oxc.rs/docs/guide/usage/transformer/lowering.html#es2019)

* Optional `catch` binding (`try {} catch {}`)

### ES2018 [​](https://oxc.rs/docs/guide/usage/transformer/lowering.html#es2018)

* Rest/Spread Properties (`const foo = { a, b, ...c }`, `const { x, y, ...z } = foo;`)
* Asynchronous Iteration (`for await (const x of y) {}`, `async function* foo() {}`)
* RegExp Unicode Property Escapes (`/\p{Script=Greek}/u`)
* RegExp Lookbehind Assertions (`/(?<=foo)bar/`)
* RegExp named capture groups (`/(?<foo>bar)/`)
* `s` (`dotAll`) flag for regular expressions (`/foo./s`)

### ES2017 [​](https://oxc.rs/docs/guide/usage/transformer/lowering.html#es2017)

* Async functions (`async function foo() {}`)

### ES2016 [​](https://oxc.rs/docs/guide/usage/transformer/lowering.html#es2016)

* Exponentiation operator (`foo ** bar`)

### ES2015 [​](https://oxc.rs/docs/guide/usage/transformer/lowering.html#es2015)

* Arrow functions (`const foo = () => {}`)
* RegExp sticky flag (`/foo/y`)
* RegExp unicode flag (`/foo/u`)

## Warnings [​](https://oxc.rs/docs/guide/usage/transformer/lowering.html#warnings)

Oxc transformer emits warnings for the syntaxes below if the target runtime does not support them.

### ES2022 [​](https://oxc.rs/docs/guide/usage/transformer/lowering.html#es2022-1)

* Top level await (`await foo()`)
* Arbitrary module namespace identifiers (`import * as "f o o" from "bar"`)

### ES2020 [​](https://oxc.rs/docs/guide/usage/transformer/lowering.html#es2020-1)

* BigInt (`1n`)

## Compiler assumptions [​](https://oxc.rs/docs/guide/usage/transformer/lowering.html#compiler-assumptions)

You can specify assumptions for the compiler to make the output more smaller.

js

```
import { transform } from "oxc-transform";

const result = await transform("lib.js", "const foo = a ?? b;", {
  target: ["chrome87", "es2022"],
  assumptions: {
    noDocumentAll: true,
  },
});
```

The following assumptions are supported.

### `noDocumentAll` [​](https://oxc.rs/docs/guide/usage/transformer/lowering.html#nodocumentall)

Assume that the deprecated `document.all` with its special behavior is not used.

### `pureGetters` [​](https://oxc.rs/docs/guide/usage/transformer/lowering.html#puregetters)

Assume that getters does not have side effects.

### `setPublicClassFields` [​](https://oxc.rs/docs/guide/usage/transformer/lowering.html#setpublicclassfields)

When using public class fields, assume that they don't shadow any getter in the current class, in its subclasses or in its superclass. Thus, it's safe to assign them rather than using `Object.defineProperty`.

Note

For TypeScript, if you wanted behavior is equivalent to `useDefineForClassFields: false`, you should set both `setPublicClassFields` and `removeClassFieldsWithoutInitializer` to `true`. See [the TypeScript page](https://oxc.rs/docs/guide/usage/transformer/typescript.html#usedefineforclassfields) for more information.

## Not supported syntaxes [​](https://oxc.rs/docs/guide/usage/transformer/lowering.html#not-supported-syntaxes)

The following syntaxes are not lowered by Oxc transformer.

* ESNext
  + Decorators (tracked at [#9170](https://github.com/oxc-project/oxc/issues/9170)) (Note that [experimental decorators in TypeScript are supported](https://oxc.rs/docs/guide/usage/transformer/typescript.html#decorators))
* ES2025
  + RegExp Modifiers (tracked at [#11826](https://github.com/oxc-project/oxc/issues/11826))
  + Duplicate named capture groups (tracked at [#11827](https://github.com/oxc-project/oxc/issues/11827))
