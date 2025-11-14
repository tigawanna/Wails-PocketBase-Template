# README

---

## About

This is a Wails-PocketBase template using SvelteKit as the frontend.

## Setup

Create a new wails project:

```bash
wails init -n <NAME> -d <DIRECTORY = "./" for root>
```

Delete the contents of `./frontend` and create a new SvelteKit (or any other framework) project:

```bash
npx sv create frontend
```

Ensure that the framework builds a static site, in the case of SvelteKit we use `adapter-static` to crate a SPA and make the following changes:

```javascript
// ./frontend/svelte.config.js

import adapter from '@sveltejs/adapter-static';

/** @type {import('@sveltejs/kit').Config} */
const config = {
	kit: {
		adapter: adapter({
			pages: 'dist',
			assets: 'dist',
			fallback: 'index.html',
			precompress: false,
			strict: true
		})
	}
};

export default config;
```

```javascript
// .frontend/src/routes/+layout.js

export const prerender = false;
export const ssr = false;
```

Edit the `./wails.json` file with the following (in the case of sveltekit, helps for methods to be imported via the [`$lib`](https://svelte.dev/docs/kit/$lib) alias)

```json
"wailsjsdir": "frontend/src/lib",
```

After the initial setup, you can interact with PocketBase in several ways: by using the provided helpers and hooks within the "extend with GO" section (which lets you write backend logic in Go), by leveraging the "extends with JS" option (for extending functionality in JavaScript), or by using the regular JS SDK with ‘http://wails.localhost:8080/’.

## Live Development

To run in live development mode, run `wails dev` in the project directory. This will run a Vite development
server that will provide very fast hot reload of your frontend changes. If you want to develop in a browser
and have access to your Go methods, there is also a dev server that runs on http://localhost:34115. Connect
to this in your browser, and you can call your Go code from devtools.

## Building

To build a redistributable, production mode package, use `wails build`.

## Considerations

While this is the easiest way I have found to create a full-featured desktop app, it is still not an efficient way to build a desktop app. Using SQLite directly within Wails will be faster and more efficient, and writing your own logic for CRUD operations in the filesystem will be more effective.

That being said, I do use this combination for personal projects, but will advise against using it in a production environment.
