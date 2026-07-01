# Rust Learning Path — Claude Context

## What this folder is
Local Cargo projects for Rust learning exercises. Each subfolder is a standalone Cargo project (`src/main.rs` is the code that matters).

## Where to upload
Repo: `VextorHorizon/vh_trainingground` (public)
Path pattern: `rust/<date>/`
Date format: `M-D-YY` (e.g. `6-4-26`)

## Upload format
- Extract `src/main.rs` from each subfolder
- Name the file after the subfolder (e.g. `read_file` → `read_file.rs`)
- Group all same-day exercises under one date folder
- Upload via GitHub Contents API (`gh api PUT`) — no cloning needed

## Already uploaded
| File | Uploaded under |
|------|---------------|
| first.rs | rust/5-10-26, rust/6-2-26 |
| frot.rs | rust/6-2-26 |
| reversed.rs | rust/6-2-26 |
| second-triedon-stdin.rs | rust/6-2-26 |
| third.rs | rust/6-2-26 |
| input_name.rs | rust/6-3-26 |
| read_file.rs | rust/6-4-26 |
| sessionfirst_ownership.rs | rust/6-4-26 |
| something.rs | rust/6-4-26 |
| read_file_ext.rs | rust/6-4-26 |
| name_checker.rs | rust/6-6-26 |
| ownlearning.rs | rust/6-7-26 |
| func_wit.rs | rust/6-11-26 (updated 6-14-26) |
| password_checker.rs | rust/6-14-26 |
| add_one.rs | rust/6-14-26 |
| double_it.rs | rust/6-14-26 |

## When user asks to update
1. `ls` the folder root — find subfolders not in the "already uploaded" list above
2. Read each new subfolder's `src/main.rs`
3. Base64 encode and upload via `gh api --method PUT repos/VextorHorizon/vh_trainingground/contents/rust/<date>/<name>.rs`
4. Commit message: `add rust exercises <date>`
5. Update the "already uploaded" table above

## Do NOT upload
- `target/` directories
- `Cargo.lock`, `Cargo.toml`
- This `CLAUDE.md` file
- `.vscode/`
