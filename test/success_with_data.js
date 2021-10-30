import http from 'k6/http';
import { check, sleep } from 'k6';

const BASE_URL = "http://localhost:8080"

export const options = {
    vus: 3,
    duration: '10s',
    // httpDebug: 'full',

    thresholds: {
        http_req_duration: ['p(99)<1500'], // 99% of requests must complete below 1.5s
    },
};

const payload = JSON.stringify({
    repo: "https://github.com/brandoncate-personal/blog-content",
});

const params = {
    headers: {
        'Content-Type': 'application/json',
    },
};
const expected = [{
    "path": "src/test1.md",
    "title": "I confess for I have sinned Updated 3"
},
{
    "path": "src/test2.md",
    "title": "Work To Do"
}]

export default function () {
    const res = http.post(`${BASE_URL}/extract`, payload, params);

    check(res, {
        'success unordered data returned': (resp) =>
            JSON.stringify(resp.json('data').map(a => a.title).sort()) == JSON.stringify(expected.map(a => a.title).sort())
    });

    sleep(1);
}