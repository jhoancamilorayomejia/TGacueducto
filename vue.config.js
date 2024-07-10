module.exports = {
  devServer: {
    proxy: {
      '/api': {
        target: 'http://localhost:8081', // El backend Go debería estar corriendo en este puerto
        changeOrigin: true
      }
    },
    port: 8080
  }
};



/*module.exports = {
  devServer: {
    port: 8081
  }
};*/