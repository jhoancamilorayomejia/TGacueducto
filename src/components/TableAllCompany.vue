<template>
  <div class="invoice-container">
    <div class="invoice-header">
      <h2>Área de Administrador / Empresas</h2>
    </div>
    <div class="invoice-buttons">
      <button class="button-nuevoPrestadorSP" @click="companyForm">+ Nuevo Prestador de Servicio Público</button>
    </div>
    <table>
      <thead>
        <tr>
          <th>NIT</th>
          <th>Nombre de la Entidad</th>
          <th>Localidad</th>
          <th>Teléfono</th>
          <th>Email</th>
          <th>Acciones</th> <!-- Nueva columna para las acciones -->
        </tr>
      </thead>
      <tbody>
        <tr v-for="company in companies" :key="company.IDcompany">
          <td>{{ company.nit }}</td>
          <td>{{ company.Name }}</td>
          <td>{{ company.address }}</td>
          <td>{{ company.Phone }}</td>
          <td>{{ company.Email }}</td>
          <td>
            <button class="btn-edit" @click="editCompany(company.idcompany)">Modificar</button>
          </td>
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
        const response = await axios.get('/api/AllCompany', {
          headers: {
            Authorization: `Bearer ${token}`
          }
        });
        this.companies = response.data;
      } catch (error) {
        console.error('Error fetching companies:', error);
      }
    },
    companyForm() {
      this.$router.push('/api/register');
    },
    editCompany(idcompany) {
      this.$router.push(`/api/company/edit/${idcompany}`);
    }
  }
};
</script>

<style scoped>
body {
  background-color: #ffffff;
  margin: 0;
  padding: 0;
  font-family: 'Roboto', sans-serif;
  font-size: 11px;
}

.invoice-container {
  background-color: #f3f3f3;
  padding: 20px;
  margin: 50px auto;
  width: 80%;
  border-radius: 10px;
  box-shadow: 0px 4px 8px rgba(0, 0, 0, 0.1);
}

.invoice-header {
  background-color: #62b5ec;
  padding: 15px;
  border-radius: 5px;
  margin-bottom: 20px;
  font-size: 15px;
}

.invoice-buttons {
  display: flex;
  gap: 10px;
  margin-bottom: 20px;
}

.button-nuevoPrestadorSP {
  background-color: #007bff;
  color: white;
  border: none;
  padding: 10px 20px;
  font-size: 14px;
  border-radius: 4px;
  cursor: pointer;
  transition: background-color 0.3s ease;
}

.button-nuevoPrestadorSP:hover {
  background-color: #0056b3;
}

table {
  width: 100%;
  border-collapse: collapse;
  margin: 0 auto;
  font-size: 14px;
}

th, td {
  border: 1px solid #ddd;
  padding: 10px;
  text-align: center;
}

th {
  background-color: #f2f2f2;
}

.delete-button {
  background-color: red;
  color: white;
  border: none;
  padding: 5px 10px;
  cursor: pointer;
  border-radius: 5px;
}

.delete-button:hover {
  background-color: darkred;
}

.btn-edit {
    background-color: #3498db;
    color: white;
    border: none;
    padding: 5px 10px;
    cursor: pointer;
    border-radius: 5px;
  }
  
  .btn-edit:hover {
    background-color: #218838;
  }
</style>