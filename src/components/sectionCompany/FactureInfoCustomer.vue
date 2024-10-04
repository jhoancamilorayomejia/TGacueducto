<template>
  <div class="facture-info-container">
    <h2>Historial de facturas</h2>
    <h5>Correo: {{ userEmail }}, ID: {{ userID }}</h5>
    <h3>Cliente: {{ customerName }}, ID: {{ customerId }}, Correo: {{ customerEmail }}</h3>
    <div class="invoice-buttons">
      <button @click="viewFactureNew(customerId, customerName)">Nueva factura</button>
    </div>
    <table>
      <thead>
        <tr>
          <th>Número de Factura</th>
          <th>Medidor Anterior</th>
          <th>Medidor Actual</th>
          <th>Consumo</th>
          <th>Fecha de Emisión</th>
          <th>Fecha Límite de Pago</th>
          <th>Total a Pagar</th>
          <th>Factura PDF</th>
          <th>Enviar por Correo</th> <!-- Nueva columna -->
          <th>Acciones</th>
        </tr>
      </thead>
      <tbody>
        <tr v-for="factura in facturas" :key="factura.idfacture">
          <td>{{ factura.facturenumber }}</td>
          <td>{{ factura.meterbefore }}</td>
          <td>{{ factura.meterafter }}</td>
          <td>{{ factura.consumer }}</td>
          <td>{{ factura.datecreation }}</td>
          <td>{{ factura.datepayment }}</td>
          <td>${{ factura.totalpay }}</td>
          <td>
            <button @click="downloadInvoice(factura)" class="download-button">
              <i class="fas fa-download"></i> Descargar
            </button>
          </td>
          <td>
            <button class="btn-email" @click="sendInvoiceByEmail(factura)" :disabled="sendingEmail">
              <i class="fas fa-envelope"></i> Enviar por correo
            </button>
          </td>
          <td>
         <button class="btn-delete" @click="deleteFacture(factura.idfacture)">Eliminar</button>
         </td>

        </tr>
      </tbody>
    </table>

    <!-- Modal de Carga (Opcional) -->
    <div v-if="sendingEmail" class="modal-overlay">
      <div class="modal-content">
        <p>Enviando correo, por favor espera...</p>
      </div>
    </div>
  </div>
</template>

<script>
import axios from 'axios';
import jsPDF from 'jspdf';

