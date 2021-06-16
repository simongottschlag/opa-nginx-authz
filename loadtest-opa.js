import http from "k6/http";
import { check } from "k6";
import { Rate } from "k6/metrics";

var failureRate = new Rate("check_failure_rate");

export let options = {
    vus: 10,
    stages: [
        { duration: "5s", target: 20 },
        { duration: "5s", target: 30 },
        { duration: "5s", target: 20 },
        { duration: "5s", target: 10 },
        { duration: "5s", target: 0 },
    ]
};

export default function () {
    const input = {
        input: {
            headers: {
                Authorization: ["Bearer test"]
            }
        }
    }
    const opaInput = JSON.stringify(input)
    let response = http.post("http://localhost:8181/v1/data/nginx/authz", opaInput)

    let checkRes = check(response, {
        "status is 200": (r) => r.status === 200,
        "result is true": (r) => r.json("result") === true
    });

    failureRate.add(!checkRes);
};