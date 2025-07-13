#!/usr/bin/env bash
set -euo pipefail

for svg in "$@"; do
  png="${svg%.svg}.png"
  convert -density 300 -background none "$svg" "$png"
  echo "→ $png"
done