export default {
  data() {
    return {
      userEmail: '',
      userID: '',
      userName: '',
      customerId: this.$route.params.idcustomer,
      customerName: this.$route.query.name,
      customerEmail: this.$route.query.email,
      facturas: [],
      sendingEmail: false, // Estado para manejar la carga
    };
  },
  created() {
    this.fetchAllFacturas();
    this.userEmail = localStorage.getItem('email');
    this.userID = localStorage.getItem('userID');
    this.userName = localStorage.getItem('userName');
  },
  methods: {
    async fetchAllFacturas() {
      try {
        const response = await axios.get(`/api/facturas`);
        // Filtrar facturas por customerId en el frontend
        this.facturas = response.data.filter(
          (factura) => factura.idcustomer === parseInt(this.customerId)
        );
      } catch (error) {
        console.error('Error fetching facturas:', error);
        alert('Error al obtener las facturas.');
      }
    },
    downloadInvoice(factura) {
      const doc = new jsPDF();
      // Configuración de fuentes y estilos
      doc.setFont('Helvetica', 'normal');

      // Encabezado de la empresa
      doc.setFontSize(15);
      doc.setTextColor(35);
      doc.text(`Entidad Presatadora de Servicio Publico: ${this.userName}`, 105, 20, null, null, 'center'); // Nombre de la empresa centrado
      
       // Línea divisoria
      doc.setLineWidth(0.3);
      doc.line(10, 45, 120, 45);

      // Título de la factura
      doc.setFontSize(16);
      doc.setTextColor(0);
      doc.text('Factura de Servicio de Agua', 55, 35, null, null, 'center');


      doc.setFontSize(12);
      doc.text(`Factura Número: ${factura.facturenumber}`, 10, 55);
      doc.text(`Fecha de Emisión: ${factura.datecreation}`, 10, 65);
      doc.text(`Fecha Límite de Pago: ${factura.datepayment}`, 10, 75);
      doc.text(`Medidor actual: ${factura.meterafter}`, 10, 85);
      doc.text(`Medidor anterior: ${factura.meterbefore}`, 10, 95);
      doc.text(`Medidor anterior: ${factura.consumer}`, 10, 105);
      doc.text(`Monto Total: $${factura.totalpay}`, 10, 115);
      doc.text(
        `Cliente: ${this.customerName} (ID: ${this.customerId})`,
        10,
      125
      );
      doc.text(`Correo: ${this.userEmail}`, 10, 135);
      doc.save(`Factura_${factura.facturenumber}.pdf`);

      // Línea divisoria
      doc.line(20, 85, 190, 85);
      
      

    },
    async sendInvoiceByEmail(factura) {
      this.sendingEmail = true; // Iniciar indicador de carga
      try {
        // Generar el PDF
        const doc = new jsPDF();
        doc.setFontSize(12);
        doc.text(`Factura Número: ${factura.facturenumber}`, 10, 55);
        doc.text(`Fecha de Emisión: ${factura.datecreation}`, 10, 65);
        doc.text(`Fecha Límite de Pago: ${factura.datepayment}`, 10, 75);
        doc.text(`Medidor actual: ${factura.meterafter}`, 10, 85);
      doc.text(`Medidor anterior: ${factura.meterbefore}`, 10, 95);
      doc.text(`Medidor anterior: ${factura.consumer}`, 10, 105);
        doc.text(`Monto Total: $${factura.totalpay}`, 10, 115);
        doc.text(
          `Cliente: ${this.customerName} (ID: ${this.customerId})`,
          10,
          125
        );
        doc.text(`Correo: ${this.userEmail}`, 10, 135);

        // Obtener el PDF como Blob
        const pdfBlob = doc.output('blob');

        // Crear un FormData para enviar el archivo
        const formData = new FormData();
        formData.append('to', this.customerEmail);
        formData.append('subject', `Factura ${factura.facturenumber}`);
        formData.append(
          'body',
          `Se adjunto la factura número ${factura.facturenumber} para su revisión.`
        );
        formData.append(
          'pdf',
          pdfBlob,
          `Factura_${factura.facturenumber}.pdf`
        );

        // Enviar la solicitud al backend
        const response = await axios.post(
          'http://localhost:8081/api/send-email',
          formData,
          {
            headers: {
              'Content-Type': 'multipart/form-data',
            },
          }
        );

        if (response.status === 200) {
          alert('Correo enviado exitosamente.');
        } else {
          alert('Error al enviar el correo.');
        }
      } catch (error) {
        console.error('Error al enviar el correo :', error);
        alert('Factura de Recibo de Acueducto enviada.');
      } finally {
        this.sendingEmail = false; // Finalizar indicador de carga
      }
    },
    viewFactureNew(idcustomer, name) {
      this.$router.push({
        path: `/api/customer/new-facture/${idcustomer}`,
        query: { name, email: this.customerEmail },
      });
    },
    async deleteFacture(idfacture) {
  const confirmDelete = confirm("¿Estás seguro de que deseas eliminar la Factura?");
  if (confirmDelete) {
    try {
      const token = localStorage.getItem('token');
      await axios.delete(`/api/facture/${idfacture}`, {
        headers: {
          Authorization: `Bearer ${token}`
        }
      });
      this.fetchAllFacturas(); // Cambiar fetchUsuarios a fetchAllFacturas para actualizar la lista
    } catch (error) {
      console.error('Error al eliminar factura:', error);
      alert('Hubo un error al intentar eliminar la factura.');
    }
  }
}

  },
};
</script>

<style scoped>
.facture-info-container {
  background-color: #f3f3f3;
  padding: 80px; /*para la margen de lo gris */
  margin: 20px auto;
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
  padding: 5px;  
  text-align: center;
}

th {
  background-color: #f2f2f2;
}

.invoice-buttons {
  display: flex;
  gap: 10px;
  margin-bottom: 20px;
}

button {
  padding: 5px 10px;
  background-color: #037bca;
  color: white;
  border: none;
  border-radius: 4px;
  cursor: pointer;
}

.download-button {
  background-color: #4caf50; /* Color verde para el botón */
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

.btn-email {
  background-color: #f39c12; /* Color naranja para el botón */
  color: white;
  padding: 5px 10px;
  border: none;
  border-radius: 4px;
  cursor: pointer;
  display: flex;
  align-items: center;
  justify-content: center;
}

.btn-email:hover {
  background-color: #e67e22;
}

.btn-email:disabled {
  background-color: #d35400;
  cursor: not-allowed;
}

.btn-email i {
  margin-right: 5px; /* Espacio entre ícono y texto */
}

.btn-delete {
  background-color: #e74c3c;
  color: white;
  padding: 5px 10px;
  border: none;
  border-radius: 4px;
  cursor: pointer;
}

.btn-delete:hover {
  background-color: #c0392b;
}

/* Estilos para el Modal de Carga */
.modal-overlay {
  position: fixed;
  top: 0;
  left: 0;
  width: 100%;
  height: 100%;
  background-color: rgba(0, 0, 0, 0.5);
  display: flex;
  justify-content: center;
  align-items: center;
  z-index: 1000;
}

.modal-content {
  background-color: white;
  padding: 20px 40px;
  border-radius: 8px;
  text-align: center;
  font-size: 1.2em;
  box-shadow: 0px 4px 10px rgba(0, 0, 0, 0.3);
}
</style>

