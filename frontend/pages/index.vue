<template>
    <div>
		<h3>未完了</h3>
		<div v-for="n in range">
			<div v-if="posts[n-1].status !== 1">
				id: {{ posts[n-1].id }},
				todo: {{ posts[n-1].todo }}
			</div>
		</div>

		<h3>完了</h3>
		<div v-for="n in range">
			<div v-if="posts[n-1].status === 1">
				id: {{ posts[n-1].id }},
				todo: {{ posts[n-1].todo }}
			</div>
		</div>
		<div>
		<h3>元JSON</h3>
		{{ posts }}
		</div>
    </div>
</template>

<script lang="ts">
import Vue from 'vue'

export default Vue.extend({
    async asyncData({ $axios }) {
		const resArray: any = [];
		// 取得先のURL
		for ( let i=1; i<100; i++) {
			const url = "http://localhost:4000/v1/todo/" + i;
			// リクエスト（Get）
			const response = await $axios.$get(url);
			const resJSON = JSON.parse(JSON.stringify(response));
			if (resJSON.todo === undefined) {
				break;
			} else {
				resArray.push(response)
			}
		}
        
		// 配列で返ってくるのでJSONにして返却
		return {
			posts: JSON.parse(JSON.stringify(resArray)),
			range: resArray.length
		};
	}
})
</script>
