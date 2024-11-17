<template>
  <link rel="stylesheet" href="https://cdnjs.cloudflare.com/ajax/libs/font-awesome/6.0.0-beta3/css/all.min.css">
  <div class="background-container"> <!-- Contenedor para el fondo -->
  <div class="facture-info-container">
    <h2 class="digital-font">Historial de facturas</h2>
    <!--h5>Correo: {{ userEmail }} </h5-->
    <!--h5>ID: {{ userID }}</h5-->
    <h3 class="digital-font">Cliente: {{ customerName }}, Correo: {{ customerEmail }} </h3>
    <!--h5>ID: {{ customerId }} </h5-->
    <table>
      <thead>
        <tr>
          <th class="digital-font">Número de Factura</th>
          <th class="digital-font">Medidor Anterior</th>
          <th class="digital-font">Medidor Actual</th>
          <th class="digital-font">Consumo</th>
          <th class="digital-font">Cod. de factura</th>
          <th class="digital-font">Comienzo de Periodo</th>
          <th class="digital-font">Final de Periodo</th>
          <th class="digital-font">Fecha limite de Pago</th>
          <th class="digital-font">Total a Pagar</th>
          <th class="digital-font">Estado</th>
          <th class="digital-font">Factura PDF</th>
          <th class="digital-font">Enviar por Correo</th> <!-- Nueva columna -->
          <th class="digital-font">Acciones</th>
        </tr>
      </thead>
      <tbody>
        <tr v-for="factura in facturas" :key="factura.idfacture">
          <td class="digital-font">{{ factura.facturenumber }}</td>
          <td class="digital-font">{{ factura.meterbefore }}</td>
          <td class="digital-font">{{ factura.meterafter }}</td>
          <td class="digital-font">{{ factura.consumer }}</td>
          <td class="digital-font">{{ factura.codfacture }} </td>
          <td class="digital-font">{{ factura.datecreation }}</td>
          <td class="digital-font">{{ factura.datepayment }}</td>
          <td class="digital-font">{{ factura.datelimit }}</td>
          <td class="digital-font">${{ factura.totalpay }}</td>
          <td>
           <select v-model="factura.statusfacture" @change="updateFactureStatus(factura)" class="form-select">
           <option value="Pendiente">Pendiente</option>
           <option value="Pagada">Pagada</option>
           <option value="Vencida">Vencida</option>
          </select>
          </td>
          <td>
            <button @click="downloadInvoice(factura)" class="download-button">
              <i class="fas fa-download"></i> Descargar
            </button>
          </td>
          <td>
            <button   class="btn-email" 
    @click="sendInvoiceByEmail(factura)" 
    :disabled="factura.statusfacture === 'Pagada' || factura.statusfacture === 'Vencida'"
    :class="{'disabled-button': factura.statusfacture === 'Pagada' || factura.statusfacture === 'Vencida'}"
  >
    <i class="fas fa-envelope"></i> Enviar por correo
            </button>
          </td>
          <td>
         <button class="btn-delete" @click="deleteFacture(factura.idfacture)"><i class="fas fa-trash-alt"></i>Eliminar</button>
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
    <div class="invoice-buttons">
      <button class="new-invoice-button" @click="viewFactureNew(customerId, customerName)">Nueva factura</button>
      <button class="back-button" type="button" @click="goBack">Regresar</button>
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
      customerLastName: this.$route.query.last_name,
      customerCedula : this.$route.query.cedula,
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
    this.userNit = localStorage.getItem('userNit');
  },
  methods: {
    async fetchAllFacturas() {
      try {
        const response = await axios.get(`/api/facturas`);
        // Filtrar facturas por customerId en el frontend
        this.facturas = response.data
      .filter((factura) => factura.idcustomer === parseInt(this.customerId))
      .sort((a, b) => b.facturenumber - a.facturenumber); // Ordena de mayor a menor   
          
      } catch (error) {
        console.error('Error fetching facturas:', error);
        alert('Error al obtener las facturas.');
      }
    },

    async updateFactureStatus(factura) {
      try {
        await axios.put(`/api/facturas/${factura.idfacture}`, {
          statusfacture: factura.statusfacture,
        });
        alert('Estado de la factura actualizado.');
      } catch (error) {
        console.error('Error updating facture status:', error);
        alert('Error al actualizar el estado de la factura.');
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

    async downloadInvoice(factura) {
      try {
        const doc = new jsPDF();

        // Obtener la imagen de fondo con opacidad 0.2
        const backgroundUrl = '/img/logoTG.png'; // Ruta relativa desde la carpeta public
        const background = await this.getImageDataUrl(backgroundUrl, 0.2); // 0.2 = 20% opacidad

        const pageWidth = doc.internal.pageSize.getWidth();
        const pageHeight = doc.internal.pageSize.getHeight();

        // Añadir la imagen de fondo al PDF
        doc.addImage(background, 'PNG', 0, 0, pageWidth, pageHeight);

        // Ahora, agregar el contenido sobre la imagen de fondo
        // Configuración de fuentes y estilos
        doc.setFont('Courier', 'normal');

        // Encabezado
        doc.setFontSize(22);
        doc.setTextColor(0, 0, 0); // Negro
        doc.text(`Empresa Acueducto, ${this.userName}`, pageWidth / 2, 20, { align: 'center' }); // Nombre de la empresa centrado

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
        doc.text(`NIT: ${this.userNit}, Correo: ${this.userEmail}`, 120, 30);

        // Datos del cliente
        doc.setFontSize(12);
        doc.setTextColor(50); // Gris
        doc.text('INFORMACIÓN GENERAL', 10, 45);

        doc.setTextColor(0); // Negro
        doc.text(`Factura Número: ${factura.facturenumber}`, 10, 50);
        doc.text(`Fecha Límite de Pago: ${factura.datelimit}`, 10, 55);

        doc.setTextColor(50); // Gris
        doc.text('Datos del Cliente', pageWidth - 88, 45, { align: 'center' });

        doc.setTextColor(0); // Negro
        doc.text(`Nombre: ${this.customerName}${this.customerLastName}`, 100, 50);
        doc.text(`N° Cedula: ${this.customerCedula}`, 100, 55);

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

    async sendInvoiceByEmail(factura) {
      this.sendingEmail = true; // Iniciar indicador de carga
      try {
        const doc = new jsPDF();

        // Obtener la imagen de fondo con opacidad 0.2
        const backgroundUrl = '/img/logoTG.png'; // Ruta relativa desde la carpeta public
        const background = await this.getImageDataUrl(backgroundUrl, 0.2); // 0.2 = 20% opacidad

        const pageWidth = doc.internal.pageSize.getWidth();
        const pageHeight = doc.internal.pageSize.getHeight();

        // Añadir la imagen de fondo al PDF
        doc.addImage(background, 'PNG', 0, 0, pageWidth, pageHeight);

        // Ahora, agregar el contenido sobre la imagen de fondo
        // Configuración de fuentes y estilos
        doc.setFont('Courier', 'normal');

        // Encabezado
        doc.setFontSize(22);
        doc.setTextColor(0, 0, 0); // Negro
        doc.text(`Empresa Acueducto, ${this.userName}`, pageWidth / 2, 20, { align: 'center' }); // Nombre de la empresa centrado

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
        doc.text(`NIT: ${this.userNit}, Correo: ${this.userEmail}`, 120, 30);

        // Datos del cliente
        doc.setFontSize(12);
        doc.setTextColor(50); // Gris
        doc.text('INFORMACIÓN GENERAL', 10, 45);

        doc.setTextColor(0); // Negro
        doc.text(`Factura Número: ${factura.facturenumber}`, 10, 50);
        doc.text(`Fecha Límite de Pago: ${factura.datelimit}`, 10, 55);

        doc.setTextColor(50); // Gris
        doc.text('Datos del Cliente', pageWidth - 88, 45, { align: 'center' });

        doc.setTextColor(0); // Negro
        doc.text(`Nombre: ${this.customerName}${this.customerLastName}`, 100, 50);
        doc.text(`N° Cedula: ${this.customerCedula}`, 100, 55);

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

        // Obtener el PDF como Blob
        const pdfBlob = doc.output('blob');
        
        
        // Crear el segundo PDF
       const doc2 = new jsPDF();
       doc2.setFont('Courier', 'normal');
       doc2.setFontSize(15);
       doc.setTextColor(0, 0, 0); // Negro
       doc2.text('Pasos detallados para realizar el Pago Correspondiente.', 10, 10);
       // Línea divisoria en color azul agua
       doc2.setFontSize(8);
       doc2.text('Para efectuar el pago en la Plataforma web http://localhost:8080/ se debe estar logueado con un correo electronico', 10, 20);
       doc2.text('asignado por la Entidad asociada y especificar que tipo de usuario eres, en este caso eres un "Usuario Cliente"', 10, 25);
       // Línea divisoria en color azul agua
       
      // Obtener la imagen para el segundo PDF     
      const imageUrl = '/img/paso1.png'; // Ruta de la nueva imagen
      const image = await this.getImageDataUrl(imageUrl); // Obtener la imagen como Data URL

    // Especificar el tamaño y la posición de la imagen
      const imgWidth = 110; // Ancho de la imagen (puedes ajustar este valor)
      const imgHeight = 43; // Alto de la imagen (puedes ajustar este valor)

    // Posición de la imagen en el PDF (x, y)
      const imgX = 10;  // Posición horizontal
      const imgY = 30;  // Posición vertical

    // Añadir la imagen en el segundo PDF
      doc2.addImage(image, 'PNG', imgX, imgY, imgWidth, imgHeight);

    // Agregar contenido adicional al segundo PDF (puedes seguir agregando más texto)
     
       // Obtener la imagen para el segundo PDF     
      const imageUrl2 = '/img/userimage.png'; // Ruta de la nueva imagen
      const image2 = await this.getImageDataUrl(imageUrl2); // Obtener la imagen como Data URL

    // Especificar el tamaño y la posición de la imagen
      const imgWidth2 = 70; // Ancho de la imagen (puedes ajustar este valor)
      const imgHeight2 = 15; // Alto de la imagen (puedes ajustar este valor)

    // Posición de la imagen en el PDF (x, y)
      const imgX2 = 130;  // Posición horizontal
      const imgY2 = 30;  // Posición vertical
      // Añadir la imagen en el segundo PDF
      doc2.addImage(image2, 'PNG', imgX2, imgY2, imgWidth2, imgHeight2);
      
      doc2.text('Al momento de iniciar sesion se le mostrara una tabla con información importante sobre facturas anteriores y la', 10, 83);
      doc2.text('actual teniendo en cuenta las fechas especificadas.', 10, 87);
      doc2.text('Al momento de hacer el pago se mostrará el estado (Pendiente, Pagado o vencido) en el que se encuentra la factura y ', 10, 91);
      doc2.text('se mostrará habilidado el boton de "Proceso de Pago" para su respectivo pago de factura.', 10, 95);

      const imageUrl3 = '/img/paso2.png'; // Ruta de la nueva imagen
      const image3 = await this.getImageDataUrl(imageUrl3); // Obtener la imagen como Data URL

    // Especificar el tamaño y la posición de la imagen
      const imgWidth3 = 125; // Ancho de la imagen (puedes ajustar este valor)
      const imgHeight3 = 50; // Alto de la imagen (puedes ajustar este valor)

    // Posición de la imagen en el PDF (x, y)
      const imgX3 = 10;  // Posición horizontal
      const imgY3 = 100;  // Posición vertical
      // Añadir la imagen en el segundo PDF
      doc2.addImage(image3, 'PNG', imgX3, imgY3, imgWidth3, imgHeight3);

      const imageUrl4 = '/img/paso3.png'; // Ruta de la nueva imagen
      const image4 = await this.getImageDataUrl(imageUrl4); // Obtener la imagen como Data URL

    // el tamaño y la posición de la imagen
      const imgWidth4 = 135; // Ancho de la imagen (puedes ajustar este valor)
      const imgHeight4 = 10; // Alto de la imagen (puedes ajustar este valor)

    // Posición de la imagen en el PDF (x, y)
      const imgX4 = 10;  // Posición horizontal2
      const imgY4 = 152;  // Posición vertical
      // Añadir la imagen en el segundo PDF
      doc2.addImage(image4, 'PNG', imgX4, imgY4, imgWidth4, imgHeight4);

      doc2.text('Al momento de darle click a "Proceso de Pago" se va a mostrar una ventana la cual tendrá el valor correspondiente', 10, 168);
      doc2.text('a pagar y con una referencia de pago el cual sera unico para su recibo.', 10, 171);
      doc2.text('Encontrará un botón el cual dirá "Confirmar Pago" para habilitar un segundo botón que dirá "Pagar con Mercado pago".', 10, 175);
      
      
      const imageUrl5 = '/img/paso4.png'; // Ruta de la nueva imagen
      const image5 = await this.getImageDataUrl(imageUrl5); // Obtener la imagen como Data URL

    // el tamaño y la posición de la imagen
      const imgWidth5 = 100; // Ancho de la imagen (puedes ajustar este valor)
      const imgHeight5 = 35; // Alto de la imagen (puedes ajustar este valor)

    // Posición de la imagen en el PDF (x, y)
      const imgX5 = 10;  // Posición horizontal2
      const imgY5 = 177;  // Posición vertical
      // Añadir la imagen en el segundo PDF
      doc2.addImage(image5, 'PNG', imgX5, imgY5, imgWidth5, imgHeight5);

      //ultimo
      doc2.text('Lo redireccionará la plataforma de mercado pago en la cual podras efectuar el pago con el metodo que quieras.', 10, 220);
      
      const imageUrl6 = '/img/paso5.png'; // Ruta de la nueva imagen
      const image6 = await this.getImageDataUrl(imageUrl6); // Obtener la imagen como Data URL

    // el tamaño y la posición de la imagen
      const imgWidth6 = 100; // Ancho de la imagen (puedes ajustar este valor)
      const imgHeight6 = 35; // Alto de la imagen (puedes ajustar este valor)

    // Posición de la imagen en el PDF (x, y)
      const imgX6 = 10;  // Posición horizontal2
      const imgY6 = 225;  // Posición vertical
      // Añadir la imagen en el segundo PDF
      doc2.addImage(image6, 'PNG', imgX6, imgY6, imgWidth6, imgHeight6);

      const pdfBlob2 = doc2.output('blob'); // Exportar segundo PDF

        // Crear un FormData para enviar el archivo
        const formData = new FormData();
        formData.append('to', this.customerEmail);
        formData.append('subject', `Factura ${factura.facturenumber}`);
        formData.append(
          'body',
          `Aviso de Factura nueva del Servicio de Acueducto de la Entidad ${this.userName}, para su revisión y respectivo pago en la plataforma web, entrando desde el link https://www.localhost:/8080 usando el correo y contraseña correspondiente.`
        );
        formData.append(
          'pdf',
          pdfBlob,
          `Factura-${factura.facturenumber}.pdf`
        );
        formData.append('pdf2', pdfBlob2, 'Pasos_Pago.pdf');

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
        query: { name, email: this.customerEmail, cedula: this.customerCedula, last_name: this.customerLastName},
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
},
goBack() {
      this.$router.go(-1); // Regresa a la página anterior (cualquiera)
    }

  },
};
</script>

<style scoped>
body {
  background-color: #ffffff; /* Fondo sólido para toda la página */
  margin: 0;
  padding: 0;
  font-family: 'Roboto', sans-serif; /* Tipo de letra */
}

.facture-info-container {
  background-color: rgba(243, 243, 243, 0.9); /* Color de formulario semi-transparente */
  padding: 30px;
  margin: 90px auto;
  width: 80%; /* Hacer el contenedor más pequeño */
  border-radius: 10px;
  box-shadow: 0 8px 16px rgba(0, 0, 0, 0.2); /* Sombra más pronunciada */
  position: relative; /* Para colocar la sombra en la parte trasera */
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
  background-color: #f2f2f2; /* Color de fondo del encabezado */
}


.invoice-buttons {
  display: flex;
  justify-content: center; /* Centrar los botones */
  margin-top: 20px; /* Espacio superior para los botones */
}

.new-invoice-button, .back-button {
  padding: 12px 27px; /* Aumentar tamaño */
  background-color: #62b5ec;
  color: white;
  border: none;
  border-radius: 20px; /* Bordes redondeados */
  box-shadow: 0 4px 10px rgba(0, 0, 0, 0.2); /* Sombra */
  font-size: 16px; /* Aumentar tamaño de fuente */
  cursor: pointer;
  transition: background-color 0.3s, transform 0.2s; /* Transición suave */
  margin: 0 10px; /* Espacio entre botones */
}

.new-invoice-button:hover, .back-button:hover {
  background-color: #2980b9; /* Color al pasar el mouse */
  transform: scale(1.05); /* Efecto de aumento */
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
  background-color: #007bff; /* Azul predeterminado */
  color: white;
  border: none;
  padding: 8px 12px;
  border-radius: 4px;
  cursor: pointer;
  transition: background-color 0.3s;
}

.btn-email:hover {
  background-color: #0056b3; /* Azul más oscuro */
}

.btn-email.disabled-button {
  background-color: #cccccc; /* Gris para el estado deshabilitado */
  color: #666666; /* Texto gris claro */
  cursor: not-allowed; /* Icono de bloqueo */
  cursor: not-allowed; /* Cambiar cursor cuando está deshabilitado */
}


.btn-disabled {
  background-color: #ccc !important; /* Gris */
  color: #666 !important; /* Gris oscuro */
  cursor: not-allowed !important; /* Cursor de no permitido */
  pointer-events: none; /* Deshabilita el clic */
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
  background: rgba(0, 0, 0, 0.5); /* Fondo oscuro con opacidad */
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

.form-select {
  width: 150px; /* Ajusta el ancho según sea necesario */
  border-radius: 5px;
  background-color: #f8f9fa; /* Color de fondo claro */
  border: 1px solid #ced4da; /* Color del borde */
  transition: border-color 0.2s ease-in-out;
}

.form-select:hover {
  border-color: #80bdff; /* Cambia el color del borde al pasar el mouse */
}

.form-select:focus {
  border-color: #0056b3; /* Color del borde al hacer foco */
  box-shadow: 0 0 0 0.2rem rgba(0, 123, 255, 0.25); /* Sombra al hacer foco */
}

.digital-font {
  font-family: 'Digital', sans-serif; /* Asegúrate de que la fuente esté disponible en tu proyecto */
  font-size: 0.8em; /* Ajusta el tamaño como desees */
}

.background-container {
  background-image: url('https://cdn.leonardo.ai/users/65a8cf55-c959-4394-91b9-30d6f5167b8c/generations/6a0f6c27-9594-4c4d-8e92-8c60e0a33d2d/Leonardo_Phoenix_A_realistic_depiction_of_a_cluttered_workspac_3.jpg');
  background-size: cover; /* Ajusta la imagen para cubrir todo el contenedor */
  background-position: center; /* Centra la imagen */
  background-repeat: no-repeat; /* Evita que la imagen se repita */
  height: 100vh; /* O ajusta a la altura deseada */
  display: flex;
  justify-content: center; /* Centrar el contenido */
  align-items: center; /* Centrar verticalmente */
}


</style>

