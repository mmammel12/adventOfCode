const { getFileData } = require("../../readFile");

const data = getFileData("2023/day8/input.txt");

const instructions = data[0].split("");

const directions = {
    L: "left",
    R: "right",
};

const NODE_START_INDEX = 0;
const NODE_END_INDEX = 3;
const CONNECTIONS_START_INDEX = 7;
const CONNECTIONS_END_INDEX = 15;
const START_NODE = "AAA";
const END_NODE = "ZZZ";

const nodes = data.reduce((acc, line, i) => {
    if (i < 2) return acc;

    const node = line.substring(NODE_START_INDEX, NODE_END_INDEX);
    const [left, right] = line.substring(CONNECTIONS_START_INDEX, CONNECTIONS_END_INDEX).split(", ");

    acc[node] = { left, right };

    return acc;
}, {});

let steps = 0;
let currentNode = START_NODE;
let instructionIndex = 0;
while (currentNode !== END_NODE) {
    const currentInstruction = instructions[instructionIndex];

    currentNode = nodes[currentNode][directions[currentInstruction]];

    steps++;
    instructionIndex = (instructionIndex + 1) % instructions.length;
}

console.log(steps);
