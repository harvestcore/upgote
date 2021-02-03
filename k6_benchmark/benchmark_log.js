import http from 'k6/http';
import { sleep } from 'k6';

export const options = {
    duration: '1m',
    vus: 150,
    insecureSkipTLSVerify: false,
};

export default function () {  
    http.post('https://localhost/api/log', {quantity: 1});
    http.get('https://localhost/api/log');

    http.post('https://localhost/api/data', {database: "log"})
    http.post('https://localhost/api/data', {database: "log", quantity: 50})
    sleep(1);
}