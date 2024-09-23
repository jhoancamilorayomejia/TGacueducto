<template>
  <div class="invoice-container">
    <div class="invoice-header">
      <h2>Área de Gestión de Empresas</h2>
      <h5>Bienvenido, {{ userEmail }} con ID: {{ userID }}</h5>
    </div>
    <div class="invoice-buttons">
      <button class="button-nuevoPrestadorSP" @click="UsuarioForm">+ Nuevo Cliente</button>
    </div>
    <table>
      <thead>
        <tr>
          <th>Nombre del Titular</th>
          <th>Apellidos</th>
          <th>Localidad</th>
          <th>Teléfono</th>
          <th>Email</th>
          <th>Acción</th>
        </tr>
      </thead>
      <tbody>
        <tr v-for="usuario in usuarios" :key="usuario.idcustomer">
          <td>{{ usuario.name }}</td>
          <td>{{ usuario.last_name }}</td>
          <td>{{ usuario.address }}</td>
          <td>{{ usuario.phone }}</td>
          <td>{{ usuario.email }}</td>
          <td class="action-buttons">
            <button class="btn-edit" @click="editCustomer(usuario.idcustomer)">Modificar</button>
            <button class="btn-delete" @click="deleteCustomer(usuario.idcustomer)">Eliminar</button>
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
      usuarios: [],
      userEmail: '', // Email del usuario
      userID: '' // ID del usuario
    };
  },
  created() {
    this.fetchUsuarios();
    this.userEmail = localStorage.getItem('email'); // Obtener el correo del localStorage
    this.userID = localStorage.getItem('userID'); // Cambiar 'ID' a 'userID'
  },
  methods: {
    async fetchUsuarios() {
      try {
        const token = localStorage.getItem('token');
        const userID = localStorage.getItem('userID'); // Obtener el ID del localStorage
        const response = await axios.get(`/api/allcustomer/${userID}`, {
          headers: {
            Authorization: `Bearer ${token}`
          }
        });
        this.usuarios = response.data.usuarios; // Ajustar según el formato de respuesta
      } catch (error) {
        console.error('Error fetching usuarios:', error);
      }
    },
    UsuarioForm() {
      this.$router.push('/api/registerCustomer');
    },
    async deleteCustomer(idcustomer) {
      const confirmDelete = confirm("¿Estás seguro de que deseas eliminar este cliente?");
      if (confirmDelete) {
        try {
          const token = localStorage.getItem('token');
          await axios.delete(`/api/customer/${idcustomer}`, {
            headers: {
              Authorization: `Bearer ${token}`
            }
          });
          this.fetchUsuarios();
        } catch (error) {
          console.error('Error al eliminar cliente:', error);
          alert('Hubo un error al intentar eliminar al cliente.');
        }
      }
    },
    editCustomer(idcustomer) {
      this.$router.push(`/api/customer/edit/${idcustomer}`);
    }
  }
};
</script>

  
  
<style scoped>
body {
  background-color: #ffffff; /* Fondo sólido para toda la página */
  margin: 0;
  padding: 0;
  font-family: 'Roboto', sans-serif; /* Tipo de letra */
  font-size: 11px; /* Tamaño de letra general */
}

.invoice-container {
  background-color: #f3f3f3; /*color de formulario */
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
  /* text-align: center; */
  font-size: 15px;
}

.invoice-buttons {
  display: flex;
  /*justify-content: center; */
  gap: 10px;
  margin-bottom: 20px;
}

.btn {
  padding: 10px 20px;
  border: none;
  background-color: #62b5ec;
  color: white;
  border-radius: 4px;
  cursor: pointer;
  font-size: 12px;
}


.btn:hover {
  background-color: #2980b9;
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

button {
  padding: 5px 10px;
  background-color: #037bca;
  color: white;
  border: none;
  border-radius: 4px;
  cursor: pointer;
}

button:hover {
  background-color: #3498db;
}

.btn-edit {
  background-color: #3498db;
  color: white;
  padding: 5px 10px;
  border: none;
  border-radius: 4px;
  cursor: pointer;
  margin-right: 10px;
}

.btn-edit:hover {
  background-color: #2980b9;
}

.btn-delete {
  background-color: #e74c3c;
  color: white;
  padding: 5px 10px;
  border: none;
  border-radius: 4px;
  cursor: pointer;
  
}

.btn-delete:hover {
  background-color: #c0392b;
}

.action-buttons {
  display: flex;
  justify-content: center; 
  gap: 15px; 
}


</style>
  