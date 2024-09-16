<script lang="ts">
import SearchComponent from './SearchComponent.vue'
import TreeComponent from './TreeComponent.vue'
import Slot from '../slot'
import Toolbox from '../toolbox'
import { defineComponent } from 'vue';


export default defineComponent({
  name: 'MainComponent',
  components: {
	SearchComponent,
	TreeComponent
  },
  data() {
    return {
      data: [] as Toolbox[],
    };
  },
  created() {
    this.fetchData();
  },
  methods: {
    async fetchData() {
      try {
        const response = await fetch('http://localhost:8080/');
        if (!response.ok) {
          throw new Error(`HTTP error! status: ${response.status}`);
        }
        const jsonData: Toolbox[] = await response.json();
        this.data = jsonData;
		console.log(this.data)
      } catch (error) {
        console.error('Error fetching data:', error);
      }
    },
  },
});


</script>

<template>
	<section class="section">
		<div class="container is-fullhd">
			<div class="notification is-warning">
				<SearchComponent />
			</div>
		</div>
	</section>

	<section class="section">
		<div class="container is-fullhd">
			<div v-if="data.length" class="notification is-warning">
				<TreeComponent :toolboxes="data"/>
			</div>
		</div>

	</section>

</template>

<style scoped>

</style>