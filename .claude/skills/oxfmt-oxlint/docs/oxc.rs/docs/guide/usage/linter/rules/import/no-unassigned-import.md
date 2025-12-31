---
title: "import/no-unassigned-import | The JavaScript Oxidation Compiler"
source_url: "https://oxc.rs/docs/guide/usage/linter/rules/import/no-unassigned-import"
fetched_at: "2025-12-31T10:44:14.234719+00:00"
---



Are you an LLM? You can read better optimized documentation at /docs/guide/usage/linter/rules/import/no-unassigned-import.md for this page in Markdown format

# import/no-unassigned-import Suspicious [​](https://oxc.rs/docs/guide/usage/linter/rules/import/no-unassigned-import.html#import-no-unassigned-import)

### What it does [​](https://oxc.rs/docs/guide/usage/linter/rules/import/no-unassigned-import.html#what-it-does)

This rule aims to remove modules with side-effects by reporting when a module is imported but not assigned.

### Why is this bad? [​](https://oxc.rs/docs/guide/usage/linter/rules/import/no-unassigned-import.html#why-is-this-bad)

With both CommonJS' require and the ES modules' import syntax, it is possible to import a module but not to use its result. This can be done explicitly by not assigning the module to a variable. Doing so can mean either of the following things:

* The module is imported but not used
* The module has side-effects. Having side-effects, makes it hard to know whether the module is actually used or can be removed. It can also make it harder to test or mock parts of your application.

### Examples [​](https://oxc.rs/docs/guide/usage/linter/rules/import/no-unassigned-import.html#examples)

Examples of **incorrect** code for this rule:

js

```
import "should";
require("should");
```

Examples of **correct** code for this rule:

js

```
import _ from "foo";
import _, { foo } from "foo";
import _, { foo as bar } from "foo";
const _ = require("foo");
const { foo } = require("foo");
const { foo: bar } = require("foo");
bar(require("foo"));
```

## Configuration [​](https://oxc.rs/docs/guide/usage/linter/rules/import/no-unassigned-import.html#configuration)

This rule accepts a configuration object with the following properties:

### allow [​](https://oxc.rs/docs/guide/usage/linter/rules/import/no-unassigned-import.html#allow)

type: `string[]`

default: `[]`

A list of glob patterns to allow unassigned imports for specific modules. For example: `{ "allow": ["*.css"] }` will allow unassigned imports for any module ending with `.css`.

## How to use [​](https://oxc.rs/docs/guide/usage/linter/rules/import/no-unassigned-import.html#how-to-use)

To **enable** this rule using the config file or in the CLI, you can use:

Config (.oxlintrc.json)CLI

json

```
{
  "plugins": ["import"],
  "rules": {
    "import/no-unassigned-import": "error"
  }
}
```

bash

```
oxlint --deny import/no-unassigned-import --import-plugin
```

## References [​](https://oxc.rs/docs/guide/usage/linter/rules/import/no-unassigned-import.html#references)

* [Rule Source](https://github.com/oxc-project/oxc/blob/1bf0ffc0f6859c90409a9701e62e8957ef1286cc/crates/oxc_linter/src/rules/import/no_unassigned_import.rs)
