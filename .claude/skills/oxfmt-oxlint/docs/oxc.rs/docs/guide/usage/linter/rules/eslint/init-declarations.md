---
title: "eslint/init-declarations | The JavaScript Oxidation Compiler"
source_url: "https://oxc.rs/docs/guide/usage/linter/rules/eslint/init-declarations"
fetched_at: "2025-12-31T10:44:14.234719+00:00"
---



Are you an LLM? You can read better optimized documentation at /docs/guide/usage/linter/rules/eslint/init-declarations.md for this page in Markdown format

# eslint/init-declarations Style [​](https://oxc.rs/docs/guide/usage/linter/rules/eslint/init-declarations.html#eslint-init-declarations)

### What it does [​](https://oxc.rs/docs/guide/usage/linter/rules/eslint/init-declarations.html#what-it-does)

Require or disallow initialization in variable declarations

### Why is this bad? [​](https://oxc.rs/docs/guide/usage/linter/rules/eslint/init-declarations.html#why-is-this-bad)

In JavaScript, variables can be assigned during declaration, or at any point afterwards using an assignment statement. For example, in the following code, foo is initialized during declaration, while bar is initialized later.

### Examples [​](https://oxc.rs/docs/guide/usage/linter/rules/eslint/init-declarations.html#examples)

var foo = 1; var bar; if (foo) { bar = 1; } else { bar = 2; }

Examples of incorrect code for the default "always" option:

js

```
/*eslint init-declarations: ["error", "always"]*/
function foo() {
  var bar;
  let baz;
}
```

Examples of incorrect code for the "never" option:

js

```
/*eslint init-declarations: ["error", "never"]*/
function foo() {
  var bar = 1;
  let baz = 2;
  for (var i = 0; i < 1; i++) {}
}
```

Examples of correct code for the default "always" option:

js

```
/*eslint init-declarations: ["error", "always"]*/

function foo() {
  var bar = 1;
  let baz = 2;
  const qux = 3;
}
```

Examples of correct code for the "never" option:

js

```
/*eslint init-declarations: ["error", "never"]*/

function foo() {
  var bar;
  let baz;
  const buzz = 1;
}
```

Examples of correct code for the "never", { "ignoreForLoopInit": true } options:

js

```
/*eslint init-declarations: ["error", "never", { "ignoreForLoopInit": true }]*/
for (var i = 0; i < 1; i++) {}
```

## Configuration [​](https://oxc.rs/docs/guide/usage/linter/rules/eslint/init-declarations.html#configuration)

This rule accepts a configuration object with the following properties:

### ignoreForLoopInit [​](https://oxc.rs/docs/guide/usage/linter/rules/eslint/init-declarations.html#ignoreforloopinit)

type: `boolean`

default: `false`

When set to `true`, allows uninitialized variables in the init expression of `for`, `for-in`, and `for-of` loops. Only applies when mode is set to `"never"`.

### mode [​](https://oxc.rs/docs/guide/usage/linter/rules/eslint/init-declarations.html#mode)

type: `"always" | "never"`

When set to `"always"` (default), requires that variables be initialized on declaration. When set to `"never"`, disallows initialization during declaration.

## How to use [​](https://oxc.rs/docs/guide/usage/linter/rules/eslint/init-declarations.html#how-to-use)

To **enable** this rule using the config file or in the CLI, you can use:

Config (.oxlintrc.json)CLI

json

```
{
  "rules": {
    "init-declarations": "error"
  }
}
```

bash

```
oxlint --deny init-declarations
```

## References [​](https://oxc.rs/docs/guide/usage/linter/rules/eslint/init-declarations.html#references)

* [Rule Source](https://github.com/oxc-project/oxc/blob/1bf0ffc0f6859c90409a9701e62e8957ef1286cc/crates/oxc_linter/src/rules/eslint/init_declarations.rs)
