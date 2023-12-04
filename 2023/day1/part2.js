const { readFile } = require("../../readFile");

let data = [];
readFile(
    "2023/day1/input.txt",
    (fileData) => {
        data = fileData;
    },
    (err) => console.log(err)
);

const numbers = {
    one: "o1e",
    two: "t2o",
    three: "t3e",
    four: "f4r",
    five: "f5e",
    six: "s6x",
    seven: "s7n",
    eight: "e8t",
    nine: "n9e",
};

let sum = 0;
data.forEach((line) => {
    if (line.length >= 2) {
        let replaceLine = line;
        Object.keys(numbers).forEach((key) => {
            replaceLine = replaceLine.replaceAll(key, numbers[key]);
        });
        const nums = replaceLine.match(/(\d)/g);

        let firstNum = nums[0];
        let secondNum = nums[nums.length - 1];
        if (firstNum.length > 1) {
            firstNum = numbers[firstNum].toString();
        }
        if (secondNum.length > 1) {
            secondNum = numbers[secondNum].toString();
        }
        const num = parseInt(firstNum + secondNum);

        sum += num;
    }
});

console.log(sum);
