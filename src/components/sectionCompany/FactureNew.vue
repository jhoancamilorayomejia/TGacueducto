<template>
  <div class="facture-new-container">
    <h2>{{ userName }}, Nueva Factura Para {{ customerName }}</h2>
    <form @submit.prevent="createFacture">
      <!-- Campo para ID de Compañía -->
      <div class="form-group">
        <label for="idcompany">ID de Compañía</label>
        <input 
          type="number" 
          v-model.number="facture.idcompany" 
          required 
          placeholder="Ingresa el ID de la Compañía"
          readonly  
        />
      </div>

      <!-- Campo para ID de Cliente -->
      <div class="form-group">
        <label for="idcustomer">ID de Cliente</label>
        <input 
          type="number" 
          v-model.number="facture.idcustomer" 
          required 
          placeholder="Ingresa el ID del Cliente"
          readonly  
        />
      </div>

      <!-- Campo para Número de Factura -->
      <div class="form-group">
        <label for="facturenumber">Número de Factura</label>
        <input 
          type="text" 
          v-model="facture.facturenumber" 
          required 
          placeholder="Ingresa el Número de Factura"
        />
      </div>

      <!-- Campo para Fecha de Emisión -->
      <div class="form-group">
        <label for="datecreation">Fecha de Emisión</label>
        <input 
          type="date" 
          v-model="facture.datecreation" 
          required 
        />
      </div>

      <!-- Campo para Fecha Límite de Pago -->
      <div class="form-group">
        <label for="datepayment">Fecha Límite de Pago</label>
        <input 
          type="date" 
          v-model="facture.datepayment" 
          required 
        />
      </div>

      <!-- Campo para Monto Total -->
      <div class="form-group">
        <label for="totalpay">Monto Total</label>
        <input 
          type="text" 
          v-model="facture.totalpay" 
          required 
          placeholder="Ingresa el Monto Total"
        />
      </div>

      <button type="submit">Crear Factura</button>
    </form>
  </div>
</template>

<script>
import axios from 'axios';

export default {
  data() {
    return {
      userID: '',
      customerId: this.$route.params.idcustomer, 
      facture: { 
        idcompany: 0, // Inicializamos idcompany como número
        idcustomer: 0, // Inicializamos idcustomer como número
        facturenumber: '',
        datecreation: '',
        datepayment: '',
        totalpay: '', // Asegurarse de que totalpay sea un string
      },
      customerName: this.$route.query.name, 
      userName: '', 
    };
  },
  created() {
    this.userName = localStorage.getItem('userName'); 
    this.userID = localStorage.getItem('userID');
    this.facture.idcompany = parseInt(this.userID); // Convertimos userID a entero y asignamos a idcompany
    this.facture.idcustomer = parseInt(this.customerId); // Convertimos customerId a entero y asignamos a idcustomer
  },
  methods: {
    async createFacture() {
      const confirmed = confirm("¿Estás seguro de que deseas crear una nueva factura?");
      if (!confirmed) return;

      try {
        const response = await axios.post(`/api/facturas`, this.facture);
        console.log(response.data); // Verificar la respuesta
        this.$router.push({
          name: 'FactureInfoCustomer',
          params: { idcustomer: this.facture.idcustomer },
          query: { name: this.customerName },
        });
      } catch (error) {
        console.error('Error creando la factura:', error.response?.data); // Mostrar el error completo
        alert('Error creando la factura: ' + (error.response?.data?.error || error.message));
      }
    },
    validateNumberInput(event) {
      const value = event.target.value;
      if (!/^\d*$/.test(value)) {
        event.target.value = value.replace(/\D/g, '');
      }
    }
  }
};
</script>

<style scoped>
.facture-new-container {
  background-color: #f3f3f3;
  padding: 20px;
  margin: 50px auto;
  width: 80%;
  border-radius: 10px;
  box-shadow: 0px 4px 8px rgba(0, 0, 0, 0.1);
}

h2 {
  text-align: center;
}

.form-group {
  margin-bottom: 15px;
}

label {
  display: block;
  margin-bottom: 5px;
}

input {
  width: 100%;
  padding: 8px;
  margin: 5px 0;
  box-sizing: border-box;
  border: 1px solid #ccc;
  border-radius: 4px;
}

button {
  display: block;
  width: 100%;
  padding: 10px;
  background-color: #4CAF50;
  color: white;
  border: none;
  border-radius: 5px;
  cursor: pointer;
}

button:hover {
  background-color: #45a049;
}
</style>





  
  
  
  
  
  
  