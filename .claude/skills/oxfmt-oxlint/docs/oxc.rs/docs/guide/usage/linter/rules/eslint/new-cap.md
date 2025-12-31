---
title: "eslint/new-cap | The JavaScript Oxidation Compiler"
source_url: "https://oxc.rs/docs/guide/usage/linter/rules/eslint/new-cap"
fetched_at: "2025-12-31T10:44:14.234719+00:00"
---



Are you an LLM? You can read better optimized documentation at /docs/guide/usage/linter/rules/eslint/new-cap.md for this page in Markdown format

# eslint/new-cap Style [​](https://oxc.rs/docs/guide/usage/linter/rules/eslint/new-cap.html#eslint-new-cap)

🚧 An auto-fix is planned for this rule, but not implemented at this time.

### What it does [​](https://oxc.rs/docs/guide/usage/linter/rules/eslint/new-cap.html#what-it-does)

This rule requires constructor names to begin with a capital letter.

### Why is this bad? [​](https://oxc.rs/docs/guide/usage/linter/rules/eslint/new-cap.html#why-is-this-bad)

The new operator in JavaScript creates a new instance of a particular type of object. That type of object is represented by a constructor function. Since constructor functions are just regular functions, the only defining characteristic is that new is being used as part of the call. Native JavaScript functions begin with an uppercase letter to distinguish those functions that are to be used as constructors from functions that are not. Many style guides recommend following this pattern to more easily determine which functions are to be used as constructors.

