const { readFile } = require("../../readFileMac");

let data = [];
readFile(
    "2022/day5/moves.txt",
    (fileData) => {
        data = fileData;
    },
    (err) => console.log(err)
);
const instructions = [];
data.forEach((line) => {
    const [move, amount, from, stack1, to, stack2] = line.split(" ");
    instructions.push({ [move]: amount, [from]: stack1, [to]: stack2 });
});

let data2 = [];
readFile(
    "2022/day5/stacks.txt",
    (fileData) => {
        data2 = fileData;
    },
    (err) => console.log(err)
);

const STACKS = {
    1: [],
    2: [],
    3: [],
    4: [],
    5: [],
    6: [],
    7: [],
    8: [],
    9: [],
};

data2.forEach((line) => {
    let stack = 1;
    for (let i = 0; i < line.length; i += 4) {
        const char = line[i + 1];
        if (char !== " ") {
            STACKS[stack].unshift(char);
        }
        stack++;
    }
});

instructions.forEach((instruction) => {
    const { move, from, to } = instruction;
    for (let i = move; i > 0; i--) {
        const char = STACKS[from].pop();
        STACKS[to].push(char);
    }
});

let str = "";
for (let i = 1; i <= 9; i++) {
    str += STACKS[i].pop();
}

console.log(str);
