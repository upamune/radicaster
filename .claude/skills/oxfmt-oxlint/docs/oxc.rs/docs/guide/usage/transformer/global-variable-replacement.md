---
title: "Global Variable Replacement | The JavaScript Oxidation Compiler"
source_url: "https://oxc.rs/docs/guide/usage/transformer/global-variable-replacement"
fetched_at: "2025-12-31T10:44:14.234719+00:00"
---



Are you an LLM? You can read better optimized documentation at /docs/guide/usage/transformer/global-variable-replacement.md for this page in Markdown format

# Global Variable Replacement [​](https://oxc.rs/docs/guide/usage/transformer/global-variable-replacement.html#global-variable-replacement)

Oxc transformer supports replacing global variables.

## Define [​](https://oxc.rs/docs/guide/usage/transformer/global-variable-replacement.html#define)

"Define" feature provides a way to replace global variables with constant expressions. This feature is similar to [Terser](https://terser.org/)'s `global_defs` option and [esbuild's `define` option](https://esbuild.github.io/api/#define).

js

```
// input
const foo = __DEV__ ? 1 : 2;

// output
const foo = 1;
```

js

```
// Example
import { transform } from "oxc-transform";

const result = await transform("lib.js", sourceCode, {
  define: {
    __DEV__: "true",
  },
});
```

Each `define` entry maps an expression to a string of code containing an expression. The keys of it must be an identifier (e.g. `__DEV__`), or a dotted sequence of identifiers (e.g. `process.env.NODE_ENV`, `import.meta.env.MODE`). The values of it must be a valid expression.

Always quote the values

The values of `define` are the string of expressions. This means the value should always a string. If you mean a string literal, you should quote it (e.g. `__MODE__: '"development"'`, `__MODE__: JSON.stringify("development")`).

The object reference are not shared

Differently from esbuild, when passing an object to the value of the `define` option, the object reference is not shared. This means that if you change the object, the changes will not be reflected in the other places.

js

```
const foo = __OBJECT__;
foo.bar = 1;
console.log(foo.bar); // 1

const bar = __OBJECT__;
console.log(foo.bar); // undefined
```

js

```
// Example
import { transform } from "oxc-transform";

const result = await transform("lib.js", sourceCode, {
  define: {
    __OBJECT__: "{}",
  },
});
```

## Inject [​](https://oxc.rs/docs/guide/usage/transformer/global-variable-replacement.html#inject)

"Inject" feature provides a way to replace global variables with an import from a module. This feature is similar to [esbuild's `inject` option](https://esbuild.github.io/api/#inject) and [`@rollup/plugin-inject`](https://github.com/rollup/plugins/tree/master/packages/inject).

js

```
// input
const foo = new Promise((resolve) => resolve(1));

// output
import { Promise as P } from "es6-promise";
const foo = new P((resolve) => resolve(1));
```

js

```
// Example
import { transform } from "oxc-transform";

const result = await transform("lib.js", sourceCode, {
  inject: {
    P: ["es6-promise", "Promise"],
  },
});
```

Each `inject` entry maps an expression to an imported identifier. The keys of it must be an identifier (e.g. `__DEV__`), or a dotted sequence of identifiers (e.g. `process.env.NODE_ENV`). The values of it must be a string of the import source, or a tuple of strings of the import source and the import name (`*` is namespace import).

js

```
const examples = {
  // import { Promise } from 'es6-promise'
  Promise: ["es6-promise", "Promise"],
  // import { Promise as P } from 'es6-promise'
  P: ["es6-promise", "Promise"],
  // import $ from 'jquery'
  $: "jquery",
  // import * as fs from 'fs'
  fs: ["fs", "*"],
  // use a local module instead of a third-party one
  "Object.assign": path.resolve("src/helpers/object-assign.js"),
};
```
