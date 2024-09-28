// src/services/paymentService.js
import axios from 'axios';

const createPayment = async (amount) => {
  try {
    const response = await axios.post('/api/pay', { amount });
    return response.data;
  } catch (error) {
    console.error("Error creating payment:", error);
    throw error; // Lanza el error para que lo manejes en el componente
  }
};

export const paymentService = {
  createPayment
};
