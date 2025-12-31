---
title: "node/global-require | The JavaScript Oxidation Compiler"
source_url: "https://oxc.rs/docs/guide/usage/linter/rules/node/global-require"
fetched_at: "2025-12-31T10:44:14.234719+00:00"
---



Are you an LLM? You can read better optimized documentation at /docs/guide/usage/linter/rules/node/global-require.md for this page in Markdown format

# node/global-require Style [​](https://oxc.rs/docs/guide/usage/linter/rules/node/global-require.html#node-global-require)

### What it does [​](https://oxc.rs/docs/guide/usage/linter/rules/node/global-require.html#what-it-does)

Require `require()` calls to be placed at top-level module scope

### Why is this bad? [​](https://oxc.rs/docs/guide/usage/linter/rules/node/global-require.html#why-is-this-bad)

In Node.js, module dependencies are included using the `require()` function, such as:

js

```
var fs = require("fs");
```

While `require()` may be called anywhere in code, some style guides prescribe that it should be called only in the top level of a module to make it easier to identify dependencies. For instance, it's arguably harder to identify dependencies when they are deeply nested inside of functions and other statements:

js

```
function foo() {
  if (condition) {
    var fs = require("fs");
  }
}
```

Since `require()` does a synchronous load, it can cause performance problems when used in other locations. Further, ES6 modules mandate that import and export statements can only occur in the top level of the module's body.

### Examples [​](https://oxc.rs/docs/guide/usage/linter/rules/node/global-require.html#examples)

Examples of **incorrect** code for this rule:

js

```
// calling require() inside of a function is not allowed
function readFile(filename, callback) {
  var fs = require("fs");
  fs.readFile(filename, callback);
}

// conditional requires like this are also not allowed
if (DEBUG) {
  require("debug");
}

// a require() in a switch statement is also flagged
switch (x) {
  case "1":
    require("1");
    break;
}

// you may not require() inside an arrow function body
var getModule = (name) => require(name);

// you may not require() inside of a function body as well
function getModule(name) {
  return require(name);
}

// you may not require() inside of a try/catch block
try {
  require(unsafeModule);
} catch (e) {
  console.log(e);
}
```

Examples of **correct** code for this rule:

js

```
// all these variations of require() are ok
require("x");
var y = require("y");
var z;
z = require("z").initialize();

// requiring a module and using it in a function is ok
var fs = require("fs");
function readFile(filename, callback) {
  fs.readFile(filename, callback);
}

// you can use a ternary to determine which module to require
var logger = DEBUG ? require("dev-logger") : require("logger");

// if you want you can require() at the end of your module
function doSomethingA() {}
function doSomethingB() {}
var x = require("x"),
  z = require("z");
```

## How to use [​](https://oxc.rs/docs/guide/usage/linter/rules/node/global-require.html#how-to-use)

To **enable** this rule using the config file or in the CLI, you can use:

Config (.oxlintrc.json)CLI

json

```
{
  "plugins": ["node"],
  "rules": {
    "node/global-require": "error"
  }
}
```

bash

```
oxlint --deny node/global-require --node-plugin
```

## References [​](https://oxc.rs/docs/guide/usage/linter/rules/node/global-require.html#references)

* [Rule Source](https://github.com/oxc-project/oxc/blob/1bf0ffc0f6859c90409a9701e62e8957ef1286cc/crates/oxc_linter/src/rules/node/global_require.rs)
