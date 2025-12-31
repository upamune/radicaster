---
title: "Mangling | The JavaScript Oxidation Compiler"
source_url: "https://oxc.rs/docs/guide/usage/minifier/mangling"
fetched_at: "2025-12-31T10:44:14.234719+00:00"
---



Are you an LLM? You can read better optimized documentation at /docs/guide/usage/minifier/mangling.md for this page in Markdown format

# Mangling [​](https://oxc.rs/docs/guide/usage/minifier/mangling.html#mangling)

Oxc minifier supports mangling variable names and private class fields.

This feature is enabled by default and can be disabled by setting the `mangle` option to `false`.

## Top Level Variables [​](https://oxc.rs/docs/guide/usage/minifier/mangling.html#top-level-variables)

Top level variables are not mangled by default for non module code. You can enable mangling for top level variables by setting the `mangle.toplevel` option to `true`.

js

```
// input
var foo = 1;

// output
var e = 1;
```

js

```
// Example
import { minify } from "oxc-minify";

const result = await minify("lib.js", code, {
  module: false, // non-module code
  compress: {
    mangle: {
      toplevel: true,
    },
  },
});
```

## Keep `name` Property Values [​](https://oxc.rs/docs/guide/usage/minifier/mangling.html#keep-name-property-values)

Mangling variable names can change the `name` property values of functions / classes. You can keep the original `name` property values by enabling the `mangle.keepNames` option.

js

```
// input
var foo = function () {};

// output
var foo = function () {};
```

js

```
// Example
import { minify } from "oxc-minify";

const result = await minify("lib.js", code, {
  compress: {
    mangle: {
      keepNames: true, // shorthand of { function: true, class: true }
    },
  },
});
```

`compress.keepNames` option

When enabling this option, you may also want to enable [the `compress.keepNames` option](https://oxc.rs/docs/guide/usage/minifier/dead-code-elimination.html#keep-name-property-values).

## Debugging The Mangler [​](https://oxc.rs/docs/guide/usage/minifier/mangling.html#debugging-the-mangler)

To debug the mangler, you can enable the `mangle.debug` option. When this option is enabled, the mangler will use `slot_0`, `slot_1`, ... as variable names.

js

```
// input
var foo = 1;

// output
var slot_0 = 1;
```

js

```
// Example
import { minify } from "oxc-minify";

const result = await minify("lib.js", code, {
  compress: {
    mangle: {
      debug: true,
    },
  },
});
```
