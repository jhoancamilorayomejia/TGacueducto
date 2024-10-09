<template>
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
  background-color: #f9f9f9;
  padding: 30px;
  margin: 50px auto;
  width: 80%;
  border-radius: 8px;
  box-shadow: 0px 4px 12px rgba(0, 0, 0, 0.1);
  text-align: center;
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
  background-color: #007bff; /* Color azul */
  color: white;
}

.back-button:hover {
  background-color: #0056b3; /* Color azul más oscuro al pasar el ratón */
}
</style>

