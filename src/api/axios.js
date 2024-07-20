import axios from 'axios';

const instance = axios.create({
  baseURL: '/api', // Esto hará que las solicitudes a /admins se redirijan a http://localhost:8080/api/admins
  timeout: 1000,
});

export default instance;



/*
import axios from 'axios';

const instance = axios.create({
  baseURL: 'http://localhost:8000/api', // Asegúrate de que esta URL apunte a tu backend
  timeout: 1000,
  headers: {
    'Content-Type': 'application/json'
  }
});

export default instance; */