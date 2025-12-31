---
title: "nextjs/no-assign-module-variable | The JavaScript Oxidation Compiler"
source_url: "https://oxc.rs/docs/guide/usage/linter/rules/nextjs/no-assign-module-variable"
fetched_at: "2025-12-31T10:44:14.234719+00:00"
---



Are you an LLM? You can read better optimized documentation at /docs/guide/usage/linter/rules/nextjs/no-assign-module-variable.md for this page in Markdown format

# nextjs/no-assign-module-variable Correctness [​](https://oxc.rs/docs/guide/usage/linter/rules/nextjs/no-assign-module-variable.html#nextjs-no-assign-module-variable)

### What it does [​](https://oxc.rs/docs/guide/usage/linter/rules/nextjs/no-assign-module-variable.html#what-it-does)

Prevents the assignment or declaration of variables named `module` in Next.js applications.

### Why is this bad? [​](https://oxc.rs/docs/guide/usage/linter/rules/nextjs/no-assign-module-variable.html#why-is-this-bad)

The variable name `module` is reserved in Next.js for internal use and module system functionality. Declaring your own `module` variable can conflict with Next.js's internal module system, lead to unexpected behavior in your application, and cause issues with code splitting and hot module replacement.

### Examples [​](https://oxc.rs/docs/guide/usage/linter/rules/nextjs/no-assign-module-variable.html#examples)

Examples of **incorrect** code for this rule:

javascript

```
// Declaring module variable
let module = {};

// Using module in variable declaration
const module = {
  exports: {},
};

// Assigning to module
module = { id: "my-module" };
```

Examples of **correct** code for this rule:

javascript

```
// Use a different variable name
let myModule = {};

// Use a more descriptive name
const customModule = {
  exports: {},
};

// Access actual module object (when available)
console.log(module.exports);
```

## How to use [​](https://oxc.rs/docs/guide/usage/linter/rules/nextjs/no-assign-module-variable.html#how-to-use)

To **enable** this rule using the config file or in the CLI, you can use:

Config (.oxlintrc.json)CLI

json

```
{
  "plugins": ["nextjs"],
  "rules": {
    "nextjs/no-assign-module-variable": "error"
  }
}
```

bash

```
oxlint --deny nextjs/no-assign-module-variable --nextjs-plugin
```

## References [​](https://oxc.rs/docs/guide/usage/linter/rules/nextjs/no-assign-module-variable.html#references)

* [Rule Source](https://github.com/oxc-project/oxc/blob/1bf0ffc0f6859c90409a9701e62e8957ef1286cc/crates/oxc_linter/src/rules/nextjs/no_assign_module_variable.rs)
