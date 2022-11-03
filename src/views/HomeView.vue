<template>
  <div class="home">
    <div class="flex h-screen">
      <div class="m-auto">
        <Card CardTitle="Todo-App"/>
      </div>
    </div>
  </div>
</template>

<script>
// @ is an alias to /src
import Card from '@/components/Card.vue';

export default {
  components: {
    Card,
  }
}
</script>
