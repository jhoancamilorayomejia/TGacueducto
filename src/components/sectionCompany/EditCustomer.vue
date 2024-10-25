<template>
  <div class="background-container"> <!-- Contenedor para el fondo -->
  <div class="invoice-container">
    <div class="invoice-header">
      <h2>Área de Gestión de Clientes</h2>
    </div>
    <div class="invoice-buttons">
      
    </div>
    <table>
      <thead>
        <tr>
          <th>Nombre del Titular</th>
          <th>Apellidos</th>
          <th>Localidad</th>
          <th>Teléfono</th>
          <th>Email</th>
          <th>Acciones</th>
        </tr>
      </thead>
      <tbody>
        <tr v-for="usuario in usuarios" :key="usuario.idcustomer">
          <td>
            <input v-model="usuario.name" />
          </td>
          <td>
            <input v-model="usuario.last_name" />
          </td>
          <td>
            <input v-model="usuario.address" />
          </td>
          <td>
            <input v-model="usuario.phone" />
          </td>
          <td>
            <input v-model="usuario.email" />
          </td>
          <td>
            <button @click="updateCustomer(usuario)" class="btn-save">Guardar</button>
          </td>
        </tr>
      </tbody>
    </table>
    <div class="invoice-buttons">
      <button class="back-button" type="button" @click="goBack">Regresar</button>
    </div>
  </div>
</div>
</template>

<script>
import axios from 'axios';

export default {
  data() {
    return {
      usuarios: []
    };
  },
  created() {
    this.fetchUsuarios();
  },
  methods: {
    async fetchUsuarios() {
      try {
        const token = localStorage.getItem('token');
        const userID = localStorage.getItem('userID'); // Obtener el userID del localStorage
        const response = await axios.get(`/api/allcustomer/${userID}`, { // Enviar el userID en la URL
          headers: {
            Authorization: `Bearer ${token}`
          }
        });
        this.usuarios = response.data.usuarios; // Ajustar según el formato de la respuesta del backend
      } catch (error) {
        console.error('Error fetching usuarios:', error);
      }
    },
    UsuarioForm() {
      // Redireccionar a la página de registro de cliente
      this.$router.push('/api/registerCustomer');
    },
    async updateCustomer(customer) {
      try {
        const token = localStorage.getItem('token');
        await axios.put(`/api/customer/${customer.idcustomer}`, customer, {
          headers: {
            Authorization: `Bearer ${token}`
          }
        });
        alert('Cliente actualizado exitosamente.');
      } catch (error) {
        console.error('Error al actualizar cliente:', error);
        alert('Hubo un error al intentar actualizar el cliente.');
      }
    },
    goBack() {
      this.$router.go(-1); // Regresa a la página anterior (cualquiera)
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
}

.invoice-container {
  background-color: rgba(243, 243, 243, 0.9); /* Color de formulario semi-transparente */
  padding: 60px;
  margin: 90px auto;
  width: 80%; /* Hacer el contenedor más pequeño */
  border-radius: 10px;
  box-shadow: 0 8px 16px rgba(0, 0, 0, 0.2); /* Sombra más pronunciada */
  position: relative; /* Para colocar la sombra en la parte trasera */
}

.invoice-header {
  display: flex;
  justify-content: space-between; /* Alinea los elementos a los extremos */
  background-color: #b7daee;
  padding: 5px; /* Reducir el padding */
  border-radius: 15px;
  margin-bottom: 10px;
  font-size: 11px; /* Hacer el texto más pequeño */
}

.user-info {
  text-align: right; /* Alinea el texto a la derecha */
  color: rgb(0, 0, 0); /* Texto en blanco para visibilidad */
}

.usuario-table {
  width: 100%;
  border-collapse: collapse;
  margin: 0 auto;
  font-size: 11px; /* Hacer el tamaño de la letra de la tabla más pequeño */
}

th, td {
  border: 1px solid #ddd;
  padding: 4px; /* Reducir el padding */
  text-align: center;
}

th {
  background-color: #f2f2f2;
}

button {
  padding: 5px 10px;
  background-color: #18b139;
  color: white;
  border: none;
  border-radius: 4px;
  cursor: pointer;
}

button:hover {
  background-color: #31e021;
}

.btn-edit {
  background-color: #31e021;
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
}

.btn-delete:hover {
  background-color: #c0392b;
}

.action-buttons {
  display: flex;
  justify-content: center; 
  gap: 10px; 
}

.btn-info {
  background-color: #e77325;
  color: white;
  border: none;
  padding: 5px 10px;
  cursor: pointer;
  border-radius: 5px;
}

.btn-info:hover {
  background-color: #d35400;
}

/* Añadir un fondo detrás de invoice-container */
.invoice-container::before {
  content: "";
  position: absolute;
  top: 10px;
  left: 10px;
  right: 10px;
  bottom: 10px;
  background-color: #8ED6F2; /* Color azul agua */
  border-radius: 10px;
  z-index: -1; /* Asegurarse de que quede detrás del contenedor */
  box-shadow: 0 14px 121px rgba(0, 0, 0, 0.1); /* Sombra para el fondo */
}

/* Estilo para el botón "Nuevo Cliente" */
.button-container {
  display: flex;
  justify-content: center; /* Centrar el botón */
  margin-top: 20px; /* Espacio superior */
}

.button-nuevoPrestadorSP {
  padding: 15px 30px; /* Aumentar tamaño */
  background-color: #62b5ec;
  color: white;
  border: none;
  border-radius: 20px; /* Bordes redondeados */
  box-shadow: 0 4px 10px rgba(0, 0, 0, 0.2); /* Sombra */
  font-size: 16px; /* Aumentar tamaño de fuente */
  cursor: pointer;
  transition: background-color 0.3s, transform 0.2s; /* Transición suave */
}

.button-nuevoPrestadorSP:hover {
  background-color: #2980b9; /* Color al pasar el mouse */
  transform: scale(1.05); /* Efecto de aumento */
}

.background-container {
  background-image: url('https://cdn.leonardo.ai/users/65a8cf55-c959-4394-91b9-30d6f5167b8c/generations/6a0f6c27-9594-4c4d-8e92-8c60e0a33d2d/Leonardo_Phoenix_A_realistic_depiction_of_a_cluttered_workspac_3.jpg');
  background-size: cover; /* Ajusta la imagen para cubrir todo el contenedor */
  background-position: center; /* Centra la imagen */
  background-repeat: no-repeat; /* Evita que la imagen se repita */
  height: 100vh; /* O ajusta a la altura deseada */
  display: flex;
  justify-content: center; /* Centrar el contenido */
  align-items: center; /* Centrar verticalmente */
}

.invoice-buttons {
  display: flex;
  justify-content: center; /* Centrar los botones */
  margin-top: 20px; /* Espacio superior para los botones */
}

 .back-button {
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

.back-button:hover {
  background-color: #2980b9; /* Color al pasar el mouse */
  transform: scale(1.05); /* Efecto de aumento */
}
</style>
