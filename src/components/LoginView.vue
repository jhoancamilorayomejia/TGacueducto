<template>
  <div class="login-container">
    <div class="login-form">
      <img src="/img/logoblack.png" alt="Logo" class="logo-image" />
      <h2>Login</h2>
      <div class="input-container">
        <input
          type="text"
          v-model="email"
          id="email"
          class="input-field"
          name="email"
          placeholder="Correo electrónico"
        />
      </div>
      <div class="input-container">
        <input
          type="password"
          v-model="password"
          id="password"
          class="input-field"
          name="password"
          placeholder="Contraseña"
        />
      </div>
      <div class="input-container2">
        <select
          v-model="userType"
          id="userType"
          class="input-field"
          name="userType"
        >
          <option value="admin">Administrador</option>
          <option value="company">Empresa Acueducto</option>
          <option value="customer">Usuario Cliente</option>
        </select>
      </div>
      <button class="login-button" @click="submitForm">Iniciar sesión</button>
      <p v-if="error" class="error-message">{{ error }}</p>
      <div class="forgot-password">
        <!--a href="#" class="forgot-link">¿Olvidaste la contraseña?</a-->
      </div>
    </div>
  </div>
</template>

<script>
import { ref } from 'vue';
import { useRouter } from 'vue-router';
import { sessionService } from '../services/sessionService';
import { useToast } from 'vue-toastification';

export default {
  setup() {
    const email = ref('');
    const password = ref('');
    const userType = ref('admin'); // Tipo de usuario predeterminado
    const router = useRouter();
    const toast = useToast();
    const error = ref('');

    const submitForm = async () => {
      const user = {
        email: email.value,
        password: password.value,
        userType: userType.value,
      };

      try {
        const response = await sessionService.getSession(user);

        if (response.data) {
          const { userID, userFKcompany, access_token, userName, userNit, userCedula, userLastName } = response.data;

          // Guardar en localStorage
          localStorage.setItem('userID', userID);
          localStorage.setItem('userFKcompany', userFKcompany);
          localStorage.setItem('userName', userName); // Asegúrate de que "nombre" se envía correctamente
          localStorage.setItem('userNit', userNit);
          localStorage.setItem('userCedula', userCedula);
          localStorage.setItem('userLastName', userLastName);
          localStorage.setItem('token', access_token);
          localStorage.setItem('email', email.value);

          toast.success(`Bienvenido, ${userName}`,  { 
            timeout: 3000
          });

          // Redirigir basado en el tipo de usuario
          setTimeout(() => {
            if (userType.value === 'admin') {
              router.push(`/api/admins/${userID}`);
            } else if (userType.value === 'company') {
              router.push(`/api/company/${userID}`);
            } else if (userType.value === 'customer') {
              router.push(`/api/welcome/${userID}`);
            }
          }, 3000);
        } else {
          error.value = 'Fallo al iniciar sesión.';
          toast.error(error.value, {
            timeout: 5000
          });
        }
      } catch (err) {
        error.value = 'Ocurrió un error durante el inicio de sesión.';
        toast.error(error.value, {
          timeout: 5000
        });
      }
    };

    return {
      email,
      password,
      userType,
      submitForm,
      error
    };
  }
};
</script>



<style scoped>
body {
  margin: 0;
  padding: 0;
  font-family: Arial, sans-serif;
  background-color: #1F1F1F;
  color: #E5D5A5;
}

.login-container {
  display: flex;
  justify-content: flex-start; /* Alinea el formulario hacia la izquierda */
  align-items: center;
  height: 97.5vh;
  background-image: url('https://cdn.leonardo.ai/users/65a8cf55-c959-4394-91b9-30d6f5167b8c/generations/673bbf3e-182b-4a03-8ad6-245b8e982156/Leonardo_Phoenix_A_wideangle_image_featuring_a_worn_cementcolo_0.jpg');
  background-size: cover;
  background-position: center;
  background-repeat: no-repeat;
  padding-left: 50px; /* Añade un margen izquierdo para separarlo del borde */
}

.login-form {
  display: flex;
  flex-direction: column;
  align-items: center;
  background: linear-gradient(90deg, #00BFFF  0%, #000000 80%); /* Fondo negro semitransparente */
  padding: 40px;
  border-radius: 20px;
  box-shadow: 0 4px 8px rgba(0, 0, 0, 0.2);
  width: 100%;
  max-width: 400px; /* Ajusta el tamaño del formulario */
  z-index: 1; /* Asegura que esté por encima de la imagen */
}

.input-container {
  width: 100%;
  margin-bottom: 15px;
}

.input-container2 {
  width: 50%;
  margin-bottom: 15px;
}

.input-field {
  width: 100%;
  padding: 10px;
  border: 1px solid #00BFFF;
  border-radius: 5px;
  font-size: 16px;
  background-color: #1F1F1F;
  color: #E5D5A5;
}

.login-button {
  width: 100%;
  padding: 10px;
  background: linear-gradient(90deg, #000000 0%, #e77325 100%);
  color: white;
  border: none;
  border-radius: 15px;
  font-size: 16px;
  cursor: pointer;
}

.login-button:hover {
  opacity: 0.8;
}

.error-message {
  color: red;
  margin-bottom: 10px;
}

.forgot-password {
  text-align: center;
}

.forgot-link {
  color: #E5D5A5;
  text-decoration: none;
}

.forgot-link:hover {
  text-decoration: underline;
}

/* Añadir un degradado semitransparente sobre la imagen de fondo */
.login-container::before {
  content: "";
  position: absolute;
  top: 0;
  left: 0;
  width: 100%;
  height: 100%;
  background: rgba(0, 0, 0, 0.6); /* Degradado oscuro */
  z-index: 0;
}

.logo-image {
  position: absolute;
  top: 1.5px;
  right: 10px;
  width: 75px;
  height: auto;
  z-index: 2; /* Asegura que esté sobre el formulario */
}

.login-form {
  position: relative; /* Esto es necesario para que la posición absoluta del logo funcione correctamente */
}

</style>
