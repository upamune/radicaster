---
title: "eslint/no-empty-function | The JavaScript Oxidation Compiler"
source_url: "https://oxc.rs/docs/guide/usage/linter/rules/eslint/no-empty-function"
fetched_at: "2025-12-31T10:44:14.234719+00:00"
---



Are you an LLM? You can read better optimized documentation at /docs/guide/usage/linter/rules/eslint/no-empty-function.md for this page in Markdown format

# eslint/no-empty-function Restriction [​](https://oxc.rs/docs/guide/usage/linter/rules/eslint/no-empty-function.html#eslint-no-empty-function)

### What it does [​](https://oxc.rs/docs/guide/usage/linter/rules/eslint/no-empty-function.html#what-it-does)

Disallows the usages of empty functions

### Why is this bad? [​](https://oxc.rs/docs/guide/usage/linter/rules/eslint/no-empty-function.html#why-is-this-bad)

Empty functions can reduce readability because readers need to guess whether it's intentional or not. So writing a clear comment for empty functions is a good practice.

### Options [​](https://oxc.rs/docs/guide/usage/linter/rules/eslint/no-empty-function.html#options)

#### allow [​](https://oxc.rs/docs/guide/usage/linter/rules/eslint/no-empty-function.html#allow)

`{ type: string[], default: [] }`

You may pass a list of allowed function kinds, which will allow functions of these kinds to be empty.

Example:

json

```
{
  "no-empty-function": ["error", { "allow": ["functions"] }]
}
```

`allow` accepts the following values:

* `"functions"`
* `"arrowFunctions"`
* `"generatorFunctions"`
* `"methods"`
* `"generatorMethods"`
* `"getters"`
* `"setters"`
* `"constructors"`
* `"privateConstructors"`
* `"protectedConstructors"`
* `"asyncFunctions"`
* `"asyncMethods"`
* `"decoratedFunctions"`
* `"overrideMethods"`

### Examples [​](https://oxc.rs/docs/guide/usage/linter/rules/eslint/no-empty-function.html#examples)

Examples of **incorrect** code for this rule:

typescript

```
function foo() {}

const bar = () => {};

class Foo {
  constructor();
  someMethod() {}
  set bar(value) {}
}
```

Examples of **correct** code for this rule:

typescript

```
function foo() {
  // do nothing
}

function foo() {
  return;
}
const add = (a, b) => a + b;

class Foo {
  // constructor body is empty, but it declares a private property named
  // `_name`
  constructor(private _name: string) {}

  public get name() {
    return this._name;
  }
}
```

## How to use [​](https://oxc.rs/docs/guide/usage/linter/rules/eslint/no-empty-function.html#how-to-use)

To **enable** this rule using the config file or in the CLI, you can use:

Config (.oxlintrc.json)CLI

json

```
{
  "rules": {
    "no-empty-function": "error"
  }
}
```

bash

```
oxlint --deny no-empty-function
```

## References [​](https://oxc.rs/docs/guide/usage/linter/rules/eslint/no-empty-function.html#references)

* [Rule Source](https://github.com/oxc-project/oxc/blob/1bf0ffc0f6859c90409a9701e62e8957ef1286cc/crates/oxc_linter/src/rules/eslint/no_empty_function.rs)
