<template>
  <Layout>
    <div class="container mx-auto px-6 md:px-12 lg:px-24 xl:px-48 mt-6">
      <h1 class="text-4xl mb-8 text-center">Welcome to Larafiber</h1>
      <p class="text-center mb-8">Create, see and edit your tasks</p>
      <p class="text-center"><inertia-link href="/task" class="underline text-blue-500">See your tasks</inertia-link></p>
    </div>
  </Layout>
</template>

<script>
import Layout from '../Shared/Layout'

import { computed } from 'vue'
import { usePage } from '@inertiajs/inertia-vue3'

export default {
  name: "Home",
  setup() {
    const auth = computed(() => usePage().props.value.auth)
    return {auth}
  },
  components: {
    Layout,
  },
  data() {
    return {
      timer: null,
      currentIndex: 0,
    }
  },
  props: {
    title: String,
    tags: Array,
    artists: Array,
    lecturers: Array,
    news: Array,
  },
  methods: {
    startSlide() {
      this.timer = setInterval(this.next, 5000)
    },
    next() {
      this.currentIndex += 1;
    },
    prev() {
      this.currentIndex -= 1;
    },
    pauseSlide() {
      clearInterval(this.timer)
    },
  },
  computed: {
    currentImg() {
      return this.artists[Math.abs(this.currentIndex) % this.artists.length]
    }
  },
  mounted() {
    this.startSlide()
  },
}
</script>

<style scoped>

</style>