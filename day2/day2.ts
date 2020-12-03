import { scanFileLinesToStrings } from "../utils/utils";

const run = () => {
  const rows = scanFileLinesToStrings("day2/input.txt", " ");

  console.time("Day 2");
  console.log(`Part 1 answer: ${part1(rows)}`);
  console.log(`Part 2 answer: ${part2(rows)}`);
  console.timeEnd("Day 2");
};

const part1 = (rows: string[][]): number => {
  console.time("Part 1");

  const res = rows.reduce((r, row) => {
    const [min, max] = row[0].split("-").map((x) => parseInt(x, 10));
    const char = row[1].charAt(0);
    const password = row[2];

    const count = password
      .split("")
      .reduce((acc, val) => (val === char ? acc + 1 : acc), 0);

    return count >= min && count <= max ? r + 1 : r;
  }, 0);

  console.timeEnd("Part 1");
  return res;
};

const part2 = (rows: string[][]): number => {
  console.time("Part 2");

  const res = rows.reduce((r, row) => {
    const [firstIndex, secondIndex] = row[0]
      .split("-")
      .map((x) => parseInt(x, 10));
    const char = row[1].charAt(0);
    const password = row[2].split("");

    const first = password[firstIndex - 1];
    const second = password[secondIndex - 1];

    return first != second && (first === char || second === char) ? r + 1 : r;
  }, 0);

  console.timeEnd("Part 2");
  return res;
};

run();
