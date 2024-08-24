<template>
  <div class="invoice-container">
    <div class="invoice-header">
      <h2>Area de Administrador / Empresas</h2>
    </div>
    <div class="invoice-buttons">
      <button class="button-nuevoPrestadorSP" @click="companyForm">+ Nuevo Prestador de Servicio Publico</button>
    </div>
    <table>
      <thead>
        <tr>
          <th>NIT</th>
          <th>Nombre de la Empresa</th>
          <th>Localidad</th>
          <th>Teléfono</th>
          <th>Email</th>
          <th>Datos</th>
        </tr>
      </thead>
      <tbody>
        <tr v-for="company in companies" :key="company.IDcompany">
          <td>{{ company.nit }}</td>
          <td>{{ company.Name }}</td>
          <td>{{ company.address }}</td>
          <td>{{ company.Phone }}</td>
          <td>{{ company.Email }}</td>
        </tr>
      </tbody>
    </table>
  </div>
</template>

<script>
import axios from 'axios';

export default {
  data() {
    return {
      companies: []
    };
  },
  created() {
    this.fetchCompanies();
  },
  methods: {
    async fetchCompanies() {
      try {
        const token = localStorage.getItem('token');
        const response = await axios.get('/api/AllCompany', {  // Ruta corregida
          headers: {
            Authorization: `Bearer ${token}`
          }
        });
        console.log(response.data); // Para verificar los datos recibidos
        this.companies = response.data;
      } catch (error) {
        console.error('Error fetching companies:', error);
      }      
    },
    companyForm() {
      // Redireccionar a la página de registro de empresa
      this.$router.push('/api/register');
    }
  }
};
</script>



<style>

.invoice-container {
  background-color: #f9f9f9;
  padding: 20px;
}

.invoice-header {
  background-color: #b4d8f1;
  padding: 15px;
  border-radius: 5px;
  margin-bottom: 20px;
}

.invoice-buttons {
  display: flex;
  gap: 10px;
  margin-bottom: 20px;
}

.btn {
  padding: 10px 15px;
  border: none;
  background-color: #3498db;
  color: white;
  border-radius: 4px;
  cursor: pointer;
}

.btn:hover {
  background-color: #2980b9;
}

.container {
  display: flex;
  flex-direction: column;
  align-items: center;
}

.buttons {
  width: 80%;
  display: flex;
  justify-content: flex-start;
  margin-bottom: 20px;
}

.buttons button {
  background-color: blue;
  color: white;
  border: none;
  padding: 10px 20px;
  margin-right: 10px;
  cursor: pointer;
}

table {
  width: 80%;
  border-collapse: collapse;
  margin: 0 auto;
}

th, td {
  border: 1px solid #ddd;
  padding: 8px;
  text-align: center;
}

th {
  background-color: #f2f2f2;
}
</style>
