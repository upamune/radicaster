---
title: "import/no-named-as-default-member | The JavaScript Oxidation Compiler"
source_url: "https://oxc.rs/docs/guide/usage/linter/rules/import/no-named-as-default-member"
fetched_at: "2025-12-31T10:44:14.234719+00:00"
---



Are you an LLM? You can read better optimized documentation at /docs/guide/usage/linter/rules/import/no-named-as-default-member.md for this page in Markdown format

# import/no-named-as-default-member Suspicious [​](https://oxc.rs/docs/guide/usage/linter/rules/import/no-named-as-default-member.html#import-no-named-as-default-member)

### What it does [​](https://oxc.rs/docs/guide/usage/linter/rules/import/no-named-as-default-member.html#what-it-does)

Reports the use of an exported name (named export) as a property on the default export. This occurs when trying to access a named export through the default export, which is incorrect.

### Why is this bad? [​](https://oxc.rs/docs/guide/usage/linter/rules/import/no-named-as-default-member.html#why-is-this-bad)

Accessing a named export via the default export is incorrect and will not work as expected. Named exports should be imported directly, while default exports are accessed without properties. This mistake can lead to runtime errors or undefined behavior.

### Examples [​](https://oxc.rs/docs/guide/usage/linter/rules/import/no-named-as-default-member.html#examples)

Given

javascript

```
// ./bar.js
export function bar() {
  return null;
}
export default () => {
  return 1;
};
```

Examples of **incorrect** code for this rule:

javascript

```
// ./foo.js
import foo from "./bar";
const bar = foo.bar; // Incorrect: trying to access named export via default
```

Examples of **correct** code for this rule:

javascript

```
// ./foo.js
import { bar } from "./bar"; // Correct: accessing named export directly
```

## How to use [​](https://oxc.rs/docs/guide/usage/linter/rules/import/no-named-as-default-member.html#how-to-use)

To **enable** this rule using the config file or in the CLI, you can use:

Config (.oxlintrc.json)CLI

json

```
{
  "plugins": ["import"],
  "rules": {
    "import/no-named-as-default-member": "error"
  }
}
```

bash

```
oxlint --deny import/no-named-as-default-member --import-plugin
```

## References [​](https://oxc.rs/docs/guide/usage/linter/rules/import/no-named-as-default-member.html#references)

* [Rule Source](https://github.com/oxc-project/oxc/blob/1bf0ffc0f6859c90409a9701e62e8957ef1286cc/crates/oxc_linter/src/rules/import/no_named_as_default_member.rs)
