<template>
  <link rel="stylesheet" href="https://cdnjs.cloudflare.com/ajax/libs/font-awesome/6.0.0-beta3/css/all.min.css">
  <div class="background-container"> <!-- Contenedor para el fondo -->
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
          <td>{{ factura.codfacture }}</td>
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
    <div class="button-container">
      <button class="button-allcustomerEmail" @click="AllCusEmail">Cambio de Clave</button>
    </div>
  </div>
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
        alert('Error al obtener las facturas.');
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
          (company) => company.idcompany === parseInt(this.userFKcompany)
        );
        //this.companies = response.data;
      } catch (error) {
        console.error('Error fetching companies:', error);
        alert('Error al obtener las empresas.');
      }
    },

    // Función auxiliar para convertir una imagen a base64 con opacidad ajustable
    getImageDataUrl(url, opacity = 1) {
      return new Promise((resolve, reject) => {
        const img = new Image();
        img.setAttribute('crossOrigin', 'anonymous'); // Evita problemas de CORS
        img.onload = () => {
          const canvas = document.createElement('canvas');
          canvas.width = img.width;
          canvas.height = img.height;
          const ctx = canvas.getContext('2d');

          // Establecer la opacidad
          ctx.globalAlpha = opacity;

          // Dibujar la imagen con la opacidad establecida
          ctx.drawImage(img, 0, 0, canvas.width, canvas.height);

          // Convertir el contenido del canvas a Data URL
          const dataURL = canvas.toDataURL('image/png');
          resolve(dataURL);
        };
        img.onerror = (error) => reject(error);
        img.src = url;
      });
    },

    // Función para descargar una factura en formato PDF
    async downloadInvoice(factura, company) {
      try {
        const doc = new jsPDF();

        // Obtener la imagen de fondo con opacidad 0.2
        const backgroundUrl = '/img/logoTG.png'; // Ruta relativa desde la carpeta public
        const background = await this.getImageDataUrl(backgroundUrl, 0.2); // 0.2 = 20% opacidad

        const pageWidth = doc.internal.pageSize.getWidth();
        const pageHeight = doc.internal.pageSize.getHeight();

        // Añadir la imagen de fondo al PDF
        doc.addImage(background, 'PNG', 0, 0, pageWidth, pageHeight);

        // Configuración de fuentes y estilos
        doc.setFont('Courier', 'normal');

        // Encabezado
        doc.setFontSize(25);
        doc.setTextColor(0, 0, 0); // Negro
        doc.text(`Empresa acueducto, ${company.Name}`, pageWidth / 2, 20, { align: 'center' }); // Nombre de la empresa centrado

        // Línea divisoria en color azul agua
        doc.setDrawColor(0, 102, 102); // Establecer el color de la línea a azul agua
        doc.setLineWidth(2);//gruesor
        doc.line(205, 25, 25, 25);

        doc.setFontSize(15);
        doc.setTextColor(0);
        doc.text('Factura de Agua', 30, 10, null, null, 'center');

        // Título de la factura en negrita
        doc.setFont('Courier', 'bold');
        doc.setFontSize(11);
        doc.text(`Pago de Referencia: #${factura.codfacture}`, pageWidth - 30, 10, { align: 'right' });

        // Info de Empresa
        doc.setFont('Courier', 'normal');
        doc.setFontSize(8);
        doc.setTextColor(0);
        doc.text(`NIT: ${company.nit}, Correo: ${company.Email}`, 118, 30);

        // Datos del cliente
        doc.setFontSize(12);
        doc.setTextColor(50); // Gris
        doc.text('INFORMACIÓN GENERAL', 10, 45);

        doc.setTextColor(0); // Negro
        doc.text(`Factura Número: ${factura.facturenumber}`, 10, 50);
        doc.text(`Fecha Límite de Pago: ${factura.datelimit}`, 10, 55);

        doc.setTextColor(50); // Gris
        doc.text('Datos del Cliente', 126, 45, { align: 'center' });

        doc.setTextColor(0); // Negro
        doc.text(`Nombre: ${this.userName} ${this.userLastName}`, 104, 50);
        doc.text(`N° Cedula: ${this.userCedula}`, 104, 55);

        // Línea divisoria en color azul agua
        doc.setDrawColor(0, 102, 102); // Establecer el color de la línea a azul agua
        doc.setLineWidth(0.5);//gruesor
        doc.line(205, 65, 65, 65);

        // Historial de consumo
        doc.setTextColor(50); // Gris
        doc.text('Historial de consumo', pageWidth / 2, 75, { align: 'center' });

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

        // Guardar el PDF
        doc.save(`Factura_${factura.facturenumber}.pdf`);
      } catch (error) {
        console.error('Error generando el PDF:', error);
        alert('Hubo un error al generar la factura.');
      }
    },
    AllCusEmail() {
      this.$router.push('/api/customersPassword');
    },

    goToPayment(totalPay, codfacture) {
      console.log("Monto total al hacer clic:", totalPay); // Imprimir para depurar
      this.$router.push({ 
        path: `/api/payment/${totalPay}`,
        query: { codfacture }
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
  background-color: rgba(243, 243, 243, 0.9); /* Color de formulario semi-transparente */
  padding: 30px;
  margin: 90px auto;
  width: 80%; /* Hacer el contenedor más pequeño */
  border-radius: 10px;
  box-shadow: 0 8px 16px rgba(0, 0, 0, 0.2); /* Sombra más pronunciada */
  position: relative; /* Para colocar la sombra en la parte trasera */
}

.welcome-header {
  display: flex;  
  justify-content: space-between; /* Alinea los elementos a los extremos */
  background-color: #b7daee;
  padding: 5px; /* Reducir el padding */
  border-radius: 15px;
  margin-bottom: 10px;
  font-size: 12px; /* Hacer el texto más pequeño */
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

/* Responsive Table */
@media (max-width: 768px) {
  table, thead, tbody, th, td, tr {
    display: block;
  }

  th, td {
    box-sizing: border-box;
    width: 100%;
    padding: 10px;
  }

  th {
    background-color: #f2f2f2;
    position: sticky;
    top: 0;
  }

  tr {
    margin-bottom: 15px;
  }
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

.button-container {
  display: flex;
  justify-content: center; /* Centrar el botón */
  margin-top: 20px; /* Espacio superior */
}

.button-allcustomerEmail {
  padding: 12px 27px; /* Aumentar tamaño */
  background-color: #62b5ec;
  color: white;
  border: none;
  border-radius: 20px; /* Bordes redondeados */
  box-shadow: 0 4px 10px rgba(0, 0, 0, 0.2); /* Sombra */
  font-size: 16px; /* Aumentar tamaño de fuente */
  cursor: pointer;
  transition: background-color 0.3s, transform 0.2s; /* Transición suave */
}

.button-allcustomerEmail:hover {
  background-color: #2980b9; /* Color al pasar el mouse */
  transform: scale(1.05); /* Efecto de aumento */
}
</style>
