const { readFile } = require("../../readFileMac");

let data = [];
readFile(
    "2022/day6/input.txt",
    (fileData) => {
        data = fileData;
    },
    (err) => console.log(err)
);

const MAX_SIZE = 100000;

const dir = { "/": {} };
const currentDir = "/";
data.forEach((line) => {
    const lineArr = line.split(" ");
    if (lineArr[0] === "$") {
        // command
        const command = lineArr[1];
        if (command === "cd") {
            // change dir
        } else {
            // list
        }
    } else if (lineArr[0] === "dir") {
        // dir
        const name = lineArr[1];
    } else {
        // file
        const size = Number(lineArr[0]);
        const name = lineArr[1];
    }
});
