const fs = require('fs').promises; // 使用promises版本的fs模块

let deviceIds = []; // 全局变量设为空数组

// 读取txt文件内容并将数据保存在deviceIds数组中
fs.readFile('device_ids.txt', 'utf8')
    .then(data => {
        deviceIds = data.split(','); // 读取到数据后赋值给deviceIds
    })
    .catch(error => {
        console.error('Error reading file:', error);
    });

const dataCache = {};
let i = 0;

const generateStatusData = (faker) => {
    const randomElement = (arr) => arr[Math.floor(Math.random() * arr.length)];
    const baseData = {
        PowerState: randomElement(['ON', 'OFF']),
        OperationalStatus: randomElement(['Running','Stopped', 'Error']),
        Mode: randomElement(['Auto','Manual'])
    };
    return baseData;
};

const generatePerformanceMetricsData = (faker) => {
    const baseData = {
        temperature: faker.datatype.number({ min: 0, max: 300 }),
        pressure: faker.datatype.number({ min: 100, max: 200 }),
        speed: faker.datatype.number({ min: 0, max: 2000 }),
        output: faker.datatype.number({ min: 0, max: 500 }),
    };
    return baseData;
};

const generator = (faker, options) => {
    const { clientId } = options;
    if (!dataCache[clientId]) {
        // 如果没有缓存数据，将设备ID从数组中取出，并放入缓存
        dataCache[clientId] = {
            deviceId: deviceIds[i],
        };
        i = (i + 1) % deviceIds.length; // 循环使用deviceIds数组中的值
    }
    let tmLoc = new Date();
    let now = tmLoc.getTime() + tmLoc.getTimezoneOffset() * 60000;

    const data = {
        ...dataCache[clientId],
        status: generateStatusData(faker),
        performanceMetrics: generatePerformanceMetricsData(faker),
        timeStamp: Math.floor(now/ 1000),
    };
    return {
        message: JSON.stringify(data),
    };
};

module.exports = {
    generator
};
