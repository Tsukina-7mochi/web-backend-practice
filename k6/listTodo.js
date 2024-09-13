import http from "k6/http";
import { sleep } from "k6";

export const options = {
  stages: [
    { duration: "10s", target: 100 },
    { duration: "10s", target: 200 },
    { duration: "10s", target: 0 },
  ],
};

const numUsers = 10;
const numTodos = 10;

export function setup() {
  const userIds = [];
  for (let i = 0; i < numUsers; i++) {
    const res = http.post(`http://localhost:3000/api/users?name=user${i}`);
    const userId = res.json().id;

    if (typeof userId !== "number") {
      throw new Error("userId is not a number");
    }

    for (let i = 0; i < numTodos; i++) {
      http.post(`http://localhost:3000/api/users/${userId}/todos?title=aaa`);
    }

    userIds.push(userId);
  }

  return { userIds };
}

export default function ({ userIds }) {
  const userId = userIds[Math.floor(Math.random() * userIds.length)];

  http.get(`http://localhost:3000/api/users/${userId}/todos`);

  sleep(1);
}
