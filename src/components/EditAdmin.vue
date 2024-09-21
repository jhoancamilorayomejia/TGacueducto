<template>
  <div class="invoice-container">
    <div class="invoice-header">
      <h2>Área de Administrador</h2>
    </div>
    <div class="invoice-buttons">
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
          <th>Acción</th> <!-- Nueva columna para acciones -->
        </tr>
      </thead>
      <tbody>
        <tr v-for="admin in admins" :key="admin.idadmin">
          <td>{{ admin.idadmin }}</td>
          <td>
            <input v-model="admin.nombre" />
          </td>
          <td>
            <input v-model="admin.apellido" />
          </td>
          <td>
            <input v-model="admin.email" />
          </td>
          <td>
            <button @click="updateAdmin(admin)" class="btn-edit">Guardar</button> <!-- Botón para guardar cambios -->
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
    async updateAdmin(admin) {
      try {
        const token = localStorage.getItem('token');
        await axios.put(`/api/admins/${admin.idadmin}`, admin, {
          headers: {
            Authorization: `Bearer ${token}`
          }
        });
        alert('Administrador actualizado exitosamente.');
      } catch (error) {
        console.error('Error al actualizar administrador:', error);
        alert('Hubo un error al intentar actualizar el administrador.');
      }
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
          this.fetchAdmins(); // Volver a obtener la lista de admins después de eliminar
        } catch (error) {
          console.error('Error al eliminar administrador:', error);
          alert('Hubo un error al intentar eliminar al administrador.');
        }
      }
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
  background-color: #28a745;
  color: white;
  padding: 5px 10px;
  border: none;
  border-radius: 4px;
  cursor: pointer;
}

.btn-edit:hover {
  background-color: #218838;
}



</style>