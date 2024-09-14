<template>
    <div class="register-company-container">
      <div class="step-1">
        <h2>Registro de Empresa - Paso 1</h2>
        <form @submit.prevent="nextStep">
          <div>
            <label for="companyName">Nombre de la empresa:</label>
            <input type="text" id="companyName" v-model="companyName" required>
          </div>
          <div>
            <label for="nit">NIT:</label>
            <input type="text" id="nit" v-model="nit" required>
          </div>
          <div>
            <label for="city">Ciudad:</label>
            <select id="city" v-model="city" required>
              <option value="" disabled>Seleccione una ciudad</option>
              <option value="Cali">Cali</option>
              <option value="Palmira">Palmira</option>
              <option value="Buenaventura">Buenaventura</option>
              <!-- Agregar más ciudades del Valle del Cauca aquí -->
            </select>
          </div>
          <div>
            <label for="address">Dirección:</label>
            <input type="text" id="address" v-model="address" required>
          </div>
          <div>
            <label for="postalCode">Código postal:</label>
            <input type="text" id="postalCode" v-model="postalCode" required>
          </div>
          <button type="submit">Siguiente</button>
        </form>
      </div>
      <div class="step-2" v-if="step === 2">
        <h2>Registro de Empresa - Paso 2</h2>
        <form @submit.prevent="submitForm">
          <div>
            <label for="companyName">Nombre de la empresa:</label>
            <input type="text" id="companyName" :value="companyName" disabled>
          </div>
          <div>
            <label for="nit">NIT:</label>
            <input type="text" id="nit" :value="nit" disabled>
          </div>
          <div>
            <label for="city">Ciudad:</label>
            <input type="text" id="city" :value="city" disabled>
          </div>
          <div>
            <label for="address">Dirección:</label>
            <input type="text" id="address" :value="address" disabled>
          </div>
          <div>
            <label for="postalCode">Código postal:</label>
            <input type="text" id="postalCode" :value="postalCode" disabled>
          </div>
          <div>
            <label for="email">Correo:</label>
            <input type="email" id="email" v-model="email" required>
          </div>
          <div>
            <label for="password">Contraseña:</label>
            <input type="password" id="password" v-model="password" required>
          </div>
          <div>
          <label for="confirmPassword">Confirmar Contraseña:</label>
          <input type="password" id="confirmPassword" v-model="confirmPassword" required>
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
        step: 1,
        companyName: '',
        nit: '',
        city: '',
        address: '',
        postalCode: '',
        email: '',
        password: ''
      };
    },
    methods: {
      nextStep() {
        if (this.companyName && this.nit && this.city && this.address && this.postalCode) {
          this.step = 2;
        }
      },
      async submitForm() {
        try {
          const registrationData = {
            companyName: this.companyName,
            nit: this.nit,
            city: this.city,
            address: this.address,
            postalCode: this.postalCode,
            email: this.email,
            password: this.password
          };
          await axios.post('/api/register-company', registrationData);
          alert('Empresa registrada con éxito');
          // Redirigir o limpiar el formulario según sea necesario
        } catch (error) {
          console.error('Error registrando la empresa:', error);
          alert('Hubo un error registrando la empresa');
        }
      }
    }
  };
  </script>
  
  <style scoped>
  .register-company-container {
    display: flex;
    justify-content: space-between;
    padding: 20px;
  }
  
  .step-1, .step-2 {
    width: 48%;
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
  
  input, select {
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
  