<template>
<Layout>
  <template v-slot:main>
    <h1>About page!</h1>
  </template>

</Layout>
</template>

<script>
// @ is an alias to /src

import Layout from "../views/Layout.vue";

export default {
  name: 'Home',
  components: {
    Layout
  }
}
</script>
