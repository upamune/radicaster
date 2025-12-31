---
title: "unicorn/filename-case | The JavaScript Oxidation Compiler"
source_url: "https://oxc.rs/docs/guide/usage/linter/rules/unicorn/filename-case"
fetched_at: "2025-12-31T10:44:14.234719+00:00"
---



Are you an LLM? You can read better optimized documentation at /docs/guide/usage/linter/rules/unicorn/filename-case.md for this page in Markdown format

# unicorn/filename-case Style [​](https://oxc.rs/docs/guide/usage/linter/rules/unicorn/filename-case.html#unicorn-filename-case)

### What it does [​](https://oxc.rs/docs/guide/usage/linter/rules/unicorn/filename-case.html#what-it-does)

Enforces a consistent case style for filenames to improve project organization and maintainability. By default, `kebab-case` is enforced, but other styles can be configured.

Files named `index.js`, `index.ts`, etc. are exempt from this rule as they cannot reliably be renamed to other casings (mainly just a problem with PascalCase).

### Why is this bad? [​](https://oxc.rs/docs/guide/usage/linter/rules/unicorn/filename-case.html#why-is-this-bad)

Inconsistent file naming conventions make it harder to locate files, navigate projects, and enforce consistency across a codebase. Standardizing naming conventions improves readability, reduces cognitive overhead, and aligns with best practices in large-scale development.

### Examples [​](https://oxc.rs/docs/guide/usage/linter/rules/unicorn/filename-case.html#examples)

Examples of **correct** filenames for each case:

#### `kebabCase` [​](https://oxc.rs/docs/guide/usage/linter/rules/unicorn/filename-case.html#kebabcase)

* `some-file-name.js`
* `some-file-name.test.js`
* `some-file-name.test-utils.js`

#### `camelCase` [​](https://oxc.rs/docs/guide/usage/linter/rules/unicorn/filename-case.html#camelcase)

* `someFileName.js`
* `someFileName.test.js`
* `someFileName.testUtils.js`

#### `snakeCase` [​](https://oxc.rs/docs/guide/usage/linter/rules/unicorn/filename-case.html#snakecase)

* `some_file_name.js`
* `some_file_name.test.js`
* `some_file_name.test_utils.js`

#### `pascalCase` [​](https://oxc.rs/docs/guide/usage/linter/rules/unicorn/filename-case.html#pascalcase)

* `SomeFileName.js`
* `SomeFileName.Test.js`
* `SomeFileName.TestUtils.js`

## Configuration [​](https://oxc.rs/docs/guide/usage/linter/rules/unicorn/filename-case.html#configuration)

This rule accepts a configuration object with the following properties:

### case [​](https://oxc.rs/docs/guide/usage/linter/rules/unicorn/filename-case.html#case)

type: `"kebabCase" | "camelCase" | "snakeCase" | "pascalCase"`

default: `"kebabCase"`

The case style to enforce for filenames.

You can set the `case` option like this:

json

```
"unicorn/filename-case": [
"error",
{
"case": "kebabCase"
}
]
```

### cases [​](https://oxc.rs/docs/guide/usage/linter/rules/unicorn/filename-case.html#cases)

type: `object`

default: `{"kebabCase":true, "camelCase":false, "snakeCase":false, "pascalCase":false}`

The case style(s) to allow/enforce for filenames. `true` means the case style is allowed, `false` means it is banned.

You can set the `cases` option like this:

json

```
"unicorn/filename-case": [
"error",
{
"cases": {
"camelCase": true,
"pascalCase": true
}
}
]
```

#### cases.camelCase [​](https://oxc.rs/docs/guide/usage/linter/rules/unicorn/filename-case.html#cases-camelcase)

type: `boolean`

default: `false`

Whether camel case is allowed, e.g. `someFileName.js`.

#### cases.kebabCase [​](https://oxc.rs/docs/guide/usage/linter/rules/unicorn/filename-case.html#cases-kebabcase)

type: `boolean`

default: `true`

Whether kebab case is allowed, e.g. `some-file-name.js`.

#### cases.pascalCase [​](https://oxc.rs/docs/guide/usage/linter/rules/unicorn/filename-case.html#cases-pascalcase)

type: `boolean`

default: `false`

Whether pascal case is allowed, e.g. `SomeFileName.js`.

#### cases.snakeCase [​](https://oxc.rs/docs/guide/usage/linter/rules/unicorn/filename-case.html#cases-snakecase)

type: `boolean`

default: `false`

Whether snake case is allowed, e.g. `some_file_name.js`.

### ignore [​](https://oxc.rs/docs/guide/usage/linter/rules/unicorn/filename-case.html#ignore)

type: `string | null`

A regular expression pattern for filenames to ignore.

You can set the `ignore` option like this:

json

```
"unicorn/filename-case": [
"error",
{
"ignore": "^foo.*$"
}
]
```

### multipleFileExtensions [​](https://oxc.rs/docs/guide/usage/linter/rules/unicorn/filename-case.html#multiplefileextensions)

type: `boolean`

default: `true`

Whether to treat additional, `.`-separated parts of a filename as parts of the extension rather than parts of the filename.

## How to use [​](https://oxc.rs/docs/guide/usage/linter/rules/unicorn/filename-case.html#how-to-use)

To **enable** this rule using the config file or in the CLI, you can use:

Config (.oxlintrc.json)CLI

json

```
{
  "rules": {
    "unicorn/filename-case": "error"
  }
}
```

bash

```
oxlint --deny unicorn/filename-case
```

## References [​](https://oxc.rs/docs/guide/usage/linter/rules/unicorn/filename-case.html#references)

* [Rule Source](https://github.com/oxc-project/oxc/blob/1bf0ffc0f6859c90409a9701e62e8957ef1286cc/crates/oxc_linter/src/rules/unicorn/filename_case.rs)
