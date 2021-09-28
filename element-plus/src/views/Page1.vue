<template>
<Layout>
  <template v-slot:main>
    <h1>Page1 ......</h1>
  </template>

</Layout>
</template>

<script>
// @ is an alias to /src

import Layout from "../views/Layout.vue";

export default {
  name: 'Page1',
  components: {
    Layout
  }
}
</script>
