import http from 'k6/http';
import { sleep } from 'k6';

export const options = {
    duration: '1m',
    vus: 150,
    insecureSkipTLSVerify: false,
};

export default function () {
    // Status and heartbeat
    http.get('https://localhost/api/status');
    http.get('https://localhost/api/healthcheck');

    // Updaters
    http.get('https://localhost/api/updater');
    http.post('https://localhost/api/updater', {});
    http.post('https://localhost/api/updater', {
        database: "k6",
        schema: {
            my: "schema"
        },
        interval: 60,
        source: "https://ipinfo.io/json",
        method: "GET",
        timeout: 30
    });

    // Data
    http.post('https://localhost/api/data', {database: "k6"})
    http.post('https://localhost/api/data', {database: "k6", quantity: 50})

    // 404
    http.get('https://localhost/api/yikes');

    sleep(1);
}