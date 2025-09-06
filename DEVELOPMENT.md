# Development environment

This repo uses 'just' as a 'make' replacement.

Versions of tools used are in the .tool-versions file.

# Development workflow

## Tools

As an example upgrade golang:

```bash
asdf install golang latest
```

If successful, edit .tool-versions accordingly.

## Changing code

Edit or add code or other development activity.

Quality check code:

```bash
just qa
```

