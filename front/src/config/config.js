import Env from './env';

let config = {
    env: Env
};
const proxy = require('http-proxy-middleware');
export default config;