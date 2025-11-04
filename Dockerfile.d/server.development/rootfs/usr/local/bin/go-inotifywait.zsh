#!/usr/bin/env -S zsh -eu

while true
do
  go test
  inotifywait --recursive --event modify,close_write,move,create,delete .
done
