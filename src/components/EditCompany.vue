<template>
  <div class="background-container"> <!-- Contenedor para el fondo -->
    <div class="invoice-container">
      <div class="invoice-header">
        <h2>Área de Administrador / Empresas</h2>
      </div>
      <div class="invoice-buttons">
      
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
      <div class="invoice-buttons">
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
      },
      goBack() {
      this.$router.go(-1); // Regresa a la página anterior (cualquiera)
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
    background-color: #f3f3f3; /* Color de formulario */
  padding: 50px;
  margin: 90px auto;
  width: 80%; /* Hacer el contenedor más pequeño */
  border-radius: 10px;
  box-shadow: 0 8px 16px rgba(0, 0, 0, 0.2); /* Sombra más pronunciada */
  position: relative; /* Para colocar la sombra en la parte trasera */
  }
  
  .invoice-header {
    display: flex;
  justify-content: space-between; /* Alinea los elementos a los extremos */
  background-color: #b7daee;
  padding: 5px; /* Reducir el padding */
  border-radius: 15px;
  margin-bottom: 10px;
  font-size: 11px; /* Hacer el texto más pequeño */
  }
  
  .invoice-buttons {
  display: flex;
  justify-content: center; /* Centrar los botones */
  margin-top: 20px; /* Espacio superior para los botones */
}

.back-button {
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

 .back-button:hover {
  background-color: #2980b9; /* Color al pasar el mouse */
  transform: scale(1.05); /* Efecto de aumento */
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