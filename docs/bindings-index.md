# Language Bindings Index

soroauth provides bindings for multiple languages over its C ABI core and WebAssembly target. This page lists all six bindings, their installation methods, and platform support.

## Overview of Bindings

| Binding | Language / Runtime | Package Registry / Install Method | Platform Support | Notes |
| --- | --- | --- | --- | --- |
| **wasm** | WebAssembly / JS | Published to npm (`@soroauth/wasm`) | All browsers and JS runtimes supporting WASM | Built from source or installed via npm. |
| **node** | Node.js (N-API) | Published to npm (`@soroauth/node`) | Linux and macOS only (`x86_64`, `aarch64`/`arm64`). Windows is not supported. | macOS thread-safety limitation under `node:test` is documented in the [Node README](../bindings/node/README.md) and linked to [#518](https://github.com/soroauth/soroauth-go/issues/518). |
| **c** | C / C++ | Build from source (`bindings/c/build.sh`) | Linux and macOS only (`x86_64`, `aarch64`/`arm64`) | C ABI shared library (`libsoroauth.so` / `libsoroauth.dylib`). |
| **jvm** | Java / JVM (FFM API) | Maven Central (`org.soroauth:soroauth-jvm`) | Linux and macOS for native symbols; pure JVM signing is platform-independent | JDK 22+ required for Foreign Function and Memory API. |
| **python** | Python (CFFI) | Build from source or local wheel (`bindings/python`) | Linux and macOS supporting CFFI and Python 3.10+ | Built from source. |
| **rust** | Rust | Cargo crate (`bindings/rust`) | Platform-independent (Rust targets) | Add to `Cargo.toml`. |

## Platform Support Summary

- **Native Bindings (c, node, python JVM native symbols)**: Linux and macOS only (`x86_64`, `aarch64` / `arm64`). Windows is not supported.
- **Pure / WASM / JVM / Rust**: Platform-independent or dependent on standard runtime targets.

## Node Binding macOS Limitation

For details regarding the known macOS thread-safety limitations and concurrency handling in the Node binding, please refer to the discussion in [#518](https://github.com/soroauth/soroauth-go/issues/518).
