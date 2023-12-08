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
const START_NODE_ENDING = "A";
const END_NODE_ENDING = "Z";

const nodes = data.reduce((acc, line, i) => {
    if (i < 2) return acc;

    const node = line.substring(NODE_START_INDEX, NODE_END_INDEX);
    const [left, right] = line.substring(CONNECTIONS_START_INDEX, CONNECTIONS_END_INDEX).split(", ");

    acc[node] = { left, right };

    return acc;
}, {});

const numberOfSteps = (node) => {
    let steps = 0;
    let currentNode = node;
    let instructionIndex = 0;
    while (!currentNode.endsWith(END_NODE_ENDING)) {
        const currentInstruction = instructions[instructionIndex];

        currentNode = nodes[currentNode][directions[currentInstruction]];

        steps++;
        instructionIndex = (instructionIndex + 1) % instructions.length;
    }

    return steps;
};

const gcd = (a, b) => {
    if (b === 0) return a;

    return gcd(b, a % b);
};

const lcm = (a, b) => {
    return (a * b) / gcd(a, b);
};

let startNodes = Object.keys(nodes).filter((node) => node.endsWith(START_NODE_ENDING));
const steps = startNodes.map(numberOfSteps).reduce((acc, current) => lcm(acc, current), 1);

console.log(steps);
