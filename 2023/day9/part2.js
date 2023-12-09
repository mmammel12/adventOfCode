const { getFileData } = require("../../readFile");

const data = getFileData("2023/day9/input.txt");

const sequences = data.map((line) => line.split(" ").map(Number));

const createStepList = (sequence) => {
    const stepList = sequence.reduce((acc, num, i) => {
        if (i + 1 < sequence.length) {
            acc.push(sequence[i + 1] - num);
        }

        return acc;
    }, []);

    return stepList;
};

let sum = 0;

sequences.forEach((sequence) => {
    const stepLists = [[...sequence]];
    let index = 0;
    while (stepLists[index].some((num) => num !== stepLists[index][0])) {
        stepLists.push(createStepList(stepLists[index]));
        index++;
    }

    let nextNumber = 0;
    for (let i = stepLists.length - 1; i >= 0; i--) {
        nextNumber = stepLists[i][0] - nextNumber;
    }

    sum += nextNumber;
});

console.log(sum);
