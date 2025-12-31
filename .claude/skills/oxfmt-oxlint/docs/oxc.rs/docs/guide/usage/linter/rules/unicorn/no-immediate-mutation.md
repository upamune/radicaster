---
title: "unicorn/no-immediate-mutation | The JavaScript Oxidation Compiler"
source_url: "https://oxc.rs/docs/guide/usage/linter/rules/unicorn/no-immediate-mutation"
fetched_at: "2025-12-31T10:44:14.234719+00:00"
---



Are you an LLM? You can read better optimized documentation at /docs/guide/usage/linter/rules/unicorn/no-immediate-mutation.md for this page in Markdown format

# unicorn/no-immediate-mutation Pedantic [​](https://oxc.rs/docs/guide/usage/linter/rules/unicorn/no-immediate-mutation.html#unicorn-no-immediate-mutation)

🚧 An auto-fix is planned for this rule, but not implemented at this time.

### What it does [​](https://oxc.rs/docs/guide/usage/linter/rules/unicorn/no-immediate-mutation.html#what-it-does)

Disallows mutating a variable immediately after initialization.

### Why is this bad? [​](https://oxc.rs/docs/guide/usage/linter/rules/unicorn/no-immediate-mutation.html#why-is-this-bad)

When you initialize a variable and immediately mutate it, it's cleaner to include the mutation in the initialization. This makes the code more readable and reduces the number of statements.

### Examples [​](https://oxc.rs/docs/guide/usage/linter/rules/unicorn/no-immediate-mutation.html#examples)

Examples of **incorrect** code for this rule:

js

```
const array = [1, 2];
array.push(3);

const object = { foo: 1 };
object.bar = 2;

const set = new Set([1, 2]);
set.add(3);

const map = new Map([["foo", 1]]);
map.set("bar", 2);
```

Examples of **correct** code for this rule:

js

```
const array = [1, 2, 3];

const object = { foo: 1, bar: 2 };

const set = new Set([1, 2, 3]);

const map = new Map([
  ["foo", 1],
  ["bar", 2],
]);
```

## How to use [​](https://oxc.rs/docs/guide/usage/linter/rules/unicorn/no-immediate-mutation.html#how-to-use)

To **enable** this rule using the config file or in the CLI, you can use:

Config (.oxlintrc.json)CLI

json

```
{
  "rules": {
    "unicorn/no-immediate-mutation": "error"
  }
}
```

bash

```
oxlint --deny unicorn/no-immediate-mutation
```

## References [​](https://oxc.rs/docs/guide/usage/linter/rules/unicorn/no-immediate-mutation.html#references)

* [Rule Source](https://github.com/oxc-project/oxc/blob/1bf0ffc0f6859c90409a9701e62e8957ef1286cc/crates/oxc_linter/src/rules/unicorn/no_immediate_mutation.rs)
