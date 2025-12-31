---
title: "eslint/arrow-body-style | The JavaScript Oxidation Compiler"
source_url: "https://oxc.rs/docs/guide/usage/linter/rules/eslint/arrow-body-style"
fetched_at: "2025-12-31T10:44:14.234719+00:00"
---



Are you an LLM? You can read better optimized documentation at /docs/guide/usage/linter/rules/eslint/arrow-body-style.md for this page in Markdown format

# eslint/arrow-body-style Style [​](https://oxc.rs/docs/guide/usage/linter/rules/eslint/arrow-body-style.html#eslint-arrow-body-style)

🛠️ An auto-fix is available for this rule.

### What it does [​](https://oxc.rs/docs/guide/usage/linter/rules/eslint/arrow-body-style.html#what-it-does)

This rule can enforce or disallow the use of braces around arrow function body. Arrow functions can use either:

* a block body `() => { ... }`
* or a concise body `() => expression` with an implicit return.

### Why is this bad? [​](https://oxc.rs/docs/guide/usage/linter/rules/eslint/arrow-body-style.html#why-is-this-bad)

Inconsistent use of block vs. concise bodies makes code harder to read. Concise bodies are limited to a single expression, whose value is implicitly returned.

### Options [​](https://oxc.rs/docs/guide/usage/linter/rules/eslint/arrow-body-style.html#options)

First option:

* Type: `string`
* Enum: `"always"`, `"as-needed"`, `"never"`
* Default: `"as-needed"`

Possible values:

* `never` enforces no braces around the function body (constrains arrow functions to the role of returning an expression)
* `always` enforces braces around the function body
* `as-needed` enforces no braces where they can be omitted (default)

Second option:

* Type: `object`
* Properties:
  + `requireReturnForObjectLiteral`: `boolean` (default: `false`) - requires braces and an explicit return for object literals.

Note: This option only applies when the first option is `"as-needed"`.

Example configuration:

json

```
{
  "arrow-body-style": ["error", "as-needed", { "requireReturnForObjectLiteral": true }]
}
```

### Examples [​](https://oxc.rs/docs/guide/usage/linter/rules/eslint/arrow-body-style.html#examples)

#### `"never"` (default) [​](https://oxc.rs/docs/guide/usage/linter/rules/eslint/arrow-body-style.html#never-default)

Examples of **incorrect** code for this rule with the `never` option:

js

```
/* arrow-body-style: ["error", "never"] */

/* ✘ Bad: */
const foo = () => {
  return 0;
};
```

Examples of **correct** code for this rule with the `never` option:

js

```
/* arrow-body-style: ["error", "never"] */

/* ✔ Good: */
const foo = () => 0;
const bar = () => ({ foo: 0 });
```

#### `"always"` [​](https://oxc.rs/docs/guide/usage/linter/rules/eslint/arrow-body-style.html#always)

Examples of **incorrect** code for this rule with the `always` option:

js

```
/* arrow-body-style: ["error", "always"] */

/* ✘ Bad: */
const foo = () => 0;
```

Examples of **correct** code for this rule with the `always` option:

js

```
/* arrow-body-style: ["error", "always"] */

/* ✔ Good: */
const foo = () => {
  return 0;
};
```

#### `"as-needed"` [​](https://oxc.rs/docs/guide/usage/linter/rules/eslint/arrow-body-style.html#as-needed)

Examples of **incorrect** code for this rule with the `as-needed` option:

js

```
/* arrow-body-style: ["error", "as-needed"] */

/* ✘ Bad: */
const foo = () => {
  return 0;
};
```

Examples of **correct** code for this rule with the `as-needed` option:

js

```
/* arrow-body-style: ["error", "as-needed"] */

/* ✔ Good: */
const foo1 = () => 0;

const foo2 = (retv, name) => {
  retv[name] = true;
  return retv;
};

const foo3 = () => {
  bar();
};
```

#### `"as-needed"` with `requireReturnForObjectLiteral` [​](https://oxc.rs/docs/guide/usage/linter/rules/eslint/arrow-body-style.html#as-needed-with-requirereturnforobjectliteral)

Examples of **incorrect** code for this rule with the `{ "requireReturnForObjectLiteral": true }` option:

js

```
/* arrow-body-style: ["error", "as-needed", { "requireReturnForObjectLiteral": true }]*/

/* ✘ Bad: */
const foo = () => ({});
const bar = () => ({ bar: 0 });
```

Examples of **correct** code for this rule with the `{ "requireReturnForObjectLiteral": true }` option:

js

```
/* arrow-body-style: ["error", "as-needed", { "requireReturnForObjectLiteral": true }]*/

/* ✔ Good: */
const foo = () => {};
const bar = () => {
  return { bar: 0 };
};
```

## How to use [​](https://oxc.rs/docs/guide/usage/linter/rules/eslint/arrow-body-style.html#how-to-use)

To **enable** this rule using the config file or in the CLI, you can use:

Config (.oxlintrc.json)CLI

json

```
{
  "rules": {
    "arrow-body-style": "error"
  }
}
```

bash

```
oxlint --deny arrow-body-style
```

## References [​](https://oxc.rs/docs/guide/usage/linter/rules/eslint/arrow-body-style.html#references)

* [Rule Source](https://github.com/oxc-project/oxc/blob/1bf0ffc0f6859c90409a9701e62e8957ef1286cc/crates/oxc_linter/src/rules/eslint/arrow_body_style.rs)
