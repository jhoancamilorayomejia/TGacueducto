<template>
    <div class="invoice-container">
      <div class="invoice-header">
        <h2>Área de Administrador / Empresas</h2>
      </div>
      <div class="invoice-buttons">
        <button class="btn" @click="companyForm">Registrar Empresa de Acueducto</button>
        <button class="btn" @click="AllCompany">Ver Registro de Empresas</button>
      </div>
      <table>
        <thead>
          <tr>
            <th>NIT</th>
            <th>Nombre</th>
            <th>Dirección</th>
            <th>Teléfono</th>
            <th>Email</th>
            <th>Acción</th> <!-- Nueva columna para acciones -->
          </tr>
        </thead>
        <tbody>
          <tr v-for="company in companies" :key="company.IDcompany">
            <td>{{ company.nit }}</td>
            <td>
              <input v-model="company.Name" />
            </td>
            <td>
              <input v-model="company.address" />
            </td>
            <td>
              <input v-model="company.Phone" />
            </td>
            <td>
              <input v-model="company.Email" />
            </td>
            <td>
              <button @click="updateCompany(company)" class="btn-edit">Guardar</button> <!-- Botón para guardar cambios -->
            
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
      async updateCompany(company) {
        try {
          const token = localStorage.getItem('token');
          await axios.put(`/api/companies/${company.IDcompany}`, company, {
            headers: {
              Authorization: `Bearer ${token}`
            }
          });
          alert('Empresa actualizada exitosamente.');
        } catch (error) {
          console.error('Error al actualizar empresa:', error);
          alert('Hubo un error al intentar actualizar la empresa.');
        }
      },
      async deleteCompany(IDcompany) {
        const confirmDelete = confirm("¿Estás seguro de que deseas eliminar esta empresa?");
        if (confirmDelete) {
          try {
            const token = localStorage.getItem('token');
            await axios.delete(`/api/companies/${IDcompany}`, {
              headers: {
                Authorization: `Bearer ${token}`
              }
            });
            this.fetchCompanies(); // Volver a obtener la lista de empresas después de eliminar
          } catch (error) {
            console.error('Error al eliminar empresa:', error);
            alert('Hubo un error al intentar eliminar la empresa.');
          }
        }
      },
      companyForm() {
        this.$router.push('/api/register');
      },
      AllCompany() {
        this.$router.push('/api/AllCompany');
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
  
  .btn {
    background-color: #007bff;
    color: white;
    border: none;
    padding: 10px 20px;
    font-size: 14px;
    border-radius: 4px;
    cursor: pointer;
    transition: background-color 0.3s ease;
  }
  
  .btn:hover {
    background-color: #0056b3;
  }
  
  table {
    width: 100%;
    border-collapse: collapse;
    margin: 0 auto;
    font-size: 14px;
  }
  
  th, td {
    border: 1px solid #ddd;
    padding: 10px;
    text-align: center;
  }
  
  th {
    background-color: #f2f2f2;
  }
  
  .btn-edit {
    background-color: #28a745;
    color: white;
    border: none;
    padding: 5px 10px;
    cursor: pointer;
    border-radius: 5px;
  }
  
  .btn-edit:hover {
    background-color: #218838;
  }
  
 
  </style>