---
title: "eslint/func-names | The JavaScript Oxidation Compiler"
source_url: "https://oxc.rs/docs/guide/usage/linter/rules/eslint/func-names"
fetched_at: "2025-12-31T10:44:14.234719+00:00"
---



Are you an LLM? You can read better optimized documentation at /docs/guide/usage/linter/rules/eslint/func-names.md for this page in Markdown format

# eslint/func-names Style [​](https://oxc.rs/docs/guide/usage/linter/rules/eslint/func-names.html#eslint-func-names)

🛠️💡 An auto-fix and a suggestion are available for this rule for some violations.

### What it does [​](https://oxc.rs/docs/guide/usage/linter/rules/eslint/func-names.html#what-it-does)

Require or disallow named function expressions.

### Why is this bad? [​](https://oxc.rs/docs/guide/usage/linter/rules/eslint/func-names.html#why-is-this-bad)

Leaving the name off a function will cause `<anonymous>` to appear in stack traces of errors thrown in it or any function called within it. This makes it more difficult to find where an error is thrown. Providing an explicit name also improves readability and consistency.

### Options [​](https://oxc.rs/docs/guide/usage/linter/rules/eslint/func-names.html#options)

First option:

* Type: `string`
* Default: `"always"`
* Possible values:
  + `"always"` - requires all function expressions to have a name.
  + `"as-needed"` - requires a name only if one is not automatically inferred.
  + `"never"` - disallows names for function expressions.

Second option:

* Type: `object`
* Properties:
  + `generators`: `("always" | "as-needed" | "never")` (default: falls back to first option)
    - `"always"` - require named generator function expressions.
    - `"as-needed"` - require a name only when not inferred.
    - `"never"` - disallow names for generator function expressions.

Example configuration:

json

```
{
  "func-names": ["error", "as-needed", { "generators": "never" }]
}
```

### Examples [​](https://oxc.rs/docs/guide/usage/linter/rules/eslint/func-names.html#examples)

Examples of **incorrect** code for this rule:

js

```
/* func-names: ["error", "always"] */

Foo.prototype.bar = function () {};
const cat = { meow: function () {} };
(function () {
  /* ... */
})();
export default function () {}
```

Examples of **correct** code for this rule:

js

```
/* func-names: ["error", "always"] */

Foo.prototype.bar = function bar() {};
const cat = { meow() {} };
(function bar() {
  /* ... */
})();
export default function foo() {}
```

#### `as-needed` [​](https://oxc.rs/docs/guide/usage/linter/rules/eslint/func-names.html#as-needed)

Examples of **incorrect** code for this rule with the `"as-needed"` option:

js

```
/* func-names: ["error", "as-needed"] */

Foo.prototype.bar = function () {};
(function () {
  /* ... */
})();
export default function () {}
```

Examples of **correct** code for this rule with the `"as-needed"` option:

js

```
/* func-names: ["error", "as-needed"] */

const bar = function () {};
const cat = { meow: function () {} };
class C {
  #bar = function () {};
  baz = function () {};
}
quux ??= function () {};
(function bar() {
  /* ... */
})();
export default function foo() {}
```

#### `never` [​](https://oxc.rs/docs/guide/usage/linter/rules/eslint/func-names.html#never)

Examples of **incorrect** code for this rule with the `"never"` option:

js

```
/* func-names: ["error", "never"] */

Foo.prototype.bar = function bar() {};
(function bar() {
  /* ... */
})();
```

Examples of **correct** code for this rule with the `"never"` option:

js

```
/* func-names: ["error", "never"] */

Foo.prototype.bar = function () {};
(function () {
  /* ... */
})();
```

#### `generators` [​](https://oxc.rs/docs/guide/usage/linter/rules/eslint/func-names.html#generators)

Examples of **incorrect** code for this rule with the `"always", { "generators": "as-needed" }` options:

js

```
/* func-names: ["error", "always", { "generators": "as-needed" }] */

(function* () {
  /* ... */
})();
```

Examples of **correct** code for this rule with the `"always", { "generators": "as-needed" }` options:

js

```
/* func-names: ["error", "always", { "generators": "as-needed" }] */

const foo = function* () {};
```

Examples of **incorrect** code for this rule with the `"always", { "generators": "never" }` options:

js

```
/* func-names: ["error", "always", { "generators": "never" }] */

const foo = bar(function* baz() {});
```

Examples of **correct** code for this rule with the `"always", { "generators": "never" }` options:

js

```
/* func-names: ["error", "always", { "generators": "never" }] */

const foo = bar(function* () {});
```

Examples of **incorrect** code for this rule with the `"as-needed", { "generators": "never" }` options:

js

```
/* func-names: ["error", "as-needed", { "generators": "never" }] */

const foo = bar(function* baz() {});
```

Examples of **correct** code for this rule with the `"as-needed", { "generators": "never" }` options:

js

```
/* func-names: ["error", "as-needed", { "generators": "never" }] */

const foo = bar(function* () {});
```

Examples of **incorrect** code for this rule with the `"never", { "generators": "always" }` options:

js

```
/* func-names: ["error", "never", { "generators": "always" }] */

const foo = bar(function* () {});
```

Examples of **correct** code for this rule with the `"never", { "generators": "always" }` options:

js

```
/* func-names: ["error", "never", { "generators": "always" }] */

const foo = bar(function* baz() {});
```

## How to use [​](https://oxc.rs/docs/guide/usage/linter/rules/eslint/func-names.html#how-to-use)

To **enable** this rule using the config file or in the CLI, you can use:

Config (.oxlintrc.json)CLI

json

```
{
  "rules": {
    "func-names": "error"
  }
}
```

bash

```
oxlint --deny func-names
```

## References [​](https://oxc.rs/docs/guide/usage/linter/rules/eslint/func-names.html#references)

* [Rule Source](https://github.com/oxc-project/oxc/blob/1bf0ffc0f6859c90409a9701e62e8957ef1286cc/crates/oxc_linter/src/rules/eslint/func_names.rs)
