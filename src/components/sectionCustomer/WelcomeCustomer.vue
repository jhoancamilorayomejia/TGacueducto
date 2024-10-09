<template>
  <div class="welcome-container">
    <div class="welcome-header">
      <h2>Sistema de Facturación Digital de Pequeños Acueductos</h2>
      <h5>Bienvenido, {{ userName }} {{ userLastName }} - N° Cédula: {{ userCedula }}</h5>
      <!--h4>{{ userFKcompany }}</h4-->
    </div>
    <table>
      <thead>
        <tr>
          <th>Número de Factura</th>
          <th>Medidor Anterior</th>
          <th>Medidor Actual</th>
          <th>Consumo</th>
          <th>Cod. de factura</th>
          <th>Comienzo de Periodo</th>
          <th>Final de Periodo</th>
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
          <td>{{ factura.meterbefore }}</td>
          <td>{{ factura.meterafter }}</td>
          <td>{{ factura.consumer }}</td>
          <td>{{ factura.codfacture }} </td>
          <td>{{ factura.datecreation }}</td>
          <td>{{ factura.datepayment }}</td>
          <td>{{ factura.datelimit }}</td>
          <td>${{ factura.totalpay }}</td>
          <td>{{ factura.statusfacture }}</td>
          <!--td>Pendiente/Pagado</td-->
          <td>
            <button @click="downloadInvoice(factura, companies[0])" class="download-button">
              <i class="fas fa-download"></i> Descargar
            </button>
          </td>
          <td>
            <button
              @click="goToPayment(factura.totalpay, factura.codfacture)"
              class="payment-button"
              :class="{ 'paid': factura.statusfacture === 'Pagada' || factura.statusfacture === 'Vencida' }"
              :disabled="factura.statusfacture === 'Pagada' || factura.statusfacture === 'Vencida'"
            >
              Proceso de pago
            </button>
          </td>
        </tr>
      </tbody>
    </table>
    <!--h3>Empresas Asociada</h3>
    <table>
      <thead>
        <tr>
          <th>ID Empresa</th>
          <th>Nombre de la Empresa</th>
          <th>NIT</th>
        </tr>
      </thead>
      <tbody>
        <tr v-for="company in companies" :key="company.idcompany">
          <td>{{ company.idcompany }}</td>
          <td>{{ company.Name }}</td>
          <td>{{ company.nit }}</td>
        </tr>
      </tbody>
    </table-->
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
      userFKcompany: '',
      userName: '', // Nombre del usuario
      userCedula: '', // Cédula del usuario
      userLastName: '',
      facturas: [], // Lista de facturas (todas)
      companies: []
    };
  },
  created() {
    this.userEmail = localStorage.getItem('email'); // Obtener el correo del localStorage
    this.userID = localStorage.getItem('userID'); // Obtener ID del usuario cliente
    this.userFKcompany = localStorage.getItem('userFKcompany'); // Obtener ID del company del usuario cliente
    this.userName = localStorage.getItem('userName'); // Obtener el nombre del usuario cliente
    this.userCedula = localStorage.getItem('userCedula'); // Obtener la cédula del usuario cliente
    this.userLastName = localStorage.getItem('userLastName'); // Obtener el apellido del usuario cliente
    
    //console.log('userLastName:', this.userLastName);

    this.fetchFacturas(); // Llama a la función para obtener facturas al crear el componente
    this.fetchCompanies();
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
    this.facturas = response.data
      .filter(factura => factura.idcustomer == this.userID)
      .sort((a, b) => b.facturenumber - a.facturenumber); // Ordenar por facturenumber de mayor a menor

      } catch (error) {
        console.error('Error fetching facturas:', error);
      }
    },

    async fetchCompanies() {
      try {
        const token = localStorage.getItem('token');
        const response = await axios.get('/api/AllCompany/', {
          headers: {
            Authorization: `Bearer ${token}`
          }
        });
        this.companies = response.data.filter(
          (companies) => companies.idcompany === parseInt(this.userFKcompany )
        );
        //this.companies = response.data;
      } catch (error) {
        console.error('Error fetching companies:', error);
      }
    },

    // Función para descargar una factura en formato PDF
    downloadInvoice(factura, company) {
      const doc = new jsPDF();
        // Configuración de fuentes y estilos
        doc.setFont('Courier', 'normal');

// Encabezado
doc.setFontSize(25);
doc.setTextColor(0, 0, 0); // Negro
doc.text(`Empresa acueducto,${company.Name}`, 130, 20, null, null, 'center'); // Nombre de la empresa centrado

// Línea divisoria en color azul agua
doc.setDrawColor(0, 102, 102); // Establecer el color de la línea a azul agua
doc.setLineWidth(2);//gruesor
doc.line(205, 25, 25, 25);

// Titulo
doc.setFontSize(15);
doc.setTextColor(0);
doc.text('Factura de Agua', 30, 10, null, null, 'center');

// Título de la factura en negrita
doc.setFont('Courier', 'bold');
doc.setFontSize(11);
doc.text(`Pago de Referencia: #${factura.codfacture}`, 170, 10, 'center');

// Configuración de fuentes y estilos
doc.setFont('Courier', 'normal');

//Info de Empresa
doc.setFontSize(8);
doc.text(`NIT: ${company.nit},`, 126, 30);
doc.text(`Correo: ${company.Email}`, 160, 30);


// Datos del cliente
doc.setFontSize(12);
doc.setTextColor(50); // Color gris para datos del cliente
doc.text('INFORMACIÓN GENERAL', 10, 45);

doc.setTextColor(0); // Restablecer el color a negro
doc.text(`Factura Número: ${factura.facturenumber}`, 10, 50);
doc.text(`Fecha Límite de Pago: ${factura.datelimit}`, 10, 55);

doc.setTextColor(50); // Color gris para datos del cliente
doc.text('Datos del Cliente', 129, 45, 'center');

doc.setTextColor(0); // Restablecer el color a negro
doc.text(`Nombre: ${this.userName} ${this.userLastName}`, 108, 50);
//doc.text(`${this.userLastName}`, 170, 50); 
doc.text(`N° Cedula: ${this.userCedula}`, 108, 55);


// Línea divisoria en color azul agua
doc.setDrawColor(0, 102, 102); // Establecer el color de la línea a azul agua
doc.setLineWidth(0.5);//gruesor
doc.line(205, 65, 65, 65);

doc.setTextColor(50); // Color gris para datos del cliente
doc.text('Historial de consumo', 36, 75, 'center');

doc.setTextColor(0); // Restablecer el color a negro
doc.text('Se establece un consumo de ', 10, 80);
doc.setFont('Courier', 'bold');
doc.text(`${factura.consumer}`, 78, 80);
doc.setFont('Courier', 'normal');
doc.text('Mt3 a partir del proceso de calcular:', 85, 80);
doc.text(`(Lectura actual: ${factura.meterafter}) - (Lectura anterior: ${factura.meterbefore}),Para un periodo comprendido`, 10, 85);
doc.text(`entre el ${factura.datecreation} y ${factura.datepayment}. Como observacion se tiene NORMAL con el`, 10, 90);
doc.text(`metodo de calculo DIFERENCIA LECTURAS`, 10, 95);

doc.setFont('Courier', 'bold');
doc.text('TOTALES DE COBRO', 30, 115, 'center');

doc.setFont('Courier', 'normal');
doc.text(`Total del mes: $${factura.totalpay}`, 10, 120);
//doc.text(`Cliente: ${this.customerName} (ID: ${this.customerId})`,10,125);

//doc.text(`ID: ${this.customerId}`, 10, 75);
      doc.save(`Factura_${factura.facturenumber}.pdf`);
    },

    goToPayment(totalPay, codfacture) {
  console.log("Monto total al hacer clic:", totalPay); // Imprimir para depurar
  this.$router.push({ path: `/api/payment/${totalPay}`,
    query: {codfacture}
   }); // Redirige a la vista de pago con el monto
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

.payment-button.paid {
  background-color: gray; /* Cambiar a gris cuando esté "Pagada" */
}

.payment-button:disabled {
  cursor: not-allowed; /* Cambiar cursor cuando está deshabilitado */
  opacity: 0.6; /* Añadir opacidad cuando está deshabilitado */
}
</style>

