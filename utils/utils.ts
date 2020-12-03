import { readFileSync } from "fs";

export const scanFile = (path: string): string => {
  return readFileSync(path, "utf-8");
};

export const scanFileLinesToNumber = (path: string): number[] =>
  scanFile(path)
    .split("\n")
    .map((line) => parseInt(line, 10));

export const scanFileLinesToNumbers = (
  path: string,
  sep: string = "\t"
): number[][] =>
  scanFile(path)
    .split("\n")
    .map((line) => line.split(sep).map((num) => parseInt(num, 10)));

export const scanFileLinesToStrings = (path: string, sep: string): string[][] =>
  scanFile(path)
    .split("\n")
    .map((line) => line.split(sep));
