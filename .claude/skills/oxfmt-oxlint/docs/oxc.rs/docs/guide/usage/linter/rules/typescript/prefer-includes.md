---
title: "typescript/prefer-includes | The JavaScript Oxidation Compiler"
source_url: "https://oxc.rs/docs/guide/usage/linter/rules/typescript/prefer-includes"
fetched_at: "2025-12-31T10:44:14.234719+00:00"
---



Are you an LLM? You can read better optimized documentation at /docs/guide/usage/linter/rules/typescript/prefer-includes.md for this page in Markdown format

# typescript/prefer-includes Pedantic [​](https://oxc.rs/docs/guide/usage/linter/rules/typescript/prefer-includes.html#typescript-prefer-includes)

💭 This rule requires [type information](https://oxc.rs/docs/guide/usage/linter/type-aware.html).

🚧 An auto-fix is planned for this rule, but not implemented at this time.

### What it does [​](https://oxc.rs/docs/guide/usage/linter/rules/typescript/prefer-includes.html#what-it-does)

Enforce using `.includes()` instead of `.indexOf() !== -1` or `/regex/.test()`.

### Why is this bad? [​](https://oxc.rs/docs/guide/usage/linter/rules/typescript/prefer-includes.html#why-is-this-bad)

`.includes()` is more readable and expressive than checking `.indexOf() !== -1`. It clearly communicates the intent to check for the presence of a value. Additionally, for simple string searches, `.includes()` is often preferred over regex `.test()` for better performance and clarity.

### Examples [​](https://oxc.rs/docs/guide/usage/linter/rules/typescript/prefer-includes.html#examples)

Examples of **incorrect** code for this rule:

ts

```
// Using indexOf
const str = "hello world";
if (str.indexOf("world") !== -1) {
  console.log("found");
}

if (str.indexOf("world") != -1) {
  console.log("found");
}

if (str.indexOf("world") > -1) {
  console.log("found");
}

// Using regex test for simple strings
if (/world/.test(str)) {
  console.log("found");
}

// Arrays
const arr = [1, 2, 3];
if (arr.indexOf(2) !== -1) {
  console.log("found");
}
```

Examples of **correct** code for this rule:

ts

```
// Using includes for strings
const str = "hello world";
if (str.includes("world")) {
  console.log("found");
}

// Using includes for arrays
const arr = [1, 2, 3];
if (arr.includes(2)) {
  console.log("found");
}

// Complex regex patterns are allowed
if (/wo+rld/.test(str)) {
  console.log("found");
}

// Regex with flags
if (/world/i.test(str)) {
  console.log("found");
}
```

## How to use [​](https://oxc.rs/docs/guide/usage/linter/rules/typescript/prefer-includes.html#how-to-use)

To **enable** this rule using the config file or in the CLI, you can use:

Config (.oxlintrc.json)CLI

json

```
{
  "rules": {
    "typescript/prefer-includes": "error"
  }
}
```

bash

```
oxlint --type-aware --deny typescript/prefer-includes
```

## References [​](https://oxc.rs/docs/guide/usage/linter/rules/typescript/prefer-includes.html#references)

* [Rule Source](https://github.com/oxc-project/oxc/blob/1bf0ffc0f6859c90409a9701e62e8957ef1286cc/crates/oxc_linter/src/rules/typescript/prefer_includes.rs)
* [Rule Source (tsgolint)](https://github.com/oxc-project/tsgolint/blob/main/internal/rules/prefer_includes/prefer_includes.go)
