<template>
    <header>
      <nav>
        <ul>
          <li><a href="#hero">Home</a></li>
          <li><a href="#features">Features</a></li>
          <li><a href="#contact">Contact</a></li>
        </ul>
      </nav>
    </header>
  </template>
  
  <script>
  export default {
    name: 'Header'
  }
  </script>
  
  <style scoped>
  header {
    background: #333;
    color: white;
    padding: 10px 20px;
  }
  
  nav ul {
    list-style: none;
    padding: 0;
    display: flex;
  }
  
  nav li {
    margin-right: 20px;
  }
  
  nav a {
    color: white;
    text-decoration: none;
  }
  </style>