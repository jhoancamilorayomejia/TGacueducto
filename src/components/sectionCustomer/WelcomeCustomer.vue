<template>
  <div class="welcome-container">
    <div class="welcome-header">
      <h2>Sistema de Facturación Digital de Pequeños Acueductos</h2>
      <h5>Bienvenido, {{ userName }} - N° Cédula: {{ userCedula }}</h5>
    </div>
    <table>
      <thead>
        <tr>
          <th>Número de Factura</th>
          <th>Fecha de Emisión</th>
          <th>Fecha Límite de Pago</th>
          <th>Monto Total</th>
          <th>Estado</th>
          <th>Factura PDF</th>
          <th>Proceso de Pago</th> <!-- Nueva columna -->
        </tr>
      </thead>
      <tbody>
        <tr v-for="factura in facturas" :key="factura.idfacture">
          <td>{{ factura.facturenumber }}</td>
          <td>{{ factura.datecreation }}</td>
          <td>{{ factura.datepayment }}</td>
          <td>${{ factura.totalpay }}</td>
          <td>Pendiente/Pagado</td>
          <td>
            <button @click="downloadInvoice(factura)" class="download-button">
              <i class="fas fa-download"></i> Descargar
            </button>
          </td>
          <td>
            <button @click="goToPayment(factura.totalpay)" class="payment-button">
              Proceso de pago
            </button>
          </td>
        </tr>
      </tbody>
    </table>
  </div>
</template>

<script>
import axios from 'axios';
import jsPDF from 'jspdf'; // Manteniendo jsPDF para la descarga de facturas en PDF

export default {
  data() {
    return {
      userEmail: '', // Email del usuario
      userID: '', // ID del usuario
      userName: '', // Nombre del usuario
      userCedula: '', // Cédula del usuario
      facturas: [] // Lista de facturas (todas)
    };
  },
  created() {
    this.userEmail = localStorage.getItem('email'); // Obtener el correo del localStorage
    this.userID = localStorage.getItem('userID'); // Obtener ID del usuario cliente
    this.userName = localStorage.getItem('userName'); // Obtener el nombre del usuario cliente
    this.userCedula = localStorage.getItem('userCedula'); // Obtener la cédula del usuario cliente

    this.fetchFacturas(); // Llama a la función para obtener facturas al crear el componente
  },
  methods: {
    // Función para obtener todas las facturas del cliente
    async fetchFacturas() {
      try {
        const token = localStorage.getItem('token'); // Obtener token de autenticación
        const response = await axios.get('/api/facturas', {
          headers: {
            Authorization: `Bearer ${token}` // Autorización con token
          }
        });

        // Filtrar las facturas para mostrar solo las que correspondan al usuario actual
        this.facturas = response.data.filter(factura => factura.idcustomer == this.userID);

      } catch (error) {
        console.error('Error fetching facturas:', error);
      }
    },

    // Función para descargar una factura en formato PDF
    downloadInvoice(factura) {
      const doc = new jsPDF();
      doc.setFontSize(12);
      doc.text(`Factura Número: ${factura.facturenumber}`, 10, 10);
      doc.text(`Fecha de Emisión: ${factura.datecreation}`, 10, 20);
      doc.text(`Fecha Límite de Pago: ${factura.datepayment}`, 10, 30);
      doc.text(`Monto Total: $${factura.totalpay}`, 10, 40);
      doc.text(`Cliente: ${this.userName} - Cédula: ${this.userCedula}`, 10, 50);
      doc.text(`Correo: ${this.userEmail}`, 10, 60);
      doc.save(`Factura_${factura.facturenumber}.pdf`);
    },

    goToPayment(totalPay) {
  console.log("Monto total al hacer clic:", totalPay); // Imprimir para depurar
  this.$router.push({ path: `/api/payment/${totalPay}` }); // Redirige a la vista de pago con el monto
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

table {
  width: 100%;
  border-collapse: collapse;
  margin-top: 20px;
}

th, td {
  padding: 12px;
  text-align: left;
  border-bottom: 1px solid #ddd;
}

.download-button {
  padding: 5px 10px;
  border: none;
  background-color: #62b5ec;
  color: white;
  border-radius: 4px;
  cursor: pointer;
}

.download-button:hover {
  background-color: #2980b9;
}

.payment-button {
  padding: 5px 10px;
  border: none;
  background-color: #4CAF50; /* Color verde */
  color: white;
  border-radius: 4px;
  cursor: pointer;
}

.payment-button:hover {
  background-color: #45a049; /* Color verde más oscuro al pasar el ratón */
}
</style>

