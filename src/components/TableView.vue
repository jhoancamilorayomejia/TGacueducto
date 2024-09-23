<template>
  <div class="invoice-container">
    <div class="invoice-header">
      <h2>Área de Administrador</h2>
      <h5>Bienvenido, {{ userEmail }} con ID: {{ userID }} nombre: {{ userName }}</h5>
    </div>
    <div class="invoice-buttons">
      <button class="btn" @click="adminForm">Nuevo Admin</button>
      <button class="btn" @click="Allcompany">Ver Registro de Prestadores de Servicio Público</button>
      <button class="btn" @click="companyForm">Registrar Empresa de Acueducto</button>
    </div>
    <table>
      <thead>
        <tr>
          <th>ID</th>
          <th>Nombre</th>
          <th>Apellido</th>
          <th>Email</th>
          <th>Acción</th>
        </tr>
      </thead>
      <tbody>
        <tr v-for="admin in admins" :key="admin.idadmin">
          <td>{{ admin.idadmin }}</td>
          <td>{{ admin.nombre }}</td>
          <td>{{ admin.apellido }}</td>
          <td>{{ admin.email }}</td>
          <td>
            <button class="btn-edit" @click="editAdmin(admin.idadmin)">Modificar</button>
            <button @click="deleteAdmin(admin.idadmin)">Eliminar</button>
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
      admins: [],
      userEmail: '', // Correo del usuario
      userID: '',    // ID del usuario
      userName: ''   // Nombre del usuario
    };
  },
  created() {
    this.fetchAdmins();
    this.userEmail = localStorage.getItem('email'); // Obtener el correo del localStorage
    this.userID = localStorage.getItem('userID');   // Obtener el ID del usuario
    this.userName = localStorage.getItem('userName');  // Obtener el nombre del usuario
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
    adminForm() {
      this.$router.push('/api/registerAdmin');
    },
    companyForm() {
      this.$router.push('/api/register');
    },
    Allcompany() {
      this.$router.push('/api/AllCompany');
    },
    async deleteAdmin(idadmin) {
      const confirmDelete = confirm("¿Estás seguro de que deseas eliminar este administrador?");
      if (confirmDelete) {
        try {
          const token = localStorage.getItem('token');
          await axios.delete(`/api/admins/${idadmin}`, {
            headers: {
              Authorization: `Bearer ${token}`
            }
          });
          this.fetchAdmins();
        } catch (error) {
          console.error('Error al eliminar administrador:', error);
          alert('Hubo un error al intentar eliminar al administrador.');
        }
      }
    },
    editAdmin(idadmin) {
      this.$router.push(`/api/admin/edit/${idadmin}`);
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
  background-color: #e74c3c;
  color: white;
  border: none;
  border-radius: 4px;
  cursor: pointer;
}

button:hover {
  background-color: #c0392b;
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

</style>
