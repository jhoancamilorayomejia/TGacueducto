<template>
    <div class="invoice-container">
      <div class="invoice-header">
        <h2>Área de Gestión de Clientes</h2>
      </div>
      <div class="invoice-buttons">
        <button class="button-nuevoPrestadorSP" @click="UsuarioForm">+ Nuevo Cliente</button>
      </div>
      <table>
        <thead>
          <tr>
            <th>Nombre del Titular</th>
            <th>Apellidos</th>
            <th>Localidad</th>
            <th>Teléfono</th>
            <th>Email</th>
            <th>Acciones</th>
          </tr>
        </thead>
        <tbody>
          <tr v-for="usuario in usuarios" :key="usuario.idcustomer">
            <td>
              <input v-model="usuario.name" />
            </td>
            <td>
              <input v-model="usuario.last_name" />
            </td>
            <td>
              <input v-model="usuario.address" />
            </td>
            <td>
              <input v-model="usuario.phone" />
            </td>
            <td>
              <input v-model="usuario.email" />
            </td>
            <td>
              <button @click="updateCustomer(usuario)" class="btn-save">Guardar</button>
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
        usuarios: []
      };
    },
    created() {
      this.fetchUsuarios();
    },
    methods: {
      async fetchUsuarios() {
        try {
          const token = localStorage.getItem('token');
          const response = await axios.get('/api/customer', {
            headers: {
              Authorization: `Bearer ${token}`
            }
          });
          this.usuarios = response.data;
        } catch (error) {
          console.error('Error fetching usuarios:', error);
        }      
      },
      UsuarioForm() {
        // Redireccionar a la página de registro de cliente
        this.$router.push('/api/registerCustomer');
      },
      async updateCustomer(customer) {
        try {
          const token = localStorage.getItem('token');
          await axios.put(`/api/customer/${customer.idcustomer}`, customer, {
            headers: {
              Authorization: `Bearer ${token}`
            }
          });
          alert('Cliente actualizado exitosamente.');
        } catch (error) {
          console.error('Error al actualizar cliente:', error);
          alert('Hubo un error al intentar actualizar el cliente.');
        }
      },
      async deleteCustomer(idcustomer, idcompany) {
        const confirmDelete = confirm("¿Estás seguro de que deseas eliminar este cliente?");
        if (confirmDelete) {
          try {
            const token = localStorage.getItem('token');
            await axios.delete(`/api/customer/${idcustomer}/${idcompany}`, {
              headers: {
                Authorization: `Bearer ${token}`
              }
            });
            this.fetchUsuarios(); // Volver a obtener la lista de clientes después de eliminar
          } catch (error) {
            console.error('Error al eliminar cliente:', error);
            alert('Hubo un error al intentar eliminar al cliente.');
          }
        }
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
  
  .invoice-container {
    background-color: #f3f3f3; /*color de formulario */
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
    transition: background-color 0.3s ease;
  }
  
  .button-nuevoPrestadorSP:hover {
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
  
  input {
    border: 1px solid #ccc;
    padding: 5px;
    border-radius: 4px;
    width: 100%;
    box-sizing: border-box;
  }
  
  .btn-save {
    background-color: #28a745;
    color: white;
    border: none;
    padding: 5px 10px;
    cursor: pointer;
    border-radius: 5px;
  }
  
  .btn-save:hover {
    background-color: #218838;
  }
  

  </style>