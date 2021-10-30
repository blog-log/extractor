import http from 'k6/http';
import { check, sleep } from 'k6';

const BASE_URL = "http://localhost:8080"

export const options = {
    vus: 3,
    duration: '10s',
    httpDebug: 'full',

    thresholds: {
        http_req_duration: ['p(99)<1500'], // 99% of requests must complete below 1.5s
    },
};

const payload = JSON.stringify({
    repo: "https://github.com/brandoncate-personal/blog-user",
});

const params = {
    headers: {
        'Content-Type': 'application/json',
    },
};

export default function () {
    const res = http.post(`${BASE_URL}/extract`, payload, params);

    check(res, {
        'success no data returned': (resp) => resp.json('data') === null,
    });

    check(res, {
        'success warning returned': (resp) => resp.json('warning') !== undefined,
    });

    sleep(1);
}