<!-- # README

## About

This is the official Wails Svelte template.

## Live Development

To run in live development mode, run `wails dev` in the project directory. This will run a Vite development
server that will provide very fast hot reload of your frontend changes. If you want to develop in a browser
and have access to your Go methods, there is also a dev server that runs on http://localhost:34115. Connect
to this in your browser, and you can call your Go code from devtools.

## Building

To build a redistributable, production mode package, use `wails build`. -->

# DezeekeesDesktopList App

DezeekeesDesktopList is a desktop implementation of the [myanimelist website](https://myanimelist.net) . This README provides essential information on installing, developing, and building the application.

# Installation
To install DezeekeesDesktopList, you can use the dedicated installer tool, dezeekeesdesktoplist-installer. Follow the steps below:

[Download Installer](https://github.com/DeZeeKees/dezeekeesdesktoplist-installer/releases/latest)

Execute the installer to install DezeekeesDesktopList on your system.
# Development
DezeekeesDesktopList is built using Wails and SvelteKit. Follow the steps below to start developing the project:

## Prerequisites
Make sure you have the following dependencies installed:

- Go 1.21.6 (https://golang.org/dl)
- Node.js >= 20.7.0 (https://nodejs.org)
- Npm >= 10.4.0 (https://www.npmjs.com)
- Wails Cli 2.7.1 (https://wails.io)

not using the recommended go version might break the wails cli

## Live Development

To run in live development mode, run `wails dev` in the project directory. This will run a Vite development
server that will provide very fast hot reload of your frontend changes. If you want to develop in a browser
and have access to your Go methods, there is also a dev server that runs on http://localhost:34115. Connect
to this in your browser, and you can call your Go code from devtools.

```
wails dev -loglevel "warning" -ldflags "-X main.Version=0.2"
```

## Building

To build a redistributable, production mode package, use `wails build`.

```
wails build -ldflags "-X main.Version=0.3.1"
```
