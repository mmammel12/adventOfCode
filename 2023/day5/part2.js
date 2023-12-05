const { getFileData } = require("../../readFile");

const data = getFileData("2023/day5/input.txt");

const SEEDS_TO_SOIL_START = 3;
const SEEDS_TO_SOIL_END = 25;

const SOIL_TO_FERTILIZER_START = SEEDS_TO_SOIL_END + 2;
const SOIL_TO_FERTILIZER_END = 55;

const FERTILIZER_TO_WATER_START = SOIL_TO_FERTILIZER_END + 2;
const FERTILIZER_TO_WATER_END = 105;

const WATER_TO_LIGHT_START = FERTILIZER_TO_WATER_END + 2;
const WATER_TO_LIGHT_END = 145;

const LIGHT_TO_TEMP_START = WATER_TO_LIGHT_END + 2;
const LIGHT_TO_TEMP_END = 194;

const TEMP_TO_HUMIDITY_START = LIGHT_TO_TEMP_END + 2;
const TEMP_TO_HUMIDITY_END = 236;

const HUMIDITY_TO_LOCATION_START = TEMP_TO_HUMIDITY_END + 2;
const HUMIDITY_TO_LOCATION_END = 261;

const DESTINATION_INDEX = 0;
const SOURCE_INDEX = 1;
const RANGE_INDEX = 2;

const seedsData = data[0]
    .split(":")[1]
    .trim()
    .match(/(\d+ \d+)/g);

const seedRanges = seedsData.map((data) => {
    const [source, range] = data.split(" ");
    return { source: parseInt(source), range: parseInt(range) };
});

const createRanges = (start, end) => {
    const ranges = [];

    for (let i = start; i < end; i++) {
        const line = data[i].split(" ");
        const destination = parseInt(line[DESTINATION_INDEX]);
        const source = parseInt(line[SOURCE_INDEX]);
        const range = parseInt(line[RANGE_INDEX]);

        ranges.push({
            destination,
            source,
            range,
        });
    }

    return ranges;
};

const seedsToSoilRanges = createRanges(SEEDS_TO_SOIL_START, SEEDS_TO_SOIL_END);
const soilToFertilizerRanges = createRanges(SOIL_TO_FERTILIZER_START, SOIL_TO_FERTILIZER_END);
const fertilizerToWaterRanges = createRanges(FERTILIZER_TO_WATER_START, FERTILIZER_TO_WATER_END);
const waterToLightRanges = createRanges(WATER_TO_LIGHT_START, WATER_TO_LIGHT_END);
const lightToTempRanges = createRanges(LIGHT_TO_TEMP_START, LIGHT_TO_TEMP_END);
const tempToHumidityRanges = createRanges(TEMP_TO_HUMIDITY_START, TEMP_TO_HUMIDITY_END);
const humidityToLocationRanges = createRanges(HUMIDITY_TO_LOCATION_START, HUMIDITY_TO_LOCATION_END);

const getSource = (ranges, destination) => {
    const range = ranges.find((range) => range.destination <= destination && destination < range.destination + range.range);

    if (range === undefined) {
        return destination;
    }
    return range.source + (destination - range.destination);
};

const sources = [
    humidityToLocationRanges,
    tempToHumidityRanges,
    lightToTempRanges,
    waterToLightRanges,
    fertilizerToWaterRanges,
    soilToFertilizerRanges,
    seedsToSoilRanges,
];

const findSeed = (location) => {
    let previousValue = location;
    for (source of sources) {
        const value = getSource(source, previousValue);
        if (value === undefined) {
            return undefined;
        }
        previousValue = value;
    }

    return previousValue;
};

const findLocation = (start, end, inc) => {
    let output = 0;
    for (let i = start; i < end; i += inc) {
        const seed = findSeed(i);

        if (seedRanges.find((range) => range.source <= seed && seed < range.source + range.range)) {
            console.log(`Found seed: ${seed}`);
            output = i;
            break;
        }
    }

    return output;
};

humidityToLocationRanges.sort((a, b) => a.destination - b.destination);

const INCREMENT = 10_000_000;

const maxLocationWithSeed = findLocation(
    0,
    humidityToLocationRanges[humidityToLocationRanges.length - 1].destination +
        humidityToLocationRanges[humidityToLocationRanges.length - 1].range,
    INCREMENT
);

const minLocation = findLocation(
    maxLocationWithSeed - INCREMENT > 0 ? maxLocationWithSeed - INCREMENT : 0,
    maxLocationWithSeed + INCREMENT,
    1
);

console.log(minLocation);
