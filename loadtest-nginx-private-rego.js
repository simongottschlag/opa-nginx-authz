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
    const params = {
        headers: {
            "Authorization": "Bearer test",
        },
    };

    let response = http.get("http://localhost:8080/private/rego", params)

    let checkRes = check(response, {
        "status is 200": (r) => r.status === 200
    });

    failureRate.add(!checkRes);
};