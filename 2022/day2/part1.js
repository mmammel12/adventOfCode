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
    A: ROCK,
    B: PAPER,
    C: SCISSORS,
};

const SELF = {
    X: WIN,
    Y: PAPER,
    Z: SCISSORS,
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
    let score = POINTS[SELF[self]];
    const oppChoice = OPPONENT[opp];
    const selfChoice = SELF[self];
    if (oppChoice === selfChoice) {
        score += POINTS[DRAW];
    } else {
        if (oppChoice === ROCK) {
            if (selfChoice === PAPER) {
                score += POINTS[WIN];
            } else {
                score += POINTS[LOSE];
            }
        } else if (oppChoice === PAPER) {
            if (selfChoice === SCISSORS) {
                score += POINTS[WIN];
            } else {
                score += POINTS[LOSE];
            }
        } else if (oppChoice === SCISSORS) {
            if (selfChoice === ROCK) {
                score += POINTS[WIN];
            } else {
                score += POINTS[LOSE];
            }
        }
    }

    return score;
};

console.log(
    data.reduce((acc, line) => {
        const [opp, self] = line.split(" ");
        return acc + compare(opp, self);
    }, 0)
);
