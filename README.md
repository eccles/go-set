# go-set

Generic set for Golang

Loosely based on the python3 sets operations using generics and iterators.

Uses a map as a synonym of of a set. Members of a set must be comparable. 
Unfortunately this precludes a 'Set of sets' as maps (on which Sets are based) are not comparable.


# Preparation

This repo uses asdf to manage tool versions and just as a 'make' replacement.

## asdf

https://asdf-vm.com/guide/getting-started.html

Install asdf following the above instructions
and ensure the following is in ~/.bashrc:

```bash
export PATH="${ASDF_DATA_DIR:-$HOME/.asdf}/shims:$PATH"
. <(asdf completion bash)
```

## Reset environment

In order to make these settings permanent, logout or reboot your PC.

## Installing tools

This wil install go and other tools including 'just'.

```bash
./scripts/asdf-install.sh
```

and then install go tools:

```bash
just tools
```

# Development workflow

## On a rebase

Initialise tools and modules

```bash
just tools
just qa
```
## Changing code

Edit or add code or other development activity.

Quality check code:

```bash
just qa
```

