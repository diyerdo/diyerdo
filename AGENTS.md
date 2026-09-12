# Repository Guidelines

## Project Overview

Diyerdo is a fan project made for the famous french MMORPG Dofus. It's a tool meant to replace other famous fan projects such as _Dofusbook_, _Dofuslab_, _Nokazu_, _Ganymède_, etc. Diyerdo has features like a simple and intuitive build planner, a fine-grained yet powerful item recipe resolver, an integrated instant messaging based on Matrix, and more.

Diyerdo mainly targets players who want to play Dofus in an SF (Self-found) mode and like to do everything themselves (DIY), especially item crafting.

## Expected Behavior

### Code Operations (MANDATORY)

**For ANY code-related task** — reading, searching, editing, analyzing, navigating, or modifying code files — **you MUST use serena-lsp's MCP tools exclusively.**
This includes, but is not limited to:

- Reading, searching, or analyzing **source files** (`.dart`, `.ts`, `.js`, `.py`, `.go`, etc.)
- Modifying, refactoring, or deleting **code logic** or **symbols**
- Debugging or validating **code structure** or **behavior**
  Use the appropriate serena-lsp MCP tool:
- **Symbol lookup**: `mcp__serena_find_symbol`, `mcp__serena_find_declaration`, `mcp__serena_find_implementations`, `mcp__serena_find_referencing_symbols`
- **File/pattern search**: `mcp__serena_search_for_pattern`
- **Content modification**: `mcp__serena_replace_content`, `mcp__serena_replace_in_files`, `mcp__serena_replace_symbol_body`, `mcp__serena_insert_after_symbol`, `mcp__serena_insert_before_symbol`, `mcp__serena_rename_symbol`, `mcp__serena_safe_delete_symbol`
- **Diagnostics**: `mcp__serena_get_diagnostics_for_file`, `mcp__serena_get_symbols_overview`
- **Memory**: `mcp__serena_read_memory`, `mcp__serena_write_memory`, `mcp__serena_list_memories`, `mcp__serena_delete_memory`, `mcp__serena_rename_memory`, `mcp__serena_edit_memory`
- **Setup**: `mcp__serena_onboarding`, `mcp__serena_initial_instructions`

### Enforcement

- **Automatic Blocking**: If an agent attempts to use non-MCP tools (`read`, `grep`, `glob`, `bash`, `write`, `edit`) for code-related tasks, **immediately raise a blocker** to halt execution and enforce MCP compliance. Except for `git`.
- **Tool Restrictions**: Non-MCP tools (`read`, `write`, `edit`, `glob`, etc.) **MUST NEVER** be used for code operations. Use only the **approved MCP tools** listed above.

### General Behavior

- _Always_ search the web for information before answering
- I don't want you to make things up, if you can't find an information, just say "I don't know", it's OK, we'll find it together
- I don't expect an answer all the time, I prefer you to say that you can't find the information, rather than making things up
- You _must_ make the shortest possible answer
- _Never_ give detailed answer, unless the user asks for it (e.g The user said "can you explain me how <subject> works?", or something similar)

## Git

Whenever you make a commit, you must follow the following rules:

- Follow the angular commit convention (e.g `fix(scope): fix something`, `feat(scope): add something`, `build: did something related to the build`, etc.)
- You _must_ use the following commit identity: `bot-omp <bot.omp@lahventure.org>`
- You **MUST NOT** change the local git configuration, use the `git` command line parameter to commit with your special identity without modifying the `git config`
- **NEVER** use the `feat` commit type, it's only used as merge commit

Here's a breakdown of the git commit type enforced by the angular conventional commit:

| Type         | Description                                                                                         |
| ------------ | --------------------------------------------------------------------------------------------------- |
| **build**    | Changes that affect the build system or external dependencies (example scopes: gulp, broccoli, npm) |
| **ci**       | Changes to our CI configuration files and scripts (examples: GitHub Actions)                        |
| **docs**     | Documentation only changes                                                                          |
| **fix**      | A bug fix                                                                                           |
| **perf**     | A code change that improves performance                                                             |
| **refactor** | A code change that neither fixes a bug nor adds a feature                                           |
| **test**     | Adding missing tests or correcting existing tests                                                   |

## Architecture

TODO

### Proto Definition

TODO

## Key Directories

TODO

### Key Files

TODO

## Conventions & Common Patterns

TODO
