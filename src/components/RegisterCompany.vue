<template>
  <div class="register-company-container">
    <h2>Registro de Empresa</h2>
    <form @submit.prevent="submitForm">
      <!-- Datos de la Empresa -->
      <div>
        <label for="nit">NIT:</label>
        <input type="text" id="nit" v-model="nit" required>
      </div>
      <div>
        <label for="Nombre">Nombre de la Empresa:</label>
        <input type="text" id="name" v-model="name" required>
      </div>
      <div>
        <label for="address">Dirección:</label>
        <input type="text" id="address" v-model="address" required>
      </div>
      <div>
        <label for="phone">Teléfono:</label>
        <input type="text" id="phone" v-model="phone" required>
      </div>

      <!-- Datos de Usuario -->
      <div>
        <label for="email">Correo:</label>
        <input type="email" id="email" v-model="email" required>
      </div>
      <div>
        <label for="password">Contraseña:</label>
        <input type="password" id="password" v-model="password" required>
      </div>

      <div>
        <button type="submit">Registrar</button>
      </div>
      
      <!-- Botón Regresar -->
      <div style="margin-top: 20px;"> <!-- Espacio entre botones -->
        <button type="button" @click="goBack">Regresar</button>
      </div>
      
    </form>
  </div>
</template>

<script>
import axios from 'axios';

export default {
  data() {
    return {
      nit: '',
      name: '',
      address: '',
      phone: '',
      email: '',
      password: ''
    };
  },
  methods: {
    async submitForm() {
      try {
        const registrationData = {
          nit: this.nit,
          name: this.name,
          address: this.address,
          phone: this.phone,
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
    },
    goBack() {
      this.$router.go(-1); // Regresa a la página anterior (cualquiera)
    }
  }
};
</script>

<style scoped>
.register-company-container {
  padding: 20px;
  background-color: #f9f9f9;
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
