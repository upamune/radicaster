---
title: "promise/no-nesting | The JavaScript Oxidation Compiler"
source_url: "https://oxc.rs/docs/guide/usage/linter/rules/promise/no-nesting"
fetched_at: "2025-12-31T10:44:14.234719+00:00"
---



Are you an LLM? You can read better optimized documentation at /docs/guide/usage/linter/rules/promise/no-nesting.md for this page in Markdown format

# promise/no-nesting Style [​](https://oxc.rs/docs/guide/usage/linter/rules/promise/no-nesting.html#promise-no-nesting)

🚧 An auto-fix is planned for this rule, but not implemented at this time.

### What it does [​](https://oxc.rs/docs/guide/usage/linter/rules/promise/no-nesting.html#what-it-does)

Disallow nested then() or catch() statements.

### Why is this bad? [​](https://oxc.rs/docs/guide/usage/linter/rules/promise/no-nesting.html#why-is-this-bad)

Nesting promises makes code harder to read and understand.

### Examples [​](https://oxc.rs/docs/guide/usage/linter/rules/promise/no-nesting.html#examples)

Examples of **incorrect** code for this rule:

javascript

```
doThing().then(() => a.then());

doThing().then(function () {
  a.then();
});

doThing().then(() => {
  b.catch();
});

doThing().catch((val) => doSomething(val).catch(errors));
```

Examples of **correct** code for this rule:

javascript

```
doThing().then(() => 4);

doThing().then(function () {
  return 4;
});

doThing().catch(() => 4);
```

javascript

```
doThing()
  .then(() => Promise.resolve(1))
  .then(() => Promise.resolve(2));
```

This example is not a rule violation as unnesting here would result in `a` being undefined in the expression `getC(a, b)`.

javascript

```
doThing().then((a) => getB(a).then((b) => getC(a, b)));
```

## How to use [​](https://oxc.rs/docs/guide/usage/linter/rules/promise/no-nesting.html#how-to-use)

To **enable** this rule using the config file or in the CLI, you can use:

Config (.oxlintrc.json)CLI

json

```
{
  "plugins": ["promise"],
  "rules": {
    "promise/no-nesting": "error"
  }
}
```

bash

```
oxlint --deny promise/no-nesting --promise-plugin
```

## References [​](https://oxc.rs/docs/guide/usage/linter/rules/promise/no-nesting.html#references)

* [Rule Source](https://github.com/oxc-project/oxc/blob/1bf0ffc0f6859c90409a9701e62e8957ef1286cc/crates/oxc_linter/src/rules/promise/no_nesting.rs)
