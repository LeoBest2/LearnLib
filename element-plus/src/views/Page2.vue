<template>
<Layout>
  <template v-slot:main>
    <h1>Page2 ......</h1>
  </template>

</Layout>
</template>

<script>
// @ is an alias to /src

import Layout from "../views/Layout.vue";

export default {
  name: 'Page2',
  components: {
    Layout
  }
}
</script>
