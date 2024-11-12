<template>
  <link rel="stylesheet" href="https://cdnjs.cloudflare.com/ajax/libs/font-awesome/6.0.0-beta3/css/all.min.css">
  <div class="background-container"> <!-- Contenedor para el fondo -->
  <div class="invoice-container">
    <div class="invoice-header">
      <h2>Área de Gestión de Empresa</h2>
      <div class="user-info">
        <h4>Bienvenido, {{ userName }}</h4>
        <h5>Correo: {{ userEmail }} | NIT: {{ userNit }}</h5>
      </div>
    </div>
    <table class="usuario-table">
      <thead>
        <tr>
          <th>Cédula</th>
          <th>Nombre del Titular</th>
          <th>Apellidos</th>
          <th>Localidad</th>
          <th>Teléfono</th>
          <th>Email</th>
          <th>Acción</th>
          <th>Ver Información</th>
        </tr>
      </thead>
      <tbody>
        <tr v-for="usuario in usuarios" :key="usuario.idcustomer">
          <td>{{ usuario.cedula }}</td>
          <td>{{ usuario.name }}</td>
          <td>{{ usuario.last_name }}</td>
          <td>{{ usuario.address }}</td>
          <td>(+57) {{ usuario.phone }}</td>
          <td>{{ usuario.email }}</td>
          <td class="action-buttons">
            <button class="btn-edit" @click="editCustomer(usuario.idcustomer)"><i class="fas fa-edit"></i> Modificar</button>
            <button class="btn-delete" @click="deleteCustomer(usuario.idcustomer)"><i class="fas fa-trash-alt"></i>Eliminar</button>
          </td>
          <td>
            <button class="btn-info" @click="viewCustomerInfo(usuario.idcustomer, usuario.name, usuario.email, usuario.last_name, usuario.cedula)">Ver Información</button>
          </td>
        </tr>
      </tbody>
    </table>
    <div class="button-container">
      <button class="button-nuevoPrestadorSP" @click="UsuarioForm">+ Nuevo Cliente</button>
      <!--button class="button-allcustomerEmail" @click="AllCusEmail">+ Envios de PDF</button-->
    </div>
  </div>
</div>
</template>

<script>
import axios from 'axios';

export default {
  data() {
    return {
      usuarios: [],
      userEmail: '', // Email del usuario
      userID: '', // ID del usuario
      userName: '', // Nombre del usuario
      userNit: ''
    };
  },
  created() {
    this.fetchUsuarios();
    this.userEmail = localStorage.getItem('email');
    this.userID = localStorage.getItem('userID');
    this.userName = localStorage.getItem('userName');
    this.userNit = localStorage.getItem('userNit');
  },
  methods: {
    async fetchUsuarios() {
      try {
        const token = localStorage.getItem('token');
        const userID = localStorage.getItem('userID');
        const response = await axios.get(`/api/allcustomer/${userID}`, {
          headers: {
            Authorization: `Bearer ${token}`
          }
        });
        this.usuarios = response.data.usuarios; // Asegúrate de que la API devuelve el formato esperado
      } catch (error) {
        console.error('Error fetching usuarios:', error);
      }
    },
    UsuarioForm() {
      this.$router.push('/api/registerCustomer');
    },
    AllCusEmail() {
      this.$router.push('/api/customers');
    },
    async deleteCustomer(idcustomer) {
      const confirmDelete = confirm("¿Estás seguro de que deseas eliminar el cliente?");
      if (confirmDelete) {
        try {
          const token = localStorage.getItem('token');
          await axios.delete(`/api/customer/${idcustomer}`, {
            headers: {
              Authorization: `Bearer ${token}`
            }
          });
          this.fetchUsuarios(); // Cambiar fetchAllFacturas a fetchUsuarios para actualizar la lista
        } catch (error) {
          console.error('Error al eliminar el cliente:', error);
          alert('Hubo un error al intentar eliminar el cliente.');
        }
      }
    },
    editCustomer(idcustomer) {
      this.$router.push(`/api/customer/edit/${idcustomer}`);
    },
    viewCustomerInfo(idcustomer, name, email, last_name, cedula) {
      this.$router.push({
        path: `/api/customer/info-facture/${idcustomer}`,
        query: { name, email, last_name, cedula } // Aquí pasas el email como parámetro de consulta
      });
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
  padding: 30px;
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
  font-size: 12px; /* Hacer el texto más pequeño */
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
  padding: 12px 27px; /* Aumentar tamaño */
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

.button-allcustomerEmail {
  padding: 12px 27px; /* Aumentar tamaño */
  background-color: #62b5ec;
  color: white;
  border: none;
  border-radius: 20px; /* Bordes redondeados */
  box-shadow: 0 4px 10px rgba(0, 0, 0, 0.2); /* Sombra */
  font-size: 16px; /* Aumentar tamaño de fuente */
  cursor: pointer;
  transition: background-color 0.3s, transform 0.2s; /* Transición suave */
}

.button-allcustomerEmail:hover {
  background-color: #2980b9; /* Color al pasar el mouse */
  transform: scale(1.05); /* Efecto de aumento */
}
</style>
