<template>
   <div class="background-container"> <!-- Contenedor para el fondo -->
  <div class="receipt-container">
    <h1 class="receipt-title">Recibo de Pago</h1>
    <div class="receipt-details">
      <h4>Referencia de Pago: {{ FactureCod }}</h4>
      <p class="total-label">Monto Total:</p>
      <p class="total-amount">${{ totalPay.toFixed(2) }}</p>
    </div>
    <div class="button-container">
      <button @click="confirmPayment" class="pay-button">Confirmar Pago</button>
      <button @click="goBack" class="back-button">Volver</button>
    </div>

    <!-- Contenedor del botón de pago -->
    <div id="wallet_container"></div>
  </div>
</div>
</template>

<script>
export default {
  data() {
    return {
      totalPay: 0.0, // Inicializa el valor
      FactureCod: this.$route.query.codfacture || '',
      preferenceId: "" // Almacenar el ID de la preferencia generada
    };
  },
  created() {
    const payAmount = this.$route.params.totalpay; // Obtener el total desde los parámetros de la ruta
    if (!isNaN(payAmount) && payAmount.trim() !== "") {
      this.totalPay = parseFloat(payAmount); // Convertir a float
    } else {
      console.error("El valor recibido no es válido:", payAmount);
      this.totalPay = 0.0; // Establecer un valor por defecto si no es válido
    }
  },
  methods: {
    confirmPayment() {
      // Llamar a la función para crear la preferencia en el backend
      this.createPaymentPreference(this.totalPay);
    },
    createPaymentPreference(amount) {
      // Hacer una solicitud al backend para generar la preferencia
      fetch("http://localhost:8080/api/payment/preference", {
        method: "POST",
        headers: {
          "Content-Type": "application/json",
        },
        body: JSON.stringify({ amount: amount,
          FactureCod: this.FactureCod
         }),
      })
        .then((response) => response.json())
        .then((data) => {
          this.preferenceId = data.preferenceId; // Obtener el ID de la preferencia
          this.initializeMercadoPago();
        })
        .catch((error) => {
          console.error("Error al crear la preferencia:", error);
        });
    },
    initializeMercadoPago() {
      // Asegúrate de haber cargado el SDK de MercadoPago antes
      const mp = new window.MercadoPago("APP_USR-7bf75337-d896-48f0-ae16-ab08fe26882f", {
        locale: "es-CO", // Cambia esto al país que corresponda (falta cambiarlo)
      });

      // Crear el Brick para el botón de pago
      const bricksBuilder = mp.bricks();

      bricksBuilder.create("wallet", "wallet_container", {
        initialization: {
          preferenceId: this.preferenceId, // Usar la preferencia generada
        },
        customization: {
          texts: {
            valueProp: "smart_option", // Texto personalizado (opcional)
          },
        },
      });
    },
    goBack() {
      this.$router.go(-1); // Volver a la página anterior
    },
  },
};
</script>

<style scoped>
.receipt-container {
  background-color: rgba(243, 243, 243, 0.9); /* Color de formulario semi-transparente */
  padding: 30px;
  margin: 90px auto;
  width: 80%; /* Hacer el contenedor más pequeño */
  border-radius: 10px;
  box-shadow: 0 8px 16px rgba(0, 0, 0, 0.2); /* Sombra más pronunciada */
  position: relative; /* Para colocar la sombra en la parte trasera */
}

.receipt-title {
  font-size: 24px;
  color: #333;
  margin-bottom: 20px;
}

.receipt-details {
  margin-bottom: 30px;
}

.total-label {
  font-size: 18px;
  color: #555;
}

.total-amount {
  font-size: 36px;
  color: #4CAF50; /* Verde para resaltar el monto total */
  font-weight: bold;
}

.button-container {
  display: flex;
  justify-content: center;
  gap: 20px; /* Espacio entre botones */
}

.pay-button,
.back-button {
  padding: 10px 20px;
  border: none;
  border-radius: 5px;
  cursor: pointer;
  font-size: 16px;
  transition: background-color 0.3s;
}

.pay-button {
  background-color: #4caf50; /* Color verde */
  color: white;
}

.pay-button:hover {
  background-color: #45a049; /* Color verde más oscuro al pasar el ratón */
}

.back-button {
  background-color: #7f8081; /* Color azul */
  color: white;
}

.back-button:hover {
  background-color: #b3b3b3; /* Color azul más oscuro al pasar el ratón */
}

.background-container {
  background-image: url('https://cdn.leonardo.ai/users/65a8cf55-c959-4394-91b9-30d6f5167b8c/generations/3379ab07-14d2-4dd8-ba58-13c1f38e6617/Leonardo_Phoenix_A_cluttered_desk_with_multiple_papers_and_doc_0.jpg');
  background-size: cover; /* Ajusta la imagen para cubrir todo el contenedor */
  background-position: center; /* Centra la imagen */
  background-repeat: no-repeat; /* Evita que la imagen se repita */
  height: 100vh; /* O ajusta a la altura deseada */
  display: flex;
  justify-content: center; /* Centrar el contenido */
  align-items: center; /* Centrar verticalmente */
}
</style>

