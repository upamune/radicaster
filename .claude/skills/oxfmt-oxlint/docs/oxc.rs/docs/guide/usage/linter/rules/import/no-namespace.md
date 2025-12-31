---
title: "import/no-namespace | The JavaScript Oxidation Compiler"
source_url: "https://oxc.rs/docs/guide/usage/linter/rules/import/no-namespace"
fetched_at: "2025-12-31T10:44:14.234719+00:00"
---



Are you an LLM? You can read better optimized documentation at /docs/guide/usage/linter/rules/import/no-namespace.md for this page in Markdown format

# import/no-namespace Style [​](https://oxc.rs/docs/guide/usage/linter/rules/import/no-namespace.html#import-no-namespace)

🚧 An auto-fix is planned for this rule, but not implemented at this time.

### What it does [​](https://oxc.rs/docs/guide/usage/linter/rules/import/no-namespace.html#what-it-does)

Enforce a convention of not using namespaced (a.k.a. "wildcard" \*) imports.

### Why is this bad? [​](https://oxc.rs/docs/guide/usage/linter/rules/import/no-namespace.html#why-is-this-bad)

Namespaced imports, while sometimes used, are generally considered less ideal in modern JavaScript development for several reasons:

1. **Specificity and Namespace Pollution**:

* **Specificity**: Namespaced imports import the entire module, bringing in everything, even if you only need a few specific functions or classes. This can lead to potential naming conflicts if different modules have the same names for different functions.
* **Pollution**: Importing an entire namespace pollutes your current scope with potentially unnecessary functions and variables. It increases the chance of accidental use of an unintended function or variable, leading to harder-to-debug errors.

2. **Maintainability**:

* **Clarity**: Namespaced imports can make it harder to understand which specific functions or classes are being used in your code. This is especially true in larger projects with numerous imports.
* **Refactoring**: If a function or class name changes within the imported module, you might need to update several parts of your code if you are using namespaced imports. This becomes even more challenging when dealing with multiple namespaces.

3. **Modern Practice**:

* **Explicit Imports**: Modern JavaScript practices encourage explicit imports for specific components. This enhances code readability and maintainability.
* **Tree-Shaking**: Tools like Webpack and Rollup use tree-shaking to remove unused code from your bundles. Namespaced imports can prevent efficient tree-shaking, leading to larger bundle sizes.

### Examples [​](https://oxc.rs/docs/guide/usage/linter/rules/import/no-namespace.html#examples)

Examples of **incorrect** code for this rule:

js

```
import * as user from "user-lib";

import some, * as user from "./user";
```

Examples of **correct** code for this rule:

js

```
import { getUserName, isUser } from "user-lib";

import user from "user-lib";
import defaultExport, { isUser } from "./user";
```

## Configuration [​](https://oxc.rs/docs/guide/usage/linter/rules/import/no-namespace.html#configuration)

This rule accepts a configuration object with the following properties:

### ignore [​](https://oxc.rs/docs/guide/usage/linter/rules/import/no-namespace.html#ignore)

type: `string[]`

default: `[]`

An array of glob strings for modules that should be ignored by the rule. For example, `["*.json"]` will ignore all JSON imports.

## How to use [​](https://oxc.rs/docs/guide/usage/linter/rules/import/no-namespace.html#how-to-use)

To **enable** this rule using the config file or in the CLI, you can use:

Config (.oxlintrc.json)CLI

json

```
{
  "plugins": ["import"],
  "rules": {
    "import/no-namespace": "error"
  }
}
```

bash

```
oxlint --deny import/no-namespace --import-plugin
```

## References [​](https://oxc.rs/docs/guide/usage/linter/rules/import/no-namespace.html#references)

* [Rule Source](https://github.com/oxc-project/oxc/blob/1bf0ffc0f6859c90409a9701e62e8957ef1286cc/crates/oxc_linter/src/rules/import/no_namespace.rs)
