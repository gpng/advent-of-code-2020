import { scanFileLinesToStrings } from "../utils/utils";

const run = () => {
  const rows = scanFileLinesToStrings("day3/input.txt", "");

  console.time("Day 3");
  console.log(`Part 1 answer: ${part1(rows)}`);
  console.log(`Part 2 answer: ${part2(rows)}`);
  console.timeEnd("Day 3");
};

const traverse = (rows: string[][], dx: number, dy: number): number => {
  let x = 0,
    y = 0,
    trees = 0;

  while (y < rows.length - dy) {
    y += dy;
    x += dx;

    const row = rows[y];
    if (row[x % row.length] === "#") trees++;
  }
  return trees;
};

const part1 = (rows: string[][]): number => {
  console.time("Part 1");

  const res = traverse(rows, 3, 1);

  console.timeEnd("Part 1");
  return res;
};

const part2 = (rows: string[][]): number => {
  console.time("Part 2");

  const res =
    traverse(rows, 1, 1) *
    traverse(rows, 3, 1) *
    traverse(rows, 5, 1) *
    traverse(rows, 7, 1) *
    traverse(rows, 1, 2);

  console.timeEnd("Part 2");
  return res;
};

run();
