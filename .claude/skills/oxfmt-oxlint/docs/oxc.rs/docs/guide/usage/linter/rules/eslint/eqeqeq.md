---
title: "eslint/eqeqeq | The JavaScript Oxidation Compiler"
source_url: "https://oxc.rs/docs/guide/usage/linter/rules/eslint/eqeqeq"
fetched_at: "2025-12-31T10:44:14.234719+00:00"
---



Are you an LLM? You can read better optimized documentation at /docs/guide/usage/linter/rules/eslint/eqeqeq.md for this page in Markdown format

# eslint/eqeqeq Pedantic [​](https://oxc.rs/docs/guide/usage/linter/rules/eslint/eqeqeq.html#eslint-eqeqeq)

⚠️🛠️️ A dangerous auto-fix is available for this rule for some violations.

### What it does [​](https://oxc.rs/docs/guide/usage/linter/rules/eslint/eqeqeq.html#what-it-does)

Requires the use of the `===` and `!==` operators, disallowing the use of `==` and `!=`.

### Why is this bad? [​](https://oxc.rs/docs/guide/usage/linter/rules/eslint/eqeqeq.html#why-is-this-bad)

Using non-strict equality operators leads to unexpected behavior due to type coercion, which can cause hard-to-find bugs.

### Options [​](https://oxc.rs/docs/guide/usage/linter/rules/eslint/eqeqeq.html#options)

First option:

* Type: `string`
* Default: `"always"`

Possible values:

* `"always"` - always require `===`/`!==`
* `"smart"` - allow safe comparisons (`typeof`, literals, nullish)

Second option (only used with `"always"`):

* Type: `object`
* Properties:
  + `null`: `string` (default: `"always"`) - `"ignore"` allows `== null` and `!= null`.

Possible values for `null`:

* `"always"` - always require `=== null`/`!== null`
* `"never"` - always require `== null`/`!= null`
* `"ignore"` - allow both `== null`/`!= null` and `=== null`/`!== null`

Example JSON configuration:

json

```
{
  "eqeqeq": ["error", "always", { "null": "ignore" }]
}
```

### Examples [​](https://oxc.rs/docs/guide/usage/linter/rules/eslint/eqeqeq.html#examples)

#### `"always"` (default) [​](https://oxc.rs/docs/guide/usage/linter/rules/eslint/eqeqeq.html#always-default)

Examples of **incorrect** code for this rule:

js

```
/* eslint eqeqeq: "error" */

if (x == 42) {
}
if ("" == text) {
}
if (obj.getStuff() != undefined) {
}
```

Examples of **correct** code for this rule:

js

```
/* eslint eqeqeq: "error" */

if (x === 42) {
}
if ("" === text) {
}
if (obj.getStuff() !== undefined) {
}
```

#### `"smart"` [​](https://oxc.rs/docs/guide/usage/linter/rules/eslint/eqeqeq.html#smart)

Examples of **incorrect** code for this rule with the `"smart"` option:

js

```
/* eslint eqeqeq: ["error", "smart"] */

if (x == 42) {
}
if ("" == text) {
}
```

Examples of **correct** code for this rule with the `"smart"` option:

js

```
/* eslint eqeqeq: ["error", "smart"] */

if (typeof foo == "undefined") {
}
if (foo == null) {
}
if (foo != null) {
}
```

#### `{"null": "ignore"}` (with `"always"` first option) [​](https://oxc.rs/docs/guide/usage/linter/rules/eslint/eqeqeq.html#null-ignore-with-always-first-option)

Examples of **incorrect** code for this rule with the `{ "null": "ignore" }` option:

js

```
/* eslint eqeqeq: ["error", "always", { "null": "ignore" }] */
if (x == 42) {
}
if ("" == text) {
}
```

Examples of **correct** code for this rule with the `{ "null": "ignore" }` option:

js

```
/* eslint eqeqeq: ["error", "always", { "null": "ignore" }] */
if (foo == null) {
}
if (foo != null) {
}
```

#### `{"null": "always"}` (default - with `"always"` first option) [​](https://oxc.rs/docs/guide/usage/linter/rules/eslint/eqeqeq.html#null-always-default-with-always-first-option)

Examples of **incorrect** code for this rule with the `{ "null": "always" }` option:

js

```
/* eslint eqeqeq: ["error", "always", { "null": "always" }] */

if (foo == null) {
}
if (foo != null) {
}
```

Examples of **correct** code for this rule with the `{ "null": "always" }` option:

js

```
/* eslint eqeqeq: ["error", "always", { "null": "always" }] */

if (foo === null) {
}
if (foo !== null) {
}
```

#### `{"null": "never"}` (with `"always"` first option) [​](https://oxc.rs/docs/guide/usage/linter/rules/eslint/eqeqeq.html#null-never-with-always-first-option)

Examples of **incorrect** code for this rule with the `{ "null": "never" }` option:

js

```
/* eslint eqeqeq: ["error", "always", { "null": "never" }] */

if (x == 42) {
}
if ("" == text) {
}
if (foo === null) {
}
if (foo !== null) {
}
```

Examples of **correct** code for this rule with the `{ "null": "never" }` option:

js

```
/* eslint eqeqeq: ["error", "always", { "null": "never" }] */

if (x === 42) {
}
if ("" === text) {
}
if (foo == null) {
}
if (foo != null) {
}
```

## Configuration [​](https://oxc.rs/docs/guide/usage/linter/rules/eslint/eqeqeq.html#configuration)

This rule accepts a configuration object with the following properties:

### compareType [​](https://oxc.rs/docs/guide/usage/linter/rules/eslint/eqeqeq.html#comparetype)

type: `"always" | "smart"`

### nullType [​](https://oxc.rs/docs/guide/usage/linter/rules/eslint/eqeqeq.html#nulltype)

type: `"always" | "never" | "ignore"`

## How to use [​](https://oxc.rs/docs/guide/usage/linter/rules/eslint/eqeqeq.html#how-to-use)

To **enable** this rule using the config file or in the CLI, you can use:

Config (.oxlintrc.json)CLI

json

```
{
  "rules": {
    "eqeqeq": "error"
  }
}
```

bash

```
oxlint --deny eqeqeq
```

## References [​](https://oxc.rs/docs/guide/usage/linter/rules/eslint/eqeqeq.html#references)

* [Rule Source](https://github.com/oxc-project/oxc/blob/1bf0ffc0f6859c90409a9701e62e8957ef1286cc/crates/oxc_linter/src/rules/eslint/eqeqeq.rs)
