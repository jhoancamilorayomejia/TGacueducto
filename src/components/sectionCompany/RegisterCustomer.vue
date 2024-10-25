<template>
  <div class="register-user-container">
    <h2>Registro de Cliente {{ userID }}</h2>      
    <form @submit.prevent="submitForm">
      <div>
        <label for="entity">Entidad:</label>
        <!-- Mostrar el nombre de la entidad en un campo de texto bloqueado -->
        <input
          type="text"
          id="entity"
          :value="entityName"
          disabled
        />
      </div>
      <div>
        <label for="cedula">Cédula:</label>
        <input type="text" id="cedula" v-model="cedula" required>
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
      <div>
        <label for="email">Correo electrónico:</label>
        <input type="email" id="email" v-model="email" required>
      </div>
      <div>
        <label for="password">Contraseña:</label>
        <input type="password" id="password" v-model="password" required>
      </div>
      <div class="form-group">    
      <button type="submit" >Registrar</button>
    </div>
    <div class="form-group"> 
      <button class="back-button" type="button" @click="goBack">Regresar</button>
    </div>
    </form>
  </div>
</template>

<script>
import axios from 'axios';

export default {
  data() {
    return {
      companies: [],
      selectedEntity: '', // Para almacenar el ID de la entidad seleccionada
      entityName: '', // Para almacenar el nombre de la entidad
      cedula: '', 
      name: '',
      lastName: '',
      address: '',
      phone: '',
      email: '',
      password: '',
      userID: '', // ID del usuario
    };
  },
  created() {
    this.userID = localStorage.getItem('userID');
  },
  methods: {
    async fetchCompanies() {
      try {
        const response = await axios.get('http://localhost:8080/api/get-switch');
        this.companies = response.data;

        // Buscar la compañía cuyo idcompany coincide con userID
        const matchingCompany = this.companies.find(company => company.idcompany == this.userID);
        
        if (matchingCompany) {
          this.selectedEntity = matchingCompany.idcompany;
          this.entityName = matchingCompany.name; // Guardar el nombre de la entidad
        }
      } catch (error) {
        console.error("Error fetching companies:", error);
      }
    },
    async submitForm() {
      try {
        const registrationData = {
          idcompany: this.selectedEntity, // Enviar el ID de la entidad seleccionada
          cedula: this.cedula,
          name: this.name,
          last_name: this.lastName,
          address: this.address,
          phone: this.phone,
          email: this.email,
          password: this.password
        };

        await axios.post('http://localhost:8080/api/registerCustomer', registrationData);
        alert('Usuario registrado con éxito');
        this.resetForm();
      } catch (error) {
        console.error('Error registrando el usuario:', error);
        alert('Hubo un error, verifica tus datos.');
      }
    },
    goBack() {
      this.$router.go(-1); // Regresa a la página anterior (cualquiera)
    },
    resetForm() {
      this.selectedEntity = '';
      this.entityName = ''; // Limpiar el campo de entidad
      this.cedula = '';
      this.name = '';
      this.lastName = '';
      this.address = '';
      this.phone = '';
      this.email = '';
      this.password = '';
    }
  },
  mounted() {
    this.fetchCompanies();
  }
};
</script>

  
  <style scoped>
  .register-user-container {
    width: 100%;
    max-width: 600px;
    margin: 0 auto;
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
    padding: 10px 280px;
    border: none;
    background-color: #3498db;
    color: white;
    border-radius: 4px;
    cursor: pointer;
  }
  
  button:hover {
    background-color: #2980b9;
  }

  .back-button {
  background-color: #837d7d; /* Color azul */
  color: white;
}


.back-button:hover {
  background-color: #b3b3b3; /* Color azul más oscuro al pasar el ratón */
}
.form-group {
  margin-bottom: 20px;
}
  </style>
  
  
  