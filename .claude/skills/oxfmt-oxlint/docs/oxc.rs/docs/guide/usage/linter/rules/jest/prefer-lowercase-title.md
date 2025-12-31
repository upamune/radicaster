---
title: "jest/prefer-lowercase-title | The JavaScript Oxidation Compiler"
source_url: "https://oxc.rs/docs/guide/usage/linter/rules/jest/prefer-lowercase-title"
fetched_at: "2025-12-31T10:44:14.234719+00:00"
---



Are you an LLM? You can read better optimized documentation at /docs/guide/usage/linter/rules/jest/prefer-lowercase-title.md for this page in Markdown format

# jest/prefer-lowercase-title Style [​](https://oxc.rs/docs/guide/usage/linter/rules/jest/prefer-lowercase-title.html#jest-prefer-lowercase-title)

🛠️ An auto-fix is available for this rule.

### What it does [​](https://oxc.rs/docs/guide/usage/linter/rules/jest/prefer-lowercase-title.html#what-it-does)

Enforce `it`, `test`, `describe`, and `bench` to have descriptions that begin with a lowercase letter. This provides more readable test failures. This rule is not enabled by default.

### Examples [​](https://oxc.rs/docs/guide/usage/linter/rules/jest/prefer-lowercase-title.html#examples)

Examples of **incorrect** code for this rule:

javascript

```
it("Adds 1 + 2 to equal 3", () => {
  expect(sum(1, 2)).toBe(3);
});
```

Examples of **correct** code for this rule:

javascript

```
it("adds 1 + 2 to equal 3", () => {
  expect(sum(1, 2)).toBe(3);
});
```

This rule is compatible with [eslint-plugin-vitest](https://github.com/vitest-dev/eslint-plugin-vitest/blob/main/docs/rules/prefer-lowercase-title.md), to use it, add the following configuration to your `.oxlintrc.json`:

json

```
{
  "rules": {
    "vitest/prefer-lowercase-title": "error"
  }
}
```

## Configuration [​](https://oxc.rs/docs/guide/usage/linter/rules/jest/prefer-lowercase-title.html#configuration)

This rule accepts a configuration object with the following properties:

### allowedPrefixes [​](https://oxc.rs/docs/guide/usage/linter/rules/jest/prefer-lowercase-title.html#allowedprefixes)

type: `string[]`

default: `[]`

This array option allows specifying prefixes, which contain capitals that titles can start with. This can be useful when writing tests for API endpoints, where you'd like to prefix with the HTTP method. By default, nothing is allowed (the equivalent of `{ "allowedPrefixes": [] }`).

Example of **correct** code for the `{ "allowedPrefixes": ["GET"] }` option:

js

```
/* jest/prefer-lowercase-title: ["error", { "allowedPrefixes": ["GET"] }] */
describe("GET /live");
```

### ignore [​](https://oxc.rs/docs/guide/usage/linter/rules/jest/prefer-lowercase-title.html#ignore)

type: `string[]`

default: `[]`

This array option controls which Jest or Vitest functions are checked by this rule. There are four possible values:

* `"describe"`
* `"test"`
* `"it"`
* `"bench"`

By default, none of these options are enabled (the equivalent of `{ "ignore": [] }`).

Example of **correct** code for the `{ "ignore": ["describe"] }` option:

js

```
/* jest/prefer-lowercase-title: ["error", { "ignore": ["describe"] }] */
describe("Uppercase description");
```

Example of **correct** code for the `{ "ignore": ["test"] }` option:

js

```
/* jest/prefer-lowercase-title: ["error", { "ignore": ["test"] }] */
test("Uppercase description");
```

Example of **correct** code for the `{ "ignore": ["it"] }` option:

js

```
/* jest/prefer-lowercase-title: ["error", { "ignore": ["it"] }] */
it("Uppercase description");
```

### ignoreTopLevelDescribe [​](https://oxc.rs/docs/guide/usage/linter/rules/jest/prefer-lowercase-title.html#ignoretopleveldescribe)

type: `boolean`

default: `false`

This option can be set to allow only the top-level `describe` blocks to have a title starting with an upper-case letter.

Example of **correct** code for the `{ "ignoreTopLevelDescribe": true }` option:

js

```
/* jest/prefer-lowercase-title: ["error", { "ignoreTopLevelDescribe": true }] */
describe("MyClass", () => {
  describe("#myMethod", () => {
    it("does things", () => {
      //
    });
  });
});
```

### lowercaseFirstCharacterOnly [​](https://oxc.rs/docs/guide/usage/linter/rules/jest/prefer-lowercase-title.html#lowercasefirstcharacteronly)

type: `boolean`

default: `true`

This option can be set to only validate that the first character of a test name is lowercased.

Example of **correct** code for the `{ "lowercaseFirstCharacterOnly": true }` option:

js

```
/* vitest/prefer-lowercase-title: ["error", { "lowercaseFirstCharacterOnly": true }] */
describe("myClass", () => {
  describe("myMethod", () => {
    it("does things", () => {
      //
    });
  });
});
```

Example of **incorrect** code for the `{ "lowercaseFirstCharacterOnly": true }` option:

js

```
/* vitest/prefer-lowercase-title: ["error", { "lowercaseFirstCharacterOnly": true }] */
describe("MyClass", () => {
  describe("MyMethod", () => {
    it("does things", () => {
      //
    });
  });
});
```

## How to use [​](https://oxc.rs/docs/guide/usage/linter/rules/jest/prefer-lowercase-title.html#how-to-use)

To **enable** this rule using the config file or in the CLI, you can use:

Config (.oxlintrc.json)CLI

json

```
{
  "plugins": ["jest"],
  "rules": {
    "jest/prefer-lowercase-title": "error"
  }
}
```

bash

```
oxlint --deny jest/prefer-lowercase-title --jest-plugin
```

## References [​](https://oxc.rs/docs/guide/usage/linter/rules/jest/prefer-lowercase-title.html#references)

* [Rule Source](https://github.com/oxc-project/oxc/blob/1bf0ffc0f6859c90409a9701e62e8957ef1286cc/crates/oxc_linter/src/rules/jest/prefer_lowercase_title/mod.rs)
