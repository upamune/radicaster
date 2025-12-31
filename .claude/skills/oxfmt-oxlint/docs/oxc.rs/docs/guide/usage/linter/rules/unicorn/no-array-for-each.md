---
title: "unicorn/no-array-for-each | The JavaScript Oxidation Compiler"
source_url: "https://oxc.rs/docs/guide/usage/linter/rules/unicorn/no-array-for-each"
fetched_at: "2025-12-31T10:44:14.234719+00:00"
---



Are you an LLM? You can read better optimized documentation at /docs/guide/usage/linter/rules/unicorn/no-array-for-each.md for this page in Markdown format

# unicorn/no-array-for-each Restriction [​](https://oxc.rs/docs/guide/usage/linter/rules/unicorn/no-array-for-each.html#unicorn-no-array-for-each)

🚧 An auto-fix is planned for this rule, but not implemented at this time.

### What it does [​](https://oxc.rs/docs/guide/usage/linter/rules/unicorn/no-array-for-each.html#what-it-does)

Forbids the use of `Array#forEach` in favor of a for loop.

### Why is this bad? [​](https://oxc.rs/docs/guide/usage/linter/rules/unicorn/no-array-for-each.html#why-is-this-bad)

Benefits of [`for…of` statement](https://developer.mozilla.org/en-US/docs/Web/JavaScript/Reference/Statements/for...of) over the `forEach` method can include:

* Faster
* Better readability
* Ability to exit early with `break` or `return`

Additionally, using `for…of` has great benefits if you are using TypeScript, because it does not cause a function boundary to be crossed. This means that type-narrowing earlier on in the current scope will work properly while inside of the loop (without having to re-type-narrow). Furthermore, any mutated variables inside of the loop will picked up on for the purposes of determining if a variable is being used.

### Examples [​](https://oxc.rs/docs/guide/usage/linter/rules/unicorn/no-array-for-each.html#examples)

Examples of **incorrect** code for this rule:

javascript

```
const foo = [1, 2, 3];
foo.forEach((element) => {
  /* ... */
});
```

Examples of **correct** code for this rule:

javascript

```
const foo = [1, 2, 3];
for (const element of foo) {
  /* ... */
}
```

## How to use [​](https://oxc.rs/docs/guide/usage/linter/rules/unicorn/no-array-for-each.html#how-to-use)

To **enable** this rule using the config file or in the CLI, you can use:

Config (.oxlintrc.json)CLI

json

```
{
  "rules": {
    "unicorn/no-array-for-each": "error"
  }
}
```

bash

```
oxlint --deny unicorn/no-array-for-each
```

## References [​](https://oxc.rs/docs/guide/usage/linter/rules/unicorn/no-array-for-each.html#references)

* [Rule Source](https://github.com/oxc-project/oxc/blob/1bf0ffc0f6859c90409a9701e62e8957ef1286cc/crates/oxc_linter/src/rules/unicorn/no_array_for_each.rs)
