import axios from "axios";

const baseURL = process.env.NEXT_PUBLIC_BACKEND_URL ?? "http://localhost:8080";

export const api = axios.create({
  baseURL,
  timeout: 5000,
});
