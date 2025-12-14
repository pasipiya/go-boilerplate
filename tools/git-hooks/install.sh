#!/bin/bash
HOOKS_DIR=$(git rev-parse --git-dir)/hooks

echo "Installing Git hooks to $HOOKS_DIR ..."

cp tools/git-hooks/pre-commit $HOOKS_DIR/pre-commit
cp tools/git-hooks/pre-push $HOOKS_DIR/pre-push

chmod +x $HOOKS_DIR/pre-commit
chmod +x $HOOKS_DIR/pre-push

echo "Hooks installed successfully!"
