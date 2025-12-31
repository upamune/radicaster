---
title: "eslint/grouped-accessor-pairs | The JavaScript Oxidation Compiler"
source_url: "https://oxc.rs/docs/guide/usage/linter/rules/eslint/grouped-accessor-pairs"
fetched_at: "2025-12-31T10:44:14.234719+00:00"
---



Are you an LLM? You can read better optimized documentation at /docs/guide/usage/linter/rules/eslint/grouped-accessor-pairs.md for this page in Markdown format

# eslint/grouped-accessor-pairs Style [​](https://oxc.rs/docs/guide/usage/linter/rules/eslint/grouped-accessor-pairs.html#eslint-grouped-accessor-pairs)

🚧 An auto-fix is planned for this rule, but not implemented at this time.

### What it does [​](https://oxc.rs/docs/guide/usage/linter/rules/eslint/grouped-accessor-pairs.html#what-it-does)

Require grouped accessor pairs in object literals and classes

### Why is this bad? [​](https://oxc.rs/docs/guide/usage/linter/rules/eslint/grouped-accessor-pairs.html#why-is-this-bad)

While it is allowed to define the pair for a getter or a setter anywhere in an object or class definition, it’s considered a best practice to group accessor functions for the same property.

### Examples [​](https://oxc.rs/docs/guide/usage/linter/rules/eslint/grouped-accessor-pairs.html#examples)

Examples of **incorrect** code for this rule:

js

```
const foo = {
  get a() {
    return this.val;
  },
  b: 1,
  set a(value) {
    this.val = value;
  },
};
```

Examples of **correct** code for this rule:

js

```
const foo = {
  get a() {
    return this.val;
  },
  set a(value) {
    this.val = value;
  },
  b: 1,
};
```

Examples of **incorrect** code for this rule with the `getBeforeSet` option:

js

```
const foo = {
  set a(value) {
    this.val = value;
  },
  get a() {
    return this.val;
  },
};
```

Examples of **correct** code for this rule with the `getBeforeSet` option:

js

```
const foo = {
  get a() {
    return this.val;
  },
  set a(value) {
    this.val = value;
  },
};
```

Examples of **incorrect** code for this rule with the `setBeforeGet` option:

js

```
const foo = {
  get a() {
    return this.val;
  },
  set a(value) {
    this.val = value;
  },
};
```

Examples of **correct** code for this rule with the `setBeforeGet` option:

js

```
const foo = {
  set a(value) {
    this.val = value;
  },
  get a() {
    return this.val;
  },
};
```

## Configuration [​](https://oxc.rs/docs/guide/usage/linter/rules/eslint/grouped-accessor-pairs.html#configuration)

This rule accepts a configuration object with the following properties:

### enforceForTSTypes [​](https://oxc.rs/docs/guide/usage/linter/rules/eslint/grouped-accessor-pairs.html#enforcefortstypes)

type: `boolean`

default: `false`

When `enforceForTSTypes` is enabled, this rule also applies to TypeScript interfaces and type aliases:

Examples of **incorrect** TypeScript code:

ts

```
interface Foo {
  get a(): string;
  someProperty: string;
  set a(value: string);
}

type Bar = {
  get b(): string;
  someProperty: string;
  set b(value: string);
};
```

Examples of **correct** TypeScript code:

ts

```
interface Foo {
  get a(): string;
  set a(value: string);
  someProperty: string;
}

type Bar = {
  get b(): string;
  set b(value: string);
  someProperty: string;
};
```

### pairOrder [​](https://oxc.rs/docs/guide/usage/linter/rules/eslint/grouped-accessor-pairs.html#pairorder)

type: `"anyOrder" | "getBeforeSet" | "setBeforeGet"`

default: `"anyOrder"`

A string value to control the order of the getter/setter pairs:

* `"anyOrder"`: Accessors can be in any order
* `"getBeforeSet"`: Getters must come before setters
* `"setBeforeGet"`: Setters must come before getters

## How to use [​](https://oxc.rs/docs/guide/usage/linter/rules/eslint/grouped-accessor-pairs.html#how-to-use)

To **enable** this rule using the config file or in the CLI, you can use:

Config (.oxlintrc.json)CLI

json

```
{
  "rules": {
    "grouped-accessor-pairs": "error"
  }
}
```

bash

```
oxlint --deny grouped-accessor-pairs
```

## References [​](https://oxc.rs/docs/guide/usage/linter/rules/eslint/grouped-accessor-pairs.html#references)

* [Rule Source](https://github.com/oxc-project/oxc/blob/1bf0ffc0f6859c90409a9701e62e8957ef1286cc/crates/oxc_linter/src/rules/eslint/grouped_accessor_pairs.rs)
