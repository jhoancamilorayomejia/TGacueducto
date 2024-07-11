<template>
    <div class="container">
      <h1>Prestadores de Servicios Públicos</h1>
      <table>
        <thead>
          <tr>
            <th>NIT</th>
            <th>Nombre</th>
            <th>Dirección</th>
            <th>Teléfono</th>
            <th>Email</th>
            <th>Password</th>
            <th>ID Usuario</th>
            <th>ID Tipo Usuario</th>
          </tr>
        </thead>
        <tbody>
          <tr v-for="cliente in clientes" :key="cliente.idusuario">
            <td>{{ cliente.nit }}</td>
            <td>{{ cliente.nombre }}</td>
            <td>{{ cliente.direccion }}</td>
            <td>{{ cliente.telefono }}</td>
            <td>{{ cliente.email }}</td>
            <td>{{ cliente.password }}</td>
            <td>{{ cliente.idusuario }}</td>
            <td>{{ cliente.idtipousuario }}</td>
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
        clientes: []
      };
    },
    created() {
      this.fetchClientes();
    },
    methods: {
      async fetchClientes() {
        try {
          const response = await axios.get('/clientes');
          this.clientes = response.data;
        } catch (error) {
          console.error('Error fetching clientes:', error);
        }
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
  