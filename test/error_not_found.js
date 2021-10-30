import http from 'k6/http';
import { check, sleep } from 'k6';

const BASE_URL = "http://localhost:8080"

export const options = {
    vus: 3,
    duration: '2s',
    httpDebug: 'full',

    thresholds: {
        http_req_duration: ['p(99)<1500'], // 99% of requests must complete below 1.5s
    },
};

const payload = JSON.stringify({
    repo: "https://github.com/brandoncate-personal/does-not-exist",
});

const params = {
    headers: {
        'Content-Type': 'application/json',
    },
};

export default function () {
    const res = http.post(`${BASE_URL}/extract`, payload, params);

    check(res, {
        'error no data returned': (resp) => resp.json('data') === undefined,
    });

    check(res, {
        'error message returned': (resp) => resp.json('error') !== undefined,
    });

    sleep(1);
}