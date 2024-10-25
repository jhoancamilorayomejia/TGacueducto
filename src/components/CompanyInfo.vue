<template>
  <div class="background-container"> <!-- Contenedor para el fondo -->
  <div class="invoice-container">
    <div class="invoice-header">
      <h2>Información de la Empresa: - {{ $route.query.name }}</h2>
      <!--h5>{{ $route.params.idcompany }}</h5-->
    </div>
    <!-- Tabla de clientes -->
    <div class="customer-table">
      <h3>Clientes Registrados</h3>
      <table>
        <thead>
          <tr>
            <!--th>ID de la Empresa</th-->
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
            <!--td>{{ usuario.idcompany }}</td-->
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
    background-color: #f3f3f3; /*color de formulario */
  padding: 20px;
  margin: 50px auto;
  width: 80%;
  border-radius: 10px;
  box-shadow: 0px 4px 8px rgba(0, 0, 0, 0.1);
  }
  
  .invoice-header {
    display: flex;  
  justify-content: space-between; /* Alinea los elementos a los extremos */
  background-color: #b7daee;
  padding: 5px; /* Reducir el padding */
  border-radius: 15px;
  margin-bottom: 10px;
  font-size: 12px; /* Hacer el texto más pequeño */
  }
  
  .customer-table {
    width: 100%;
  border-collapse: collapse;
  margin: 0 auto;
  font-size: 11px; /* Hacer el tamaño de la letra de la tabla más pequeño */
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

  .background-container {
  background-image: url('https://cdn.leonardo.ai/users/65a8cf55-c959-4394-91b9-30d6f5167b8c/generations/3b56f729-de70-45a8-a5a0-6c274211025a/Leonardo_Phoenix_A_modern_administrative_dashboard_webpage_fea_1.jpg');
  background-size: cover; /* Ajusta la imagen para cubrir todo el contenedor */
  background-position: center; /* Centra la imagen */
  background-repeat: no-repeat; /* Evita que la imagen se repita */
  height: 100vh; /* O ajusta a la altura deseada */
  display: flex;
  justify-content: center; /* Centrar el contenido */
  align-items: center; /* Centrar verticalmente */
}
  </style>