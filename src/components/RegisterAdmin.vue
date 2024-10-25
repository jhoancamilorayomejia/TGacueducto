<template>
    <div class="register-container">
      <div class="form-card">
        <h2>Registro de Administrador</h2>
        <form @submit.prevent="submitForm">
          <div class="form-group">
            <label for="nombre">Nombre:</label>
            <input type="text" id="nombre" v-model="nombre" required />
          </div>
          <div class="form-group">
            <label for="apellido">Apellido:</label>
            <input type="text" id="apellido" v-model="apellido" required />
          </div>
          <div class="form-group">
            <label for="email">Correo:</label>
            <input type="email" id="email" v-model="email" required />
          </div>
          <div class="form-group">
            <label for="password">Contraseña:</label>
            <input type="password" id="password" v-model="password" required />
          </div >
          <div class="form-group">
          <button type="submit" class="submit-btn">Registrar</button>
        </div>
        <div class="form-group">
        <button type="button" @click="goBack" class="back-btn">Regresar</button>
      </div>
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
          alert('Hubo un error, verifica tus datos');
        }
      },
      goBack() {
      this.$router.go(-1);
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
  .register-container {
    display: flex;
    justify-content: center;
    align-items: center;
    height: 100vh;
    background-color: #f0f4f7;
  }
  
  .form-card {
    background-color: #ffffff;
    padding: 40px;
    border-radius: 10px;
    box-shadow: 0 4px 20px rgba(0, 0, 0, 0.1);
    width: 400px;
  }
  
  h2 {
    text-align: center;
    color: #333;
    margin-bottom: 30px;
  }
  
  .form-group {
    margin-bottom: 20px;
  }
  
  label {
    display: block;
    margin-bottom: 8px;
    font-weight: bold;
    color: #555;
  }
  
  input {
    width: 100%;
    padding: 12px;
    border: 1px solid #ddd;
    border-radius: 6px;
    font-size: 16px;
    background-color: #f9f9f9;
    transition: border-color 0.3s;
  }
  
  input:focus {
    border-color: #3498db;
    outline: none;
  }
  
  .submit-btn {
    width: 100%;
    padding: 15px;
    background-color: #3498db;
    color: white;
    border: none;
    border-radius: 6px;
    font-size: 18px;
    cursor: pointer;
    transition: background-color 0.3s;
  }
  
  .submit-btn:hover {
    background-color: #2980b9;
  }
  
  .back-btn {
  width: 100%; /* Asegura que el botón ocupe todo el ancho */
  padding: 15px; /* Añade padding similar al de submit-btn */
  background-color: #cccccc; /* Color de fondo neutro */
  color: white; /* Texto en color blanco para mejor contraste */
  border: none;
  border-radius: 6px; /* Bordes redondeados */
  font-size: 18px; /* Mismo tamaño de fuente que submit-btn */
  cursor: pointer;
  transition: background-color 0.3s; /* Transición suave */
}

.back-btn:hover {
  background-color: #b3b3b3; /* Cambio de color en hover */
}

  
  </style>