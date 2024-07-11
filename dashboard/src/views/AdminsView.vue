<template>
  <div class="container">
    <h1>Administradores</h1>
    <div class="buttons">
      <button @click="navigateTo('/serviciosPrestadores')">Ver prestadores de servicios públicos</button>
      <button @click="navigateTo('/clientes')">Clientes de servicios públicos</button>
    </div>
    <table>
      <thead>
        <tr>
          <th>ID</th>
          <th>Nombre</th>
          <th>Apellido</th>
          <th>Email</th>
          <th>Password</th>
          <th>Tipo de Usuario</th>
        </tr>
      </thead>
      <tbody>
        <tr v-for="admin in admins" :key="admin.idadmin">
          <td>{{ admin.idadmin }}</td>
          <td>{{ admin.nombre }}</td>
          <td>{{ admin.apellido }}</td>
          <td>{{ admin.email }}</td>
          <td>{{ admin.password }}</td>
          <td>{{ admin.tipouser }}</td>
        </tr>
      </tbody>
    </table>
  </div>
</template>

<script>
import axios from '../api/axios';

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
        const response = await axios.get('/admins', {
          headers: {
            "authorization": localStorage.getItem("token")
          }
        });
        this.admins = response.data;
      } catch (error) {
        console.error('Error fetching admins:', error);
      }
    },
    navigateTo(route) {
      this.$router.push(route);
    }
  }
};
</script>

<style>
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

