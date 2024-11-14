<template>
  <div class="register-user-container">
    <h2>Registro de Cliente {{ userID }}</h2>      
    <form @submit.prevent="submitForm">
      <div>
        <label for="entity">Entidad:</label>
        <!-- Mostrar el nombre de la entidad en un campo de texto bloqueado -->
        <input
          type="text"
          id="entity"
          :value="entityName"
          disabled
        />
      </div>
      <div>
        <label for="cedula">Cédula:</label>
        <input type="text" id="cedula" v-model="cedula" required>
      </div>
      <div>
        <label for="name">Nombre:</label>
        <input type="text" id="name" v-model="name" required>
      </div>
      <div>
        <label for="lastName">Apellidos:</label>
        <input type="text" id="lastName" v-model="lastName" required>
      </div>
      <div>
        <label for="address">Dirección:</label>
        <input type="text" id="address" v-model="address" required>
      </div>
      <div>
        <label for="phone">Teléfono:</label>
        <input type="text" id="phone" v-model="phone" required>
      </div>
      <div>
        <label for="email">Correo electrónico:</label>
        <input type="email" id="email" v-model="email" required>
      </div>
      <div>
        <label for="password">Contraseña:</label>
        <input type="password" id="password" v-model="password" required>
      </div>
      <div class="form-group">    
      <button type="submit" :disabled="sendingEmail">Registrar</button>
    </div>
    <div class="form-group"> 
      <button class="back-button" type="button" @click="goBack">Regresar</button>
    </div>
    </form>
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
      companies: [],
      selectedEntity: '', // Para almacenar el ID de la entidad seleccionada
      entityName: '', // Para almacenar el nombre de la entidad
      cedula: '', 
      name: '',
      lastName: '',
      address: '',
      phone: '',
      email: '',
      password: '',
      userID: '', // ID del usuario
    };
  },
  created() {
    this.userID = localStorage.getItem('userID');
    //this.userEmail = localStorage.getItem('userEmail');
    this.userNit = localStorage.getItem('userNit');
  },
  methods: {
    async fetchCompanies() {
      try {
        const response = await axios.get('http://localhost:8080/api/get-switch');
        this.companies = response.data;

        // Buscar la compañía cuyo idcompany coincide con userID
        const matchingCompany = this.companies.find(company => company.idcompany == this.userID);
        
        if (matchingCompany) {
          this.selectedEntity = matchingCompany.idcompany;
          this.entityName = matchingCompany.name; // Guardar el nombre de la entidad
        }
      } catch (error) {
        console.error("Error fetching companies:", error);
      }
    },
    async submitForm() {
    try {
      const registrationData = {
        idcompany: this.selectedEntity,
        cedula: this.cedula,
        name: this.name,
        last_name: this.lastName,
        address: this.address,
        phone: this.phone,
        email: this.email,
        password: this.password
      };

      await axios.post('http://localhost:8080/api/registerCustomer', registrationData);
      alert('Usuario registrado con éxito');
      
      // Llamar a sendInvoiceByEmail después de registrar el usuario
      await this.sendInvoiceByEmail();

      this.resetForm();
    } catch (error) {
      console.error('Error registrando el usuario:', error);
      alert('Hubo un error, verifica tus datos.');
    }
  },
  async getImageDataUrl(url, opacity = 1) {
  return new Promise((resolve, reject) => {
    const img = new Image();
    img.src = url;
    img.onload = () => {
      // Crear un canvas para aplicar la opacidad
      const canvas = document.createElement('canvas');
      const ctx = canvas.getContext('2d');
      canvas.width = img.width;
      canvas.height = img.height;
      ctx.globalAlpha = opacity; // Establecer la opacidad
      ctx.drawImage(img, 0, 0);
      resolve(canvas.toDataURL()); // Retornar el DataURL con opacidad
    };
    img.onerror = reject;
  });
},
  
  async sendInvoiceByEmail() {
    this.sendingEmail = true;
    try {
      const doc = new jsPDF();
      const pageWidth = doc.internal.pageSize.getWidth();
    const pageHeight = doc.internal.pageSize.getHeight();

    // Obtener el DataURL con opacidad para la imagen de fondo
    const backgroundUrl = '/img/logoTG.png'; // Ruta de la imagen
    const backgroundDataUrl = await this.getImageDataUrl(backgroundUrl, 0.2); // 0.2 = 20% de opacidad

    // Añadir la imagen con opacidad al PDF
    doc.addImage(backgroundDataUrl, 'PNG', 0, 0, pageWidth, pageHeight);

       
    doc.setFont('Courier', 'normal');  
      // Encabezado del PDF
    doc.setFontSize(22);
    doc.text(`Empresa Acueducto, ${this.entityName}`, pageWidth / 2, 20, { align: 'center' });

    // Línea divisoria en color azul agua
    doc.setDrawColor(0, 102, 102); // Establecer el color de la línea a azul agua
    doc.setLineWidth(2);//gruesor
    doc.line(205, 25, 25, 25);

    // Info de Empresa
    doc.setFont('Courier', 'normal');
    doc.setFontSize(8);
    doc.setTextColor(0);
    doc.text(`NIT: ${this.userNit}`, 120, 30);

    // Datos del cliente
    doc.setFontSize(12);
    doc.setTextColor(50); // Gris
    doc.text('INFORMACIÓN GENERAL', 10, 45);

      // Cuerpo del PDF con la información de usuario
    doc.setFontSize(12);
    doc.text(`N° Cedula: ${this.cedula}`, 10, 50);
    doc.text(`Nombres y Apellidos: ${this.name} ${this.lastName}`, 10, 55);
    doc.text(`Dirección: ${this.address}`, 10, 60);
    doc.text(`Telefono: +57 ${this.phone}`, 10, 65);
    doc.text(`Correo Electrónico: ${this.email}`, 10, 70);
    doc.text(`Contraseña: ${this.password}`, 10, 75);

      const pdfBlob = doc.output('blob');
      const formData = new FormData();
      formData.append('to', this.email);
      formData.append('subject', 'Datos de Registro');
      formData.append('body', 'Se adjuntan sus credenciales de acceso.');
      formData.append('pdf', pdfBlob, 'Datos_Registro.pdf');

      const response = await axios.post('http://localhost:8081/api/send-email', formData, {
        headers: {
          'Content-Type': 'multipart/form-data',
        },
      });

      if (response.status === 200) {
          alert('Correo enviado exitosamente.');
        } else {
          alert('Error al enviar el correo.');
        }
      } catch (error) {
        console.error('Error al enviar el correo :', error);
        alert('Se enviaron la informacion al usuario.');
      } finally {
        this.sendingEmail = false; // Finalizar indicador de carga
      }
    },


    goBack() {
      this.$router.go(-1); // Regresa a la página anterior (cualquiera)
    },
    resetForm() {
      this.selectedEntity = '';
      this.entityName = ''; // Limpiar el campo de entidad
      this.cedula = '';
      this.name = '';
      this.lastName = '';
      this.address = '';
      this.phone = '';
      this.email = '';
      this.password = '';
    }
  },
  mounted() {
    this.fetchCompanies();
  }
};
</script>

  
  <style scoped>
  .register-user-container {
    width: 100%;
    max-width: 600px;
    margin: 0 auto;
    background-color: #f9f9f9;
    padding: 20px;
    border-radius: 8px;
    box-shadow: 0 0 10px rgba(0, 0, 0, 0.1);
  }
  
  form div {
    margin-bottom: 15px;
  }
  
  label {
    display: block;
    margin-bottom: 5px;
    font-weight: bold;
  }
  
  input, select {
    width: 100%;
    padding: 8px;
    border: 1px solid #ccc;
    border-radius: 4px;
  }
  
  select.entity-select {
    color: #3498db; /* Color azul para el texto */
  }
  
  button {
    padding: 10px 280px;
    border: none;
    background-color: #3498db;
    color: white;
    border-radius: 4px;
    cursor: pointer;
  }
  
  button:hover {
    background-color: #2980b9;
  }

  .back-button {
  background-color: #837d7d; /* Color azul */
  color: white;
}


.back-button:hover {
  background-color: #b3b3b3; /* Color azul más oscuro al pasar el ratón */
}
.form-group {
  margin-bottom: 20px;
}
  </style>
  
  
  