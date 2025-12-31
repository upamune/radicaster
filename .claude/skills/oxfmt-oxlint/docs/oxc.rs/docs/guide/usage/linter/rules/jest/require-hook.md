---
title: "jest/require-hook | The JavaScript Oxidation Compiler"
source_url: "https://oxc.rs/docs/guide/usage/linter/rules/jest/require-hook"
fetched_at: "2025-12-31T10:44:14.234719+00:00"
---



Are you an LLM? You can read better optimized documentation at /docs/guide/usage/linter/rules/jest/require-hook.md for this page in Markdown format

# jest/require-hook Style [​](https://oxc.rs/docs/guide/usage/linter/rules/jest/require-hook.html#jest-require-hook)

### What it does [​](https://oxc.rs/docs/guide/usage/linter/rules/jest/require-hook.html#what-it-does)

This rule flags any expression that is either at the toplevel of a test file or directly within the body of a `describe`, *except* for the following:

* `import` statements
* `const` variables
* `let` *declarations*, and initializations to `null` or `undefined`
* Classes
* Types
* Calls to the standard Jest globals

### Why is this bad? [​](https://oxc.rs/docs/guide/usage/linter/rules/jest/require-hook.html#why-is-this-bad)

Having setup and teardown code outside of hooks can lead to unpredictable test behavior. Code that runs at the top level executes when the test file is loaded, not when tests run, which can cause issues with test isolation and make tests dependent on execution order. Using proper hooks like `beforeEach`, `beforeAll`, `afterEach`, and `afterAll` ensures that setup and teardown code runs at the correct time and maintains test isolation.

### Examples [​](https://oxc.rs/docs/guide/usage/linter/rules/jest/require-hook.html#examples)

Examples of **incorrect** code for this rule:

javascript

```
import { database, isCity } from "../database";
import { Logger } from "../../../src/Logger";
import { loadCities } from "../api";

jest.mock("../api");

const initializeCityDatabase = () => {
  database.addCity("Vienna");
  database.addCity("San Juan");
  database.addCity("Wellington");
};

const clearCityDatabase = () => {
  database.clear();
};

initializeCityDatabase();

test("that persists cities", () => {
  expect(database.cities.length).toHaveLength(3);
});
test("city database has Vienna", () => {
  expect(isCity("Vienna")).toBeTruthy();
});

test("city database has San Juan", () => {
  expect(isCity("San Juan")).toBeTruthy();
});

describe("when loading cities from the api", () => {
  let consoleWarnSpy = jest.spyOn(console, "warn");
  loadCities.mockResolvedValue(["Wellington", "London"]);

  it("does not duplicate cities", async () => {
    await database.loadCities();
    expect(database.cities).toHaveLength(4);
  });
});
clearCityDatabase();
```

Examples of **correct** code for this rule:

javascript

```
import { database, isCity } from "../database";
import { Logger } from "../../../src/Logger";
import { loadCities } from "../api";

jest.mock("../api");
const initializeCityDatabase = () => {
  database.addCity("Vienna");
  database.addCity("San Juan");
  database.addCity("Wellington");
};

const clearCityDatabase = () => {
  database.clear();
};

beforeEach(() => {
  initializeCityDatabase();
});

test("that persists cities", () => {
  expect(database.cities.length).toHaveLength(3);
});

test("city database has Vienna", () => {
  expect(isCity("Vienna")).toBeTruthy();
});

test("city database has San Juan", () => {
  expect(isCity("San Juan")).toBeTruthy();
});

describe("when loading cities from the api", () => {
  let consoleWarnSpy;
  beforeEach(() => {
    consoleWarnSpy = jest.spyOn(console, "warn");
    loadCities.mockResolvedValue(["Wellington", "London"]);
  });

  it("does not duplicate cities", async () => {
    await database.loadCities();
    expect(database.cities).toHaveLength(4);
  });
});
afterEach(() => {
  clearCityDatabase();
});
```

This rule is compatible with [eslint-plugin-vitest](https://github.com/vitest-dev/eslint-plugin-vitest/blob/main/docs/rules/require-hook.md), to use it, add the following configuration to your `.oxlintrc.json`:

json

```
{
  "rules": {
    "vitest/require-hook": "error"
  }
}
```

## Configuration [​](https://oxc.rs/docs/guide/usage/linter/rules/jest/require-hook.html#configuration)

This rule accepts a configuration object with the following properties:

### allowedFunctionCalls [​](https://oxc.rs/docs/guide/usage/linter/rules/jest/require-hook.html#allowedfunctioncalls)

type: `string[]`

default: `[]`

An array of function names that are allowed to be called outside of hooks.

## How to use [​](https://oxc.rs/docs/guide/usage/linter/rules/jest/require-hook.html#how-to-use)

To **enable** this rule using the config file or in the CLI, you can use:

Config (.oxlintrc.json)CLI

json

```
{
  "plugins": ["jest"],
  "rules": {
    "jest/require-hook": "error"
  }
}
```

bash

```
oxlint --deny jest/require-hook --jest-plugin
```

## References [​](https://oxc.rs/docs/guide/usage/linter/rules/jest/require-hook.html#references)

* [Rule Source](https://github.com/oxc-project/oxc/blob/1bf0ffc0f6859c90409a9701e62e8957ef1286cc/crates/oxc_linter/src/rules/jest/require_hook.rs)
