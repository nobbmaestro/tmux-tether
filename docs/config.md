# User Config

Default path for the global configuration file:

- `~/.config/tmux-tether/tmux-tether.yml`

## Default

```yml
session:
  # List of directories to search in
  dirs:
    - ~/repos

  # Max depth to recurse below each dir. 0 = only the directory itself, -1 = unlimited
  depth: 0

  # Files or directories that mark a valid session root
  markers:
    - .git

  # Directory names or glob patterns to exclude from scanning
  exclude: []

picker:
  # Pointer icon
  pointer: " "

  # Prompt icon
  prompt: "  "
```
