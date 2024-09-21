<template>
  <div class="invoice-container">
    <div class="invoice-header">
      <h2>Área de Administrador / Empresas</h2>
    </div>
    <div class="invoice-buttons">
      <button class="button-nuevoPrestadorSP" @click="companyForm">+ Nuevo Prestador de Servicio Público</button>
    </div>
    <table>
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
            <button class="btn-edit" @click="editCompany(company.idcompany)">Modificar</button>
            <button class= "btn-delete" @click="deleteCompany(company.idcompany)">Eliminar</button>
          </div>
          </td>
        </tr>
      </tbody>
    </table>
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
    }
  }
};
</script>

<style scoped>
/* Estilos personalizados */
.invoice-container {
  background-color: #f3f3f3;
  padding: 20px;
  margin: 50px auto;
  width: 80%;
  border-radius: 10px;
  box-shadow: 0px 4px 8px rgba(0, 0, 0, 0.1);
}

.invoice-header {
  background-color: #62b5ec;
  padding: 15px;
  border-radius: 5px;
  margin-bottom: 20px;
  font-size: 15px;
}

.invoice-buttons {
  display: flex;
  gap: 10px;
  margin-bottom: 20px;
}

.button-nuevoPrestadorSP {
  background-color: #007bff;
  color: white;
  border: none;
  padding: 10px 20px;
  font-size: 14px;
  border-radius: 4px;
  cursor: pointer;
}

table {
  width: 100%;
  border-collapse: collapse;
  font-size: 14px;
}

th, td {
  border: 1px solid #ddd;
  padding: 10px;
  text-align: center;
}

.btn-info {
  background-color: #d35400;
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
}

.btn-edit:hover {
  background-color: #d35400;
}

.btn-delete:hover {
  background-color: #c0392b;
}
</style>