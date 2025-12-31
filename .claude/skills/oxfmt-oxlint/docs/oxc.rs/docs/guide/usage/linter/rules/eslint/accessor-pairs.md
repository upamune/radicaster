---
title: "eslint/accessor-pairs | The JavaScript Oxidation Compiler"
source_url: "https://oxc.rs/docs/guide/usage/linter/rules/eslint/accessor-pairs"
fetched_at: "2025-12-31T10:44:14.234719+00:00"
---



Are you an LLM? You can read better optimized documentation at /docs/guide/usage/linter/rules/eslint/accessor-pairs.md for this page in Markdown format

# eslint/accessor-pairs Pedantic [​](https://oxc.rs/docs/guide/usage/linter/rules/eslint/accessor-pairs.html#eslint-accessor-pairs)

### What it does [​](https://oxc.rs/docs/guide/usage/linter/rules/eslint/accessor-pairs.html#what-it-does)

Enforces getter/setter pairs in objects and classes.

### Why is this bad? [​](https://oxc.rs/docs/guide/usage/linter/rules/eslint/accessor-pairs.html#why-is-this-bad)

It's a common mistake in JavaScript to create an object with just a setter for a property but never have a corresponding getter defined for it. Without a getter, you cannot read the property, so it ends up not being used.

### Examples [​](https://oxc.rs/docs/guide/usage/linter/rules/eslint/accessor-pairs.html#examples)

Examples of **incorrect** code for this rule:

js

```
var o = {
  set a(value) {
    this.val = value;
  },
};

class C {
  set a(value) {
    this.val = value;
  }
}
```

Examples of **correct** code for this rule:

js

```
var o = {
  set a(value) {
    this.val = value;
  },
  get a() {
    return this.val;
  },
};

class C {
  set a(value) {
    this.val = value;
  }
  get a() {
    return this.val;
  }
}
```

## Configuration [​](https://oxc.rs/docs/guide/usage/linter/rules/eslint/accessor-pairs.html#configuration)

This rule accepts a configuration object with the following properties:

### enforceForClassMembers [​](https://oxc.rs/docs/guide/usage/linter/rules/eslint/accessor-pairs.html#enforceforclassmembers)

type: `boolean`

default: `true`

Enforce the rule for class members.

### enforceForTSTypes [​](https://oxc.rs/docs/guide/usage/linter/rules/eslint/accessor-pairs.html#enforcefortstypes)

type: `boolean`

default: `false`

Enforce the rule for TypeScript interfaces and types.

### getWithoutSet [​](https://oxc.rs/docs/guide/usage/linter/rules/eslint/accessor-pairs.html#getwithoutset)

type: `boolean`

default: `false`

Report a getter without a setter.

### setWithoutGet [​](https://oxc.rs/docs/guide/usage/linter/rules/eslint/accessor-pairs.html#setwithoutget)

type: `boolean`

default: `true`

Report a setter without a getter.

## How to use [​](https://oxc.rs/docs/guide/usage/linter/rules/eslint/accessor-pairs.html#how-to-use)

To **enable** this rule using the config file or in the CLI, you can use:

Config (.oxlintrc.json)CLI

json

```
{
  "rules": {
    "accessor-pairs": "error"
  }
}
```

bash

```
oxlint --deny accessor-pairs
```

## References [​](https://oxc.rs/docs/guide/usage/linter/rules/eslint/accessor-pairs.html#references)

* [Rule Source](https://github.com/oxc-project/oxc/blob/1bf0ffc0f6859c90409a9701e62e8957ef1286cc/crates/oxc_linter/src/rules/eslint/accessor_pairs.rs)
