const { getFileData } = require('../../readFile');

const data = getFileData('2023/day12/test.txt');

const getRowParts = (row) => {
    const [record, valuesStr] = row.split(' ');
    const values = valuesStr.split(',').map((value) => parseInt(value));

    return { record, values };
};

let sum = 0;

data.forEach((row) => {
    const { record, values } = getRowParts(row);

    console.log('hi');
});

console.log(test);
