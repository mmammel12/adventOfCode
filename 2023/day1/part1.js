const { readFile } = require("../../readFile");

let data = [];
readFile(
    "2023/day1/input.txt",
    (fileData) => {
        data = fileData;
    },
    (err) => console.log(err)
);

let sum = 0;

data.forEach((line) => {
    if (line.length >= 2) {
        const nums = line.match(/(\d)/g);
        const num = parseInt(nums[0] + nums[nums.length - 1]);
        sum += num;
    }
});

console.log(sum);
