<template>
  <div class="register-company-container">
    <h2>Registro de Empresa</h2>
    <form @submit.prevent="submitForm">
      <!-- Datos de la Empresa -->
      <div class="form-group">
        <label for="nit">NIT:</label>
        <input type="text" id="nit" v-model="nit" placeholder="Ingrese el NIT de la empresa" required />
      </div>
      <div class="form-group">
        <label for="name">Nombre de la Empresa:</label>
        <input type="text" id="name" v-model="name" placeholder="Ingrese el nombre de la empresa" required />
      </div>
      <div class="form-group">
        <label for="address">Dirección:</label>
        <input type="text" id="address" v-model="address" placeholder="Ingrese la dirección de la empresa" required />
      </div>
      <div class="form-group">
        <label for="phone">Teléfono:</label>
        <input type="text" id="phone" v-model="phone" placeholder="Ingrese el teléfono de la empresa" required />
      </div>

      <!-- Datos de Usuario -->
      <div class="form-group">
        <label for="email">Correo:</label>
        <input type="email" id="email" v-model="email" placeholder="Ingrese el correo electrónico" required />
      </div>
      <div class="form-group">
        <label for="password">Contraseña:</label>
        <input type="password" id="password" v-model="password" placeholder="Ingrese una contraseña segura" required />
      </div>

      <!-- Botón Registrar -->
      <div class="form-group">
        <button type="submit" class="submit-btn">Registrar</button>
      </div>
      
      <!-- Botón Regresar -->
      <div class="form-group">
        <button type="button" @click="goBack" class="back-btn">Regresar</button>
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
        alert('Hubo un error, verifica tus datos');
      }
    },
    goBack() {
      this.$router.go(-1);
    }
  }
};
</script>

<style scoped>
.register-company-container {
  padding: 40px;
  max-width: 600px;
  margin: 0 auto;
  background-color: #ffffff;
  border-radius: 12px;
  box-shadow: 0 4px 20px rgba(0, 0, 0, 0.1);
}

h2 {
  text-align: center;
  color: #333;
  margin-bottom: 20px;
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

.submit-btn,
.back-btn {
  width: 100%;
  padding: 15px;
  font-size: 16px;
  border-radius: 6px;
  border: none;
  cursor: pointer;
  transition: background-color 0.3s;
}

.submit-btn {
  background-color: #3498db;
  color: white;
}

.submit-btn:hover {
  background-color: #2980b9;
}

.back-btn {
  background-color: #cccccc;
}

.back-btn:hover {
  background-color: #b3b3b3;
}


</style>
