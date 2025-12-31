---
title: "import/no-absolute-path | The JavaScript Oxidation Compiler"
source_url: "https://oxc.rs/docs/guide/usage/linter/rules/import/no-absolute-path"
fetched_at: "2025-12-31T10:44:14.234719+00:00"
---



Are you an LLM? You can read better optimized documentation at /docs/guide/usage/linter/rules/import/no-absolute-path.md for this page in Markdown format

# import/no-absolute-path Suspicious [​](https://oxc.rs/docs/guide/usage/linter/rules/import/no-absolute-path.html#import-no-absolute-path)

🚧 An auto-fix is planned for this rule, but not implemented at this time.

### What it does [​](https://oxc.rs/docs/guide/usage/linter/rules/import/no-absolute-path.html#what-it-does)

This rule forbids the import of modules using absolute paths.

### Why is this bad? [​](https://oxc.rs/docs/guide/usage/linter/rules/import/no-absolute-path.html#why-is-this-bad)

Node.js allows the import of modules using an absolute path such as `/home/xyz/file.js`. That is a bad practice as it ties the code using it to your computer, and therefore makes it unusable in packages distributed on npm for instance.

### Examples [​](https://oxc.rs/docs/guide/usage/linter/rules/import/no-absolute-path.html#examples)

Examples of **incorrect** code for this rule:

js

```
import f from "/foo";
import f from "/some/path";
var f = require("/foo");
var f = require("/some/path");
```

Examples of **correct** code for this rule:

js

```
import _ from "lodash";
import foo from "foo";
import foo from "./foo";

var _ = require("lodash");
var foo = require("foo");
var foo = require("./foo");
```

Examples of **incorrect** code for the `{ amd: true }` option:

js

```
define("/foo", function (foo) {});
require("/foo", function (foo) {});
```

Examples of **correct** code for the `{ amd: true }` option:

js

```
define("./foo", function (foo) {});
require("./foo", function (foo) {});
```

## Configuration [​](https://oxc.rs/docs/guide/usage/linter/rules/import/no-absolute-path.html#configuration)

This rule accepts a configuration object with the following properties:

### amd [​](https://oxc.rs/docs/guide/usage/linter/rules/import/no-absolute-path.html#amd)

type: `boolean`

default: `false`

If set to `true`, dependency paths for AMD-style define and require calls will be resolved:

js

```
/* eslint import/no-absolute-path: ['error', { commonjs: false, amd: true }] */
define(["/foo"], function (foo) {
  /*...*/
}); // reported
require(["/foo"], function (foo) {
  /*...*/
}); // reported

const foo = require("/foo"); // ignored because of explicit `commonjs: false`
```

### commonjs [​](https://oxc.rs/docs/guide/usage/linter/rules/import/no-absolute-path.html#commonjs)

type: `boolean`

default: `true`

If set to `true`, dependency paths for CommonJS-style require calls will be resolved:

js

```
var foo = require("/foo"); // reported
```

### esmodule [​](https://oxc.rs/docs/guide/usage/linter/rules/import/no-absolute-path.html#esmodule)

type: `boolean`

default: `true`

If set to `true`, dependency paths for ES module import statements will be resolved:

js

```
import foo from "/foo"; // reported
```

## How to use [​](https://oxc.rs/docs/guide/usage/linter/rules/import/no-absolute-path.html#how-to-use)

To **enable** this rule using the config file or in the CLI, you can use:

Config (.oxlintrc.json)CLI

json

```
{
  "plugins": ["import"],
  "rules": {
    "import/no-absolute-path": "error"
  }
}
```

bash

```
oxlint --deny import/no-absolute-path --import-plugin
```

## References [​](https://oxc.rs/docs/guide/usage/linter/rules/import/no-absolute-path.html#references)

* [Rule Source](https://github.com/oxc-project/oxc/blob/1bf0ffc0f6859c90409a9701e62e8957ef1286cc/crates/oxc_linter/src/rules/import/no_absolute_path.rs)
