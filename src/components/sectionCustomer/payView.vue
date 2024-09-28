<template>
  <div>
    <h1>Factura a Pagar</h1>
    <div class="price">{{ price }}</div>
    <button @click="pay">Pagar con PayPal</button>
  </div>
</template>

<script>
import { paymentService } from '../../services/paymentService'; // Ruta de importación correcta

export default {
  data() {
    return {
      price: '100.00' // Aquí puedes poner el valor a pagar
    };
  },
  methods: {
    async pay() {
      try {
        const response = await paymentService.createPayment(this.price);
        window.location.href = response.approvalUrl; // Redirigir a PayPal
      } catch (error) {
        console.error("Error al procesar el pago:", error);
      }
    }
  }
};
</script>

<style scoped>
.price {
  font-size: 24px;
  font-weight: bold;
}
</style>
