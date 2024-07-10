import axios from 'axios';

const instance = axios.create({
  baseURL: '/api', // Esto hará que las solicitudes a /admins se redirijan a http://localhost:8080/api/admins
  timeout: 1000,
});

export default instance;
