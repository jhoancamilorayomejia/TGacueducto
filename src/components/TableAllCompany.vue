<template>
  <link rel="stylesheet" href="https://cdnjs.cloudflare.com/ajax/libs/font-awesome/6.0.0-beta3/css/all.min.css">
  <div class="background-container"> <!-- Contenedor para el fondo -->
  <div class="invoice-container">
    <div class="invoice-header">
      <h2>Área de Administrador / Empresas</h2>
    </div>
    <table class="usuario-table">
      <thead>
        <tr>
          <th>NIT</th>
          <th>Nombre de la Entidad</th>
          <th>Localidad</th>
          <th>Teléfono</th>
          <th>Email</th>
          <th>Información</th>
          <th>Acciones</th>
        </tr>
      </thead>
      <tbody>
        <tr v-for="company in companies" :key="company.idcompany">
          <td>{{ company.nit }}</td>
          <td>{{ company.Name }}</td>
          <td>{{ company.address }}</td>
          <td>{{ company.Phone }}</td>
          <td>{{ company.Email }}</td>
          <td>
            <button class="btn-info" @click="viewCompanyInfo(company.idcompany, company.Name)">Ver Información</button>
          </td>
          <td>
            <div class="button-group">
            <button class="btn-edit" @click="editCompany(company.idcompany)"><i class="fas fa-edit"></i>Modificar</button>
            <button class= "btn-delete" @click="deleteCompany(company.idcompany)" > 
              <i class="fas fa-trash-alt"></i>Eliminar</button>
          </div>
          </td>
        </tr>
      </tbody>
    </table>
    <div class="invoice-buttons">
      <button class="button-nuevoPrestadorSP" @click="companyForm">+ Nuevo Prestador de Servicio Público</button>
      <button class="back-button" type="button" @click="goBack">Regresar</button>
    </div>
  </div>
</div>
</template>

<script>
import axios from 'axios';

export default {
  data() {
    return {
      companies: []
    };
  },
  created() {
    this.fetchCompanies();
  },
  methods: {
    async fetchCompanies() {
      try {
        const token = localStorage.getItem('token');
        const response = await axios.get('/api/AllCompany', {
          headers: {
            Authorization: `Bearer ${token}`
          }
        });
        this.companies = response.data;
      } catch (error) {
        console.error('Error fetching companies:', error);
      }
    },
    companyForm() {
      this.$router.push('/api/register');
    },
    async deleteCompany(idcompany) {
      const confirmDelete = confirm("¿Estás seguro de que deseas eliminar esta Entidad?");
      if (confirmDelete) {
        try {
          const token = localStorage.getItem('token');
          await axios.delete(`/api/companies/${idcompany}`, {
            headers: {
              Authorization: `Bearer ${token}`
            }
          });
          // Actualizar la lista de empresas en el frontend sin hacer otra solicitud
      this.companies = this.companies.filter(company => company.idcompany !== idcompany);
        } catch (error) {
          console.error('Error al eliminar esta Entidad:', error);
          alert('Hubo un error al intentar eliminar la entidad.');
        }
      }
    },
    editCompany(idcompany) {
      this.$router.push(`/api/company/edit/${idcompany}`);
    },
    viewCompanyInfo(idcompany, name) {
      this.$router.push({
        path: `/api/company/info/${idcompany}`,
        query: { name }
      });
    },
    goBack() {
      this.$router.go(-1); // Regresa a la página anterior (cualquiera)
    }
  }
};
</script>

<style scoped>

body {
  background-color: #ffffff; /* Fondo sólido para toda la página */
  margin: 0;
  padding: 0;
  font-family: 'Roboto', sans-serif; /* Tipo de letra */
  font-size: 11px; /* Tamaño de letra general */
}

/* Estilos personalizados */
.invoice-container {
  background-color: #f3f3f3; /*color de formulario */
  padding: 20px;
  margin: 50px auto;
  width: 80%;
  border-radius: 10px;
  box-shadow: 0px 4px 8px rgba(0, 0, 0, 0.1);
}

.invoice-header {
  display: flex;  
  justify-content: space-between; /* Alinea los elementos a los extremos */
  background-color: #b7daee;
  padding: 5px; /* Reducir el padding */
  border-radius: 15px;
  margin-bottom: 10px;
  font-size: 12px; /* Hacer el texto más pequeño */
}

.invoice-buttons {
  display: flex;
  justify-content: center; /* Centrar los botones */
  margin-top: 20px; /* Espacio superior para los botones */
}

.button-nuevoPrestadorSP, .back-button {
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

.button-nuevoPrestadorSP:hover, .back-button:hover {
  background-color: #2980b9; /* Color al pasar el mouse */
  transform: scale(1.05); /* Efecto de aumento */
}

table {
  width: 100%;
  border-collapse: collapse;
  margin: 0 auto;
  font-size: 14px;
}
.usuario-table {
  width: 100%;
  border-collapse: collapse;
  margin: 0 auto;
  font-size: 11px; /* Hacer el tamaño de la letra de la tabla más pequeño */
}

th, td {
  border: 1px solid #ddd;
  padding: 10px;
  text-align: center;
}

.btn-info {
  background-color: #e77325;
  color: white;
  border: none;
  padding: 5px 10px;
  cursor: pointer;
  border-radius: 5px;
}

.btn-info:hover {
  background-color: #d35400;
}

.button-group {
  display: flex;
  gap: 15px; /* Ajusta el espacio entre los botones aquí */
}

.btn-edit {
  background-color: #2980b9;
  color: white;
  border: none;
  padding: 5px 10px;
  cursor: pointer;
  border-radius: 5px;
}

.btn-delete {
  background-color: #e74c3c;
  color: white;
  padding: 5px 10px;
  border: none;
  border-radius: 4px;
  cursor: pointer;
  margin-right: 5px;
}

.btn-edit:hover {
  background-color: #022ebd;
}

.btn-delete:hover {
  background-color: #c0392b;
}

.background-container {
  background-image: url('https://cdn.leonardo.ai/users/65a8cf55-c959-4394-91b9-30d6f5167b8c/generations/3b56f729-de70-45a8-a5a0-6c274211025a/Leonardo_Phoenix_A_modern_administrative_dashboard_webpage_fea_1.jpg');
  background-size: cover; /* Ajusta la imagen para cubrir todo el contenedor */
  background-position: center; /* Centra la imagen */
  background-repeat: no-repeat; /* Evita que la imagen se repita */
  height: 100vh; /* O ajusta a la altura deseada */
  display: flex;
  justify-content: center; /* Centrar el contenido */
  align-items: center; /* Centrar verticalmente */
}
</style>