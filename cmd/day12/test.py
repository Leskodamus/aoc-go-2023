#!/usr/bin/env python3

import pathlib
import re
import sys

from functools import cache

@cache
def arrangements(pattern: str, *res: re.Pattern) -> int:
  if len(res) == 0:
    return 0 if "#" in pattern else 1

  matches = 0
  index = 0

  while (m := res[0].search(pattern[index:])) and "#" not in pattern[:index + m.start()]:
    matches += arrangements(pattern[index + m.end() - 1:], *res[1:])
    index += m.start() + 1

  return matches

def run() -> None:
  with open('input.txt') as f:
    lines = [line.strip() for line in f.readlines()]

  total_arrangements = 0
  folded_arrangements = 0

  for line in lines:
    pattern, spring_lens = line.split()
    spring_res = [re.compile(f"[.?][#?]{{{int(s)}}}[.?]") for s in spring_lens.split(",")]
    total_arrangements += arrangements(f".{pattern}.", *spring_res)

    pattern = '?'.join((pattern, ) * 5)
    spring_res = [re.compile(f"[.?][#?]{{{int(s)}}}[.?]") for s in (spring_lens.split(",") * 5)]
    folded_arrangements += arrangements(f".{pattern}.", *spring_res)

  print(f"Sum of simple arrangements: {total_arrangements}")
  print(f"Sum of folded arrangements: {folded_arrangements}")

if __name__ == '__main__':
  run()
  sys.exit(0)
