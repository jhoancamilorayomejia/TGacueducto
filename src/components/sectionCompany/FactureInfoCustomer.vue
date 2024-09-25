<template>
    <div class="facture-info-container">
      <h2>Historial de facturas</h2>
      <h5>Correo: {{ userEmail }} </h5>
      <h3>Cliente: {{ customerName }}, ID: {{ customerId }}</h3>
      <table>
        <thead>
          <tr>
            <th>Número de Factura</th>
            <th>Fecha de Emisión</th>
            <th>Fecha Límite de Pago</th>
            <th>Monto Total</th>
            <th>Factura PDF</th>
          </tr>
        </thead>
        <tbody>
          <tr v-for="factura in facturas" :key="factura.idfacture">
            <td>{{ factura.facturenumber }}</td>
            <td>{{ factura.datecreation }}</td>
            <td>{{ factura.datepayment }}</td>
            <td>${{ factura.totalpay }}</td>
            <td>
              <button @click="downloadInvoice(factura)" class="download-button">
                <i class="fas fa-download"></i> Descargar
              </button>
            </td>
          </tr>
        </tbody>
      </table>
    </div>
  </template>
  
  <script>
  import axios from 'axios';
  import jsPDF from 'jspdf';
  
  export default {
    data() {
      return {
        //datos de company
        userEmail: '',
        // datos de customer
        customerId: this.$route.params.idcustomer, // Obtener ID del cliente desde la ruta
        customerName: this.$route.query.name, // Obtener nombre del cliente desde la consulta
        facturas: []
      };
    },
    created() {
      this.fetchAllFacturas(); // Llamar a la función al crear el componente
      this.userEmail = localStorage.getItem('email');
    },
    methods: {
      async fetchAllFacturas() {
        try {
          const response = await axios.get(`/api/facturas/${this.customerId}`); // Obtener facturas por ID de cliente
          // Filtrar las facturas para mostrar solo las que coinciden con el idcustomer
          this.facturas = response.data.filter(factura => factura.idcustomer == this.customerId);
        } catch (error) {
          console.error('Error fetching facturas:', error);
        }
      },
      downloadInvoice(factura) {
        // Crear una nueva instancia de jsPDF
        const doc = new jsPDF();
  
        // Añadir contenido al PDF
        doc.setFontSize(12);
        doc.text(`Factura Número: ${factura.facturenumber}`, 10, 10);
        doc.text(`Fecha de Emisión: ${factura.datecreation}`, 10, 20);
        doc.text(`Fecha Límite de Pago: ${factura.datepayment}`, 10, 30);
        doc.text(`Monto Total: $${factura.totalpay}`, 10, 40);
        doc.text(`Cliente: ${this.customerName} (ID: ${this.customerId})`, 10, 50);
        doc.text(`Correo: ${this.userEmail}`, 10, 60);
  
        // Descargar el PDF con el nombre de la factura
        doc.save(`Factura_${factura.facturenumber}.pdf`);
      }
    }
  };
  </script>
  
  <style scoped>
  .facture-info-container {
    background-color: #f3f3f3;
    padding: 20px;
    margin: 50px auto;
    width: 80%;
    border-radius: 10px;
    box-shadow: 0px 4px 8px rgba(0, 0, 0, 0.1);
  }
  
  h2 {
    text-align: center;
  }
  
  table {
    width: 100%;
    border-collapse: collapse;
    margin-top: 20px;
  }
  
  th,
  td {
    border: 1px solid #ddd;
    padding: 10px;
    text-align: center;
  }
  
  th {
    background-color: #f2f2f2;
  }
  
  .download-button {
    background-color: #4CAF50; /* Color verde para el botón */
    color: white; /* Texto blanco */
    border: none; /* Sin borde */
    border-radius: 5px; /* Bordes redondeados */
    padding: 10px 15px; /* Espaciado */
    cursor: pointer; /* Cambiar cursor al pasar */
    transition: background-color 0.3s; /* Transición suave */
  }
  
  .download-button:hover {
    background-color: #45a049; /* Color más oscuro al pasar el mouse */
  }
  
  .download-button i {
    margin-right: 5px; /* Espacio entre ícono y texto */
  }
  </style>
  
  
  