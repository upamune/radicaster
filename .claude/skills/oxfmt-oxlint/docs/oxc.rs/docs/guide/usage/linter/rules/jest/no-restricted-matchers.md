---
title: "jest/no-restricted-matchers | The JavaScript Oxidation Compiler"
source_url: "https://oxc.rs/docs/guide/usage/linter/rules/jest/no-restricted-matchers"
fetched_at: "2025-12-31T10:44:14.234719+00:00"
---



Are you an LLM? You can read better optimized documentation at /docs/guide/usage/linter/rules/jest/no-restricted-matchers.md for this page in Markdown format

# jest/no-restricted-matchers Style [​](https://oxc.rs/docs/guide/usage/linter/rules/jest/no-restricted-matchers.html#jest-no-restricted-matchers)

### What it does [​](https://oxc.rs/docs/guide/usage/linter/rules/jest/no-restricted-matchers.html#what-it-does)

Ban specific matchers & modifiers from being used, and can suggest alternatives.

### Why is this bad? [​](https://oxc.rs/docs/guide/usage/linter/rules/jest/no-restricted-matchers.html#why-is-this-bad)

Some matchers or modifiers might be discouraged in your codebase for various reasons: they might be deprecated, cause confusion, have performance implications, or there might be better alternatives available. This rule allows you to enforce consistent testing patterns by restricting certain Jest matchers and providing guidance on preferred alternatives.

### Examples [​](https://oxc.rs/docs/guide/usage/linter/rules/jest/no-restricted-matchers.html#examples)

Bans are expressed in the form of a map, with the value being either a string message to be shown, or null if only the default rule message should be used. Bans are checked against the start of the expect chain - this means that to ban a specific matcher entirely you must specify all six permutations, but allows you to ban modifiers as well. By default, this map is empty, meaning no matchers or modifiers are banned.

Example configuration:

json

```
{
  "jest/no-restricted-matchers": [
    "error",
    {
      "toBeFalsy": null,
      "resolves": "Use `expect(await promise)` instead.",
      "toHaveBeenCalledWith": null,
      "not.toHaveBeenCalledWith": null,
      "resolves.toHaveBeenCalledWith": null,
      "rejects.toHaveBeenCalledWith": null,
      "resolves.not.toHaveBeenCalledWith": null,
      "rejects.not.toHaveBeenCalledWith": null
    }
  ]
}
```

Examples of **incorrect** code for this rule with the above configuration:

javascript

```
it("is false", () => {
  // if this has a modifier (i.e. `not.toBeFalsy`), it would be considered fine
  expect(a).toBeFalsy();
});

it("resolves", async () => {
  // all uses of this modifier are disallowed, regardless of matcher
  await expect(myPromise()).resolves.toBe(true);
});

describe("when an error happens", () => {
  it("does not upload the file", async () => {
    // all uses of this matcher are disallowed
    expect(uploadFileMock).not.toHaveBeenCalledWith("file.name");
  });
});
```

This rule is compatible with [eslint-plugin-vitest](https://github.com/vitest-dev/eslint-plugin-vitest/blob/main/docs/rules/no-restricted-matchers.md), to use it, add the following configuration to your `.oxlintrc.json`:

json

```
{
  "rules": {
    "vitest/no-restricted-matchers": "error"
  }
}
```

## Configuration [​](https://oxc.rs/docs/guide/usage/linter/rules/jest/no-restricted-matchers.html#configuration)

This rule accepts a configuration object with the following properties:

### restrictedMatchers [​](https://oxc.rs/docs/guide/usage/linter/rules/jest/no-restricted-matchers.html#restrictedmatchers)

type: `Record<string, string>`

default: `{}`

A map of restricted matchers/modifiers to custom messages. The key is the matcher/modifier name (e.g., "toBeFalsy", "resolves", "not.toHaveBeenCalledWith"). The value is an optional custom message to display when the matcher/modifier is used.

## How to use [​](https://oxc.rs/docs/guide/usage/linter/rules/jest/no-restricted-matchers.html#how-to-use)

To **enable** this rule using the config file or in the CLI, you can use:

Config (.oxlintrc.json)CLI

json

```
{
  "plugins": ["jest"],
  "rules": {
    "jest/no-restricted-matchers": "error"
  }
}
```

bash

```
oxlint --deny jest/no-restricted-matchers --jest-plugin
```

## References [​](https://oxc.rs/docs/guide/usage/linter/rules/jest/no-restricted-matchers.html#references)

* [Rule Source](https://github.com/oxc-project/oxc/blob/1bf0ffc0f6859c90409a9701e62e8957ef1286cc/crates/oxc_linter/src/rules/jest/no_restricted_matchers.rs)
