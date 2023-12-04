const fs = require("fs");

const readFile = (path, callback, errorCallback) => {
    try {
        console.log(`${__dirname}/${path}`);
        const data = fs.readFileSync(`${__dirname}/${path}`, "utf-8");
        if (callback != null) callback(data.split("\n"));
    } catch (err) {
        if (errorCallback != null) errorCallback(err);
    }
};

const getFileData = (path) => {
    let data = [];
    readFile(
        path,
        (fileData) => {
            data = fileData;
        },
        (err) => console.log(err)
    );
    return data;
};

exports.readFile = readFile;
exports.getFileData = getFileData;
