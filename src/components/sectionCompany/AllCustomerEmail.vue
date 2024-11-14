<template>
  <div class="background-container">
    <div class="facture-info-container">
      <h2 class="digital-font">Proceso de Cambio de Clave</h2>
      <h5>Correo: {{ userEmail }} | Para su seguridad las Claves se mostrarán encryptadas</h5>

      <!-- Tabla para mostrar la información de email y contraseña -->
      <table class="company-info-table">
        <thead>
          <tr>
            <th>Correo Electrónico</th>
            <th>Contraseña</th>
          </tr>
        </thead>
        <tbody>
          <tr>
            <td>{{ userEmail }}</td>
            <td>{{ userPassword }}</td>
          </tr>
        </tbody>
      </table>

      <div class="input-container">
  <label for="newPassword" class="input-label">Nueva Clave:</label>
  <div class="password-input-wrapper">
    <input
      type="password"
      id="newPassword"
      class="input-field"
      placeholder="Escribe tu nueva contraseña"
      v-model="newPassword"
    />
    <span class="toggle-password" id="togglePassword">
      <i class="fas fa-eye"></i>
    </span>
  </div>
</div>


      <div class="invoice-buttons">
        <button class="back-button" @click="goBack">Regresar</button>
        <button class="new-invoice-button" @click="changePassword" :disabled="sendingEmail">
          {{ sendingEmail ? 'Cambiando...' : 'Cambiar Contraseña' }}
        </button>
      </div>
    </div>
  </div>
</template>

<script>
import axios from 'axios';

export default {
  data() {
    return {
      userEmail: '',
      userPassword: '',
      newPassword: '',
      sendingEmail: false,
    };
  },
  created() {
    this.userEmail = localStorage.getItem('email');
    this.fetchCompanyData();
  },
  methods: {
    async fetchCompanyData() {
      try {
        const response = await axios.get('/api/company/get-email-password', {
          params: { email: this.userEmail },
        });
        this.userPassword = response.data.password;
      } catch (error) {
        console.error("Error al obtener los datos de la compañía:", error);
      }
    },
    async changePassword() {
      if (!this.newPassword) {
        alert("Por favor, ingresa una nueva contraseña.");
        return;
      }

      this.sendingEmail = true;

      try {
        await axios.post('/api/company/change-password', {
          email: this.userEmail,
          newPassword: this.newPassword,
        });
        alert("Contraseña cambiada exitosamente.");
        this.goBack();
      } catch (error) {
        console.error("Error al cambiar la contraseña:", error);
        alert("Error al cambiar la contraseña. Verifica tu información.");
      } finally {
        this.sendingEmail = false;
      }
    },
    goBack() {
      this.$router.go(-1);
    },
  },
};
</script>

<style scoped>
body {
  background-color: #ffffff;
  margin: 0;
  padding: 0;
  font-family: 'Roboto', sans-serif;
}

.facture-info-container {
  background-color: rgba(243, 243, 243, 0.9);
  padding: 30px;
  margin: 90px auto;
  width: 80%;
  border-radius: 10px;
  box-shadow: 0 8px 16px rgba(0, 0, 0, 0.2);
  position: relative;
}

h2 {
  text-align: center;
}

.company-info-table {
  width: 100%;
  border-collapse: collapse;
  margin: 20px 0;
  font-size: 1em;
  text-align: left;
}

.company-info-table th,
.company-info-table td {
  padding: 12px 15px;
  border: 1px solid #ddd;
}

.company-info-table th {
  background-color: #62b5ec;
  color: white;
}

.company-info-table td {
  background-color: #f9f9f9;
}

.invoice-buttons {
  display: flex;
  justify-content: center;
  margin-top: 20px;
}

.input-container {
  position: relative;
  margin-bottom: 20px;
  width: 100%;
}

.new-invoice-button, .back-button {
  padding: 12px 27px;
  background-color: #62b5ec;
  color: white;
  border: none;
  border-radius: 20px;
  box-shadow: 0 4px 10px rgba(0, 0, 0, 0.2);
  font-size: 16px;
  cursor: pointer;
  transition: background-color 0.3s, transform 0.2s;
  margin: 0 10px;
}

.new-invoice-button:hover, .back-button:hover {
  background-color: #2980b9;
  transform: scale(1.05);
}

.digital-font {
  font-family: 'Digital', sans-serif;
  font-size: 0.8em;
}

.background-container {
  background-image: url('https://cdn.leonardo.ai/users/65a8cf55-c959-4394-91b9-30d6f5167b8c/generations/6a0f6c27-9594-4c4d-8e92-8c60e0a33d2d/Leonardo_Phoenix_A_realistic_depiction_of_a_cluttered_workspac_3.jpg');
  background-size: cover;
  background-position: center;
  background-repeat: no-repeat;
  height: 100vh;
  display: flex;
  justify-content: center;
  align-items: center;
}

/* Estilo del campo de entrada */
.input-field {
  width: 100%;
  padding: 12px 15px;
  font-size: 16px;
  border: 2px solid #ccc;
  border-radius: 5px;
  outline: none;
  background-color: #f9f9f9;
  transition: all 0.3s ease;
}

/* Estilo cuando el campo de entrada está enfocado */
.input-field:focus {
  border-color: #5c6bc0;
  box-shadow: 0 0 5px rgba(92, 107, 192, 0.6);
}

/* Icono para mostrar/ocultar contraseña */
.password-input-wrapper {
  position: relative;
  display: flex;
  align-items: center;
}

/* Estilo del icono de visibilidad */
.toggle-password {
  position: absolute;
  right: 10px;
  top: 50%;
  transform: translateY(-50%);
  cursor: pointer;
  color: #aaa;
  transition: color 0.3s ease;
}

.toggle-password:hover {
  color: #5c6bc0;
}

/* Efecto cuando el campo tiene texto */
.input-field:valid {
  border-color: #5c6bc0;
}

/* Efecto cuando el campo está enfocado y la etiqueta flota */
.input-field:focus + .input-label,
.input-field:not(:placeholder-shown) + .input-label {
  top: -20px;
  left: 10px;
  font-size: 12px;
  font-weight: 600;
  color: #5c6bc0;
}
</style>
