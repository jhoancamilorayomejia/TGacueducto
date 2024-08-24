<template>
  <div class="invoice-container">
    <div class="invoice-header">
      <h2>Area de Administrador</h2>
    </div>
    <div class="invoice-buttons">
      <button class="button-registerPrestadorSP" @click="Allcompany">Ver Registro de Prestadores de Servicio Publico</button>
      <button class="button-nuevoPrestadorSP" @click="companyForm">Registrar Empresa de Acueducto</button>
    </div>
    <table>
      <thead>
        <tr>
          <th>ID</th>
          <th>Nombre</th>
          <th>Apellido</th>
          <th>Email</th>
          <th>Perfil</th>
        </tr>
      </thead>
      <tbody>
        <tr v-for="admin in admins" :key="admin.idadmin">
          <td>{{ admin.idadmin }}</td>
          <td>{{ admin.nombre }}</td>
          <td>{{ admin.apellido }}</td>
          <td>{{ admin.email }}</td>
          <td>{{ admin.tipouser }}</td>
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
      admins: []
    };
  },
  created() {
    this.fetchAdmins();
  },
  methods: {
    async fetchAdmins() {
      try {
        const token = localStorage.getItem('token');
        const response = await axios.get('/api/admins', {
          headers: {
            Authorization: `Bearer ${token}`
          }
        });
        this.admins = response.data;
      } catch (error) {
        console.error('Error fetching admins:', error);
      }      
    },
    companyForm() {
      // Redireccionar a la página de registro de empresa
      this.$router.push('/api/register');
    },
    Allcompany() {
      //Redireccionar a la pagina de empresas registradas
      this.$router.push('/api/AllCompany');
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

