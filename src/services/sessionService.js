// src/services/sessionService.js
import axios from 'axios';

const getSession = async (user) => {
  return axios.post('/api/login', user);
};

export const sessionService = {
  getSession
};
