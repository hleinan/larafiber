<template>
  <Layout>
    <div class="container mx-auto px-6 md:px-12 lg:px-24 xl:px-48 mt-6">
      <h1 class="text-4xl mb-8 text-center">Your tasks</h1>
      <form class="pt-6 pb-8 mb-4 mr-2" @submit.prevent="submit">
        <input v-model="form.text"
           name="task"
           class="mb-2 shadow appearance-none border rounded w-full py-4 px-6 text-2xl text-gray-700 leading-tight focus:outline-none focus:shadow-outline"
           id="task"
           type="text"
           placeholder="Add task"
        >
      </form>
      <ul v-if="tasks?.length > 0" class="ml-8">
        <li @click="updateStatus(task.uuid)" v-for="task in tasks" :key="task.uuid" class="text-xl text-gray-600 mb-2 cursor-pointer select-none transition-all duration-250" :class="{'line-through text-orange-500': task.done }">
          {{task.text}}
        </li>
      </ul>
      <div v-else class="text-gray-500 text-center mg-8 text-xl">
        You have no tasks. Start by creating one
      </div>
    </div>
  </Layout>
</template>

<script>
import Layout from '../../Shared/Layout'

import { computed } from 'vue'
import { usePage } from '@inertiajs/inertia-vue3'

export default {
  name: "Task",
  setup() {
    const auth = computed(() => usePage().props.value.auth)
    return {auth}
  },
  components: {
    Layout,
  },
  data() {
    return {
      form: {
        text: ""
      }
    }
  },
  props: {
    title: String,
    tasks: Array,
  },
  methods: {
    submit() {
      if (this.form.text.trim()) {
        this.$inertia.post("/task", this.form).then(() => {
          this.form.text = "";
        });
      }
    },
    updateStatus(uuid) {
      this.$inertia.post("/task/change/" + uuid, this.form)
    },
  }
}
</script>

<style scoped>

</style>