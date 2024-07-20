import axios from 'axios';

export const sessionService = {
  async getSession(user) {
    const url = '/api/login';
    return await axios.post(url, user);
  }
};
