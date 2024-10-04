<template>
  <div class="invoice-container">
    <div class="invoice-header">
      <h2>Información de la Empresa: {{ $route.params.idcompany }} - {{ $route.query.name }}</h2>
    </div>
    <!-- Tabla de clientes -->
    <div class="customer-table">
      <h3>Clientes Registrados</h3>
      <table>
        <thead>
          <tr>
            <th>ID de la Empresa</th>
            <th>Cedula</th>
            <th>Nombre</th>
            <th>Apellido</th>
            <th>Localidad</th>
            <th>Teléfono</th>
            <th>Email</th>
          </tr>
        </thead>
        <tbody>
          <tr v-for="usuario in usuarios" :key="usuario.idcustomer">
            <td>{{ usuario.idcompany }}</td>
            <td>{{ usuario.cedula }}</td>
            <td>{{ usuario.name }}</td>
            <td>{{ usuario.last_name }}</td>
            <td>{{ usuario.address }}</td>
            <td>{{ usuario.phone }}</td>
            <td>{{ usuario.email }}</td>
          </tr>
        </tbody>
      </table>
    </div>
  </div>
</template>

<script>
import axios from 'axios';

export default {
  data() {
    return {
      usuarios: [], // Inicializa la variable para almacenar clientes
    };
  },
  methods: {
    async fetchUsuarios() {
  try {
    const idcompany = this.$route.params.idcompany; // Obtiene el id de la empresa de la ruta
    const token = localStorage.getItem('token');
    const response = await axios.get(`/api/customer?idcompany=${idcompany}`, {
      headers: {
        Authorization: `Bearer ${token}`
      }
    });
    this.usuarios = response.data; // Almacena los usuarios en el frontend
  } catch (error) {
    console.error('Error fetching usuarios:', error);
  }
},

},
  mounted() {
    this.fetchUsuarios(); // Llama a la función cuando se monta el componente
  }
};
</script>
  
  <style scoped>
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
    font-size: 20px;
    color: white;
    text-align: center;
  }
  
  .customer-table {
    margin-top: 20px;
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
  </style>