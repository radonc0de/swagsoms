<template>
	<aside class="menu">
	<p class="menu-label">Toolboxes</p>
	<ul class="menu-list">
		<li v-for="toolbox in toolboxes">
			<a>
				{{  toolbox.address }}
			</a>
			<ul v-if="toolbox.slots.length">
				<li v-for="slot in toolbox.slots">
					<a>
						{{slot.address }} {{ slot.tags}}
					</a>
					<ul v-if="slot.slots.length">
						<li v-for="subslot in slot.slots">
							<a>
								{{subslot.address }} {{ subslot.tags}}
							</a>
						</li>
					</ul>
				</li>
			</ul>
		</li>
	</ul>
	</aside>	

</template>

<script lang="ts">
import { defineComponent, PropType } from 'vue';
import Slot from '../slot'
import Toolbox from '../toolbox'

export default defineComponent({
  name: 'TreeComponent',
  props: {
    toolboxes: {
      type: Array as PropType<Toolbox[]>, // Define the prop type
      required: true, // This prop must be provided
    },
  },
});
</script>