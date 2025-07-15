# Hot Module Reloading (HMR) in Aether 0.4.0-nightly 🍕🔥

Welcome to the future of rapid development in Aether! With HMR, you can script, build games, and develop engines with the speed of a dynamic language and the power of native code. No more restarts, no more lost state—just pure, hot-reloadable code vibes.

---

## 1. What is HMR?

**Hot Module Reloading (HMR)** lets you update your code while your app, server, or game is running. Change a file, and Aether swaps in the new code instantly—no restart, no downtime, no lost progress.

- Perfect for scripting with a compiled language
- Game and engine development
- Live servers and tools

---

## 2. How HMR Works in Aether

- **HMR Off (default):**
  - Your project is compiled as a single, static binary for max performance and zero bloat.
  - Great for production, CLI tools, and simple deployment.

- **HMR On:**
  - Every file is compiled as a hot-reloadable module (dynamic/shared library).
  - Any change to any file can be reloaded instantly, without restarting the process.
  - State is preserved, so you can tweak code and see results immediately—ideal for games, live servers, and rapid prototyping.
  - Switching HMR on or off requires a restart to keep things safe and simple.

---

## 3. Use Cases

- **Scripting with a compiled language:**
  - Write scripts that are compiled, but can be swapped out at runtime for fast feedback.
- **Game development:**
  - Instantly update gameplay logic, AI, or UI without restarting your game or losing state.
- **Game engines:**
  - Reload engine subsystems or plugins on the fly.
- **Live servers:**
  - Patch logic or add features without downtime.
- **Tooling:**
  - Build editors or plugins that reload instantly as you work.

---

## 4. How to Enable HMR

- Use the CLI flag or environment variable:
  - `aether run mygame.ae --hmr`
  - Or: `AETHER_HMR=1 aether run mygame.ae`
- **Note:** Toggling HMR requires a restart.

---

## 5. Best Practices

- **State Management:**
  - Keep persistent state outside hot-reloaded modules, or use hooks to save/restore state.
- **Module Boundaries:**
  - In HMR mode, every file is a module. Design your code to take advantage of this for maximum flexibility.
- **Opt-Out:**
  - (Optional) Exclude files from HMR if needed via config or naming convention.
- **Debugging:**
  - If a reload fails (e.g., compilation error), the old module stays active and your app keeps running.

---

## 6. FAQ

**Q: What happens if a reload fails?**
A: The old module stays loaded, and your app/game/server keeps running. You’ll get an error message, but no crash.

**Q: Can I use HMR in production?**
A: HMR is designed for development. For production, keep HMR off for best performance and stability.

**Q: Does HMR add bloat?**
A: Only when enabled! With HMR off, you get a lean, static binary.

**Q: Why does toggling HMR require a restart?**
A: To keep things safe and predictable. The runtime needs to know if it should load modules dynamically or not.

---

Ready to build games, tools, and servers with the speed of scripting and the power of native code? HMR in Aether 0.4.0-nightly has you covered. 🍕🚀 