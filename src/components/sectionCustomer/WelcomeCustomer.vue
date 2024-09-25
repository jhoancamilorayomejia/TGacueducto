<template>
  <div class="welcome-container">
    <div class="welcome-header">
      <h2>Sistema de Facturación Digital de Pequeños Acueductos</h2>
      <h5>Bienvenido, {{ userName }} - N° Cedula: {{ userCedula }}</h5>
    </div>
    <div class="welcome-buttons">
      <!-- aqui para añadir botones funcionales-->
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
      userName: '',
      userCedula: ''
    };
  },
  created() {
    this.fetchUsuarios();
    this.userEmail = localStorage.getItem('email'); // Obtener el correo del localStorage
    this.userID = localStorage.getItem('userID'); // OBtener ID del usuario cliente
    this.userName = localStorage.getItem('userName'); // Obtener el nombre del usuario cliente'
    this.userCedula = localStorage.getItem('userCedula'); // Obtener la Cedula del usuario cliente'
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
    }    
  }
};

</script>

<style scoped>
body {
  background-color: #ffffff;
  margin: 0;
  padding: 0;
  font-family: 'Roboto', sans-serif;
  font-size: 11px;
}

.welcome-container {
  background-color: #f3f3f3;
  padding: 20px;
  margin: 50px auto;
  width: 80%;
  border-radius: 10px;
  box-shadow: 0px 4px 8px rgba(0, 0, 0, 0.1);
}

.welcome-header {
  background-color: #62b5ec;
  padding: 15px;
  border-radius: 5px;
  margin-bottom: 20px;
  font-size: 18px;
  text-align: center;
}

.welcome-buttons {
  display: flex;
  justify-content: center;
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
</style>