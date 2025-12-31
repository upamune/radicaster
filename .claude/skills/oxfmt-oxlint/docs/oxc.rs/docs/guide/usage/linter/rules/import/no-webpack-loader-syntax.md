---
title: "import/no-webpack-loader-syntax | The JavaScript Oxidation Compiler"
source_url: "https://oxc.rs/docs/guide/usage/linter/rules/import/no-webpack-loader-syntax"
fetched_at: "2025-12-31T10:44:14.234719+00:00"
---



Are you an LLM? You can read better optimized documentation at /docs/guide/usage/linter/rules/import/no-webpack-loader-syntax.md for this page in Markdown format

# import/no-webpack-loader-syntax Restriction [​](https://oxc.rs/docs/guide/usage/linter/rules/import/no-webpack-loader-syntax.html#import-no-webpack-loader-syntax)

### What it does [​](https://oxc.rs/docs/guide/usage/linter/rules/import/no-webpack-loader-syntax.html#what-it-does)

Forbids using Webpack loader syntax directly in import or require statements.

### Why is this bad? [​](https://oxc.rs/docs/guide/usage/linter/rules/import/no-webpack-loader-syntax.html#why-is-this-bad)

This loader syntax is non-standard, so it couples the code to Webpack. The recommended way to specify Webpack loader configuration is in a [Webpack configuration file](https://webpack.js.org/concepts/loaders/#configuration).

### Examples [​](https://oxc.rs/docs/guide/usage/linter/rules/import/no-webpack-loader-syntax.html#examples)

Examples of **incorrect** code for this rule:

javascript

```
import myModule from "my-loader!my-module";
import theme from "style!css!./theme.css";

var myModule = require("my-loader!./my-module");
var theme = require("style!css!./theme.css");
```

Examples of **correct** code for this rule:

javascript

```
import myModule from "./my-module";
import theme from "./theme.css";

var myModule = require("./my-module");
var theme = require("./theme.css");
```

## How to use [​](https://oxc.rs/docs/guide/usage/linter/rules/import/no-webpack-loader-syntax.html#how-to-use)

To **enable** this rule using the config file or in the CLI, you can use:

Config (.oxlintrc.json)CLI

json

```
{
  "plugins": ["import"],
  "rules": {
    "import/no-webpack-loader-syntax": "error"
  }
}
```

bash

```
oxlint --deny import/no-webpack-loader-syntax --import-plugin
```

## References [​](https://oxc.rs/docs/guide/usage/linter/rules/import/no-webpack-loader-syntax.html#references)

* [Rule Source](https://github.com/oxc-project/oxc/blob/1bf0ffc0f6859c90409a9701e62e8957ef1286cc/crates/oxc_linter/src/rules/import/no_webpack_loader_syntax.rs)
