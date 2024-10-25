<template>
  <div class="background-container"> <!-- Contenedor para el fondo -->
  <div class="invoice-container">
    <div class="invoice-header">
      <h2>Área de Administrador</h2>
      
      <div class="user-info">
        <h4>Bienvenido, {{ userName }} </h4>
        <!--h5>ID: {{ userID }}</h5-->
      </div>
      <!--h5>Bienvenido, {{ userEmail }} con ID: {{ userID }} nombre: {{ userName }}</h5-->
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
    <div class="invoice-buttons">
      <button class="new-invoice-button"  @click="adminForm">Nuevo Admin</button>
      <button class="new-invoice-button"  @click="Allcompany">Ver Entidades</button>
      <button class="new-invoice-button"  @click="companyForm">Registro de Entidades</button>
    </div>
  </div>
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
  display: flex;  
  justify-content: space-between; /* Alinea los elementos a los extremos */
  background-color: #b7daee;
  padding: 5px; /* Reducir el padding */
  border-radius: 15px;
  margin-bottom: 10px;
  font-size: 12px; /* Hacer el texto más pequeño */
}

.user-info {
  text-align: right; /* Alinea el texto a la derecha */
  color: rgb(0, 0, 0); /* Texto en blanco para visibilidad */
}


.invoice-buttons {
  display: flex;
  justify-content: center; /* Centrar los botones */
  margin-top: 20px; /* Espacio superior para los botones */
}

.new-invoice-button, .back-button {
  padding: 12px 27px; /* Aumentar tamaño */
  background-color: #62b5ec;
  color: white;
  border: none;
  border-radius: 20px; /* Bordes redondeados */
  box-shadow: 0 4px 10px rgba(0, 0, 0, 0.2); /* Sombra */
  font-size: 16px; /* Aumentar tamaño de fuente */
  cursor: pointer;
  transition: background-color 0.3s, transform 0.2s; /* Transición suave */
  margin: 0 10px; /* Espacio entre botones */
}

.new-invoice-button:hover, .back-button:hover {
  background-color: #2980b9; /* Color al pasar el mouse */
  transform: scale(1.05); /* Efecto de aumento */
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
