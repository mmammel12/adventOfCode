const { readFile } = require("../../readFileMac");

let data = [];
readFile(
    "2022/day4/input.txt",
    (fileData) => {
        data = fileData;
    },
    (err) => console.log(err)
);

const groupedData = [];
data.forEach((line) => {
    const [left, right] = line.split(",");
    const [leftStart, leftEnd] = left.split("-");
    const [rightStart, rightEnd] = right.split("-");
    groupedData.push({
        leftStart: Number(leftStart),
        leftEnd: Number(leftEnd),
        rightStart: Number(rightStart),
        rightEnd: Number(rightEnd),
    });
});

console.log(
    groupedData.reduce((acc, group) => {
        let num = 0;
        const { leftStart, leftEnd, rightStart, rightEnd } = group;
        if (leftStart <= rightStart && leftEnd >= rightEnd) {
            // left fully contains right
            num = 1;
        } else if (rightStart <= leftStart && rightEnd >= leftEnd) {
            // right fully contains left
            num = 1;
        } else if (leftStart <= rightStart && leftEnd >= rightStart) {
            // left overlaps right
            num = 1;
        } else if (rightStart <= leftStart && rightEnd >= leftStart) {
            // right overlaps left
            num = 1;
        }
        return acc + num;
    }, 0)
);
