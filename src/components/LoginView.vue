<template>
  <div class="login-container">
    <div class="login-form">
      <h2>Login</h2>
      <div class="input-container">
        <input type="text" v-model="email" id="email" class="input-field" name="email" placeholder="Correo electrónico">
      </div>
      <div class="input-container2">
        <input type="password" v-model="password" id="password" class="input-field" name="password" placeholder="Contraseña">
      </div>
      <div class="input-container2">
        <select v-model="userType" id="userType" class="input-field" name="userType">
          <option value="admin">Soy Administrador</option>
          <option value="company">Prestador de Servicio publico</option>
        </select>
      </div>
      <button class="login-button" @click="submitFrom">Iniciar sesión</button>
      <p v-if="error" class="error-message">{{ error }}</p>
      <div class="forgot-password">
        <a href="#" class="forgot-link">¿Olvidaste la contraseña?</a>
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
    const userType = ref('admin'); // Default to admin
    const router = useRouter();
    const toast = useToast();

    const submitFrom = async () => {
      const user = {
        email: email.value,
        password: password.value,
        userType: userType.value
      };

      try {
        const response = await sessionService.getSession(user);

        if (response.data) {
          const { ID, token } = response.data;
          localStorage.setItem('ID', ID);
          localStorage.setItem('token', token);

          toast.success('Inicio de sesión exitoso!', {
            timeout: 3000
          });
          setTimeout(() => {
            router.push('/api/admins');
          }, 3000);
        } else {
          toast.error('Fallo al iniciar sesión.', {
            timeout: 5000
          });
        }

      } catch (error) {
        toast.error('Ocurrió un error durante el inicio de sesión.', {
          timeout: 5000
        });
      }
    };

    return {
      email,
      password,
      userType,
      submitFrom
    };
  }
};
</script>




<style scoped>
body {
  background-color: #1F1F1F;
  color: #E5D5A5;
  font-family: Arial, sans-serif;
}

.login-container {
  display: flex;
  justify-content: center;
  align-items: center;
  height: 100vh;
}

.login-form {
  display: flex;
  flex-direction: column;
  align-items: center;
  background-color: #121212;
  padding: 40px;
  border-radius: 80px;
  box-shadow: 0 4px 8px rgba(0, 0, 0, 0.1);
  width: 90%;
  max-width: 200px;
  height: 85vh;
}

.welcome-text {
  font-size: 24px;
  margin-bottom: 20px;
  color: #E5D5A5;
}

.icon-container {
  margin-bottom: 20px;
}

.icon {
  font-size: 40px;
  background-color: #ccc;
  border-radius: 50%;
  padding: 10px;
}

.input-container {
  width: 110%;
  margin-bottom: 15px;
  position: relative;
  margin-bottom: 10px;
  height: 10vh;
  margin-top: 80px;
}

.input-container2 {
  width: 110%;
  margin-bottom: 15px;
  position: center;
  margin-bottom: 80px;
}

.input-field {
  width: 100%;
  padding: 10px;
  border: 1px solid #404040;
  border-radius: 5px;
  font-size: 16px;
  background-color: #1F1F1F;
  color: #E5D5A5;
}

.password-icon {
  position: absolute;
  right: 10px;
  top: 10px;
  cursor: pointer;
  color: #E5D5A5;
}

.login-button {
  width: 100%;
  padding: 10px;
  background: linear-gradient(90deg, #00BFFF 0%, #8A2BE2 100%);
  color: white;
  border: none;
  border-radius: 15px;
  font-size: 16px;
  cursor: pointer;
  margin-bottom: 10px;
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
</style>