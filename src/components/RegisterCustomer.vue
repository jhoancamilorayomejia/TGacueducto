<template>
    <div class="register-user-container">
      <!-- Paso 1: Datos personales -->
      <div class="step-1" v-if="step === 1">
        <h2>Registro de Usuario - Paso 1</h2>
        <form @submit.prevent="nextStep">
          <div>
            <label for="entity">Entidad:</label>
            <select id="entity" v-model="selectedEntity" class="entity-select" required>
              <option value="" disabled>Seleccione una entidad</option>
              <option v-for="company in companies" :key="company.idcompany" :value="company.idcompany">
                {{ company.name }}
              </option>
            </select>
          </div>
          <div>
            <label for="name">Nombre:</label>
            <input type="text" id="name" v-model="name" required>
          </div>
          <div>
            <label for="lastName">Apellidos:</label>
            <input type="text" id="lastName" v-model="lastName" required>
          </div>
          <div>
            <label for="address">Dirección:</label>
            <input type="text" id="address" v-model="address" required>
          </div>
          <div>
            <label for="phone">Teléfono:</label>
            <input type="text" id="phone" v-model="phone" required>
          </div>
          <button type="submit">Siguiente</button>
        </form>
      </div>
      <!-- Paso 2: Correo y contraseña -->
      <div class="step-2" v-if="step === 2">
        <h2>Registro de Usuario - Paso 2</h2>
        <form @submit.prevent="submitForm">
          <div>
            <label for="entity">Entidad:</label>
            <input type="text" id="entity" :value="selectedEntityName" disabled>
          </div>
          <div>
            <label for="name">Nombre:</label>
            <input type="text" id="name" :value="name" disabled>
          </div>
          <div>
            <label for="lastName">Apellidos:</label>
            <input type="text" id="lastName" :value="lastName" disabled>
          </div>
          <div>
            <label for="address">Dirección:</label>
            <input type="text" id="address" :value="address" disabled>
          </div>
          <div>
            <label for="phone">Teléfono:</label>
            <input type="text" id="phone" :value="phone" disabled>
          </div>
          <div>
            <label for="email">Correo electrónico:</label>
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
        companies: [],
        selectedEntity: '',
        name: '',
        lastName: '',
        address: '',
        phone: '',
        email: '',
        password: '',
        confirmPassword: ''
      };
    },
    computed: {
      selectedEntityName() {
        const company = this.companies.find(c => c.idcompany === this.selectedEntity);
        return company ? company.name : '';
      }
    },
    methods: {
      async fetchCompanies() {
        try {
          const response = await axios.get('http://localhost:8080/api/get-switch');
          this.companies = response.data;
        } catch (error) {
          console.error("Error fetching companies:", error);
        }
      },
      nextStep() {
        if (this.selectedEntity && this.name && this.lastName && this.address && this.phone) {
          this.step = 2;
        } else {
          alert("Por favor, completa todos los campos.");
        }
      },
      async submitForm() {
        if (this.password !== this.confirmPassword) {
          alert("Las contraseñas no coinciden.");
          return;
        }
  
        try {
          const registrationData = {
            idcompany: this.selectedEntity,
            name: this.name,
            last_name: this.lastName,
            address: this.address,
            phone: this.phone,
            email: this.email,
            password: this.password
          };
          await axios.post('http://localhost:8080/api/registerCustomer', registrationData);
          alert('Usuario registrado con éxito');
          // Redirigir o limpiar el formulario según sea necesario
        } catch (error) {
          console.error('Error registrando el usuario:', error);
          alert('Hubo un error registrando el usuario.');
        }
      }
    },
    mounted() {
      this.fetchCompanies();
    }
  };
  </script>
  
  <style scoped>
  .register-user-container {
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
  
  select.entity-select {
    color: #3498db; /* Color azul para el texto */
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
  
  