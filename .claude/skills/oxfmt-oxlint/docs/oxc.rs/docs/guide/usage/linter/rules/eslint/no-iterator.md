---
title: "eslint/no-iterator | The JavaScript Oxidation Compiler"
source_url: "https://oxc.rs/docs/guide/usage/linter/rules/eslint/no-iterator"
fetched_at: "2025-12-31T10:44:14.234719+00:00"
---



Are you an LLM? You can read better optimized documentation at /docs/guide/usage/linter/rules/eslint/no-iterator.md for this page in Markdown format

# eslint/no-iterator Restriction [​](https://oxc.rs/docs/guide/usage/linter/rules/eslint/no-iterator.html#eslint-no-iterator)

💡 A suggestion is available for this rule.

### What it does [​](https://oxc.rs/docs/guide/usage/linter/rules/eslint/no-iterator.html#what-it-does)

Disallow the use of the `__iterator__` property

### Why is this bad? [​](https://oxc.rs/docs/guide/usage/linter/rules/eslint/no-iterator.html#why-is-this-bad)

The `__iterator__` property was a SpiderMonkey extension to JavaScript that could be used to create custom iterators that are compatible with JavaScript’s for in and for each constructs. However, this property is now obsolete, so it should not be used. Here’s an example of how this used to work:

js

```
Foo.prototype.__iterator__ = function () {
  return new FooIterator(this);
};
```

### Examples [​](https://oxc.rs/docs/guide/usage/linter/rules/eslint/no-iterator.html#examples)

Examples of **incorrect** code for this rule:

javascript

```
Foo.prototype.__iterator__ = function () {
  return new FooIterator(this);
};

foo.__iterator__ = function () {};

foo["__iterator__"] = function () {};
```

Examples of **correct** code for this rule:

js

```
const __iterator__ = 42; // not using the __iterator__ property

Foo.prototype[Symbol.iterator] = function () {
  return new FooIterator(this);
};
```

## How to use [​](https://oxc.rs/docs/guide/usage/linter/rules/eslint/no-iterator.html#how-to-use)

To **enable** this rule using the config file or in the CLI, you can use:

Config (.oxlintrc.json)CLI

json

```
{
  "rules": {
    "no-iterator": "error"
  }
}
```

bash

```
oxlint --deny no-iterator
```

## References [​](https://oxc.rs/docs/guide/usage/linter/rules/eslint/no-iterator.html#references)

* [Rule Source](https://github.com/oxc-project/oxc/blob/1bf0ffc0f6859c90409a9701e62e8957ef1286cc/crates/oxc_linter/src/rules/eslint/no_iterator.rs)
