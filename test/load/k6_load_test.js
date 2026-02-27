// Load testing script for the Go backend API using k6
import http from 'k6/http';
import { check, sleep } from 'k6';

export const options = {
    vus: 10, // Number of virtual users
    duration: '30s', // Duration of the test
};

export default function () {
    // User registration
    let registerResponse = http.post('http://localhost:8081/api/v1/users', JSON.stringify({
        username: 'test_user',
        email: 'test@example.com',
        password_hash: 'password123',
        role: 'user'
    }), {
        headers: { 'Content-Type': 'application/json' },
    });

    check(registerResponse, {
        'registration status is 201': (r) => r.status === 201,
    });

    // User login
    let loginResponse = http.post('http://localhost:8081/api/v1/login', JSON.stringify({
        username: 'test_user',
        password: 'password123',
    }), {
        headers: { 'Content-Type': 'application/json' },
    });

    check(loginResponse, {
        'login status is 200': (r) => r.status === 200,
    });

    // Money transfer
    let transferResponse = http.post('http://localhost:8081/api/v1/transactions', JSON.stringify({
        from_user_id: 1,
        to_user_id: 2,
        amount: 100.50,
        description: 'Payment for services',
    }), {
        headers: { 'Content-Type': 'application/json' },
    });

    check(transferResponse, {
        'transfer status is 201': (r) => r.status === 201,
    });

    sleep(1); // Sleep for 1 second between iterations
}