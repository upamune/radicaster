---
title: "oxc/no-async-endpoint-handlers | The JavaScript Oxidation Compiler"
source_url: "https://oxc.rs/docs/guide/usage/linter/rules/oxc/no-async-endpoint-handlers"
fetched_at: "2025-12-31T10:44:14.234719+00:00"
---



Are you an LLM? You can read better optimized documentation at /docs/guide/usage/linter/rules/oxc/no-async-endpoint-handlers.md for this page in Markdown format

# oxc/no-async-endpoint-handlers Suspicious [​](https://oxc.rs/docs/guide/usage/linter/rules/oxc/no-async-endpoint-handlers.html#oxc-no-async-endpoint-handlers)

### What it does [​](https://oxc.rs/docs/guide/usage/linter/rules/oxc/no-async-endpoint-handlers.html#what-it-does)

Disallows the use of `async` functions as Express endpoint handlers.

### Why is this bad? [​](https://oxc.rs/docs/guide/usage/linter/rules/oxc/no-async-endpoint-handlers.html#why-is-this-bad)

Before v5, Express will not automatically handle Promise rejections from handler functions with your application's error handler. You must instead explicitly pass the rejected promise to `next()`.

js

```
const app = express();
app.get("/", (req, res, next) => {
  new Promise((resolve, reject) => {
    return User.findById(req.params.id);
  })
    .then((user) => res.json(user))
    .catch(next);
});
```

If this is not done, your server will crash with an unhandled promise rejection.

js

```
const app = express();
app.get("/", async (req, res) => {
  // Server will crash if User.findById rejects
  const user = await User.findById(req.params.id);
  res.json(user);
});
```

See [Express' Error Handling Guide](https://expressjs.com/en/guide/error-handling.html) for more information.

### Examples [​](https://oxc.rs/docs/guide/usage/linter/rules/oxc/no-async-endpoint-handlers.html#examples)

Examples of **incorrect** code for this rule:

js

```
const app = express();
app.get("/", async (req, res) => {
  const user = await User.findById(req.params.id);
  res.json(user);
});

const router = express.Router();
router.use(async (req, res, next) => {
  const user = await User.findById(req.params.id);
  req.user = user;
  next();
});

const createUser = async (req, res) => {
  const user = await User.create(req.body);
  res.json(user);
};
app.post("/user", createUser);

// Async handlers that are imported will not be detected because each
// file is checked in isolation. This does not trigger the rule, but still
// violates it and _will_ result in server crashes.
const asyncHandler = require("./asyncHandler");
app.get("/async", asyncHandler);
```

Examples of **correct** code for this rule:

js

```
const app = express();
// not async
app.use((req, res, next) => {
  req.receivedAt = Date.now();
});

app.get("/", (req, res, next) => {
  fs.readFile("/file-does-not-exist", (err, data) => {
    if (err) {
      next(err); // Pass errors to Express.
    } else {
      res.send(data);
    }
  });
});

const asyncHandler = async (req, res) => {
  const user = await User.findById(req.params.id);
  res.json(user);
};
app.get("/user", (req, res, next) => asyncHandler(req, res).catch(next));
```

## Configuration [​](https://oxc.rs/docs/guide/usage/linter/rules/oxc/no-async-endpoint-handlers.html#configuration)

This rule accepts a configuration object with the following properties:

### allowedNames [​](https://oxc.rs/docs/guide/usage/linter/rules/oxc/no-async-endpoint-handlers.html#allowednames)

type: `string[]`

default: `[]`

An array of names that are allowed to be async.

## How to use [​](https://oxc.rs/docs/guide/usage/linter/rules/oxc/no-async-endpoint-handlers.html#how-to-use)

To **enable** this rule using the config file or in the CLI, you can use:

Config (.oxlintrc.json)CLI

json

```
{
  "rules": {
    "oxc/no-async-endpoint-handlers": "error"
  }
}
```

bash

```
oxlint --deny oxc/no-async-endpoint-handlers
```

## References [​](https://oxc.rs/docs/guide/usage/linter/rules/oxc/no-async-endpoint-handlers.html#references)

* [Rule Source](https://github.com/oxc-project/oxc/blob/1bf0ffc0f6859c90409a9701e62e8957ef1286cc/crates/oxc_linter/src/rules/oxc/no_async_endpoint_handlers.rs)
