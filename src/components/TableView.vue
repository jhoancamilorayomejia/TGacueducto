<template>
  <div>
    <h1>Administradores</h1>
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
        const response = await axios.get('/admins');
        this.admins = response.data;
      } catch (error) {
        console.error('Error fetching admins:', error);
      }
    }
  }
};
</script>

<style scoped>
table {
  width: 100%;
  border-collapse: collapse;
}

th, td {
  border: 1px solid #ddd;
  padding: 8px;
}

th {
  background-color: #f2f2f2;
}
</style>