**Warning**: The option `newIsCapExceptionPattern` and `capIsNewExceptionPattern` are implemented with the [rust regex syntax](https://docs.rs/regex/latest/regex/). Many JavaScript features are not supported (Lookahead, Lookbehinds, ...).

### Examples [​](https://oxc.rs/docs/guide/usage/linter/rules/eslint/new-cap.html#examples)

Examples of **incorrect** code for this rule:

js

```
function foo(arg) {
  return Boolean(arg);
}
```

Examples of **incorrect** code for this rule with the default `{ "newIsCap": true }` option:

js

```
/*eslint new-cap: ["error", { "newIsCap": true }]*/

var friend = new person();
```

Examples of **correct** code for this rule with the default `{ "newIsCap": true }` option:

js

```
/*eslint new-cap: ["error", { "newIsCap": true }]*/

var friend = new Person();
```

Examples of **correct** code for this rule with the `{ "newIsCap": false }` option:

js

```
/*eslint new-cap: ["error", { "newIsCap": false }]*/

var friend = new person();
```

Examples of **incorrect** code for this rule with the default `{ "capIsNew": true }` option:

js

```
/*eslint new-cap: ["error", { "capIsNew": true }]*/

var colleague = Person();
```

Examples of **correct** code for this rule with the default `{ "capIsNew": true }` option:

js

```
/*eslint new-cap: ["error", { "capIsNew": true }]*/

var colleague = new Person();
```

Examples of **correct** code for this rule with the `{ "capIsNew": false }` option:

js

```
/*eslint new-cap: ["error", { "capIsNew": false }]*/

var colleague = Person();
```

Examples of additional **correct** code for this rule with the `{ "newIsCapExceptions": ["events"] }` option:

js

```
/*eslint new-cap: ["error", { "newIsCapExceptions": ["events"] }]*/

var events = require("events");

var emitter = new events();
```

Examples of additional **correct** code for this rule with the `{ "newIsCapExceptionPattern": "^person\\.." }` option:

js

```
/*eslint new-cap: ["error", { "newIsCapExceptionPattern": "^person\\.." }]*/

var friend = new person.acquaintance();

var bestFriend = new person.friend();
```

Examples of additional **correct** code for this rule with the `{ "newIsCapExceptionPattern": "\\.bar$" }` option:

js

```
/*eslint new-cap: ["error", { "newIsCapExceptionPattern": "\\.bar$" }]*/

var friend = new person.bar();
```

Examples of additional **correct** code for this rule with the `{ "capIsNewExceptions": ["Person"] }` option:

js

```
/*eslint new-cap: ["error", { "capIsNewExceptions": ["Person"] }]*/

function foo(arg) {
  return Person(arg);
}
```

Examples of additional **correct** code for this rule with the `{ "capIsNewExceptionPattern": "^person\\.." }` option:

js

```
/*eslint new-cap: ["error", { "capIsNewExceptionPattern": "^person\\.." }]*/

var friend = person.Acquaintance();
var bestFriend = person.Friend();
```

Examples of additional **correct** code for this rule with the `{ "capIsNewExceptionPattern": "\\.Bar$" }` option:

js

```
/*eslint new-cap: ["error", { "capIsNewExceptionPattern": "\\.Bar$" }]*/

foo.Bar();
```

Examples of additional **correct** code for this rule with the `{ "capIsNewExceptionPattern": "^Foo" }` option:

js

```
/*eslint new-cap: ["error", { "capIsNewExceptionPattern": "^Foo" }]*/

var x = Foo(42);

var y = Foobar(42);

var z = Foo.Bar(42);
```

### properties [​](https://oxc.rs/docs/guide/usage/linter/rules/eslint/new-cap.html#properties)

Examples of **incorrect** code for this rule with the default `{ "properties": true }` option:

js

```
/*eslint new-cap: ["error", { "properties": true }]*/

var friend = new person.acquaintance();
```

Examples of **correct** code for this rule with the default `{ "properties": true }` option:

js

```
/*eslint new-cap: ["error", { "properties": true }]*/

var friend = new person.Acquaintance();
```

Examples of **correct** code for this rule with the `{ "properties": false }` option:

js

```
/*eslint new-cap: ["error", { "properties": false }]*/

var friend = new person.acquaintance();
```

Examples of **incorrect** code for this rule with the default `{ "newIsCap": true }` option:

js

```
/*eslint new-cap: ["error", { "newIsCap": true }]*/

var friend = new person();
```

Examples of **correct** code for this rule with the default `{ "newIsCap": true }` option:

js

```
/*eslint new-cap: ["error", { "newIsCap": true }]*/

var friend = new Person();
```

Examples of **correct** code for this rule with the `{ "newIsCap": false }` option:

js

```
/*eslint new-cap: ["error", { "newIsCap": false }]*/

var friend = new person();
```

Examples of **incorrect** code for this rule with the default `{ "capIsNew": true }` option:

js

```
/*eslint new-cap: ["error", { "capIsNew": true }]*/

var colleague = Person();
```

Examples of **correct** code for this rule with the default `{ "capIsNew": true }` option:

js

```
/*eslint new-cap: ["error", { "capIsNew": true }]*/

var colleague = new Person();
```

Examples of **correct** code for this rule with the `{ "capIsNew": false }` option:

js

```
/*eslint new-cap: ["error", { "capIsNew": false }]*/

var colleague = Person();
```

Examples of additional **correct** code for this rule with the `{ "newIsCapExceptions": ["events"] }` option:

js

```
/*eslint new-cap: ["error", { "newIsCapExceptions": ["events"] }]*/

var events = require("events");

var emitter = new events();
```

Examples of additional **correct** code for this rule with the `{ "newIsCapExceptionPattern": "^person\\.." }` option:

js

```
/*eslint new-cap: ["error", { "newIsCapExceptionPattern": "^person\\.." }]*/

var friend = new person.acquaintance();

var bestFriend = new person.friend();
```

Examples of additional **correct** code for this rule with the `{ "newIsCapExceptionPattern": "\\.bar$" }` option:

js

```
/*eslint new-cap: ["error", { "newIsCapExceptionPattern": "\\.bar$" }]*/

var friend = new person.bar();
```

Examples of additional **correct** code for this rule with the `{ "capIsNewExceptions": ["Person"] }` option:

::: correct

js

```
/*eslint new-cap: ["error", { "capIsNewExceptions": ["Person"] }]*/

function foo(arg) {
  return Person(arg);
}
```

Examples of additional **correct** code for this rule with the `{ "capIsNewExceptionPattern": "^person\\.." }` option:

js

```
/*eslint new-cap: ["error", { "capIsNewExceptionPattern": "^person\\.." }]*/

var friend = person.Acquaintance();
var bestFriend = person.Friend();
```

Examples of additional **correct** code for this rule with the `{ "capIsNewExceptionPattern": "\\.Bar$" }` option:

js

```
/*eslint new-cap: ["error", { "capIsNewExceptionPattern": "\\.Bar$" }]*/

foo.Bar();
```

Examples of additional **correct** code for this rule with the `{ "capIsNewExceptionPattern": "^Foo" }` option:

js

```
/*eslint new-cap: ["error", { "capIsNewExceptionPattern": "^Foo" }]*/

var x = Foo(42);

var y = Foobar(42);

var z = Foo.Bar(42);
```

Examples of **incorrect** code for this rule with the default `{ "properties": true }` option:

js

```
/*eslint new-cap: ["error", { "properties": true }]*/

var friend = new person.acquaintance();
```

Examples of **correct** code for this rule with the default `{ "properties": true }` option:

js

```
/*eslint new-cap: ["error", { "properties": true }]*/

var friend = new person.Acquaintance();
```

Examples of **correct** code for this rule with the `{ "properties": false }` option:

js

```
/*eslint new-cap: ["error", { "properties": false }]*/

var friend = new person.acquaintance();
```

## Configuration [​](https://oxc.rs/docs/guide/usage/linter/rules/eslint/new-cap.html#configuration)

This rule accepts a configuration object with the following properties:

### capIsNew [​](https://oxc.rs/docs/guide/usage/linter/rules/eslint/new-cap.html#capisnew)

type: `boolean`

default: `true`

`true` to require that all functions with names starting with an uppercase letter to be called with `new`.

### capIsNewExceptionPattern [​](https://oxc.rs/docs/guide/usage/linter/rules/eslint/new-cap.html#capisnewexceptionpattern)

type: `string | null`

A regex pattern to match exceptions for functions with names starting with an uppercase letter.

### capIsNewExceptions [​](https://oxc.rs/docs/guide/usage/linter/rules/eslint/new-cap.html#capisnewexceptions)

type: `string[]`

default: `[]`

Exceptions to ignore for functions with names starting with an uppercase letter.

### newIsCap [​](https://oxc.rs/docs/guide/usage/linter/rules/eslint/new-cap.html#newiscap)

type: `boolean`

default: `true`

`true` to require that all constructor names start with an uppercase letter, e.g. `new Person()`.

### newIsCapExceptionPattern [​](https://oxc.rs/docs/guide/usage/linter/rules/eslint/new-cap.html#newiscapexceptionpattern)

type: `string | null`

A regex pattern to match exceptions for constructor names starting with an uppercase letter.

### newIsCapExceptions [​](https://oxc.rs/docs/guide/usage/linter/rules/eslint/new-cap.html#newiscapexceptions)

type: `string[]`

default: `["Array", "Boolean", "Date", "Error", "Function", "Number", "Object", "RegExp", "String", "Symbol", "BigInt"]`

Exceptions to ignore for constructor names starting with an uppercase letter.

### properties [​](https://oxc.rs/docs/guide/usage/linter/rules/eslint/new-cap.html#properties-1)

type: `boolean`

default: `true`

`true` to require capitalization for object properties (e.g., `new obj.Method()`).

## How to use [​](https://oxc.rs/docs/guide/usage/linter/rules/eslint/new-cap.html#how-to-use)

To **enable** this rule using the config file or in the CLI, you can use:

Config (.oxlintrc.json)CLI

json

```
{
  "rules": {
    "new-cap": "error"
  }
}
```

bash

```
oxlint --deny new-cap
```

## References [​](https://oxc.rs/docs/guide/usage/linter/rules/eslint/new-cap.html#references)

* [Rule Source](https://github.com/oxc-project/oxc/blob/1bf0ffc0f6859c90409a9701e62e8957ef1286cc/crates/oxc_linter/src/rules/eslint/new_cap.rs)
