<template>
  <Layout>
    <div class="container mx-auto  md:px-24 md:py-16" xmlns:x-transition="http://www.w3.org/1999/xhtml">
      <div class="flex flex-col justify-center items-center">
        <h1 class="text-6xl my-16">404</h1>
        <h2>Kunne dessverre ikke finne siden du leter etter.</h2>
        <h4 class="underline text-pink-500"><inertia-link href="/">Gå til forsiden</inertia-link></h4>
      </div>
    </div>
  </Layout>

</template>

<script>
import Layout from '../../Shared/Layout'

export default {
  components: {
    Layout
  },

}
</script>

<style scoped>

</style>