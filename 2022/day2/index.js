const { readFile } = require("../../readFileMac");

let data = [];
readFile(
    "2022/day2/input.txt",
    (fileData) => {
        data = fileData;
    },
    (err) => console.log(err)
);

const ROCK = "ROCK";
const PAPER = "PAPER";
const SCISSORS = "SCISSORS";
const WIN = "WIN";
const LOSE = "LOSE";
const DRAW = "DRAW";

const OPPONENT = {
    A: { WIN: PAPER, LOSE: SCISSORS, DRAW: ROCK },
    B: { WIN: SCISSORS, LOSE: ROCK, DRAW: PAPER },
    C: { WIN: ROCK, LOSE: PAPER, DRAW: SCISSORS },
};

const SELF = {
    X: LOSE,
    Y: DRAW,
    Z: WIN,
};

const POINTS = {
    [ROCK]: 1,
    [PAPER]: 2,
    [SCISSORS]: 3,
    [WIN]: 6,
    [LOSE]: 0,
    [DRAW]: 3,
};

const compare = (opp, self) => {
    return POINTS[OPPONENT[opp][SELF[self]]] + POINTS[SELF[self]];
};

console.log(
    data.reduce((acc, line) => {
        const [opp, self] = line.split(" ");
        return acc + compare(opp, self);
    }, 0)
);
