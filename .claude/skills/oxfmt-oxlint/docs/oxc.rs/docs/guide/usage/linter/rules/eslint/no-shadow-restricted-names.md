---
title: "eslint/no-shadow-restricted-names | The JavaScript Oxidation Compiler"
source_url: "https://oxc.rs/docs/guide/usage/linter/rules/eslint/no-shadow-restricted-names"
fetched_at: "2025-12-31T10:44:14.234719+00:00"
---



Are you an LLM? You can read better optimized documentation at /docs/guide/usage/linter/rules/eslint/no-shadow-restricted-names.md for this page in Markdown format

# eslint/no-shadow-restricted-names Correctness [​](https://oxc.rs/docs/guide/usage/linter/rules/eslint/no-shadow-restricted-names.html#eslint-no-shadow-restricted-names)

✅ This rule is turned on by default.

### What it does [​](https://oxc.rs/docs/guide/usage/linter/rules/eslint/no-shadow-restricted-names.html#what-it-does)

Disallows the redefining of global variables such as `undefined`, `NaN`, `Infinity`, `eval`, and `arguments`.

### Why is this bad? [​](https://oxc.rs/docs/guide/usage/linter/rules/eslint/no-shadow-restricted-names.html#why-is-this-bad)

Value properties of the Global Object `NaN`, `Infinity`, `undefined` as well as the strict mode restricted identifiers `eval` and `arguments` are considered to be restricted names in JavaScript. Defining them to mean something else can have unintended consequences and confuse others reading the code. For example, there’s nothing preventing you from writing:

javascript

```
var undefined = "foo";
```

Then any code used within the same scope would not get the global undefined, but rather the local version with a very different meaning.

### Examples [​](https://oxc.rs/docs/guide/usage/linter/rules/eslint/no-shadow-restricted-names.html#examples)

Examples of **incorrect** code for this rule:

javascript

```
function NaN() {}

!function (Infinity) {};

var undefined = 5;

try {
} catch (eval) {}
```

javascript

```
import NaN from "foo";

import { undefined } from "bar";

class Infinity {}
```

Examples of **correct** code for this rule:

javascript

```
var Object;

function f(a, b) {}

// Exception: `undefined` may be shadowed if the variable is never assigned a value.
var undefined;
```

javascript

```
import { undefined as undef } from "bar";
```

## Configuration [​](https://oxc.rs/docs/guide/usage/linter/rules/eslint/no-shadow-restricted-names.html#configuration)

This rule accepts a configuration object with the following properties:

### reportGlobalThis [​](https://oxc.rs/docs/guide/usage/linter/rules/eslint/no-shadow-restricted-names.html#reportglobalthis)

type: `boolean`

default: `false`

If true, also report shadowing of `globalThis`.

## How to use [​](https://oxc.rs/docs/guide/usage/linter/rules/eslint/no-shadow-restricted-names.html#how-to-use)

To **enable** this rule using the config file or in the CLI, you can use:

Config (.oxlintrc.json)CLI

json

```
{
  "rules": {
    "no-shadow-restricted-names": "error"
  }
}
```

bash

```
oxlint --deny no-shadow-restricted-names
```

## References [​](https://oxc.rs/docs/guide/usage/linter/rules/eslint/no-shadow-restricted-names.html#references)

* [Rule Source](https://github.com/oxc-project/oxc/blob/1bf0ffc0f6859c90409a9701e62e8957ef1286cc/crates/oxc_linter/src/rules/eslint/no_shadow_restricted_names.rs)
