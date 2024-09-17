<script lang="ts">
import { defineComponent, type PropType } from 'vue';
import type Slot from '../slot'
import type Toolbox from '../toolbox'

export default defineComponent({
  name: 'SearchComponent',
  props: {
    toolboxes: {
      type: Array as PropType<Toolbox[]>, // Define the prop type
      required: true, // This prop must be provided
    },
  },
});
</script>

<template>
	<input class="input is-large" type="text" placeholder="Search">
</template>

<style scoped>

</style>