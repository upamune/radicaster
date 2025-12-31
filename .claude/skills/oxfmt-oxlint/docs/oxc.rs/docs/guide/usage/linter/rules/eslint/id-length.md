---
title: "eslint/id-length | The JavaScript Oxidation Compiler"
source_url: "https://oxc.rs/docs/guide/usage/linter/rules/eslint/id-length"
fetched_at: "2025-12-31T10:44:14.234719+00:00"
---



Are you an LLM? You can read better optimized documentation at /docs/guide/usage/linter/rules/eslint/id-length.md for this page in Markdown format

# eslint/id-length Style [​](https://oxc.rs/docs/guide/usage/linter/rules/eslint/id-length.html#eslint-id-length)

### What it does [​](https://oxc.rs/docs/guide/usage/linter/rules/eslint/id-length.html#what-it-does)

This rule enforces a minimum and/or maximum identifier length convention by counting the graphemes for a given identifier.

### Why is this bad? [​](https://oxc.rs/docs/guide/usage/linter/rules/eslint/id-length.html#why-is-this-bad)

Very short identifier names like e, x, \_t or very long ones like hashGeneratorResultOutputContainerObject can make code harder to read and potentially less maintainable. To prevent this, one may enforce a minimum and/or maximum identifier length.

### Examples [​](https://oxc.rs/docs/guide/usage/linter/rules/eslint/id-length.html#examples)

Examples of **incorrect** code for this rule:

js

```
/*eslint id-length: "error"*/ // default is minimum 2-chars ({ "min": 2 })

const x = 5;
obj.e = document.body;
const foo = function (e) {};
try {
  dangerousStuff();
} catch (e) {
  // ignore as many do
}
const myObj = { a: 1 };
(a) => {
  a * a;
};
class y {}
class Foo {
  x() {}
}
class Bar {
  #x() {}
}
class Baz {
  x = 1;
}
class Qux {
  #x = 1;
}
function bar(...x) {}
function baz([x]) {}
const [z] = arr;
const {
  prop: [i],
} = {};
function qux({ x }) {}
const { j } = {};
const { prop: a } = {};
({ prop: obj.x } = {});
```

Examples of **correct** code for this rule:

js

```
/*eslint id-length: "error"*/ // default is minimum 2-chars ({ "min": 2 })

const num = 5;
function _f() {
  return 42;
}
function _func() {
  return 42;
}
obj.el = document.body;
const foo = function (evt) {
  /* do stuff */
};
try {
  dangerousStuff();
} catch (error) {
  // ignore as many do
}
const myObj = { apple: 1 };
(num) => {
  num * num;
};
function bar(num = 0) {}
class MyClass {}
class Foo {
  method() {}
}
class Bar {
  #method() {}
}
class Baz {
  field = 1;
}
class Qux {
  #field = 1;
}
function baz(...args) {}
function qux([longName]) {}
const { prop } = {};
const {
  prop: [name],
} = {};
const [longName] = arr;
function foobar({ prop }) {}
function foobaz({ a: prop }) {}
const { a: property } = {};
({ prop: obj.longName } = {});
const data = { x: 1 }; // excused because of quotes
data["y"] = 3; // excused because of calculated property access
```

## Configuration [​](https://oxc.rs/docs/guide/usage/linter/rules/eslint/id-length.html#configuration)

This rule accepts a configuration object with the following properties:

### exceptionPatterns [​](https://oxc.rs/docs/guide/usage/linter/rules/eslint/id-length.html#exceptionpatterns)

type: `string[]`

An array of regex patterns for identifiers to exclude from the rule. For example, `["^x.*"]` would exclude all identifiers starting with "x".

### exceptions [​](https://oxc.rs/docs/guide/usage/linter/rules/eslint/id-length.html#exceptions)

type: `string[]`

default: `[]`

An array of identifier names that are excluded from the rule. For example, `["x", "y", "z"]` would allow single-letter identifiers "x", "y", and "z".

### max [​](https://oxc.rs/docs/guide/usage/linter/rules/eslint/id-length.html#max)

type: `integer`

default: `18446744073709551615`

The maximum number of graphemes allowed in an identifier. Defaults to no maximum (effectively unlimited).

### min [​](https://oxc.rs/docs/guide/usage/linter/rules/eslint/id-length.html#min)

type: `integer`

default: `2`

The minimum number of graphemes required in an identifier.

### properties [​](https://oxc.rs/docs/guide/usage/linter/rules/eslint/id-length.html#properties)

type: `"always" | "never"`

default: `"always"`

When set to `"never"`, property names are not checked for length. When set to `"always"` (default), property names are checked just like other identifiers.

## How to use [​](https://oxc.rs/docs/guide/usage/linter/rules/eslint/id-length.html#how-to-use)

To **enable** this rule using the config file or in the CLI, you can use:

Config (.oxlintrc.json)CLI

json

```
{
  "rules": {
    "id-length": "error"
  }
}
```

bash

```
oxlint --deny id-length
```

## References [​](https://oxc.rs/docs/guide/usage/linter/rules/eslint/id-length.html#references)

* [Rule Source](https://github.com/oxc-project/oxc/blob/1bf0ffc0f6859c90409a9701e62e8957ef1286cc/crates/oxc_linter/src/rules/eslint/id_length.rs)
