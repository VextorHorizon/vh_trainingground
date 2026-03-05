let word:string = "Hello World"

console.log(word);

import { getApiVersion } from './getApiVersion.js';

const currentVersion = getApiVersion();
console.log(`System Status: Online | Version: ${currentVersion}`)
