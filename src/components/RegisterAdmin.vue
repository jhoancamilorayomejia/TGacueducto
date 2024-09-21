<template>
    <div class="register-user-container">
      <div class="step-1">
        <h2>Registro de Administrador</h2>
        <form @submit.prevent="submitForm">
          <div>
            <label for="nombre">Nombre:</label>
            <input type="text" id="nombre" v-model="nombre" required />
          </div>
          <div>
            <label for="apellido">Apellido:</label>
            <input type="text" id="apellido" v-model="apellido" required />
          </div>
          <div>
            <label for="email">Correo:</label>
            <input type="email" id="email" v-model="email" required />
          </div>
          <div>
            <label for="password">Contraseña:</label>
            <input type="password" id="password" v-model="password" required />
          </div>
          <button type="submit">Registrar</button>
        </form>
      </div>
    </div>
  </template>
  
  <script>
  import axios from 'axios';
  
  export default {
    data() {
      return {
        nombre: '',
        apellido: '',
        email: '',
        password: ''
      };
    },
    methods: {
      async submitForm() {
        try {
          const registrationData = {
            nombre: this.nombre,
            apellido: this.apellido,
            email: this.email,
            password: this.password // Enviar la contraseña en texto claro
          };
  
          const response = await axios.post('/api/register-user', registrationData);
          alert(`Administrador registrado con éxito. Secret Key: ${response.data.secret_key}`);
          this.resetForm();
        } catch (error) {
          console.error('Error registrando el administrador:', error);
          alert('Hubo un error registrando el administrador');
        }
      },
      resetForm() {
        this.nombre = '';
        this.apellido = '';
        this.email = '';
        this.password = '';
      }
    }
  };
  </script>
  
  <style scoped>
  .register-user-container {
    display: flex;
    justify-content: center;
    padding: 20px;
  }
  
  .step-1 {
    width: 50%;
    background-color: #f9f9f9;
    padding: 20px;
    border-radius: 8px;
    box-shadow: 0 0 10px rgba(0, 0, 0, 0.1);
  }
  
  form div {
    margin-bottom: 15px;
  }
  
  label {
    display: block;
    margin-bottom: 5px;
    font-weight: bold;
  }
  
  input {
    width: 100%;
    padding: 8px;
    border: 1px solid #ccc;
    border-radius: 4px;
  }
  
  button {
    padding: 10px 20px;
    border: none;
    background-color: #3498db;
    color: white;
    border-radius: 4px;
    cursor: pointer;
  }
  
  button:hover {
    background-color: #2980b9;
  }
  </style>