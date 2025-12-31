---
title: "eslint/no-useless-computed-key | The JavaScript Oxidation Compiler"
source_url: "https://oxc.rs/docs/guide/usage/linter/rules/eslint/no-useless-computed-key"
fetched_at: "2025-12-31T10:44:14.234719+00:00"
---



Are you an LLM? You can read better optimized documentation at /docs/guide/usage/linter/rules/eslint/no-useless-computed-key.md for this page in Markdown format

# eslint/no-useless-computed-key Style [​](https://oxc.rs/docs/guide/usage/linter/rules/eslint/no-useless-computed-key.html#eslint-no-useless-computed-key)

🚧 An auto-fix is planned for this rule, but not implemented at this time.

### What it does [​](https://oxc.rs/docs/guide/usage/linter/rules/eslint/no-useless-computed-key.html#what-it-does)

Disallow unnecessary computed property keys in objects and classes

### Why is this bad? [​](https://oxc.rs/docs/guide/usage/linter/rules/eslint/no-useless-computed-key.html#why-is-this-bad)

It’s unnecessary to use computed properties with literals such as:

js

```
const foo = { ["a"]: "b" };
```

The code can be rewritten as:

js

```
const foo = { a: "b" };
```

### Examples [​](https://oxc.rs/docs/guide/usage/linter/rules/eslint/no-useless-computed-key.html#examples)

Examples of **incorrect** code for this rule:

js

```
const a = { ["0"]: 0 };
const b = { ["0+1,234"]: 0 };
const c = { [0]: 0 };
const e = { ["x"]() {} };

class Foo {
  ["foo"] = "bar";
  [0]() {}
  static ["foo"] = "bar";
  get ["b"]() {}
  set ["c"](https://oxc.rs/docs/guide/usage/linter/rules/eslint/value) {}
}
```

Examples of **correct** code for this rule:

js

```
const a = { a: 0 };
const b = { 0: 0 };
const c = { x() {} };
const e = { "0+1,234": 0 };

class Foo {
  foo = "bar";
  0() {}
  a() {}
  static foo = "bar";
}
```

Examples of additional **correct** code for this rule:

js

```
const c = {
  __proto__: foo, // defines object's prototype
  ["__proto__"]: bar, // defines a property named "__proto__"
};
class Foo {
  ["constructor"]; // instance field named "constructor"
  constructor() {} // the constructor of this class
  static ["constructor"]; // static field named "constructor"
  static ["prototype"]; // runtime error, it would be a parsing error without `[]`
}
```

## Configuration [​](https://oxc.rs/docs/guide/usage/linter/rules/eslint/no-useless-computed-key.html#configuration)

This rule accepts a configuration object with the following properties:

### enforceForClassMembers [​](https://oxc.rs/docs/guide/usage/linter/rules/eslint/no-useless-computed-key.html#enforceforclassmembers)

type: `boolean`

default: `true`

The `enforceForClassMembers` option controls whether the rule applies to class members (methods and properties).

Examples of **correct** code for this rule with the `{ "enforceForClassMembers": false }` option:

js

```
class SomeClass {
  ["foo"] = "bar";
  [42] = "baz";
  get ["b"]() {}
  set ["c"](https://oxc.rs/docs/guide/usage/linter/rules/eslint/value) {}
  static ["foo"] = "bar";
}
```

## How to use [​](https://oxc.rs/docs/guide/usage/linter/rules/eslint/no-useless-computed-key.html#how-to-use)

To **enable** this rule using the config file or in the CLI, you can use:

Config (.oxlintrc.json)CLI

json

```
{
  "rules": {
    "no-useless-computed-key": "error"
  }
}
```

bash

```
oxlint --deny no-useless-computed-key
```

## References [​](https://oxc.rs/docs/guide/usage/linter/rules/eslint/no-useless-computed-key.html#references)

* [Rule Source](https://github.com/oxc-project/oxc/blob/1bf0ffc0f6859c90409a9701e62e8957ef1286cc/crates/oxc_linter/src/rules/eslint/no_useless_computed_key.rs)
