import { scanFileLinesToNumber } from "../utils/utils";

const run = () => {
  const nums = scanFileLinesToNumber("day1/input.txt");

  console.time("Day 1");
  console.log(`Part 1 answer: ${part1(nums)}`);
  console.log(`Part 2 answer: ${part2(nums)}`);
  console.timeEnd("Day 1");
};

const part1 = (nums: number[]): number => {
  console.time("Part 1");
  for (let i = 0; i < nums.length - 1; i += 1) {
    for (let j = 0; j < nums.length; j += 1) {
      if (nums[i] + nums[j] === 2020) {
        console.timeEnd("Part 1");
        return nums[i] * nums[j];
      }
    }
  }
  return 0;
};

const part2 = (nums: number[]): number => {
  console.time("Part 2");
  for (let i = 0; i < nums.length - 2; i += 1) {
    for (let j = 0; j < nums.length - 1; j += 1) {
      for (let k = 0; k < nums.length; k += 1) {
        if (nums[i] + nums[j] + nums[k] === 2020) {
          console.timeEnd("Part 2");
          return nums[i] * nums[j] * nums[k];
        }
      }
    }
  }
  return 0;
};

run();
