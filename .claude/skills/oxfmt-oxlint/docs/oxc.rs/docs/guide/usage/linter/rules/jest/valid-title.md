---
title: "jest/valid-title | The JavaScript Oxidation Compiler"
source_url: "https://oxc.rs/docs/guide/usage/linter/rules/jest/valid-title"
fetched_at: "2025-12-31T10:44:14.234719+00:00"
---



Are you an LLM? You can read better optimized documentation at /docs/guide/usage/linter/rules/jest/valid-title.md for this page in Markdown format

# jest/valid-title Correctness [​](https://oxc.rs/docs/guide/usage/linter/rules/jest/valid-title.html#jest-valid-title)

🛠️ An auto-fix is available for this rule for some violations.

### What it does [​](https://oxc.rs/docs/guide/usage/linter/rules/jest/valid-title.html#what-it-does)

Checks that the titles of Jest and Vitest blocks are valid.

Titles must be:

* not empty,
* strings,
* not prefixed with their block name,
* have no leading or trailing spaces.

### Why is this bad? [​](https://oxc.rs/docs/guide/usage/linter/rules/jest/valid-title.html#why-is-this-bad)

Titles that are not valid can be misleading and make it harder to understand the purpose of the test.

### Examples [​](https://oxc.rs/docs/guide/usage/linter/rules/jest/valid-title.html#examples)

Examples of **incorrect** code for this rule:

javascript

```
describe("", () => {});
describe("foo", () => {
  it("", () => {});
});
it("", () => {});
test("", () => {});
xdescribe("", () => {});
xit("", () => {});
xtest("", () => {});
```

Examples of **correct** code for this rule:

javascript

```
describe("foo", () => {});
it("bar", () => {});
test("baz", () => {});
```

### Options [​](https://oxc.rs/docs/guide/usage/linter/rules/jest/valid-title.html#options)

typescript

```
interface Options {
  ignoreSpaces?: boolean;
  ignoreTypeOfTestName?: boolean;
  ignoreTypeOfDescribeName?: boolean;
  allowArguments?: boolean;
  disallowedWords?: string[];
  mustNotMatch?: Partial<Record<"describe" | "test" | "it", string>> | string;
  mustMatch?: Partial<Record<"describe" | "test" | "it", string>> | string;
}
```

## How to use [​](https://oxc.rs/docs/guide/usage/linter/rules/jest/valid-title.html#how-to-use)

To **enable** this rule using the config file or in the CLI, you can use:

Config (.oxlintrc.json)CLI

json

```
{
  "plugins": ["jest"],
  "rules": {
    "jest/valid-title": "error"
  }
}
```

bash

```
oxlint --deny jest/valid-title --jest-plugin
```

## References [​](https://oxc.rs/docs/guide/usage/linter/rules/jest/valid-title.html#references)

* [Rule Source](https://github.com/oxc-project/oxc/blob/1bf0ffc0f6859c90409a9701e62e8957ef1286cc/crates/oxc_linter/src/rules/jest/valid_title.rs)
